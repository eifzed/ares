package order

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID       primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID    primitive.ObjectID `bson:"customer_id,omitempty"`
	BusinessID    primitive.ObjectID `bson:"business_id,omitempty"`
	OrderProducts []OrderProduct     `bson:"order_products,omitempty"`
	CustomerNote  string             `bson:"customer_note,omitempty"`
}

type OrderProduct struct {
	ProductID       primitive.ObjectID `bson:"product_id,omitempty"`
	Quantity        uint32             `bson:"quantity,omitempty"`
	PricePerItemIDR uint32             `bson:"price_per_item_IDR,omitempty"`
}

type GetOrdersParams struct {
	Status            string
	BusinessID        primitive.ObjectID
	StartOrderDate    string
	EndOrderDate      string
	SortDateAscending bool
}
