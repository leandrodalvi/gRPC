package services

import (
	"context"
	"gRPC/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	//operations like inserting on db

	return &pb.User{
		Email: req.GetEmail(),
		Id:    "1",
		Name:  req.GetName(),
	}, nil
}
