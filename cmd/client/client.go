package main

import (
	"context"
	"gRPC/pb"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		print("Could not connect to gRPC server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "JOAO",
		Email: "teste@teste.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		print("Could not make gRPC server: %v", err)
	}

	print(res)
}
