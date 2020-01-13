package test

import (
	"context"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestGrpcRegisterUser(t *testing.T) {
	log.Printf("Starting GET5 gRPC Client...")
	conn, err := grpc.Dial("127.0.0.1:50055", grpc.WithInsecure())
	if err != nil {
		t.Errorf("client connection error:%v\n", err)
	}
	defer conn.Close()
	client := pb.NewGet5Client(conn)
	req := &pb.RegisterUserRequest{Steamid: "76561198072054549"}
	res, err := client.RegisterUser(context.TODO(), req)
	if err != nil {
		t.Errorf("error::%v \n", err)
		return
	}
	log.Printf("result: %v \n", res)
}
