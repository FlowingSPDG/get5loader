package get5grpc

import (
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	pbservices "github.com/FlowingSPDG/get5-web-go/server/src/grpc/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGrpc(addr string) error {
	// create listiner
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterGet5Server(s, pbservices.Server{})
	log.Printf("Listening on %s", addr)

	// and start...
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
