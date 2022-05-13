package user

import (
	"github.com/eifzed/ares/internal/config"
	"github.com/eifzed/ares/internal/model/usecase/user"
	"github.com/eifzed/ares/pb"
)

type userHandler struct {
	pb.UserServiceServer
	UserUC user.UserUCInterface
	Config *config.Config
}

type Option struct {
	UserUC user.UserUCInterface
	Config *config.Config
}

func GetNewUserHandler(userHandlerOption *Option) *userHandler {
	if userHandlerOption == nil {
		return nil
	}
	return &userHandler{
		UserUC: userHandlerOption.UserUC,
		Config: userHandlerOption.Config,
	}
}
