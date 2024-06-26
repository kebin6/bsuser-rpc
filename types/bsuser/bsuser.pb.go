// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.3
// source: bsuser.proto

package bsuser

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

// base message
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{0}
}

type IDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *IDReq) Reset() {
	*x = IDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDReq) ProtoMessage() {}

func (x *IDReq) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDReq.ProtoReflect.Descriptor instead.
func (*IDReq) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{1}
}

func (x *IDReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IDsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *IDsReq) Reset() {
	*x = IDsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDsReq) ProtoMessage() {}

func (x *IDsReq) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDsReq.ProtoReflect.Descriptor instead.
func (*IDsReq) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{2}
}

func (x *IDsReq) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UUIDsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids"`
}

func (x *UUIDsReq) Reset() {
	*x = UUIDsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDsReq) ProtoMessage() {}

func (x *UUIDsReq) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDsReq.ProtoReflect.Descriptor instead.
func (*UUIDsReq) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{3}
}

func (x *UUIDsReq) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UUIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *UUIDReq) Reset() {
	*x = UUIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDReq) ProtoMessage() {}

func (x *UUIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDReq.ProtoReflect.Descriptor instead.
func (*UUIDReq) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{4}
}

func (x *UUIDReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MobileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile"`
}

func (x *MobileReq) Reset() {
	*x = MobileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MobileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MobileReq) ProtoMessage() {}

func (x *MobileReq) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MobileReq.ProtoReflect.Descriptor instead.
func (*MobileReq) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{5}
}

func (x *MobileReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{6}
}

func (x *BaseResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PageInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	PageSize uint64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
}

func (x *PageInfoReq) Reset() {
	*x = PageInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfoReq) ProtoMessage() {}

func (x *PageInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfoReq.ProtoReflect.Descriptor instead.
func (*PageInfoReq) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{7}
}

func (x *PageInfoReq) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageInfoReq) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type BaseIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
}

func (x *BaseIDResp) Reset() {
	*x = BaseIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseIDResp) ProtoMessage() {}

func (x *BaseIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseIDResp.ProtoReflect.Descriptor instead.
func (*BaseIDResp) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{8}
}

func (x *BaseIDResp) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseIDResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type BaseUUIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
}

func (x *BaseUUIDResp) Reset() {
	*x = BaseUUIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseUUIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseUUIDResp) ProtoMessage() {}

func (x *BaseUUIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseUUIDResp.ProtoReflect.Descriptor instead.
func (*BaseUUIDResp) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{9}
}

func (x *BaseUUIDResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BaseUUIDResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type BsUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *uint64  `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id"`
	CreatedAt   *int64   `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at"`
	UpdatedAt   *int64   `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at"`
	Status      *uint32  `protobuf:"varint,4,opt,name=status,proto3,oneof" json:"status"`
	Name        *string  `protobuf:"bytes,5,opt,name=name,proto3,oneof" json:"name"`
	Pwd         *string  `protobuf:"bytes,6,opt,name=pwd,proto3,oneof" json:"pwd"`
	Mobile      *string  `protobuf:"bytes,7,opt,name=mobile,proto3,oneof" json:"mobile"`
	InviteCode  *string  `protobuf:"bytes,8,opt,name=invite_code,json=inviteCode,proto3,oneof" json:"invite_code"`
	TotalAmount *float64 `protobuf:"fixed64,9,opt,name=total_amount,json=totalAmount,proto3,oneof" json:"total_amount"`
	ValidAmount *float64 `protobuf:"fixed64,10,opt,name=valid_amount,json=validAmount,proto3,oneof" json:"valid_amount"`
	InvitedBy   *uint64  `protobuf:"varint,11,opt,name=invited_by,json=invitedBy,proto3,oneof" json:"invited_by"`
}

func (x *BsUserInfo) Reset() {
	*x = BsUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BsUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BsUserInfo) ProtoMessage() {}

func (x *BsUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BsUserInfo.ProtoReflect.Descriptor instead.
func (*BsUserInfo) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{10}
}

