package services

import (
	"fmt"
	// "github.com/FlowingSPDG/get5-web-go/server/src/api"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	// "github.com/FlowingSPDG/get5-web-go/server/src/util"
	"context"
	// "google.golang.org/grpc"
	// "log"
)

func (s Server) RegisterGameServer(ctx context.Context, req *pb.RegisterGameServerRequest) (*pb.RegisterGameServerReply, error) {
	srcds := &db.GameServerData{}
	srcds, err := srcds.Create(int(req.GetUserid()), req.GetDisplayName(), req.GetIpString(), int(req.GetPort()), req.GetRconPassword(), req.GetPublicServer())
	if err != nil {
		return &pb.RegisterGameServerReply{
			Error:        true,
			Errormessage: "Could not register server",
			Id:           0,
		}, err
	}
	return &pb.RegisterGameServerReply{
		Error:        false,
		Errormessage: "",
		Id:           int32(srcds.ID),
	}, nil
}

func (s Server) GetGameServer(ctx context.Context, req *pb.GetGameServerRequest) (*pb.GetGameServerReply, error) {
	srcds := &db.GameServerData{}
	rec := db.SQLAccess.Gorm.First(&srcds, req.GetId())
	if rec.RecordNotFound() {
		return &pb.GetGameServerReply{
			Error:        true,
			Errormessage: "Server not found",
			Gameserver:   &pb.GameServerData{},
		}, fmt.Errorf("Server not found")
	}
	return &pb.GetGameServerReply{
		Error:        false,
		Errormessage: "",
		Gameserver: &pb.GameServerData{
			Id:           int32(srcds.ID),
			Userid:       int32(srcds.UserID),
			InUse:        srcds.InUse,
			Ipstring:     srcds.IPString,
			Port:         int32(srcds.Port),
			RconPassword: srcds.RconPassword,
			DisplayName:  srcds.DisplayName,
			PublicServer: srcds.PublicServer,
		},
	}, nil
}

func (s Server) EditGameServer(ctx context.Context, req *pb.EditGameServerRequest) (*pb.EditGameServerReply, error) {
	return nil, nil // TODO
}

func (s Server) DeleteGameServer(ctx context.Context, req *pb.DeleteGameServerRequest) (*pb.DeleteGameServerReply, error) {
	srcds := &db.GameServerData{}
	rec := db.SQLAccess.Gorm.First(&srcds, req.GetId())
	if rec.RecordNotFound() {
		return &pb.DeleteGameServerReply{
			Error:        true,
			Errormessage: "Server not found",
		}, fmt.Errorf("User not found")
	}
	db.SQLAccess.Gorm.Delete(&srcds)
	return &pb.DeleteGameServerReply{
		Error:        false,
		Errormessage: "",
	}, nil
}
