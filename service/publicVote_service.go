package service

import (
	"context"
	pb "project/genproto/public"
	"project/storage"
)

type PublicVoteService struct {
	stg storage.StorageI
	pb.UnimplementedPublicVoteServiceServer
}

func NewPublicVoteService(stg storage.StorageI) *PublicVoteService {
	return &PublicVoteService{stg: stg}
}

func (pvs *PublicVoteService)Create(ctx context.Context, req *pb.CreatePublicVoteReq)(*pb.Void, error ){
	err := pvs.stg.PublicVotes().Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{},nil
}
func (pvs *PublicVoteService) GetAll(ctx context.Context, req *pb.Filter) (*pb.PublicVotesGetAllResp, error) {
	res,err := pvs.stg.PublicVotes().GetAll(req)
	if err!= nil {
		return nil, err
	}
	return res,nil
}
