// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: auction.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartAuctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionDuration int32 `protobuf:"varint,1,opt,name=auctionDuration,proto3" json:"auctionDuration,omitempty"`
}

func (x *StartAuctionRequest) Reset() {
	*x = StartAuctionRequest{}
	mi := &file_auction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartAuctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartAuctionRequest) ProtoMessage() {}

func (x *StartAuctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartAuctionRequest.ProtoReflect.Descriptor instead.
func (*StartAuctionRequest) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{0}
}

func (x *StartAuctionRequest) GetAuctionDuration() int32 {
	if x != nil {
		return x.AuctionDuration
	}
	return 0
}

type StartAuctionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartAuctionResponse) Reset() {
	*x = StartAuctionResponse{}
	mi := &file_auction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartAuctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartAuctionResponse) ProtoMessage() {}

func (x *StartAuctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartAuctionResponse.ProtoReflect.Descriptor instead.
func (*StartAuctionResponse) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{1}
}

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BidderName string `protobuf:"bytes,2,opt,name=bidderName,proto3" json:"bidderName,omitempty"`
	Amount     int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	mi := &file_auction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{2}
}

func (x *BidRequest) GetBidderName() string {
	if x != nil {
		return x.BidderName
	}
	return ""
}

func (x *BidRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp int32  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *BidResponse) Reset() {
	*x = BidResponse{}
	mi := &file_auction_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidResponse) ProtoMessage() {}

func (x *BidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidResponse.ProtoReflect.Descriptor instead.
func (*BidResponse) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{3}
}

func (x *BidResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BidResponse) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResultRequest) Reset() {
	*x = ResultRequest{}
	mi := &file_auction_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRequest) ProtoMessage() {}

func (x *ResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRequest.ProtoReflect.Descriptor instead.
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{4}
}

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	HighestBid int32  `protobuf:"varint,2,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
	BidderName string `protobuf:"bytes,3,opt,name=bidderName,proto3" json:"bidderName,omitempty"`
	Timestamp  int32  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	mi := &file_auction_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{5}
}

func (x *ResultResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ResultResponse) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

func (x *ResultResponse) GetBidderName() string {
	if x != nil {
		return x.BidderName
	}
	return ""
}

func (x *ResultResponse) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auction   *Auction `protobuf:"bytes,1,opt,name=auction,proto3" json:"auction,omitempty"`
	Timestamp int32    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	mi := &file_auction_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{6}
}

func (x *HealthCheckRequest) GetAuction() *Auction {
	if x != nil {
		return x.Auction
	}
	return nil
}

func (x *HealthCheckRequest) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	mi := &file_auction_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{7}
}

type Auction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighestBid int32                  `protobuf:"varint,1,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
	BidderName string                 `protobuf:"bytes,2,opt,name=bidderName,proto3" json:"bidderName,omitempty"`
	EndTime    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *Auction) Reset() {
	*x = Auction{}
	mi := &file_auction_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auction) ProtoMessage() {}

func (x *Auction) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auction.ProtoReflect.Descriptor instead.
func (*Auction) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{8}
}

func (x *Auction) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

func (x *Auction) GetBidderName() string {
	if x != nil {
		return x.BidderName
	}
	return ""
}

func (x *Auction) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_auction_proto protoreflect.FileDescriptor

var file_auction_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x0f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x44, 0x0a, 0x0a, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x0b, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x0e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73,
	0x74, 0x42, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68,
	0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x69, 0x64, 0x64,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x5b, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x15, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7f, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74,
	0x42, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x84, 0x02, 0x0a, 0x0e, 0x41, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x12, 0x10, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x13, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auction_proto_rawDescOnce sync.Once
	file_auction_proto_rawDescData = file_auction_proto_rawDesc
)

func file_auction_proto_rawDescGZIP() []byte {
	file_auction_proto_rawDescOnce.Do(func() {
		file_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_auction_proto_rawDescData)
	})
	return file_auction_proto_rawDescData
}

var file_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_auction_proto_goTypes = []any{
	(*StartAuctionRequest)(nil),   // 0: grpc.StartAuctionRequest
	(*StartAuctionResponse)(nil),  // 1: grpc.StartAuctionResponse
	(*BidRequest)(nil),            // 2: grpc.BidRequest
	(*BidResponse)(nil),           // 3: grpc.BidResponse
	(*ResultRequest)(nil),         // 4: grpc.ResultRequest
	(*ResultResponse)(nil),        // 5: grpc.ResultResponse
	(*HealthCheckRequest)(nil),    // 6: grpc.HealthCheckRequest
	(*HealthCheckResponse)(nil),   // 7: grpc.HealthCheckResponse
	(*Auction)(nil),               // 8: grpc.Auction
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_auction_proto_depIdxs = []int32{
	8, // 0: grpc.HealthCheckRequest.auction:type_name -> grpc.Auction
	9, // 1: grpc.Auction.endTime:type_name -> google.protobuf.Timestamp
	0, // 2: grpc.AuctionService.startAuction:input_type -> grpc.StartAuctionRequest
	2, // 3: grpc.AuctionService.bid:input_type -> grpc.BidRequest
	4, // 4: grpc.AuctionService.result:input_type -> grpc.ResultRequest
	6, // 5: grpc.AuctionService.healthCheck:input_type -> grpc.HealthCheckRequest
	1, // 6: grpc.AuctionService.startAuction:output_type -> grpc.StartAuctionResponse
	3, // 7: grpc.AuctionService.bid:output_type -> grpc.BidResponse
	5, // 8: grpc.AuctionService.result:output_type -> grpc.ResultResponse
	7, // 9: grpc.AuctionService.healthCheck:output_type -> grpc.HealthCheckResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auction_proto_init() }
func file_auction_proto_init() {
	if File_auction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auction_proto_goTypes,
		DependencyIndexes: file_auction_proto_depIdxs,
		MessageInfos:      file_auction_proto_msgTypes,
	}.Build()
	File_auction_proto = out.File
	file_auction_proto_rawDesc = nil
	file_auction_proto_goTypes = nil
	file_auction_proto_depIdxs = nil
}
