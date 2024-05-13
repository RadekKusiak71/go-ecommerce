package types

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CategoryID  int       `json:"category_id"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductPayload struct {
	Name        string  `json:"name"`
	CategoryID  int     `json:"category_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductStore interface {
	GetProducts() ([]*Product, error)
	GetProductByID(int) (*Product, error)
	CreateProduct(Product) (string, error)
	UpdateProduct(Product, int) (*Product, error)
	DeleteProduct(int) (string, error)

	CreateCategory(Category) error
	GetCategory(int) (*Category, error)
	GetCategories() ([]*Category, error)
}
