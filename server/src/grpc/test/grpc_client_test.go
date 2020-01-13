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

func TestGrpcGetUser(t *testing.T) {
	log.Printf("Starting GET5 gRPC Client...")
	conn, err := grpc.Dial("127.0.0.1:50055", grpc.WithInsecure())
	if err != nil {
		t.Errorf("client connection error:%v\n", err)
	}
	defer conn.Close()
	client := pb.NewGet5Client(conn)
	req := &pb.GetUserRequest{Ids: &pb.GetUserRequest_Steamid{Steamid: "76561198072054549"}}
	res, err := client.GetUser(context.TODO(), req)
	if err != nil {
		t.Errorf("error::%v \n", err)
		return
	}
	log.Printf("result: %v \n", res)
}

func TestGrpcEditUser(t *testing.T) {
	/* TODO */
}

func TestGrpcDeleteUser(t *testing.T) {
	log.Printf("Starting GET5 gRPC Client...")
	conn, err := grpc.Dial("127.0.0.1:50055", grpc.WithInsecure())
	if err != nil {
		t.Errorf("client connection error:%v\n", err)
	}
	defer conn.Close()
	client := pb.NewGet5Client(conn)
	req := &pb.DeleteUserRequest{Id: 0}
	res, err := client.DeleteUser(context.TODO(), req)
	if err != nil {
		t.Errorf("error::%v \n", err)
		return
	}
	log.Printf("result: %v \n", res)
}

func TestGrpcRegisterGameServer(t *testing.T) {
	log.Printf("Starting GET5 gRPC Client...")
	conn, err := grpc.Dial("127.0.0.1:50055", grpc.WithInsecure())
	if err != nil {
		t.Errorf("client connection error:%v\n", err)
	}
	defer conn.Close()
	client := pb.NewGet5Client(conn)
	req := &pb.RegisterGameServerRequest{
		Userid:       0,
		DisplayName:  "gRPC API TEST",
		IpString:     "0.0.0.0",
		Port:         27015,
		RconPassword: "wasd",
		PublicServer: false,
	}
	res, err := client.RegisterGameServer(context.TODO(), req)
	if err != nil {
		t.Errorf("error::%v \n", err)
		return
	}
	log.Printf("result: %v \n", res)
}

func TestGrpcGetGameServer(t *testing.T) {
	log.Printf("Starting GET5 gRPC Client...")
	conn, err := grpc.Dial("127.0.0.1:50055", grpc.WithInsecure())
	if err != nil {
		t.Errorf("client connection error:%v\n", err)
	}
	defer conn.Close()
	client := pb.NewGet5Client(conn)
	req := &pb.GetGameServerRequest{
		Id: 80,
	}
	res, err := client.GetGameServer(context.TODO(), req)
	if err != nil {
		t.Errorf("error::%v \n", err)
		return
	}
	log.Printf("result: %v \n", res)
}

func TestGrpcDeleteGameServer(t *testing.T) {
	log.Printf("Starting GET5 gRPC Client...")
	conn, err := grpc.Dial("127.0.0.1:50055", grpc.WithInsecure())
	if err != nil {
		t.Errorf("client connection error:%v\n", err)
	}
	defer conn.Close()
	client := pb.NewGet5Client(conn)
	req := &pb.DeleteGameServerRequest{
		Id: 80,
	}
	res, err := client.DeleteGameServer(context.TODO(), req)
	if err != nil {
		t.Errorf("error::%v \n", err)
		return
	}
	log.Printf("result: %v \n", res)
}
