package main

import (
	"gRPC/pb"
	"gRPC/services"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		print("Could not connect: %v", err)
	}

	grcpServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grcpServer, services.NewUserService())

	if err := grcpServer.Serve(lis); err != nil {
		print("Could not serve: %v", err)
	}
}
