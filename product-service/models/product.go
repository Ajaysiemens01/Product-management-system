package models

type Product struct {
	ID string `jsonapi:"primary,product"`
	Name string `jsonapi:"attr,name" validate:"required"`
	Description string `jsonapi:"attr,description" validate:"required"`
	Price float64 `jsonapi:"attr,price" validate:"required,gt=0"`
	Quantity int `jsonapi:"attr,quantity" validate:"required,gt=0"`
}

type Response struct {
    ID      string `jsonapi:"primary,success_response"`
    Message string `jsonapi:"attr,message"` // Message for response
}
