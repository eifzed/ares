package order

import (
	"context"

	"github.com/eifzed/ares/internal/model/order"
	"github.com/eifzed/ares/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *orderHandler) RegisterBusiness(ctx context.Context, in *pb.RegisterBusinessRequest) (*pb.MessageResponse, error) {
	// TODO: validate input
	products := []order.Products{}
	for _, prod := range in.Products {
		products = append(products, order.Products{
			ID:          primitive.NewObjectID(),
			Name:        prod.Name,
			PriceIDR:    prod.Price_IDR,
			Description: prod.Description,
			PhotoURL:    prod.Photo_URL,
		})
	}
	businessData := order.BusinessData{
		Name:        in.Name,
		Address:     in.Address,
		PhoneNumber: in.PhoneNumber,
		Description: in.Description,
		PhotoURL:    in.Photo_URL,
	}
	businessDetail := order.BusinessDetail{
		Business: businessData,
		Products: products,
	}
	err := h.OrderUC.RegisterBusiness(ctx, businessDetail)
	if err != nil {
		return &pb.MessageResponse{Message: "failed to register business"}, err
	}
	return &pb.MessageResponse{Message: "OK"}, nil
}

func (h *orderHandler) GetBusinessList(ctx context.Context, in *pb.GenericFilterParams) (*pb.GetBusinessListResponse, error) {
	params := order.GenericFilterParams{
		Keyword: in.Keyword,
		Limit:   in.Limit,
		Page:    in.Page,
	}
	result, err := h.OrderUC.GetBusinessList(ctx, params)
	if err != nil {
		return nil, err
	}
	list := []*pb.BusinessList{}
	for _, r := range result.List {
		list = append(list, &pb.BusinessList{
			Id:        r.ID.Hex(),
			Name:      r.Name,
			Address:   r.Address,
			Photo_URL: r.PhotoURL,
		})
	}
	return &pb.GetBusinessListResponse{
		Total:        uint32(result.Total),
		BusinessList: list,
	}, nil
}

func (h *orderHandler) GetBusinessDetail(ctx context.Context, in *pb.GetBusinessDetailRequest) (*pb.GetBusinessDetailResponse, error) {
	businessID, err := primitive.ObjectIDFromHex(in.BusinessId)
	if err != nil {
		return nil, err
	}
	data, err := h.OrderUC.GetBusinessDetail(ctx, businessID)
	if err != nil {
		return nil, err
	}
	products := []*pb.ProductDetail{}

	for _, prod := range data.Products {
		products = append(products, &pb.ProductDetail{
			Id:          prod.ID.Hex(),
			Name:        prod.Name,
			Price_IDR:   prod.PriceIDR,
			Description: prod.Description,
			Photo_URL:   prod.PhotoURL,
		})
	}

	return &pb.GetBusinessDetailResponse{
		Id:          data.Business.ID.Hex(),
		Name:        data.Business.Name,
		Address:     data.Business.Address,
		PhoneNumber: data.Business.PhoneNumber,
		Description: data.Business.Description,
		Photo_URL:   data.Business.PhotoURL,
		Products:    products,
	}, nil
}
