package main

import (
	"log"
	"net"
	"project/config"
	"project/genproto/public"
	"project/service"
	"project/storage/postgres"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()
	db, err := postgres.ConnectDb(cfg)
	if err != nil {
		log.Println("error while connecting to postgres: ", err)
	}

	liss, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	public.RegisterCandidateServiceServer(s, service.NewCandidateService(db))
	public.RegisterElectionServiceServer(s, service.NewElectionService(db))
	public.RegisterPublicVoteServiceServer(s, service.NewPublicVoteService(db))

	reflection.Register(s)
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
