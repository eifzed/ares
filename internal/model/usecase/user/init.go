package user

import (
	"context"

	"github.com/eifzed/ares/internal/model/user"
)

type UserUCInterface interface {
	Register(ctx context.Context, user *user.User) (*user.Token, error)
	Login(ctx context.Context, email string, passwrod string) (*user.Token, error)
	Logout(ctx context.Context, user *user.User) error
}
