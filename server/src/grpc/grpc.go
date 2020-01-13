package main

import (
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	pbservices "github.com/FlowingSPDG/get5-web-go/server/src/grpc/services"
	"net"

	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// create listiner
	addr := ":50055"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterGet5Server(s, pbservices.Server{})
	log.Printf("Listening on : %s", addr)

	// and start...
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
