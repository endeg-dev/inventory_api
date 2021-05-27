// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: stock/stock.proto

package stock

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

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductsCode string `protobuf:"bytes,2,opt,name=products_code,json=productsCode,proto3" json:"products_code,omitempty"`
	ProductsName string `protobuf:"bytes,3,opt,name=products_name,json=productsName,proto3" json:"products_name,omitempty"`
	Quantity     string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	SupplierId   string `protobuf:"bytes,5,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	CreatedOn    string `protobuf:"bytes,9,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	UpdatedOn    string `protobuf:"bytes,10,opt,name=updated_on,json=updatedOn,proto3" json:"updated_on,omitempty"`
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
	mi := &file_stock_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_stock_stock_proto_rawDescGZIP(), []int{0}
}

func (x *Stock) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Stock) GetProductsCode() string {
	if x != nil {
		return x.ProductsCode
	}
	return ""
}

func (x *Stock) GetProductsName() string {
	if x != nil {
		return x.ProductsName
	}
	return ""
}

func (x *Stock) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *Stock) GetSupplierId() string {
	if x != nil {
		return x.SupplierId
	}
	return ""
}

func (x *Stock) GetCreatedOn() string {
	if x != nil {
		return x.CreatedOn
	}
	return ""
}

func (x *Stock) GetUpdatedOn() string {
	if x != nil {
		return x.UpdatedOn
	}
	return ""
}

type CreateStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stock *Stock `protobuf:"bytes,1,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *CreateStockRequest) Reset() {
	*x = CreateStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStockRequest) ProtoMessage() {}

func (x *CreateStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStockRequest.ProtoReflect.Descriptor instead.
func (*CreateStockRequest) Descriptor() ([]byte, []int) {
	return file_stock_stock_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStockRequest) GetStock() *Stock {
	if x != nil {
		return x.Stock
	}
	return nil
}

type CreateStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateStockResponse) Reset() {
	*x = CreateStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_stock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStockResponse) ProtoMessage() {}

func (x *CreateStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_stock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStockResponse.ProtoReflect.Descriptor instead.
func (*CreateStockResponse) Descriptor() ([]byte, []int) {
	return file_stock_stock_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStockResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StockByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StockByIDRequest) Reset() {
	*x = StockByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_stock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockByIDRequest) ProtoMessage() {}

func (x *StockByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_stock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockByIDRequest.ProtoReflect.Descriptor instead.
func (*StockByIDRequest) Descriptor() ([]byte, []int) {
	return file_stock_stock_proto_rawDescGZIP(), []int{3}
}

func (x *StockByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StockByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stock *Stock `protobuf:"bytes,1,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *StockByIDResponse) Reset() {
	*x = StockByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_stock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockByIDResponse) ProtoMessage() {}

func (x *StockByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_stock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockByIDResponse.ProtoReflect.Descriptor instead.
func (*StockByIDResponse) Descriptor() ([]byte, []int) {
	return file_stock_stock_proto_rawDescGZIP(), []int{4}
}

func (x *StockByIDResponse) GetStock() *Stock {
	if x != nil {
		return x.Stock
	}
	return nil
}

type StockListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StockListRequest) Reset() {
	*x = StockListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_stock_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockListRequest) ProtoMessage() {}

func (x *StockListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_stock_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockListRequest.ProtoReflect.Descriptor instead.
func (*StockListRequest) Descriptor() ([]byte, []int) {
	return file_stock_stock_proto_rawDescGZIP(), []int{5}
}

type StockListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stock []*Stock `protobuf:"bytes,1,rep,name=stock,proto3" json:"stock,omitempty"`
}

func (x *StockListResponse) Reset() {
	*x = StockListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_stock_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockListResponse) ProtoMessage() {}

