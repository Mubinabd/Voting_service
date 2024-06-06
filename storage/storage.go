package storage

import (
	pb "project/genproto/public"
)
type StorageI interface {
	Candidate() CandidateI
	Election() ElectionI
	PublicVotes() PublicVotesI
}

type CandidateI interface {
	Create(candidate *pb.CreateCandidateReq) error
	Get(id *pb.GetByIdReq) (*pb.CandidateRes, error)
	GetAll(filter *pb.Filter) (*pb.CandidatiesGetAllResp, error)
	Delete(id *pb.GetByIdReq) error
}

type ElectionI interface {
	Create(election *pb.CreateElectionReq) error
	Get(id *pb.GetByIdReq) (*pb.ElectionRes, error)
	GetAll(filter *pb.Filter) (*pb.ElectionsGetAllResp, error)
	Update(update *pb.ElectionUpdate) error
	Delete(id *pb.GetByIdReq)  error
}

type PublicVotesI interface {
	Create(publicvote *pb.CreatePublicVoteReq)  error
	GetAll(filter *pb.Filter) (*pb.PublicVotesGetAllResp, error)
}