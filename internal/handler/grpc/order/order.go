package order

import (
	"context"
	"fmt"
	"time"

	"github.com/eifzed/ares/internal/config"
	orderDomain "github.com/eifzed/ares/internal/model/order"
	"github.com/eifzed/ares/internal/model/usecase/order"
	"github.com/eifzed/ares/lib/common"
	"github.com/eifzed/ares/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type orderHandler struct {
	pb.OrderServiceServer
	OrderUC order.OrderUCInterface
	Config  *config.Config
}

type Option struct {
	OrderUC order.OrderUCInterface
	Config  *config.Config
}

func GetNewOrderHandler(orderHandlerOption *Option) *orderHandler {
	if orderHandlerOption == nil {
		return nil
	}
	return &orderHandler{
		OrderUC: orderHandlerOption.OrderUC,
		Config:  orderHandlerOption.Config,
	}
}

func (h *orderHandler) RegisterOrder(ctx context.Context, in *pb.RegisterOrderRequest) (*pb.MessageResponse, error) {

	var orderProducts []orderDomain.OrderProduct

	for _, o := range in.OrderProducts {
		oid, err := primitive.ObjectIDFromHex(o.ProductId)
		if err != nil {
			return &pb.MessageResponse{Message: fmt.Sprintf("invalid product ID [%s]", o.ProductId)}, err
		}
		orderProducts = append(orderProducts, orderDomain.OrderProduct{
			ProductID: oid,
			Quantity:  o.Quantity,
		})
	}
	businessID, err := primitive.ObjectIDFromHex(in.BusinessId)
	if err != nil {
		return &pb.MessageResponse{Message: fmt.Sprintf("invalid shop ID [%s]", in.BusinessId)}, err
	}
	err = h.OrderUC.RegisterOrder(ctx, &orderDomain.Order{
		OrderProducts: orderProducts,
		BusinessID:    businessID,
	})
	if err != nil {
		return nil, err
	}
	return &pb.MessageResponse{Message: "OK"}, nil
}
func (h *orderHandler) GetOrders(ctx context.Context, in *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	startOrder, err := time.Parse(common.TimeYYYMMDD_Dash, in.StartOrderDate)
	if err != nil {
		return nil, err
	}
	endOrder, err := time.Parse(common.TimeYYYMMDD_Dash, in.EndOrderDate)
	if err != nil {
		return nil, err
	}
	params := orderDomain.GetOrderListParams{
		Status:         in.Status,
		StartOrderDate: startOrder,
		EndOrderDate:   endOrder,
	}
	data, err := h.OrderUC.GetUserOrderList(ctx, params)
	if err != nil {
		return nil, err
	}
	var result *pb.GetOrdersResponse
	var orders []*pb.Order
	for _, d := range data {
		orders = append(orders, &pb.Order{
			OrderId: d.OrderID,
		})
	}
	return result, nil
}
