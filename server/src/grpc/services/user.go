package services

import (
	"fmt"
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
		fmt.Println("Type : *pb.GetUserRequest_Id")
		db.SQLAccess.Gorm.First(&user, req.GetId())
	case *pb.GetUserRequest_Steamid:
		fmt.Println("Type : *pb.GetUserRequest_Steamid")
		db.SQLAccess.Gorm.Where("steam_id = ?", req.GetSteamid()).First(&user)
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

func (s Server) GetOrRegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserReply, error) {
	user := &db.UserData{
		SteamID: req.GetSteamid(),
	}
	fmt.Printf("req : %v\n", req)
	user, _, err := user.GetOrCreate()
	if err != nil {
		return nil, err
	}
	fmt.Printf("user : %v\n", user)
	return &pb.RegisterUserReply{
		User: &pb.UserData{
			Id:      int32(user.ID),
			Steamid: user.SteamID,
			Name:    user.Name,
			Admin:   user.Admin,
		},
	}, nil
}

func (s Server) EditUser(ctx context.Context, req *pb.EditUserRequest) (*pb.EditUserReply, error) {
	return &pb.EditUserReply{
		Error:        true,
		Errormessage: "NOT IMPLEMENTED YET",
	}, fmt.Errorf("NOT IMPLEMENTED YET") // TODO
}

func (s Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	user := &db.UserData{}
	rec := db.SQLAccess.Gorm.First(&user, req.GetId())
	if rec.RecordNotFound() {
		return &pb.DeleteUserReply{
			Error:        true,
			Errormessage: "User not found",
		}, fmt.Errorf("User not found")
	}
	db.SQLAccess.Gorm.Delete(&user)
	return &pb.DeleteUserReply{
		Error:        false,
		Errormessage: "",
	}, nil
}
