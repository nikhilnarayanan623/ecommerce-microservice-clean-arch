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

func (c *productServiceServer) FindAllCategories(ctx context.Context, req *pb.FindAllCategoriesRequest) (*pb.FindAllCategoriesResponse, error) {
	utils.LogMessage(utils.Cyan, "FindAllCategories Invoked")

	categories, err := c.usecase.FindAllCategories(ctx)
	if err != nil {
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Error(codes.Internal, "failed to find all categories")
	}

	outputCategories := make([]*pb.FindAllCategoriesResponse_Categories, len(categories))
	for i, category := range categories {

		outputCategories[i] = &pb.FindAllCategoriesResponse_Categories{
			Id:               category.ID,
			Name:             category.Name,
			MainCategoryId:   category.CategoryID,
			MainCategoryName: category.MainCategoryName,
		}
	}
	utils.LogMessage(utils.Green, "successfully found all categories")

	return &pb.FindAllCategoriesResponse{
		Categories: outputCategories,
	}, nil
}

func (c *productServiceServer) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {

	productID, err := c.usecase.AddProduct(ctx, request.AddProduct{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		CategoryID:  req.GetCategoryId(),
		Price:       req.GetPrice(),
		Image:       req.GetImage(),
	})

	if err != nil {
		errorCode := codes.Internal
		if err == usecase.ErrProductExist {
			errorCode = codes.AlreadyExists
		}
		return nil, status.Error(errorCode, err.Error())
	}

	return &pb.AddProductResponse{ProductId: productID}, nil
}

func (c *productServiceServer) FindAllProducts(ctx context.Context, req *pb.FindAllProductsRequest) (*pb.FindAllProductsResponse, error) {

	products, err := c.usecase.FindAllProducts(ctx, request.Pagination{
		PageNumber: req.GetPageNumber(),
		Count:      req.GetCount(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to find all products")
	}

	outputProducts := make([]*pb.FindAllProductsResponse_Product, len(products))

	for i, product := range products {

		outputProducts[i] = &pb.FindAllProductsResponse_Product{
			Id:           product.ID,
			Name:         product.Name,
			Description:  product.Description,
			Price:        product.Price,
			Image:        product.Image,
			CategoryId:   product.CategoryID,
			CategoryName: product.CategoryName,
		}
	}

	return &pb.FindAllProductsResponse{Products: outputProducts}, nil

}

func (c *productServiceServer) AddProductItem(ctx context.Context, req *pb.AddProductItemRequest) (*pb.AddProductItemResponse, error) {

	productItemID, err := c.usecase.AddProductItem(ctx, request.AddProductItem{
		ProductID:         req.GetProductId(),
		QtyInStock:        req.GetQtyInStock(),
		Price:             req.GetPrice(),
		VariationOptionID: req.GetVariationOptionId(),
	})

	if err != nil {
		var errCode codes.Code

		switch err {
		case usecase.ErrInvalidProductID, usecase.ErrInvalidVariationOptionID:
			errCode = codes.InvalidArgument
		case usecase.ErrProductItemExist:
			errCode = codes.AlreadyExists
		default:
			errCode = codes.Internal
		}

		return nil, status.Error(errCode, err.Error())
	}

	return &pb.AddProductItemResponse{ProductItemId: productItemID}, nil
}

func (c *productServiceServer) FindAllProductItems(ctx context.Context, req *pb.FindAllProductItemsRequest) (*pb.FindAllProductItemsResponse, error) {

	productItems, err := c.usecase.FindAllProductItems(ctx, req.GetProductId())
	if err != nil {
		errCode := codes.Internal
		if err == usecase.ErrInvalidProductID {
			errCode = codes.InvalidArgument
		}
		return nil, status.Error(errCode, err.Error())
	}

	outputProductItems := make([]*pb.FindAllProductItemsResponse_ProductItem, len(productItems))

	for i, productItem := range productItems {
		outputProductItems[i] = &pb.FindAllProductItemsResponse_ProductItem{
			Id:             productItem.ID,
			Name:           productItem.Name,
			Price:          productItem.Price,
			QtyInStock:     productItem.QtyInStock,
			Sku:            productItem.SKU,
			DiscountPrice:  productItem.DiscountPrice,
			VariationValue: productItem.VariationValue,
		}
	}

	return &pb.FindAllProductItemsResponse{ProductItems: outputProductItems}, nil
}

func (c *productServiceServer) FindProductItem(ctx context.Context, req *pb.FindProductItemRequest) (*pb.FindProductItemResponse, error) {

	productItem, err := c.usecase.FindProductItemByID(ctx, req.GetProductItemId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.FindProductItemResponse{
		Id:             productItem.ID,
		Name:           productItem.Name,
		Price:          productItem.Price,
		QtyInStock:     productItem.QtyInStock,
		Sku:            productItem.SKU,
		DiscountPrice:  productItem.DiscountPrice,
		VariationValue: productItem.VariationValue,
	}, nil
}

func (c *productServiceServer) MultipleStockDecrease(ctx context.Context,
	req *pb.MultipleStockDecreaseRequest) (*pb.MultipleStockDecreaseResponse, error) {

	stockDetails := make([]request.StockDecrease, len(req.GetStockDecreases()))

	for i, stockDecrease := range req.GetStockDecreases() {
		stockDetails[i].SKU = stockDecrease.GetSku()
		stockDetails[i].QtyToDecrease = stockDecrease.GetQtyToDecrease()
	}

	err := c.usecase.MultipleStockDecrease(ctx, stockDetails)
	if err != nil {
		errCode := codes.Internal
		if err == usecase.ErrInvalidStockUpdateQty {
			errCode = codes.InvalidArgument
		}

		return nil, status.Error(errCode, "failed to decrease stock qty")
	}

	return &pb.MultipleStockDecreaseResponse{}, nil
}
