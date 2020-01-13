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

func (s Server) RegisterMatch(ctx context.Context, req *pb.RegisterMatchRequest) (*pb.RegisterMatchReply, error) {
	return nil, nil // TODO
}

func (s Server) GetMatch(ctx context.Context, req *pb.GetMatchRequest) (*pb.GetMatchReply, error) {
	return nil, nil // TODO
}

func (s Server) EditMatch(ctx context.Context, req *pb.EditMatchRequest) (*pb.EditMatchReply, error) {
	return nil, nil // TODO
}

func (s Server) DeleteMatch(ctx context.Context, req *pb.DeleteMatchRequest) (*pb.DeleteMatchReply, error) {
	return nil, nil // TODO
}
