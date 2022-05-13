package user

import (
	"context"
	"time"

	"github.com/eifzed/ares/internal/model/user"
	"github.com/eifzed/ares/lib/common"
	"github.com/eifzed/ares/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *userHandler) Logout(ctx context.Context, in *emptypb.Empty) (*pb.MessageResponse, error) {
	return &pb.MessageResponse{Message: "OK"}, nil
}
func (h *userHandler) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := h.UserUC.Login(ctx, in.Email, in.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: token.JWT, ValidUntil: token.ValidUntil}, nil
}

func (h *userHandler) Register(ctx context.Context, in *pb.UserRegistrationRequest) (*pb.LoginResponse, error) {
	birthDate, err := time.Parse(common.TimeYYYMMDD_Dash, in.BirthDate)
	if err != nil {
		return nil, err
	}
	token, err := h.UserUC.Register(ctx, &user.User{
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
		BirthDate:   birthDate,
		Password:    in.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: token.JWT, ValidUntil: token.ValidUntil}, nil
}
