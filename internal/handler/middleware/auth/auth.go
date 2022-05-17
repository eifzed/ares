package auth

import (
	"context"

	"github.com/eifzed/ares/internal/config"
	"github.com/eifzed/ares/internal/constant"
	userRepo "github.com/eifzed/ares/internal/model/repo/user"
	"github.com/eifzed/ares/internal/model/user"
	"github.com/eifzed/ares/lib/common/commonerr"
	"github.com/eifzed/ares/lib/utility/jwt"

	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

type authModule struct {
	JWTCertificate *jwt.JWTCertificate
	RouteRoles     map[string]jwt.RouteRoles
	Config         *config.Config
	UserDB         userRepo.UserDBInterface
}

type userContext struct{}

var (
	userContextKey = userContext{}
)

type Info struct {
	UserID int64
	Type   string
	Data   map[string]interface{}
}

type Options struct {
	JWTCertificate *jwt.JWTCertificate
	RouteRoles     map[string]jwt.RouteRoles
	Config         *config.Config
	UserDB         userRepo.UserDBInterface
}

func NewAuthModule(option *Options) *authModule {
	return &authModule{
		JWTCertificate: option.JWTCertificate,
		RouteRoles:     option.RouteRoles,
		Config:         option.Config,
		UserDB:         option.UserDB,
	}
}

func (m *authModule) AuthHandler(ctx context.Context) (context.Context, error) {
	route, _ := grpc.Method(ctx)
	roles := m.RouteRoles[route].Roles
	if isPublicRoute(roles) {
		return ctx, nil
	}
	jwtToken, err := grpcauth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, commonerr.ErrorUnauthorized(err.Error())
	}
	userPayload, err := jwt.DecodeToken(jwtToken, m.JWTCertificate.PublicKey)
	if err != nil {
		return nil, commonerr.ErrorBadRequest(err.Error())
	}
	userData, err := m.UserDB.GetUserByEmail(ctx, userPayload.Email)
	if err != nil {
		return nil, commonerr.ErrorUnauthorized(err.Error())
	}
	if !isUserAuthorized(userData.Roles, m.RouteRoles[route].Roles) {
		return nil, commonerr.ErrorForbidden("you don't have the right to access this resource")
	}
	ctx = setUserDataToContext(ctx, user.User{
		UserID:    userPayload.UserID,
		FirstName: userPayload.FirstName,
		LastName:  userPayload.LastName,
		Email:     userPayload.Email,
		Roles:     userPayload.Roles,
	})
	return ctx, nil

}

func setUserDataToContext(ctx context.Context, value interface{}) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, userContextKey, value)
}

func GetUserDataFromContext(ctx context.Context) (user.User, bool) {
	user, exist := ctx.Value(userContextKey).(user.User)
	return user, exist
}

func isUserAuthorized(userRoles []user.UserRole, authorizedRoles []user.UserRole) bool {
	if len(userRoles) == 0 || len(authorizedRoles) == 0 {
		return false
	}
	for _, userRole := range userRoles {
		for _, authRole := range authorizedRoles {
			if userRole.ID == authRole.ID {
				return true
			}
		}
	}
	return false
}
func isPublicRoute(roles []user.UserRole) bool {
	for _, role := range roles {
		if role.ID == constant.RolePublicID {
			return true
		}
	}
	return false
}
