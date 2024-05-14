package orders

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/RadekKusiak71/goEcom/types"
)

var wg sync.WaitGroup

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}
func (s *Store) GetFullOrderDetails(orderID int) (*types.OrderDetails, error) {
	ch1 := make(chan *types.Order)
	ch2 := make(chan []*types.CartItem)
	wg.Add(2)
	go func() {
		defer wg.Done()
		order, err := s.GetOrderByID(orderID)
		if err != nil {
			log.Println(err)
			return
		}
		ch1 <- order
	}()

	go func() {
		defer wg.Done()
		rows, err := s.db.Query(`
            SELECT products.ID,products.Name,order_item.quantity, order_item.total_price
            FROM order_item JOIN products
            ON products.ID = order_item.product_id
            WHERE order_item.order_id = $1; `, orderID)
		if err != nil {
			log.Println(err)
			return
		}
		items := make([]*types.CartItem, 0)
		for rows.Next() {
			cartItem, err := scanIntoOrderItem(rows)
			if err != nil {
				log.Println(err)
				return
			}
			items = append(items, cartItem)
		}
		ch2 <- items
	}()

	order := <-ch1
	items := <-ch2

	orderDetails := &types.OrderDetails{
		Order:    order,
		Products: items,
	}
	wg.Wait()
	return orderDetails, nil
}

func (s *Store) GetOrders() ([]*types.Order, error) {
	rows, err := s.db.Query("SELECT * FROM orders")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	orders := make([]*types.Order, 0)
	for rows.Next() {
		order, err := scanIntoOrder(rows)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (s *Store) GetOrderByID(orderID int) (*types.Order, error) {
	rows, err := s.db.Query("SELECT * FROM orders WHERE id = $1", orderID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		return scanIntoOrder(rows)
	}
	return nil, fmt.Errorf("order not found")
}

func (s *Store) GetUserOrders(accountID int) ([]*types.Order, error) {
	rows, err := s.db.Query("SELECT * FROM orders WHERE account_id = $1 ", accountID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	orders := make([]*types.Order, 0)
	for rows.Next() {
		order, err := scanIntoOrder(rows)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func scanIntoOrder(rows *sql.Rows) (*types.Order, error) {
	order := new(types.Order)
	err := rows.Scan(
		&order.ID,
		&order.AccountID,
		&order.Address,
		&order.TotalPrice,
		&order.CreatedAt,
	)
	return order, err
}

func scanIntoOrderItem(rows *sql.Rows) (*types.CartItem, error) {
	cartItem := new(types.CartItem)
	err := rows.Scan(
		&cartItem.ProductID,
		&cartItem.Name,
		&cartItem.Quantity,
		&cartItem.TotalPrice,
	)
	return cartItem, err
}
