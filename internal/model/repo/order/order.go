package order

import (
	"context"

	"github.com/eifzed/ares/internal/model/order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderDBInterface interface {
	GetCustomerOrderList(ctx context.Context, customerID primitive.ObjectID, filter order.GetOrdersParams) ([]order.OrderList, error)
	InsertOrder(ctx context.Context, order *order.Order) error

	InsertBusiness(ctx context.Context, businessData *order.BusinessData) error
	CheckUserAlreadyHasBusiness(ctx context.Context, userID primitive.ObjectID) (bool, error)
	GetBusinessList(ctx context.Context, filter order.GenericFilterParams) ([]order.BusinessList, error)
	GetBusinessDetail(ctx context.Context, businessID primitive.ObjectID) (*order.BusinessDetail, error)
	GetBulkProductByProductIDs(ctx context.Context, productIDs []primitive.ObjectID) ([]order.Products, error)

	InsertBulkProducts(ctx context.Context, products []order.Products) error
}