func (x *BsUserInfo) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *BsUserInfo) GetCreatedAt() int64 {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return 0
}

func (x *BsUserInfo) GetUpdatedAt() int64 {
	if x != nil && x.UpdatedAt != nil {
		return *x.UpdatedAt
	}
	return 0
}

func (x *BsUserInfo) GetStatus() uint32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *BsUserInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *BsUserInfo) GetPwd() string {
	if x != nil && x.Pwd != nil {
		return *x.Pwd
	}
	return ""
}

func (x *BsUserInfo) GetMobile() string {
	if x != nil && x.Mobile != nil {
		return *x.Mobile
	}
	return ""
}

func (x *BsUserInfo) GetInviteCode() string {
	if x != nil && x.InviteCode != nil {
		return *x.InviteCode
	}
	return ""
}

func (x *BsUserInfo) GetTotalAmount() float64 {
	if x != nil && x.TotalAmount != nil {
		return *x.TotalAmount
	}
	return 0
}

func (x *BsUserInfo) GetValidAmount() float64 {
	if x != nil && x.ValidAmount != nil {
		return *x.ValidAmount
	}
	return 0
}

func (x *BsUserInfo) GetInvitedBy() uint64 {
	if x != nil && x.InvitedBy != nil {
		return *x.InvitedBy
	}
	return 0
}

type BsUserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page           uint64   `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	PageSize       uint64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Name           *string  `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name"`
	Mobile         *string  `protobuf:"bytes,4,opt,name=mobile,proto3,oneof" json:"mobile"`
	Status         *uint32  `protobuf:"varint,5,opt,name=status,proto3,oneof" json:"status"`
	InvitedBy      *uint64  `protobuf:"varint,6,opt,name=invited_by,json=invitedBy,proto3,oneof" json:"invited_by"`
	InviteCode     *string  `protobuf:"bytes,7,opt,name=invite_code,json=inviteCode,proto3,oneof" json:"invite_code"`
	MinTotalAmount *float64 `protobuf:"fixed64,8,opt,name=min_total_amount,json=minTotalAmount,proto3,oneof" json:"min_total_amount"`
	MaxTotalAmount *float64 `protobuf:"fixed64,9,opt,name=max_total_amount,json=maxTotalAmount,proto3,oneof" json:"max_total_amount"`
	MinValidAmount *float64 `protobuf:"fixed64,10,opt,name=min_valid_amount,json=minValidAmount,proto3,oneof" json:"min_valid_amount"`
	MaxValidAmount *float64 `protobuf:"fixed64,11,opt,name=max_valid_amount,json=maxValidAmount,proto3,oneof" json:"max_valid_amount"`
	StartTime      *int64   `protobuf:"varint,12,opt,name=start_time,json=startTime,proto3,oneof" json:"start_time"`
	EndTime        *int64   `protobuf:"varint,13,opt,name=end_time,json=endTime,proto3,oneof" json:"end_time"`
}

func (x *BsUserListReq) Reset() {
	*x = BsUserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BsUserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BsUserListReq) ProtoMessage() {}

func (x *BsUserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BsUserListReq.ProtoReflect.Descriptor instead.
func (*BsUserListReq) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{11}
}

func (x *BsUserListReq) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *BsUserListReq) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *BsUserListReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *BsUserListReq) GetMobile() string {
	if x != nil && x.Mobile != nil {
		return *x.Mobile
	}
	return ""
}

func (x *BsUserListReq) GetStatus() uint32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *BsUserListReq) GetInvitedBy() uint64 {
	if x != nil && x.InvitedBy != nil {
		return *x.InvitedBy
	}
	return 0
}

func (x *BsUserListReq) GetInviteCode() string {
	if x != nil && x.InviteCode != nil {
		return *x.InviteCode
	}
	return ""
}

func (x *BsUserListReq) GetMinTotalAmount() float64 {
	if x != nil && x.MinTotalAmount != nil {
		return *x.MinTotalAmount
	}
	return 0
}

