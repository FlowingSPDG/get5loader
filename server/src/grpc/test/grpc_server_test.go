package test

import (
	"github.com/FlowingSPDG/get5-web-go/server/src/grpc"
	"log"
	"testing"
)

func TestGrpcStart(t *testing.T) {
	log.Printf("Starting GET5 gRPC Server...")
	err := get5grpc.StartGrpc(":50055")
	if err != nil {
		panic(err)
	}
}
