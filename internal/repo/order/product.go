package order

import (
	"context"

	"github.com/eifzed/ares/internal/model/order"
	"github.com/eifzed/ares/lib/database/mongodb/transaction"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *orderDB) InsertBulkProducts(ctx context.Context, products []order.Products) error {
	session := transaction.GetSessionFromContext(ctx)
	err := mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		var collection *mongo.Collection
		if session != nil {
			collection = transaction.GetCollectionFromSession(session, "products")
		} else {
			collection = db.DB.Collection("products")
		}

		insertData := []interface{}{}
		for _, product := range products {
			insertData = append(insertData, product)
		}
		result, err := collection.InsertMany(sc, insertData)
		if err != nil {
			return err
		}
		for i, id := range result.InsertedIDs {
			products[i].ID = id.(primitive.ObjectID)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil

}