func (x *StockListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_stock_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockListResponse.ProtoReflect.Descriptor instead.
func (*StockListResponse) Descriptor() ([]byte, []int) {
	return file_stock_stock_proto_rawDescGZIP(), []int{6}
}

func (x *StockListResponse) GetStock() []*Stock {
	if x != nil {
		return x.Stock
	}
	return nil
}

type UpdateStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stock *Stock `protobuf:"bytes,1,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *UpdateStockRequest) Reset() {
	*x = UpdateStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_stock_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockRequest) ProtoMessage() {}

func (x *UpdateStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_stock_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockRequest.ProtoReflect.Descriptor instead.
func (*UpdateStockRequest) Descriptor() ([]byte, []int) {
	return file_stock_stock_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateStockRequest) GetStock() *Stock {
	if x != nil {
		return x.Stock
	}
	return nil
}

type UpdateStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateStockResponse) Reset() {
	*x = UpdateStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_stock_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockResponse) ProtoMessage() {}

func (x *UpdateStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_stock_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockResponse.ProtoReflect.Descriptor instead.
func (*UpdateStockResponse) Descriptor() ([]byte, []int) {
	return file_stock_stock_proto_rawDescGZIP(), []int{8}
}

type DeleteStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteStockRequest) Reset() {
	*x = DeleteStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_stock_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStockRequest) ProtoMessage() {}

func (x *DeleteStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_stock_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStockRequest.ProtoReflect.Descriptor instead.
func (*DeleteStockRequest) Descriptor() ([]byte, []int) {
	return file_stock_stock_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteStockRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteStockResponse) Reset() {
	*x = DeleteStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_stock_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStockResponse) ProtoMessage() {}

func (x *DeleteStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_stock_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStockResponse.ProtoReflect.Descriptor instead.
func (*DeleteStockResponse) Descriptor() ([]byte, []int) {
	return file_stock_stock_proto_rawDescGZIP(), []int{10}
}

var File_stock_stock_proto protoreflect.FileDescriptor

var file_stock_stock_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0xdc, 0x01, 0x0a, 0x05, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0x38, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37,
	0x0a, 0x11, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x11, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x22, 0x38, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x15,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xe0, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x19,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stock_stock_proto_rawDescOnce sync.Once
	file_stock_stock_proto_rawDescData = file_stock_stock_proto_rawDesc
)

func file_stock_stock_proto_rawDescGZIP() []byte {
	file_stock_stock_proto_rawDescOnce.Do(func() {
		file_stock_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_stock_stock_proto_rawDescData)
	})
	return file_stock_stock_proto_rawDescData
}

var file_stock_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_stock_stock_proto_goTypes = []interface{}{
	(*Stock)(nil),               // 0: stock.Stock
	(*CreateStockRequest)(nil),  // 1: stock.CreateStockRequest
	(*CreateStockResponse)(nil), // 2: stock.CreateStockResponse
	(*StockByIDRequest)(nil),    // 3: stock.StockByIDRequest
	(*StockByIDResponse)(nil),   // 4: stock.StockByIDResponse
	(*StockListRequest)(nil),    // 5: stock.StockListRequest
	(*StockListResponse)(nil),   // 6: stock.StockListResponse
	(*UpdateStockRequest)(nil),  // 7: stock.UpdateStockRequest
	(*UpdateStockResponse)(nil), // 8: stock.UpdateStockResponse
	(*DeleteStockRequest)(nil),  // 9: stock.DeleteStockRequest
	(*DeleteStockResponse)(nil), // 10: stock.DeleteStockResponse
}
var file_stock_stock_proto_depIdxs = []int32{
	0,  // 0: stock.CreateStockRequest.stock:type_name -> stock.Stock
	0,  // 1: stock.StockByIDResponse.stock:type_name -> stock.Stock
	0,  // 2: stock.StockListResponse.stock:type_name -> stock.Stock
	0,  // 3: stock.UpdateStockRequest.stock:type_name -> stock.Stock
	1,  // 4: stock.StockService.CreateStock:input_type -> stock.CreateStockRequest
	3,  // 5: stock.StockService.StockByID:input_type -> stock.StockByIDRequest
	5,  // 6: stock.StockService.StockList:input_type -> stock.StockListRequest
	7,  // 7: stock.StockService.UpdateStock:input_type -> stock.UpdateStockRequest
	9,  // 8: stock.StockService.DeleteStock:input_type -> stock.DeleteStockRequest
	2,  // 9: stock.StockService.CreateStock:output_type -> stock.CreateStockResponse
	4,  // 10: stock.StockService.StockByID:output_type -> stock.StockByIDResponse
	6,  // 11: stock.StockService.StockList:output_type -> stock.StockListResponse
	8,  // 12: stock.StockService.UpdateStock:output_type -> stock.UpdateStockResponse
	10, // 13: stock.StockService.DeleteStock:output_type -> stock.DeleteStockResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_stock_stock_proto_init() }
func file_stock_stock_proto_init() {
	if File_stock_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stock_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stock); i {
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
		file_stock_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStockRequest); i {
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
		file_stock_stock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStockResponse); i {
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
		file_stock_stock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockByIDRequest); i {
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
		file_stock_stock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockByIDResponse); i {
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
		file_stock_stock_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockListRequest); i {
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
		file_stock_stock_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockListResponse); i {
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
		file_stock_stock_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockRequest); i {
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
		file_stock_stock_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockResponse); i {
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
		file_stock_stock_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStockRequest); i {
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
		file_stock_stock_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStockResponse); i {
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
			RawDescriptor: file_stock_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stock_stock_proto_goTypes,
		DependencyIndexes: file_stock_stock_proto_depIdxs,
		MessageInfos:      file_stock_stock_proto_msgTypes,
	}.Build()
	File_stock_stock_proto = out.File
	file_stock_stock_proto_rawDesc = nil
	file_stock_stock_proto_goTypes = nil
	file_stock_stock_proto_depIdxs = nil
}
