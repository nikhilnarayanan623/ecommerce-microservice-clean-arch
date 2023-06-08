// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pkg/proto/cart.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddToCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId uint64 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartRequest.ProtoReflect.Descriptor instead.
func (*AddToCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{0}
}

func (x *AddToCartRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddToCartRequest) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type AddToCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddToCartResponse) Reset() {
	*x = AddToCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartResponse) ProtoMessage() {}

func (x *AddToCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartResponse.ProtoReflect.Descriptor instead.
func (*AddToCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{1}
}

type FindCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FindCartRequest) Reset() {
	*x = FindCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindCartRequest) ProtoMessage() {}

func (x *FindCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindCartRequest.ProtoReflect.Descriptor instead.
func (*FindCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{2}
}

func (x *FindCartRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FindCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPrice float64                      `protobuf:"fixed64,1,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	CartItems  []*FindCartResponse_CartItem `protobuf:"bytes,2,rep,name=cart_items,json=cartItems,proto3" json:"cart_items,omitempty"`
}

func (x *FindCartResponse) Reset() {
	*x = FindCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindCartResponse) ProtoMessage() {}

func (x *FindCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindCartResponse.ProtoReflect.Descriptor instead.
func (*FindCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{3}
}

func (x *FindCartResponse) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *FindCartResponse) GetCartItems() []*FindCartResponse_CartItem {
	if x != nil {
		return x.CartItems
	}
	return nil
}

type FindCartResponse_CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductItemId  uint64  `protobuf:"varint,1,opt,name=product_item_id,json=productItemId,proto3" json:"product_item_id,omitempty"`
	ProductName    string  `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Sku            string  `protobuf:"bytes,3,opt,name=sku,proto3" json:"sku,omitempty"`
	VariationValue string  `protobuf:"bytes,4,opt,name=variation_value,json=variationValue,proto3" json:"variation_value,omitempty"`
	Price          float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Qty            uint64  `protobuf:"varint,6,opt,name=qty,proto3" json:"qty,omitempty"`
	SubTotal       float64 `protobuf:"fixed64,7,opt,name=sub_total,json=subTotal,proto3" json:"sub_total,omitempty"`
	QtyInStock     uint64  `protobuf:"varint,8,opt,name=qty_in_stock,json=qtyInStock,proto3" json:"qty_in_stock,omitempty"`
	DiscountPrice  uint64  `protobuf:"varint,9,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
}

func (x *FindCartResponse_CartItem) Reset() {
	*x = FindCartResponse_CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindCartResponse_CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindCartResponse_CartItem) ProtoMessage() {}

func (x *FindCartResponse_CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindCartResponse_CartItem.ProtoReflect.Descriptor instead.
func (*FindCartResponse_CartItem) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{3, 0}
}

func (x *FindCartResponse_CartItem) GetProductItemId() uint64 {
	if x != nil {
		return x.ProductItemId
	}
	return 0
}

func (x *FindCartResponse_CartItem) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *FindCartResponse_CartItem) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *FindCartResponse_CartItem) GetVariationValue() string {
	if x != nil {
		return x.VariationValue
	}
	return ""
}

func (x *FindCartResponse_CartItem) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *FindCartResponse_CartItem) GetQty() uint64 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *FindCartResponse_CartItem) GetSubTotal() float64 {
	if x != nil {
		return x.SubTotal
	}
	return 0
}

func (x *FindCartResponse_CartItem) GetQtyInStock() uint64 {
	if x != nil {
		return x.QtyInStock
	}
	return 0
}

func (x *FindCartResponse_CartItem) GetDiscountPrice() uint64 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

var File_pkg_proto_cart_proto protoreflect.FileDescriptor

var file_pkg_proto_cart_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x4a, 0x0a, 0x10, 0x41, 0x64,
	0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x0a, 0x0f, 0x46,
	0x69, 0x6e, 0x64, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x92, 0x03, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a,
	0x0a, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x09, 0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x9e, 0x02, 0x0a, 0x08,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x27, 0x0a, 0x0f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0c, 0x71, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x71, 0x74, 0x79, 0x49, 0x6e,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x32, 0x82, 0x01, 0x0a,
	0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09,
	0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_cart_proto_rawDescOnce sync.Once
	file_pkg_proto_cart_proto_rawDescData = file_pkg_proto_cart_proto_rawDesc
)

func file_pkg_proto_cart_proto_rawDescGZIP() []byte {
	file_pkg_proto_cart_proto_rawDescOnce.Do(func() {
		file_pkg_proto_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_cart_proto_rawDescData)
	})
	return file_pkg_proto_cart_proto_rawDescData
}

var file_pkg_proto_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_proto_cart_proto_goTypes = []interface{}{
	(*AddToCartRequest)(nil),          // 0: pb.AddToCartRequest
	(*AddToCartResponse)(nil),         // 1: pb.AddToCartResponse
	(*FindCartRequest)(nil),           // 2: pb.FindCartRequest
	(*FindCartResponse)(nil),          // 3: pb.FindCartResponse
	(*FindCartResponse_CartItem)(nil), // 4: pb.FindCartResponse.CartItem
}
var file_pkg_proto_cart_proto_depIdxs = []int32{
	4, // 0: pb.FindCartResponse.cart_items:type_name -> pb.FindCartResponse.CartItem
	0, // 1: pb.CartService.AddToCart:input_type -> pb.AddToCartRequest
	2, // 2: pb.CartService.FindCart:input_type -> pb.FindCartRequest
	1, // 3: pb.CartService.AddToCart:output_type -> pb.AddToCartResponse
	3, // 4: pb.CartService.FindCart:output_type -> pb.FindCartResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_cart_proto_init() }
func file_pkg_proto_cart_proto_init() {
	if File_pkg_proto_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindCartResponse_CartItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_cart_proto_goTypes,
		DependencyIndexes: file_pkg_proto_cart_proto_depIdxs,
		MessageInfos:      file_pkg_proto_cart_proto_msgTypes,
	}.Build()
	File_pkg_proto_cart_proto = out.File
	file_pkg_proto_cart_proto_rawDesc = nil
	file_pkg_proto_cart_proto_goTypes = nil
	file_pkg_proto_cart_proto_depIdxs = nil
}
