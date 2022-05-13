package order

import "go.mongodb.org/mongo-driver/bson/primitive"

type BusinessData struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	OwnerID     primitive.ObjectID   `bson:"owner_id,omitempty"`
	Name        string               `bson:"name,omitempty"`
	Address     string               `bson:"address,omitempty"`
	PhoneNumber string               `bson:"phone_number,omitempty"`
	Description string               `bson:"description,omitempty"`
	PhotoURL    string               `bson:"photo_URL,omitempty"`
	ProductIDs  []primitive.ObjectID `bson:"product_ids,omitempty"`
}

type BusinessDetail struct {
	Business BusinessData `bson:",inline"`
	Products []Products   `bson:"products,omitempty"`
}

type Products struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	PriceIDR    uint32             `bson:"price_IDR,omitempty"`
	Description string             `bson:"description,omitempty"`
	PhotoURL    string             `bson:"photo_URL,omitempty"`
}

type GenericFilterParams struct {
	Keyword string
	Limit   uint32
	Page    uint32
}

type BusinessListData struct {
	Total int
	List  []BusinessList
}

type BusinessList struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Address  string             `bson:"address,omitempty"`
	PhotoURL string             `bson:"photo_URL,omitempty"`
}
