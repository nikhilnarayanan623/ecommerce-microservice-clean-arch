package client

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/request"
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
