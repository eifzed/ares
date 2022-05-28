package user

import (
	"github.com/eifzed/ares/internal/config"
	"github.com/eifzed/ares/internal/model/repo/transaction"
	"github.com/eifzed/ares/internal/model/repo/user"
)

type userUC struct {
	UserDB user.UserDBInterface
	Config *config.Config
	TX     transaction.TransactionInterface
}

type Options struct {
	UserDB user.UserDBInterface
	Config *config.Config
	TX     transaction.TransactionInterface
}

func GetNewUserUC(option *Options) *userUC {
	if option == nil || option.UserDB == nil {
		return nil
	}
	return &userUC{
		UserDB: option.UserDB,
		Config: option.Config,
		TX:     option.TX,
	}
}
