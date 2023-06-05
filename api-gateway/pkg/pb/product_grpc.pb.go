// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/proto/product.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProductService_AddCategory_FullMethodName         = "/pb.ProductService/AddCategory"
	ProductService_FindAllCategories_FullMethodName   = "/pb.ProductService/FindAllCategories"
	ProductService_AddVariation_FullMethodName        = "/pb.ProductService/AddVariation"
	ProductService_AddVariationOption_FullMethodName  = "/pb.ProductService/AddVariationOption"
	ProductService_AddProduct_FullMethodName          = "/pb.ProductService/AddProduct"
	ProductService_FindAllProducts_FullMethodName     = "/pb.ProductService/FindAllProducts"
	ProductService_AddProductItem_FullMethodName      = "/pb.ProductService/AddProductItem"
	ProductService_FindAllProductItems_FullMethodName = "/pb.ProductService/FindAllProductItems"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*AddCategoryResponse, error)
	FindAllCategories(ctx context.Context, in *FindAllCategoriesRequest, opts ...grpc.CallOption) (*FindAllCategoriesResponse, error)
	AddVariation(ctx context.Context, in *AddVariationRequest, opts ...grpc.CallOption) (*AddVariationResponse, error)
	AddVariationOption(ctx context.Context, in *AddVariationOptionRequest, opts ...grpc.CallOption) (*AddVariationOptionResponse, error)
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error)
	FindAllProducts(ctx context.Context, in *FindAllProductsRequest, opts ...grpc.CallOption) (*FindAllProductsResponse, error)
	AddProductItem(ctx context.Context, in *AddProductItemRequest, opts ...grpc.CallOption) (*AddProductItemResponse, error)
	FindAllProductItems(ctx context.Context, in *FindAllProductItemsRequest, opts ...grpc.CallOption) (*FindAllProductItemsResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*AddCategoryResponse, error) {
	out := new(AddCategoryResponse)
	err := c.cc.Invoke(ctx, ProductService_AddCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindAllCategories(ctx context.Context, in *FindAllCategoriesRequest, opts ...grpc.CallOption) (*FindAllCategoriesResponse, error) {
	out := new(FindAllCategoriesResponse)
	err := c.cc.Invoke(ctx, ProductService_FindAllCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddVariation(ctx context.Context, in *AddVariationRequest, opts ...grpc.CallOption) (*AddVariationResponse, error) {
	out := new(AddVariationResponse)
	err := c.cc.Invoke(ctx, ProductService_AddVariation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddVariationOption(ctx context.Context, in *AddVariationOptionRequest, opts ...grpc.CallOption) (*AddVariationOptionResponse, error) {
	out := new(AddVariationOptionResponse)
	err := c.cc.Invoke(ctx, ProductService_AddVariationOption_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error) {
	out := new(AddProductResponse)
	err := c.cc.Invoke(ctx, ProductService_AddProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindAllProducts(ctx context.Context, in *FindAllProductsRequest, opts ...grpc.CallOption) (*FindAllProductsResponse, error) {
	out := new(FindAllProductsResponse)
	err := c.cc.Invoke(ctx, ProductService_FindAllProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddProductItem(ctx context.Context, in *AddProductItemRequest, opts ...grpc.CallOption) (*AddProductItemResponse, error) {
	out := new(AddProductItemResponse)
	err := c.cc.Invoke(ctx, ProductService_AddProductItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindAllProductItems(ctx context.Context, in *FindAllProductItemsRequest, opts ...grpc.CallOption) (*FindAllProductItemsResponse, error) {
	out := new(FindAllProductItemsResponse)
	err := c.cc.Invoke(ctx, ProductService_FindAllProductItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	AddCategory(context.Context, *AddCategoryRequest) (*AddCategoryResponse, error)
	FindAllCategories(context.Context, *FindAllCategoriesRequest) (*FindAllCategoriesResponse, error)
	AddVariation(context.Context, *AddVariationRequest) (*AddVariationResponse, error)
	AddVariationOption(context.Context, *AddVariationOptionRequest) (*AddVariationOptionResponse, error)
	AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error)
	FindAllProducts(context.Context, *FindAllProductsRequest) (*FindAllProductsResponse, error)
	AddProductItem(context.Context, *AddProductItemRequest) (*AddProductItemResponse, error)
	FindAllProductItems(context.Context, *FindAllProductItemsRequest) (*FindAllProductItemsResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) AddCategory(context.Context, *AddCategoryRequest) (*AddCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedProductServiceServer) FindAllCategories(context.Context, *FindAllCategoriesRequest) (*FindAllCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllCategories not implemented")
}
func (UnimplementedProductServiceServer) AddVariation(context.Context, *AddVariationRequest) (*AddVariationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVariation not implemented")
}
func (UnimplementedProductServiceServer) AddVariationOption(context.Context, *AddVariationOptionRequest) (*AddVariationOptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVariationOption not implemented")
}
func (UnimplementedProductServiceServer) AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedProductServiceServer) FindAllProducts(context.Context, *FindAllProductsRequest) (*FindAllProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllProducts not implemented")
}
func (UnimplementedProductServiceServer) AddProductItem(context.Context, *AddProductItemRequest) (*AddProductItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProductItem not implemented")
}
func (UnimplementedProductServiceServer) FindAllProductItems(context.Context, *FindAllProductItemsRequest) (*FindAllProductItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllProductItems not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddCategory(ctx, req.(*AddCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindAllCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindAllCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_FindAllCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindAllCategories(ctx, req.(*FindAllCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddVariation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVariationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddVariation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddVariation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddVariation(ctx, req.(*AddVariationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddVariationOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVariationOptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddVariationOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddVariationOption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddVariationOption(ctx, req.(*AddVariationOptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_FindAllProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindAllProducts(ctx, req.(*FindAllProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddProductItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProductItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddProductItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProductItem(ctx, req.(*AddProductItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindAllProductItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllProductItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindAllProductItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_FindAllProductItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindAllProductItems(ctx, req.(*FindAllProductItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCategory",
			Handler:    _ProductService_AddCategory_Handler,
		},
		{
			MethodName: "FindAllCategories",
			Handler:    _ProductService_FindAllCategories_Handler,
		},
		{
			MethodName: "AddVariation",
			Handler:    _ProductService_AddVariation_Handler,
		},
		{
			MethodName: "AddVariationOption",
			Handler:    _ProductService_AddVariationOption_Handler,
		},
		{
			MethodName: "AddProduct",
			Handler:    _ProductService_AddProduct_Handler,
		},
		{
			MethodName: "FindAllProducts",
			Handler:    _ProductService_FindAllProducts_Handler,
		},
		{
			MethodName: "AddProductItem",
			Handler:    _ProductService_AddProductItem_Handler,
		},
		{
			MethodName: "FindAllProductItems",
			Handler:    _ProductService_FindAllProductItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/product.proto",
}
