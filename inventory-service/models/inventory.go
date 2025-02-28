package models

type InventoryUpdate struct {
	ProductID   string  `jsonapi:"attr,product_id"`
	StockAdded  int     `jsonapi:"attr,stock_added" validate:"gt=0"` // Positive for restock, Negative for sales //New
	Name        string  `jsonapi:"attr,name"`
	Description string  `jsonapi:"attr,description" `
	Price       float64 `jsonapi:"attr,price" validate:"gt=0"`
	Quantity    int     `jsonapi:"attr,quantity" validate:"gt=0"`
}

// NewInventoryUpdate initializes InventoryUpdate with default values
func NewInventoryUpdate() *InventoryUpdate {
	return &InventoryUpdate{
		StockAdded: 0,
		Price:      0.0,
		Quantity:   0,
	}
}

type Response struct {
	ID      string `jsonapi:"primary,success_response"`
	Message string `jsonapi:"attr,message"` // Message for response
}
