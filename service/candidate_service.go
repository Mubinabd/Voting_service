package service

import (
	"context"
	pb "project/genproto/public"
	"project/storage"
)

type CandidateService struct {
	stg storage.StorageI
	pb.UnimplementedCandidateServiceServer
}

func NewCandidateService(stg storage.StorageI) *CandidateService {
	return &CandidateService{stg: stg}
}

func (cs *CandidateService) Create(ctx context.Context, req *pb.CreateCandidateReq) (*pb.Void, error) {
	err := cs.stg.Candidate().Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (cs *CandidateService) Get(ctx context.Context, req *pb.GetByIdReq) (*pb.CandidateRes, error) {
	res, err := cs.stg.Candidate().Get(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (cs *CandidateService) GetAll(ctx context.Context, req *pb.Filter) (*pb.CandidatiesGetAllResp, error) {
	res, err := cs.stg.Candidate().GetAll(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cs *CandidateService) Delete(ctx context.Context, req *pb.GetByIdReq) (*pb.Void, error) {
	err := cs.stg.Candidate().Delete(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
