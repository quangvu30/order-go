// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: order.proto

package grpc

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID          string `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	KindOfOrder      string `protobuf:"bytes,2,opt,name=KindOfOrder,proto3" json:"KindOfOrder,omitempty"`
	KindOfOffer      string `protobuf:"bytes,3,opt,name=KindOfOffer,proto3" json:"KindOfOffer,omitempty"`
	OrderType        string `protobuf:"bytes,4,opt,name=OrderType,proto3" json:"OrderType,omitempty"`
	Commodity        string `protobuf:"bytes,5,opt,name=Commodity,proto3" json:"Commodity,omitempty"`
	AccountCode      string `protobuf:"bytes,6,opt,name=AccountCode,proto3" json:"AccountCode,omitempty"`
	Volume1Lot       string `protobuf:"bytes,7,opt,name=Volume1Lot,proto3" json:"Volume1Lot,omitempty"`
	Packing          string `protobuf:"bytes,8,opt,name=Packing,proto3" json:"Packing,omitempty"`
	DepositRate      int32  `protobuf:"varint,9,opt,name=DepositRate,proto3" json:"DepositRate,omitempty"`
	TransactionType  string `protobuf:"bytes,10,opt,name=TransactionType,proto3" json:"TransactionType,omitempty"`
	DeliveryTime     uint64 `protobuf:"varint,11,opt,name=DeliveryTime,proto3" json:"DeliveryTime,omitempty"`
	DeliveryLocation string `protobuf:"bytes,12,opt,name=DeliveryLocation,proto3" json:"DeliveryLocation,omitempty"`
	Assessor         string `protobuf:"bytes,13,opt,name=Assessor,proto3" json:"Assessor,omitempty"`
	Price            uint64 `protobuf:"varint,14,opt,name=Price,proto3" json:"Price,omitempty"`
	CurrencyUnit     string `protobuf:"bytes,15,opt,name=CurrencyUnit,proto3" json:"CurrencyUnit,omitempty"`
	OrderVolume      string `protobuf:"bytes,16,opt,name=OrderVolume,proto3" json:"OrderVolume,omitempty"`
	OrderValidity    int32  `protobuf:"varint,17,opt,name=OrderValidity,proto3" json:"OrderValidity,omitempty"`
	Status           string `protobuf:"bytes,18,opt,name=Status,proto3" json:"Status,omitempty"`
	CreatedAt        int64  `protobuf:"varint,19,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	CreatedBy        string `protobuf:"bytes,20,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
	MatchedOrderID   string `protobuf:"bytes,21,opt,name=MatchedOrderID,proto3" json:"MatchedOrderID,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *Order) GetKindOfOrder() string {
	if x != nil {
		return x.KindOfOrder
	}
	return ""
}

func (x *Order) GetKindOfOffer() string {
	if x != nil {
		return x.KindOfOffer
	}
	return ""
}

func (x *Order) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}

func (x *Order) GetCommodity() string {
	if x != nil {
		return x.Commodity
	}
	return ""
}

func (x *Order) GetAccountCode() string {
	if x != nil {
		return x.AccountCode
	}
	return ""
}

func (x *Order) GetVolume1Lot() string {
	if x != nil {
		return x.Volume1Lot
	}
	return ""
}

func (x *Order) GetPacking() string {
	if x != nil {
		return x.Packing
	}
	return ""
}

func (x *Order) GetDepositRate() int32 {
	if x != nil {
		return x.DepositRate
	}
	return 0
}

func (x *Order) GetTransactionType() string {
	if x != nil {
		return x.TransactionType
	}
	return ""
}

func (x *Order) GetDeliveryTime() uint64 {
	if x != nil {
		return x.DeliveryTime
	}
	return 0
}

func (x *Order) GetDeliveryLocation() string {
	if x != nil {
		return x.DeliveryLocation
	}
	return ""
}

func (x *Order) GetAssessor() string {
	if x != nil {
		return x.Assessor
	}
	return ""
}

func (x *Order) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetCurrencyUnit() string {
	if x != nil {
		return x.CurrencyUnit
	}
	return ""
}

func (x *Order) GetOrderVolume() string {
	if x != nil {
		return x.OrderVolume
	}
	return ""
}

func (x *Order) GetOrderValidity() int32 {
	if x != nil {
		return x.OrderValidity
	}
	return 0
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Order) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Order) GetMatchedOrderID() string {
	if x != nil {
		return x.MatchedOrderID
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0xb3, 0x05, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x69, 0x6e, 0x64,
	0x4f, 0x66, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4b,
	0x69, 0x6e, 0x64, 0x4f, 0x66, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x69,
	0x6e, 0x64, 0x4f, 0x66, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x4b, 0x69, 0x6e, 0x64, 0x4f, 0x66, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x31, 0x4c, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x31, 0x4c, 0x6f, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52,
	0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x55, 0x6e,
	0x69, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x32, 0x3c, 0x0a, 0x0d, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x75, 0x61, 0x6e, 0x67, 0x76, 0x75, 0x33, 0x30,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_order_proto_goTypes = []interface{}{
	(*Order)(nil), // 0: order.Order
}
var file_order_proto_depIdxs = []int32{
	0, // 0: order.OrderTransfer.CreateOrder:input_type -> order.Order
	0, // 1: order.OrderTransfer.CreateOrder:output_type -> order.Order
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
