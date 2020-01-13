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

func (s Server) RegisterGameServer(ctx context.Context, req *pb.RegisterGameServerRequest) (*pb.RegisterGameServerReply, error) {
	return nil, nil // TODO
}

func (s Server) GetGameServer(ctx context.Context, req *pb.GetGameServerRequest) (*pb.GetGameServerReply, error) {
	return nil, nil // TODO
}

func (s Server) EditGameServer(ctx context.Context, req *pb.EditGameServerRequest) (*pb.EditGameServerReply, error) {
	return nil, nil // TODO
}

func (s Server) DeleteGameServer(ctx context.Context, req *pb.DeleteGameServerRequest) (*pb.DeleteGameServerReply, error) {
	return nil, nil // TODO
}
