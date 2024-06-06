package service

import (
	"context"
	pb "project/genproto/public"
	"project/storage"
)

type ElectionService struct {
	stg storage.StorageI
	pb.UnimplementedElectionServiceServer
}

func NewElectionService(stg storage.StorageI) *ElectionService {
	return &ElectionService{stg: stg}
}

func (es *ElectionService) Create(ctx context.Context, req *pb.CreateElectionReq)(*pb.Void, error ){
	err := es.stg.Election().Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{},nil
}
func (es *ElectionService) GetById(ctx context.Context, req *pb.GetByIdReq) (*pb.ElectionRes, error) {
	res,err := es.stg.Election().Get(req)
	if err!= nil {
		return nil, err
	}
	return res,nil
}
func (es *ElectionService) GetAll(ctx context.Context, req *pb.Filter) (*pb.ElectionsGetAllResp, error) {
	res,err := es.stg.Election().GetAll(req)
	if err!= nil {
		return nil, err
	}
	return res,nil
}

func (es *ElectionService) UpdateItem(ctx context.Context, elec *pb.ElectionUpdate) (*pb.Void, error) {
	err := es.stg.Election().Update(elec)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (es *ElectionService) Delete(ctx context.Context, req *pb.GetByIdReq) (*pb.Void, error ) {
	err := es.stg.Election().Delete(req)
    if err != nil {
        return nil, err
    }
    return &pb.Void{},nil
}