package user

import (
	"context"

	"github.com/eifzed/ares/internal/model/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDBInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	InsertUser(ctx context.Context, user *user.User) error
	CheckUserExistsByEmail(ctx context.Context, email string) (bool, error)
	GetRoleByRoleName(ctx context.Context, roleName string) (*user.UserRole, error)
	UpdateUserRoles(ctx context.Context, userID primitive.ObjectID, newRoles []user.UserRole) error
}
