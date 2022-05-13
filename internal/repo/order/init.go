package order

import "go.mongodb.org/mongo-driver/mongo"

type orderDB struct {
	DB *mongo.Database
}

type OrderDBOption struct {
	DB *mongo.Database
}

func GetNewOrderDB(option *OrderDBOption) *orderDB {
	if option == nil || option.DB == nil {
		return nil
	}
	return &orderDB{
		DB: option.DB,
	}
}
