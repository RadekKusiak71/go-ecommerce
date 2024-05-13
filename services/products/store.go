package products

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/RadekKusiak71/goEcom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	products := make([]*types.Product, 0)
	for rows.Next() {
		product, err := scanIntoProduct(rows)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (s *Store) GetProductByID(productID int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id=$1", productID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		return scanIntoProduct(rows)
	}
	return nil, fmt.Errorf("product not found")
}

func (s *Store) CreateProduct(product types.Product) (string, error) {
	_, err := s.db.Query(`INSERT INTO products ( name, category_id, description, price, quantity )
		VALUES ($1, $2, $3, $4, $5)`, product.Name, product.CategoryID, product.Description, product.Price, product.Quantity)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return "product created", nil
}

func (s *Store) UpdateProduct(product types.Product, productID int) (*types.Product, error) {
	_, err := s.db.Query(`UPDATE products 
	SET name = $1 , category_id = $2 ,description = $3, price = $4 ,  quantity = $5
	WHERE id = $6`, product.Name, product.CategoryID, product.Description, product.Price, product.Quantity, productID)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return s.GetProductByID(productID)
}

func (s *Store) DeleteProduct(productID int) (string, error) {
	res, err := s.db.Exec("DELETE FROM products WHERE id = $1", productID)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if c, err := res.RowsAffected(); err != nil || c == 0 {
		return "", fmt.Errorf("product wasn't found")
	}
	return "product deleted", nil
}

func (s *Store) CreateCategory(cat types.Category) error {
	_, err := s.db.Query("INSERT INTO categories (name) VALUES ($1)", cat.Name)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (s *Store) GetCategories() ([]*types.Category, error) {
	rows, err := s.db.Query("SELECT * FROM categories")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	categories := make([]*types.Category, 0)
	for rows.Next() {
		cat, err := scanIntoCategory(rows)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, nil
}
func (s *Store) GetCategory(catID int) (*types.Category, error) {
	rows, err := s.db.Query("SELECT * FROM categories WHERE id = $1", catID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		return scanIntoCategory(rows)
	}
	return nil, fmt.Errorf("category not found")
}

func scanIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.CategoryID,
		&product.Description,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	return product, err
}

func scanIntoCategory(rows *sql.Rows) (*types.Category, error) {
	category := new(types.Category)
	err := rows.Scan(
		&category.ID,
		&category.Name,
	)
	return category, err
}
