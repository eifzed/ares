package order

import (
	"context"

	"github.com/eifzed/ares/internal/constant"
	"github.com/eifzed/ares/internal/handler/middleware/auth"
	"github.com/eifzed/ares/internal/model/order"
	"github.com/eifzed/ares/internal/model/user"
	"github.com/eifzed/ares/lib/common/commonerr"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *orderUC) RegisterBusiness(ctx context.Context, params order.BusinessDetail) error {
	userData, isExist := auth.GetUserDataFromContext(ctx)
	if !isExist {
		return commonerr.ErrorForbidden("you are forbidden to register your business")
	}

	hasBusiness, err := uc.OrderDB.CheckUserAlreadyHasBusiness(ctx, userData.UserID)
	if err != nil {
		return err
	}
	if hasBusiness {
		return commonerr.ErrorAlreadyExist("user already has registered business")
	}

	ctx, err = uc.TX.Start(ctx)
	defer uc.TX.Finish(ctx, &err)

	err = uc.OrderDB.InsertBulkProducts(ctx, params.Products)
	if err != nil {
		return err
	}

	businessData := params.Business
	for _, products := range params.Products {
		businessData.ProductIDs = append(businessData.ProductIDs, products.ID)
	}

	businessData.OwnerID = userData.UserID

	err = uc.OrderDB.InsertBusiness(ctx, &businessData)
	if err != nil {
		return err
	}

	userData.Roles = append(userData.Roles, user.UserRole{ID: constant.RoleOwnerID, Name: constant.RoleOwner})
	err = uc.UserDB.UpdateUserRoles(ctx, userData.UserID, userData.Roles)
	if err != nil {
		return err
	}

	return nil
}

func (uc *orderUC) GetBusinessList(ctx context.Context, params order.GenericFilterParams) (*order.BusinessListData, error) {
	list, err := uc.OrderDB.GetBusinessList(ctx, params)
	if err != nil {
		return nil, err
	}
	return &order.BusinessListData{
		Total: len(list),
		List:  list,
	}, nil
}

func (uc *orderUC) GetBusinessDetail(ctx context.Context, businessID primitive.ObjectID) (*order.BusinessDetail, error) {
	data, err := uc.OrderDB.GetBusinessDetail(ctx, businessID)
	if err != nil {
		return nil, err
	}
	return data, nil
}
