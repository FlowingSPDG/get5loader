package services

import (
	// "github.com/FlowingSPDG/get5-web-go/server/src/api"
	// "github.com/FlowingSPDG/get5-web-go/server/src/db"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	// "github.com/FlowingSPDG/get5-web-go/server/src/util"
	"context"
	// "google.golang.org/grpc"
	// "log"
)

func (s Server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserReply, error) {
	// return nil, nil // TODO
	// debug...
	resp := &pb.RegisterUserReply{
		Error:        false,
		Errormessage: "Debug Error Message!",
	}
	return resp, nil
}

func (s Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return nil, nil // TODO
}

func (s Server) EditUser(ctx context.Context, req *pb.EditUserRequest) (*pb.EditUserReply, error) {
	return nil, nil // TODO
}

func (s Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return nil, nil // TODO
}
