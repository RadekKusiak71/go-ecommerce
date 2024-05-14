package types

import "time"

type Order struct {
	ID         int       `json:"ID"`
	AccountID  int       `json:"account_id"`
	Address    string    `json:"address"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}

type OrderPayload struct {
	AccountID  int     `json:"account_id"`
	Address    string  `json:"address"`
	TotalPrice float64 `json:"total_price"`
}

type OrderItem struct {
	ID         int     `json:"id"`
	OrderID    int     `json:"order_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type OrderItemPayload struct {
	OrderID    int     `json:"order_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type OrderDetails struct {
	Order    *Order
	Products []*CartItem
}

type OrderStore interface {
	GetOrders() ([]*Order, error)
	GetOrderByID(int) (*Order, error)
	GetUserOrders(int) ([]*Order, error)
	GetFullOrderDetails(int) (*OrderDetails, error)
}
