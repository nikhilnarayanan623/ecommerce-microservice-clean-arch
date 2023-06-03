package service

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/usecase"
	usecaseInterface "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/request"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productServiceServer struct {
	usecase usecaseInterface.ProductUseCase
	pb.ProductServiceServer
}

func NewProductServiceServer(usecase usecaseInterface.ProductUseCase) pb.ProductServiceServer {
	return &productServiceServer{
		usecase: usecase,
	}
}

func (c *productServiceServer) AddCategory(ctx context.Context, req *pb.AddCategoryRequest) (*pb.AddCategoryResponse, error) {
	utils.LogMessage(utils.Cyan, "AddCategory Invoked")
	categoryID, err := c.usecase.AddCategory(ctx, request.AddCategory{
		MainCategoryID: req.GetMainCategoryId(),
		CategoryName:   req.GetCategoryName(),
	})

	if err != nil {
		var errCode codes.Code
		switch err {
		case usecase.ErrCategoryExist:
			errCode = codes.AlreadyExists
		case usecase.ErrInvalidCategoryID:
			errCode = codes.InvalidArgument
		default:
			errCode = codes.Internal
		}
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Error(errCode, err.Error())
	}
	utils.LogMessage(utils.Green, "category successfully added")
	return &pb.AddCategoryResponse{CategoryId: categoryID}, nil
}

// Add new variation
func (c *productServiceServer) AddVariation(ctx context.Context, req *pb.AddVariationRequest) (*pb.AddVariationResponse, error) {
	utils.LogMessage(utils.Cyan, "AddVariation Invoked")
	variationID, err := c.usecase.AddVariation(ctx, request.AddVariation{
		CategoryID:    req.GetCategoryId(),
		VariationName: req.GetVariationName(),
	})
	if err != nil {
		var errCode codes.Code
		switch err {
		case usecase.ErrInvalidCategoryID:
			errCode = codes.InvalidArgument
		case usecase.ErrVariationExist:
			errCode = codes.AlreadyExists
		default:
			errCode = codes.Internal
		}
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Error(errCode, err.Error())
	}
	utils.LogMessage(utils.Green, "variation successfully added")
	return &pb.AddVariationResponse{VariationId: variationID}, nil
}
func (c *productServiceServer) AddVariationOption(ctx context.Context, req *pb.AddVariationOptionRequest) (*pb.AddVariationOptionResponse, error) {
	utils.LogMessage(utils.Cyan, "AddVariationOption Invoked")
	variationOptionID, err := c.usecase.AddVariationOption(ctx, request.AddVariationOption{
		VariationID:    req.GetVariationId(),
		VariationValue: req.GetVariationValue(),
	})
	if err != nil {
		var errCode codes.Code
		switch err {
		case usecase.ErrInvalidVariationID:
			errCode = codes.InvalidArgument
		case usecase.ErrVariationOptionExist:
			errCode = codes.AlreadyExists
		default:
			errCode = codes.Internal
		}
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Error(errCode, err.Error())
	}
	utils.LogMessage(utils.Green, "variation option successfully added")
	return &pb.AddVariationOptionResponse{VariationOptionId: variationOptionID}, nil
}

// func (c *productServiceServer) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {

// }
