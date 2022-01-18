package main

import (
	"context"
	"fmt"
	"gRPC/pb"
	"io"
	"time"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		print("Could not connect to gRPC server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUserStreamBoth(client)
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

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "JOAO",
		Email: "teste@teste.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		print("Could not make gRPC server: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			print("msg error")
		}
		print("Status:", stream.Status, "\n")
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "JOAO",
			Email: "teste@teste.com",
		},
		&pb.User{
			Id:    "1",
			Name:  "Leandro",
			Email: "leandro@teste.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Ricardito",
			Email: "rcks@teste.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Superman",
			Email: "spm@teste.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Raj",
			Email: "raj@teste.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		print("err on users stream")
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		print("Error on closing")
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "JOAO",
			Email: "teste@teste.com",
		},
		&pb.User{
			Id:    "1",
			Name:  "Leandro",
			Email: "leandro@teste.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Ricardito",
			Email: "rcks@teste.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Superman",
			Email: "spm@teste.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Raj",
			Email: "raj@teste.com",
		},
	}

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		print("Could not make gRPC server: %v", err)
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			print("Sending user:", req.Name, "\n")
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				print("Error on client receiving")
				break
			}
			print("Receiving user:", res.GetUser().GetName(), "\n")
		}
		close(wait)
	}()

	<-wait

}
