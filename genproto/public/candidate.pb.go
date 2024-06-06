// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protos/candidate.proto

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

type Candidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ElectionId string `protobuf:"bytes,2,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	PublicId   string `protobuf:"bytes,3,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt  string `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt  string `protobuf:"bytes,6,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *Candidate) Reset() {
	*x = Candidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_candidate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candidate) ProtoMessage() {}

func (x *Candidate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_candidate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candidate.ProtoReflect.Descriptor instead.
func (*Candidate) Descriptor() ([]byte, []int) {
	return file_protos_candidate_proto_rawDescGZIP(), []int{0}
}

func (x *Candidate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Candidate) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *Candidate) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *Candidate) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Candidate) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Candidate) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type CreateCandidateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ElectionId string `protobuf:"bytes,2,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	PublicId   string `protobuf:"bytes,3,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
}

func (x *CreateCandidateReq) Reset() {
	*x = CreateCandidateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_candidate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCandidateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCandidateReq) ProtoMessage() {}

func (x *CreateCandidateReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_candidate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCandidateReq.ProtoReflect.Descriptor instead.
func (*CreateCandidateReq) Descriptor() ([]byte, []int) {
	return file_protos_candidate_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCandidateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateCandidateReq) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *CreateCandidateReq) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

type GetCandidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ElectionId string `protobuf:"bytes,2,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	PublicId   string `protobuf:"bytes,3,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
}

func (x *GetCandidate) Reset() {
	*x = GetCandidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_candidate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCandidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCandidate) ProtoMessage() {}

func (x *GetCandidate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_candidate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCandidate.ProtoReflect.Descriptor instead.
func (*GetCandidate) Descriptor() ([]byte, []int) {
	return file_protos_candidate_proto_rawDescGZIP(), []int{2}
}

func (x *GetCandidate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetCandidate) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *GetCandidate) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

type CandidateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Election *Election `protobuf:"bytes,2,opt,name=election,proto3" json:"election,omitempty"`
	Public   *Public   `protobuf:"bytes,3,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *CandidateRes) Reset() {
	*x = CandidateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_candidate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateRes) ProtoMessage() {}

func (x *CandidateRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_candidate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateRes.ProtoReflect.Descriptor instead.
func (*CandidateRes) Descriptor() ([]byte, []int) {
	return file_protos_candidate_proto_rawDescGZIP(), []int{3}
}

func (x *CandidateRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CandidateRes) GetElection() *Election {
	if x != nil {
		return x.Election
	}
	return nil
}

func (x *CandidateRes) GetPublic() *Public {
	if x != nil {
		return x.Public
	}
	return nil
}

type CandidatiesGetAllResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candidaties []*CandidateRes `protobuf:"bytes,1,rep,name=candidaties,proto3" json:"candidaties,omitempty"`
}

func (x *CandidatiesGetAllResp) Reset() {
	*x = CandidatiesGetAllResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_candidate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidatiesGetAllResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidatiesGetAllResp) ProtoMessage() {}

func (x *CandidatiesGetAllResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_candidate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidatiesGetAllResp.ProtoReflect.Descriptor instead.
func (*CandidatiesGetAllResp) Descriptor() ([]byte, []int) {
	return file_protos_candidate_proto_rawDescGZIP(), []int{4}
}

func (x *CandidatiesGetAllResp) GetCandidaties() []*CandidateRes {
	if x != nil {
		return x.Candidaties
	}
	return nil
}

var File_protos_candidate_proto protoreflect.FileDescriptor

var file_protos_candidate_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb3, 0x01, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x62, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x22, 0x5c, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x4f, 0x0a,
	0x15, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x69, 0x65, 0x73, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x69, 0x65, 0x73, 0x32, 0xca,
	0x01, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x1a,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x07, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x69, 0x65, 0x73, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_candidate_proto_rawDescOnce sync.Once
	file_protos_candidate_proto_rawDescData = file_protos_candidate_proto_rawDesc
)

func file_protos_candidate_proto_rawDescGZIP() []byte {
	file_protos_candidate_proto_rawDescOnce.Do(func() {
		file_protos_candidate_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_candidate_proto_rawDescData)
	})
	return file_protos_candidate_proto_rawDescData
}

var file_protos_candidate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_candidate_proto_goTypes = []interface{}{
	(*Candidate)(nil),             // 0: protos.Candidate
	(*CreateCandidateReq)(nil),    // 1: protos.CreateCandidateReq
	(*GetCandidate)(nil),          // 2: protos.GetCandidate
	(*CandidateRes)(nil),          // 3: protos.CandidateRes
	(*CandidatiesGetAllResp)(nil), // 4: protos.CandidatiesGetAllResp
	(*Election)(nil),              // 5: protos.Election
	(*Public)(nil),                // 6: protos.Public
	(*Filter)(nil),                // 7: Filter
	(*GetByIdReq)(nil),            // 8: GetByIdReq
	(*Void)(nil),                  // 9: Void
}
var file_protos_candidate_proto_depIdxs = []int32{
	5, // 0: protos.CandidateRes.election:type_name -> protos.Election
	6, // 1: protos.CandidateRes.public:type_name -> protos.Public
	3, // 2: protos.CandidatiesGetAllResp.candidaties:type_name -> protos.CandidateRes
	1, // 3: protos.CandidateService.Create:input_type -> protos.CreateCandidateReq
	2, // 4: protos.CandidateService.Get:input_type -> protos.GetCandidate
	7, // 5: protos.CandidateService.GetAll:input_type -> Filter
	8, // 6: protos.CandidateService.Delete:input_type -> GetByIdReq
	9, // 7: protos.CandidateService.Create:output_type -> Void
	3, // 8: protos.CandidateService.Get:output_type -> protos.CandidateRes
	4, // 9: protos.CandidateService.GetAll:output_type -> protos.CandidatiesGetAllResp
	9, // 10: protos.CandidateService.Delete:output_type -> Void
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_candidate_proto_init() }
func file_protos_candidate_proto_init() {
	if File_protos_candidate_proto != nil {
		return
	}
	file_protos_void_proto_init()
	file_protos_election_proto_init()
	file_protos_public_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_candidate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candidate); i {
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
		file_protos_candidate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCandidateReq); i {
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
		file_protos_candidate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCandidate); i {
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
		file_protos_candidate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateRes); i {
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
		file_protos_candidate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidatiesGetAllResp); i {
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
			RawDescriptor: file_protos_candidate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_candidate_proto_goTypes,
		DependencyIndexes: file_protos_candidate_proto_depIdxs,
		MessageInfos:      file_protos_candidate_proto_msgTypes,
	}.Build()
	File_protos_candidate_proto = out.File
	file_protos_candidate_proto_rawDesc = nil
	file_protos_candidate_proto_goTypes = nil
	file_protos_candidate_proto_depIdxs = nil
}