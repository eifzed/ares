package order

import (
	"context"

	"github.com/eifzed/ares/internal/model/order"
	tx "github.com/eifzed/ares/lib/database/mongodb/transaction"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *orderDB) InsertBusiness(ctx context.Context, businessData *order.BusinessData) error {
	session := tx.GetSessionFromContext(ctx)
	err := mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		var collection *mongo.Collection
		if session != nil {
			collection = tx.GetCollectionFromSession(session, "business")
		} else {
			collection = db.DB.Collection("business")
		}
		result, err := collection.InsertOne(sc, businessData)
		if err != nil {
			return err
		}
		businessData.ID = result.InsertedID.(primitive.ObjectID)
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (db *orderDB) CheckUserAlreadyHasBusiness(ctx context.Context, userID primitive.ObjectID) (bool, error) {
	count, err := db.DB.Collection("business").CountDocuments(ctx, bson.M{"owner_id": userID})

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (db *orderDB) GetBulkProductPriceByIDs(ctx context.Context, productIDs []primitive.ObjectID) (map[primitive.ObjectID]uint32, error) {
	// db.DB.Collection("business").
	return nil, nil
}
func (db *orderDB) GetBusinessList(ctx context.Context, filter order.GenericFilterParams) ([]order.BusinessList, error) {
	// TODO: handle keyword
	aggPipeline := []bson.M{}
	aggPipeline = append(aggPipeline, bson.M{"$project": bson.M{"_id": 1, "name": 1, "address": 1, "photo_URL": 1}})
	skip := filter.Limit * filter.Page
	if skip > 0 {
		aggPipeline = append(aggPipeline, bson.M{"$skip": skip})
	}
	if filter.Limit > 0 {
		aggPipeline = append(aggPipeline, bson.M{"$limit": filter.Limit})
	}
	cursor, err := db.DB.Collection("business").Aggregate(ctx, aggPipeline)
	if err != nil {
		return nil, err
	}
	result := []order.BusinessList{}
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *orderDB) GetBusinessDetail(ctx context.Context, businessID primitive.ObjectID) (*order.BusinessDetail, error) {
	data := []*order.BusinessDetail{}
	aggPipeline := []bson.M{{"$match": bson.M{"_id": businessID}}, {"$lookup": bson.M{"from": "products", "localField": "product_ids", "foreignField": "_id", "as": "products"}}}
	cursor, err := db.DB.Collection("business").Aggregate(ctx, aggPipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	return data[0], nil
}

func (db *orderDB) GetBulkProductByProductIDs(ctx context.Context, productIDs []primitive.ObjectID) ([]order.Products, error) {
	data := []order.Products{}
	cursor, err := db.DB.Collection("business").Find(ctx, bson.M{"products._id": bson.M{"$in": productIDs}})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
