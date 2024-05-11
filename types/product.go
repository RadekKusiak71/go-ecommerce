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
