// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: party.proto

package genproto

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

type Party struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedDate  string `protobuf:"bytes,4,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   string `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Party) Reset() {
	*x = Party{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Party) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Party) ProtoMessage() {}

func (x *Party) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Party.ProtoReflect.Descriptor instead.
func (*Party) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{0}
}

func (x *Party) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Party) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Party) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *Party) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *Party) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Party) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Party) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Party) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type PartyCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,2,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedDate  string `protobuf:"bytes,3,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PartyCreate) Reset() {
	*x = PartyCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyCreate) ProtoMessage() {}

func (x *PartyCreate) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyCreate.ProtoReflect.Descriptor instead.
func (*PartyCreate) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{1}
}

func (x *PartyCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PartyCreate) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *PartyCreate) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *PartyCreate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PartyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedDate  string `protobuf:"bytes,4,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PartyRes) Reset() {
	*x = PartyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyRes) ProtoMessage() {}

func (x *PartyRes) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyRes.ProtoReflect.Descriptor instead.
func (*PartyRes) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{2}
}

func (x *PartyRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PartyRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PartyRes) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *PartyRes) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *PartyRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PartyUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedDate  string `protobuf:"bytes,4,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PartyUpdate) Reset() {
	*x = PartyUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyUpdate) ProtoMessage() {}

func (x *PartyUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyUpdate.ProtoReflect.Descriptor instead.
func (*PartyUpdate) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{3}
}

func (x *PartyUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PartyUpdate) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *PartyUpdate) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *PartyUpdate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetAllPartysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Party []*PartyRes `protobuf:"bytes,1,rep,name=party,proto3" json:"party,omitempty"`
}

func (x *GetAllPartysResponse) Reset() {
	*x = GetAllPartysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPartysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPartysResponse) ProtoMessage() {}

func (x *GetAllPartysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPartysResponse.ProtoReflect.Descriptor instead.
func (*GetAllPartysResponse) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllPartysResponse) GetParty() []*PartyRes {
	if x != nil {
		return x.Party
	}
	return nil
}

var File_party_proto protoreflect.FileDescriptor

var file_party_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65,
	0x6e, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7c, 0x0a, 0x0b, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c,
	0x6f, 0x67, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x89, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67,
	0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x3e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x32, 0x9c, 0x02, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x42, 0x0b, 0x5a, 0x09, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_party_proto_rawDescOnce sync.Once
	file_party_proto_rawDescData = file_party_proto_rawDesc
)

func file_party_proto_rawDescGZIP() []byte {
	file_party_proto_rawDescOnce.Do(func() {
		file_party_proto_rawDescData = protoimpl.X.CompressGZIP(file_party_proto_rawDescData)
	})
	return file_party_proto_rawDescData
}

var file_party_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_party_proto_goTypes = []interface{}{
	(*Party)(nil),                // 0: protos.Party
	(*PartyCreate)(nil),          // 1: protos.PartyCreate
	(*PartyRes)(nil),             // 2: protos.PartyRes
	(*PartyUpdate)(nil),          // 3: protos.PartyUpdate
	(*GetAllPartysResponse)(nil), // 4: protos.GetAllPartysResponse
	(*GetByIdReq)(nil),           // 5: protos.GetByIdReq
	(*Filter)(nil),               // 6: protos.Filter
	(*Void)(nil),                 // 7: protos.Void
}
var file_party_proto_depIdxs = []int32{
	2, // 0: protos.GetAllPartysResponse.party:type_name -> protos.PartyRes
	1, // 1: protos.PartyService.CreateParty:input_type -> protos.PartyCreate
	5, // 2: protos.PartyService.GetParty:input_type -> protos.GetByIdReq
	6, // 3: protos.PartyService.GetAllParty:input_type -> protos.Filter
	3, // 4: protos.PartyService.UpdateParty:input_type -> protos.PartyUpdate
	5, // 5: protos.PartyService.DeleteParty:input_type -> protos.GetByIdReq
	7, // 6: protos.PartyService.CreateParty:output_type -> protos.Void
	2, // 7: protos.PartyService.GetParty:output_type -> protos.PartyRes
	4, // 8: protos.PartyService.GetAllParty:output_type -> protos.GetAllPartysResponse
	7, // 9: protos.PartyService.UpdateParty:output_type -> protos.Void
	7, // 10: protos.PartyService.DeleteParty:output_type -> protos.Void
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_party_proto_init() }
func file_party_proto_init() {
	if File_party_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_party_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Party); i {
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
		file_party_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyCreate); i {
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
		file_party_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyRes); i {
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
		file_party_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyUpdate); i {
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
		file_party_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPartysResponse); i {
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
			RawDescriptor: file_party_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_party_proto_goTypes,
		DependencyIndexes: file_party_proto_depIdxs,
		MessageInfos:      file_party_proto_msgTypes,
	}.Build()
	File_party_proto = out.File
	file_party_proto_rawDesc = nil
	file_party_proto_goTypes = nil
	file_party_proto_depIdxs = nil
}
