package types

type CartItem struct {
	ProductID  int
	Name       string
	Quantity   int
	TotalPrice float64
}

type CartPayload struct {
	AccountID int
	Cart      Cart
}

type Cart []CartItem
