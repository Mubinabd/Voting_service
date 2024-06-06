// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protos/election.proto

package public

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

type Election struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Date      string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt string `protobuf:"bytes,6,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *Election) Reset() {
	*x = Election{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_election_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Election) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Election) ProtoMessage() {}

func (x *Election) ProtoReflect() protoreflect.Message {
	mi := &file_protos_election_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Election.ProtoReflect.Descriptor instead.
func (*Election) Descriptor() ([]byte, []int) {
	return file_protos_election_proto_rawDescGZIP(), []int{0}
}

func (x *Election) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Election) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Election) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Election) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Election) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Election) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type CreateElectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Date string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *CreateElectionReq) Reset() {
	*x = CreateElectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_election_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateElectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateElectionReq) ProtoMessage() {}

func (x *CreateElectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_election_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateElectionReq.ProtoReflect.Descriptor instead.
func (*CreateElectionReq) Descriptor() ([]byte, []int) {
	return file_protos_election_proto_rawDescGZIP(), []int{1}
}

func (x *CreateElectionReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateElectionReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateElectionReq) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ElectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Date string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ElectionReq) Reset() {
	*x = ElectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_election_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionReq) ProtoMessage() {}

func (x *ElectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_election_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionReq.ProtoReflect.Descriptor instead.
func (*ElectionReq) Descriptor() ([]byte, []int) {
	return file_protos_election_proto_rawDescGZIP(), []int{2}
}

func (x *ElectionReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ElectionReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ElectionReq) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ElectionUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Date string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ElectionUpdate) Reset() {
	*x = ElectionUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_election_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionUpdate) ProtoMessage() {}

func (x *ElectionUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_election_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionUpdate.ProtoReflect.Descriptor instead.
func (*ElectionUpdate) Descriptor() ([]byte, []int) {
	return file_protos_election_proto_rawDescGZIP(), []int{3}
}

func (x *ElectionUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ElectionUpdate) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ElectionsGetAllResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elections []*ElectionReq `protobuf:"bytes,1,rep,name=elections,proto3" json:"elections,omitempty"`
}

func (x *ElectionsGetAllResp) Reset() {
	*x = ElectionsGetAllResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_election_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionsGetAllResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionsGetAllResp) ProtoMessage() {}

func (x *ElectionsGetAllResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_election_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionsGetAllResp.ProtoReflect.Descriptor instead.
func (*ElectionsGetAllResp) Descriptor() ([]byte, []int) {
	return file_protos_election_proto_rawDescGZIP(), []int{4}
}

func (x *ElectionsGetAllResp) GetElections() []*ElectionReq {
	if x != nil {
		return x.Elections
	}
	return nil
}

var File_protos_election_proto protoreflect.FileDescriptor

var file_protos_election_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a,
	0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x45,
	0x0a, 0x0b, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x38, 0x0a, 0x0e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22,
	0x48, 0x0a, 0x13, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x09, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x52, 0x09,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xef, 0x01, 0x0a, 0x0f, 0x45, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x07, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x29, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_election_proto_rawDescOnce sync.Once
	file_protos_election_proto_rawDescData = file_protos_election_proto_rawDesc
)

func file_protos_election_proto_rawDescGZIP() []byte {
	file_protos_election_proto_rawDescOnce.Do(func() {
		file_protos_election_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_election_proto_rawDescData)
	})
	return file_protos_election_proto_rawDescData
}

var file_protos_election_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_election_proto_goTypes = []interface{}{
	(*Election)(nil),            // 0: protos.Election
	(*CreateElectionReq)(nil),   // 1: protos.CreateElectionReq
	(*ElectionReq)(nil),         // 2: protos.ElectionReq
	(*ElectionUpdate)(nil),      // 3: protos.ElectionUpdate
	(*ElectionsGetAllResp)(nil), // 4: protos.ElectionsGetAllResp
	(*Filter)(nil),              // 5: Filter
	(*GetByIdReq)(nil),          // 6: GetByIdReq
	(*Void)(nil),                // 7: Void
}
var file_protos_election_proto_depIdxs = []int32{
	2, // 0: protos.ElectionsGetAllResp.elections:type_name -> protos.ElectionReq
	1, // 1: protos.ElectionService.Create:input_type -> protos.CreateElectionReq
	2, // 2: protos.ElectionService.Get:input_type -> protos.ElectionReq
	5, // 3: protos.ElectionService.GetAll:input_type -> Filter
	3, // 4: protos.ElectionService.Update:input_type -> protos.ElectionUpdate
	6, // 5: protos.ElectionService.Delete:input_type -> GetByIdReq
	7, // 6: protos.ElectionService.Create:output_type -> Void
	2, // 7: protos.ElectionService.Get:output_type -> protos.ElectionReq
	4, // 8: protos.ElectionService.GetAll:output_type -> protos.ElectionsGetAllResp
	7, // 9: protos.ElectionService.Update:output_type -> Void
	7, // 10: protos.ElectionService.Delete:output_type -> Void
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_election_proto_init() }
func file_protos_election_proto_init() {
	if File_protos_election_proto != nil {
		return
	}
	file_protos_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_election_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Election); i {
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
		file_protos_election_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateElectionReq); i {
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
		file_protos_election_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionReq); i {
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
		file_protos_election_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionUpdate); i {
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
		file_protos_election_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionsGetAllResp); i {
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
			RawDescriptor: file_protos_election_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_election_proto_goTypes,
		DependencyIndexes: file_protos_election_proto_depIdxs,
		MessageInfos:      file_protos_election_proto_msgTypes,
	}.Build()
	File_protos_election_proto = out.File
	file_protos_election_proto_rawDesc = nil
	file_protos_election_proto_goTypes = nil
	file_protos_election_proto_depIdxs = nil
}