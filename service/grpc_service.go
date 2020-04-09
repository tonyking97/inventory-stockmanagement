package service

import (
	"../proto/inventory"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {}

func (*server) Ping(ctx context.Context, req *inventory_pb.PingRequest) (*inventory_pb.PingResponse, error){
	log.Println("Incoming Ping Request.")
	response := "pong"
	res := &inventory_pb.PingResponse{
		Pong:                 response,
	}
	return res, nil
}

func GRPCInit() {
	go startGRPCServer()
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	inventory_pb.RegisterInventoryServiceServer(s, &server{})

	reflection.Register(s)

	log.Println("Starting gRPC Service...")

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to start gRPC Service.", err)
	}
}

