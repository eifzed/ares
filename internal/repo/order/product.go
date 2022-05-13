package order

import (
	"context"

	"github.com/eifzed/ares/internal/model/order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *orderDB) InsertBulkProducts(ctx context.Context, products []order.Products) error {
	insertData := []interface{}{}
	for _, product := range products {
		insertData = append(insertData, product)
	}
	result, err := db.DB.Collection("products").InsertMany(ctx, insertData)
	if err != nil {
		return err
	}
	for i, id := range result.InsertedIDs {
		products[i].ID = id.(primitive.ObjectID)
	}
	return nil

}
