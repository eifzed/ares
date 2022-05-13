package order

import (
	"context"

	"github.com/eifzed/ares/internal/handler/middleware/auth"
	"github.com/eifzed/ares/internal/model/order"
	"github.com/eifzed/ares/lib/common/commonerr"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *orderUC) GetUserOrderList(ctx context.Context, filter order.GetOrderListParams) ([]order.OrderList, error) {
	userData, isExist := auth.GetUserDataFromContext(ctx)
	if !isExist {
		return nil, commonerr.ErrorForbidden("you do not have permission to access this resource")
	}
	return uc.OrderDB.GetCustomerOrderList(ctx, userData.UserID, order.GetOrdersParams{})
}

func (uc *orderUC) RegisterOrder(ctx context.Context, order *order.Order) error {
	userData, isExist := auth.GetUserDataFromContext(ctx)
	if !isExist {
		return commonerr.ErrorUnauthorized("you are not allowed to register order")
	}
	order.CustomerID = userData.UserID
	productIDs := []primitive.ObjectID{}
	for _, product := range order.OrderProducts {
		productIDs = append(productIDs, product.ProductID)
	}
	products, err := uc.OrderDB.GetBulkProductByProductIDs(ctx, productIDs)
	if err != nil {
		return nil
	}
	mapProductIDPrice := uc.mapProductIDToPrice(products)
	if len(mapProductIDPrice) != len(productIDs) {
		return commonerr.ErrorBadRequest("some products are not found")
	}
	for i := range order.OrderProducts {
		order.OrderProducts[i].PricePerItemIDR = mapProductIDPrice[order.OrderProducts[i].ProductID.Hex()]
	}
	return uc.OrderDB.InsertOrder(ctx, order)
}

func (uc *orderUC) mapProductIDToPrice(products []order.Products) map[string]uint32 {
	mapProductIDPrice := map[string]uint32{}
	for _, product := range products {
		mapProductIDPrice[product.ID.Hex()] = product.PriceIDR
	}
	return mapProductIDPrice
}
