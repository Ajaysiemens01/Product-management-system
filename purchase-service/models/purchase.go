package models

type PurchaseUpdate struct {
    ProductID string `jsonapi:"attr,product_id"`
    Change    int    `jsonapi:"attr,change" validate:"required"` // Positive for restock, Negative for sales
}


type Response struct {
    ID      string `jsonapi:"primary,success_response"`
    Message string `jsonapi:"attr,message"` // Message for response
}