func (x *BsUserListReq) GetMaxTotalAmount() float64 {
	if x != nil && x.MaxTotalAmount != nil {
		return *x.MaxTotalAmount
	}
	return 0
}

func (x *BsUserListReq) GetMinValidAmount() float64 {
	if x != nil && x.MinValidAmount != nil {
		return *x.MinValidAmount
	}
	return 0
}

func (x *BsUserListReq) GetMaxValidAmount() float64 {
	if x != nil && x.MaxValidAmount != nil {
		return *x.MaxValidAmount
	}
	return 0
}

func (x *BsUserListReq) GetStartTime() int64 {
	if x != nil && x.StartTime != nil {
		return *x.StartTime
	}
	return 0
}

func (x *BsUserListReq) GetEndTime() int64 {
	if x != nil && x.EndTime != nil {
		return *x.EndTime
	}
	return 0
}

type BsUserListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint64        `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	Data  []*BsUserInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
}

func (x *BsUserListResp) Reset() {
	*x = BsUserListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bsuser_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BsUserListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BsUserListResp) ProtoMessage() {}

func (x *BsUserListResp) ProtoReflect() protoreflect.Message {
	mi := &file_bsuser_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BsUserListResp.ProtoReflect.Descriptor instead.
func (*BsUserListResp) Descriptor() ([]byte, []int) {
	return file_bsuser_proto_rawDescGZIP(), []int{12}
}

func (x *BsUserListResp) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *BsUserListResp) GetData() []*BsUserInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_bsuser_proto protoreflect.FileDescriptor

var file_bsuser_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x73, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x73, 0x75, 0x73, 0x65, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x17, 0x0a, 0x05, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x06, 0x49, 0x44, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x1c, 0x0a, 0x08, 0x55, 0x55, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x22, 0x19, 0x0a, 0x07, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a,
	0x09, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x22, 0x1c, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x3e, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x2e, 0x0a, 0x0a, 0x42, 0x61, 0x73, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x30, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0xfa, 0x03, 0x0a, 0x0a, 0x42, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x05, 0x52, 0x03, 0x70, 0x77, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52,
	0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x26,
	0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x01, 0x48, 0x08, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x48, 0x09, 0x52, 0x0b,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22,
	0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x04, 0x48, 0x0a, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x42, 0x79, 0x88,
	0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f,
	0x70, 0x77, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x22,
	0x8b, 0x05, 0x0a, 0x0d, 0x42, 0x73, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x48, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04,
	0x52, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x2d, 0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x48, 0x05, 0x52, 0x0e, 0x6d, 0x69, 0x6e,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d,
	0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x48, 0x06, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a,
	0x10, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x48, 0x07, 0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10,
	0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x48, 0x08, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x1e, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x0a, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x13, 0x0a,
	0x11, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6d, 0x69, 0x6e, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x13, 0x0a, 0x11,
	0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x4e, 0x0a,
	0x0e, 0x42, 0x73, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x73, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x73, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xed, 0x02,
	0x0a, 0x06, 0x42, 0x73, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x0d, 0x2e, 0x62, 0x73, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x62, 0x73, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x62, 0x73, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x73, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x12, 0x2e, 0x62, 0x73, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x06, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x62, 0x73, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x12, 0x2e, 0x62, 0x73, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a,
	0x07, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e, 0x62, 0x73, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x62, 0x73, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x42, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x0b, 0x67,
	0x65, 0x74, 0x42, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x62, 0x73, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x62, 0x73, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x30, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x12, 0x2e, 0x62, 0x73,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x12, 0x2e, 0x62, 0x73, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15,
	0x2e, 0x62, 0x73, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x73, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x62, 0x73, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x62, 0x73, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_bsuser_proto_rawDescOnce sync.Once
	file_bsuser_proto_rawDescData = file_bsuser_proto_rawDesc
)

func file_bsuser_proto_rawDescGZIP() []byte {
	file_bsuser_proto_rawDescOnce.Do(func() {
		file_bsuser_proto_rawDescData = protoimpl.X.CompressGZIP(file_bsuser_proto_rawDescData)
	})
	return file_bsuser_proto_rawDescData
}

var file_bsuser_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_bsuser_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: bsuser.Empty
	(*IDReq)(nil),          // 1: bsuser.IDReq
	(*IDsReq)(nil),         // 2: bsuser.IDsReq
	(*UUIDsReq)(nil),       // 3: bsuser.UUIDsReq
	(*UUIDReq)(nil),        // 4: bsuser.UUIDReq
	(*MobileReq)(nil),      // 5: bsuser.MobileReq
	(*BaseResp)(nil),       // 6: bsuser.BaseResp
	(*PageInfoReq)(nil),    // 7: bsuser.PageInfoReq
	(*BaseIDResp)(nil),     // 8: bsuser.BaseIDResp
	(*BaseUUIDResp)(nil),   // 9: bsuser.BaseUUIDResp
	(*BsUserInfo)(nil),     // 10: bsuser.BsUserInfo
	(*BsUserListReq)(nil),  // 11: bsuser.BsUserListReq
	(*BsUserListResp)(nil), // 12: bsuser.BsUserListResp
}
var file_bsuser_proto_depIdxs = []int32{
	10, // 0: bsuser.BsUserListResp.data:type_name -> bsuser.BsUserInfo
	0,  // 1: bsuser.Bsuser.initDatabase:input_type -> bsuser.Empty
	10, // 2: bsuser.Bsuser.create:input_type -> bsuser.BsUserInfo
	10, // 3: bsuser.Bsuser.update:input_type -> bsuser.BsUserInfo
	1,  // 4: bsuser.Bsuser.getById:input_type -> bsuser.IDReq
	5,  // 5: bsuser.Bsuser.getByMobile:input_type -> bsuser.MobileReq
	10, // 6: bsuser.Bsuser.getOne:input_type -> bsuser.BsUserInfo
	11, // 7: bsuser.Bsuser.getList:input_type -> bsuser.BsUserListReq
	6,  // 8: bsuser.Bsuser.initDatabase:output_type -> bsuser.BaseResp
	8,  // 9: bsuser.Bsuser.create:output_type -> bsuser.BaseIDResp
	8,  // 10: bsuser.Bsuser.update:output_type -> bsuser.BaseIDResp
	10, // 11: bsuser.Bsuser.getById:output_type -> bsuser.BsUserInfo
	10, // 12: bsuser.Bsuser.getByMobile:output_type -> bsuser.BsUserInfo
	10, // 13: bsuser.Bsuser.getOne:output_type -> bsuser.BsUserInfo
	12, // 14: bsuser.Bsuser.getList:output_type -> bsuser.BsUserListResp
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_bsuser_proto_init() }
func file_bsuser_proto_init() {
	if File_bsuser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bsuser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_bsuser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDReq); i {
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
		file_bsuser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDsReq); i {
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
		file_bsuser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUIDsReq); i {
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
		file_bsuser_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUIDReq); i {
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
		file_bsuser_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MobileReq); i {
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
		file_bsuser_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_bsuser_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfoReq); i {
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
		file_bsuser_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseIDResp); i {
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
		file_bsuser_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseUUIDResp); i {
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
		file_bsuser_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BsUserInfo); i {
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
		file_bsuser_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BsUserListReq); i {
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
		file_bsuser_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BsUserListResp); i {
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
	file_bsuser_proto_msgTypes[10].OneofWrappers = []interface{}{}
	file_bsuser_proto_msgTypes[11].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bsuser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bsuser_proto_goTypes,
		DependencyIndexes: file_bsuser_proto_depIdxs,
		MessageInfos:      file_bsuser_proto_msgTypes,
	}.Build()
	File_bsuser_proto = out.File
	file_bsuser_proto_rawDesc = nil
	file_bsuser_proto_goTypes = nil
	file_bsuser_proto_depIdxs = nil
}
