package handler

import (
	"github.com/eifzed/ares/pb"
)

type GRPCHandler struct {
	OrderHandler pb.OrderServiceServer
	UserHandler  pb.UserServiceServer
}
