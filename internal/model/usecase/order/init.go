package order

import (
	"context"

	"github.com/eifzed/ares/internal/model/order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderUCInterface interface {
	GetUserOrderList(ctx context.Context, filter order.GetOrderListParams) ([]order.OrderList, error)
	RegisterOrder(ctx context.Context, order *order.Order) error

	RegisterBusiness(ctx context.Context, params order.BusinessDetail) error
	GetBusinessList(ctx context.Context, params order.GenericFilterParams) (*order.BusinessListData, error)
	GetBusinessDetail(ctx context.Context, businessID primitive.ObjectID) (*order.BusinessDetail, error)
}
