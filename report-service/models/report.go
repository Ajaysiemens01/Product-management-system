package models

type ProductReport struct {
	ID          string  `jsonapi:"primary,product-report"`
	Name        string  `jsonapi:"attr,name"`
	Price       float64 `jsonapi:"attr,price"`
	Description string  `jsonapi:"attr,description"`
	Quantity    int     `jsonapi:"attr,quantity"`
}
