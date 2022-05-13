package order

import (
	"context"

	"github.com/eifzed/ares/internal/model/order"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *orderDB) GetCustomerOrderList(ctx context.Context, customerID primitive.ObjectID, filter order.GetOrdersParams) ([]order.OrderList, error) {
	aggPipeline := []bson.M{}
	aggPipeline = append(aggPipeline, bson.M{"customer_id": customerID})
	if !filter.BusinessID.IsZero() {
		aggPipeline = append(aggPipeline, bson.M{"$match": bson.M{"business_id": filter.BusinessID}})
	}

	if filter.Status != "" {
		aggPipeline = append(aggPipeline, bson.M{"$match": bson.M{"status": filter.Status}})
	}

	result, err := db.DB.Collection("orders").Aggregate(ctx, aggPipeline)
	if err != nil {
		return nil, err
	}
	var data []order.OrderList
	err = result.Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *orderDB) InsertOrder(ctx context.Context, order *order.Order) error {
	_, err := db.DB.Collection("orders").InsertOne(ctx, order)
	if err != nil {
		return err
	}
	return nil
}
