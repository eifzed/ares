package order

import (
	"github.com/eifzed/ares/internal/config"
	"github.com/eifzed/ares/internal/model/repo/order"
	"github.com/eifzed/ares/internal/model/repo/user"
)

type orderUC struct {
	OrderDB order.OrderDBInterface
	UserDB  user.UserDBInterface
	Config  *config.Config
}

type Options struct {
	OrderDB order.OrderDBInterface
	UserDB  user.UserDBInterface
	Config  *config.Config
}

func GetNewOrderUC(option *Options) *orderUC {
	if option == nil || option.OrderDB == nil {
		return nil
	}
	return &orderUC{
		OrderDB: option.OrderDB,
		Config:  option.Config,
		UserDB:  option.UserDB,
	}
}
