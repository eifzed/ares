package user

import (
	"context"

	"github.com/eifzed/ares/internal/model/user"
	"gopkg.in/mgo.v2/bson"
)

func (db *userDB) GetRoleByRoleName(ctx context.Context, roleName string) (*user.UserRole, error) {
	role := &user.UserRole{}
	err := db.DB.Collection("user_roles").FindOne(ctx, bson.M{"name": roleName}).Decode(role)
	if err != nil {
		return nil, err
	}
	return role, nil
}
