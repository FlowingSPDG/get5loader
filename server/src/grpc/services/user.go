package services

import (
	//"github.com/FlowingSPDG/get5-web-go/server/src/api"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	// "github.com/FlowingSPDG/get5-web-go/server/src/util"
	"context"
	// "google.golang.org/grpc"
	// "log"
)

func (s Server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserReply, error) {
	user := &db.UserData{}
	user.SteamID = req.GetSteamid()
	user, exist, err := user.GetOrCreate()
	if err != nil {
		return &pb.RegisterUserReply{
			Error:        true,
			Errormessage: err.Error(),
		}, err
	}
	if exist {
		return &pb.RegisterUserReply{
			Error:        true,
			Errormessage: "User exist",
		}, err
	}
	resp := &pb.RegisterUserReply{
		Error:        false,
		Errormessage: "",
	}
	return resp, nil
}

func (s Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user := &db.UserData{}
	switch req.Ids.(type) {
	case *pb.GetUserRequest_Id:
		user.SteamID = req.GetSteamid()
	case *pb.GetUserRequest_Steamid:
		user.ID = int(req.GetId())
		// DO gorm thing...
	}
	return &pb.GetUserReply{
		User: &pb.UserData{
			Id:      int32(user.ID),
			Steamid: user.SteamID,
			Name:    user.Name,
			Admin:   user.Admin,
		},
	}, nil
}

func (s Server) EditUser(ctx context.Context, req *pb.EditUserRequest) (*pb.EditUserReply, error) {
	return nil, nil // TODO
}

func (s Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return nil, nil // TODO
}
