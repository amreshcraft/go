package main

import (
	"context"
	"cube-rpc/proto"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

)

type server struct {
	proto.UnimplementedCubeServiceServer
}

func (s *server) CalculateCube(ctx context.Context, req *proto.CubeRequest) (*proto.CubeResponse, error) {
	fmt.Println("Request received for number:", req.Number)

	result := req.Number * req.Number * req.Number

	return &proto.CubeResponse{
		Result: result,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterCubeServiceServer(grpcServer, &server{})

	reflection.Register(grpcServer)
	fmt.Println("gRPC Server started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
