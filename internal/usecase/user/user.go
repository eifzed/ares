package user

import (
	"context"
	"errors"
	"time"

	"github.com/eifzed/ares/internal/constant"
	"github.com/eifzed/ares/internal/model/user"
	"github.com/eifzed/ares/lib/common/commonerr"
	"github.com/eifzed/ares/lib/utility/hash"
	"github.com/eifzed/ares/lib/utility/jwt"
)

func (uc *userUC) Register(ctx context.Context, userData *user.User) (*user.Token, error) {
	isExist, err := uc.UserDB.CheckUserExistsByEmail(ctx, userData.Email)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, commonerr.ErrorAlreadyExist("user already exist")
	}
	// TODO: start transaction
	userData.PasswordHashed, err = hash.HashPassword(userData.Password)
	if err != nil {
		return nil, err
	}

	userData.Roles = []user.UserRole{{ID: constant.RoleCustomerID, Name: constant.RoleCustomer}}
	err = uc.UserDB.InsertUser(ctx, userData)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	token, err := uc.generateUserToken(userData)
	if err != nil {
		return nil, err
	}
	return &user.Token{JWT: token, ValidUntil: now.Add(time.Minute * constant.MinutesInOneDay).String()}, nil
}
func (uc *userUC) Login(ctx context.Context, email string, passwrod string) (*user.Token, error) {
	userData, err := uc.UserDB.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !hash.IsCorrectPasswordHash(passwrod, userData.PasswordHashed) {
		return nil, commonerr.ErrorForbidden("invalid password")
	}
	now := time.Now()
	token, err := uc.generateUserToken(userData)
	if err != nil {
		return nil, err
	}

	return &user.Token{JWT: token, ValidUntil: now.Add(time.Minute * constant.MinutesInOneDay).String()}, nil
}

func (uc *userUC) Logout(ctx context.Context, user *user.User) error {
	return nil
}

func (uc *userUC) generateUserToken(userData *user.User) (string, error) {
	if userData == nil {
		return "", errors.New("nil user data")
	}
	userPayload := jwt.JWTPayload{
		UserID:         userData.UserID,
		FirstName:      userData.FirstName,
		LastName:       userData.LastName,
		Email:          userData.Email,
		PasswordHashed: userData.PasswordHashed,
		Roles:          userData.Roles,
	}
	token, err := jwt.GenerateToken(userPayload, uc.Config.Secrets.Data.JWTCertificate.PrivateKey, constant.MinutesInOneDay)
	if err != nil {
		return "", commonerr.SetError(err)
	}
	return token, nil
}
