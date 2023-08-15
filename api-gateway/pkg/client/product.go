package client

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type productClient struct {
	client pb.ProductServiceClient
}

func NewProductClient(cfg *config.Config) (interfaces.ProductClient, error) {
	gcc, err := grpc.Dial(cfg.ProductServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewProductServiceClient(gcc)

	return &productClient{
		client: client,
	}, nil
}

func (c *productClient) AddCategory(ctx context.Context, category request.AddCategory) (uint64, error) {
	res, err := c.client.AddCategory(ctx, &pb.AddCategoryRequest{
		MainCategoryId: category.MainCategoryID,
		CategoryName:   category.CategoryName,
	})
	if err != nil {
		return 0, err
	}
	return res.GetCategoryId(), nil
}
func (c *productClient) AddVariation(ctx context.Context, variation request.AddVariation) (uint64, error) {

	res, err := c.client.AddVariation(ctx, &pb.AddVariationRequest{
		CategoryId:    variation.CategoryID,
		VariationName: variation.VariationName,
	})
	if err != nil {
		return 0, err
	}

	return res.GetVariationId(), nil
}
func (c *productClient) AddVariationOption(ctx context.Context, variationOption request.AddVariationOption) (uint64, error) {

	res, err := c.client.AddVariationOption(ctx, &pb.AddVariationOptionRequest{
		VariationId:    variationOption.VariationID,
		VariationValue: variationOption.VariationValue,
	})
	if err != nil {
		return 0, err
	}
	return res.GetVariationOptionId(), nil
}

// Find all categories
func (c *productClient) FindAllCategories(ctx context.Context) ([]response.Category, error) {

	res, err := c.client.FindAllCategories(ctx, &pb.FindAllCategoriesRequest{})
	if err != nil {
		return nil, err
	}

	categories := make([]response.Category, len(res.GetCategories()))

	for i, category := range res.GetCategories() {
		categories[i].ID = category.GetId()
		categories[i].Name = category.GetName()
		categories[i].CategoryID = category.GetMainCategoryId()
		categories[i].MainCategoryName = category.GetMainCategoryName()
	}
	return categories, nil
}

// Add Product
func (c *productClient) AddProduct(ctx context.Context, product request.AddProduct) (uint64, error) {
	res, err := c.client.AddProduct(ctx, &pb.AddProductRequest{
		Name:        product.Name,
		Description: product.Description,
		CategoryId:  product.CategoryID,
		Price:       product.Price,
		Image:       product.Image,
	})
	if err != nil {
		return 0, err
	}

	return res.GetProductId(), nil
}

func (c *productClient) FindAllProducts(ctx context.Context, pagination request.Pagination) ([]response.Product, error) {

	res, err := c.client.FindAllProducts(ctx, &pb.FindAllProductsRequest{
		PageNumber: pagination.PageNumber,
		Count:      pagination.Count,
	})
	if err != nil {
		return nil, err
	}

	products := make([]response.Product, len((res.Products)))

	for i, resProduct := range res.GetProducts() {
		products[i] = response.Product{
			ID:           resProduct.GetId(),
			Name:         resProduct.GetName(),
			Description:  resProduct.GetDescription(),
			Price:        resProduct.GetPrice(),
			Image:        resProduct.GetImage(),
			CategoryID:   resProduct.GetCategoryId(),
			CategoryName: resProduct.GetCategoryName(),
		}
	}

	return products, nil
}

func (c *productClient) AddProductItem(ctx context.Context, productItem request.AddProductItem) (uint64, error) {

	res, err := c.client.AddProductItem(ctx, &pb.AddProductItemRequest{
		ProductId:         productItem.ProductID,
		QtyInStock:        productItem.QtyInStock,
		Price:             productItem.Price,
		VariationOptionId: productItem.VariationOptionID,
	})

	if err != nil {
		return 0, err
	}

	return res.GetProductItemId(), nil
}

func (c *productClient) FindAllProductItems(ctx context.Context, productID uint64) ([]response.ProductItem, error) {

	res, err := c.client.FindAllProductItems(ctx, &pb.FindAllProductItemsRequest{ProductId: productID})

	if err != nil {
		return nil, err
	}

	productItems := make([]response.ProductItem, len(res.GetProductItems()))

	for i, productItem := range res.GetProductItems() {
		productItems[i] = response.ProductItem{
			ID:             productItem.GetId(),
			Name:           productItem.GetName(),
			QtyInStock:     productItem.GetQtyInStock(),
			Price:          productItem.GetPrice(),
			SKU:            productItem.GetSku(),
			DiscountPrice:  productItem.GetDiscountPrice(),
			VariationValue: productItem.GetVariationValue(),
		}
	}

	return productItems, nil
}
