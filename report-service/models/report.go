package models

type ProductReport struct {
    ProductID   string  `jsonapi:"attr,product_id"`
    Name        string  `jsonapi:"attr,name"`
    Price       float64 `jsonapi:"attr,price"`
    Description string   `jsonapi:"attr,description"`
    Quantity    int     `jsonapi:"attr,quantity"`
}
