package main

import (
	"context"
	"flag"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
)

var (
	addr    *string
	matchid *int
	steamid *string
)

func init() {
	addr = flag.String("addr", "127.0.0.1:50055", "gRPC target address and port.")
	matchid = flag.Int("matchid", 0, "Streaming API MatchID.")
	steamid = flag.String("steamid", "76561198072054549", "User SteamID64")
	flag.Parse()
}

// go run main.go -addr 127.0.0.1:50055 -matchid 100
func main() {
	log.Println("Starting GET5 gRPC Client...")
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client connection error:%v\n", err)
	}
	defer conn.Close()
	client := pb.NewGet5Client(conn)
	userreq := &pb.GetUserRequest{
		Ids: &pb.GetUserRequest_Steamid{
			Steamid: *steamid,
		},
	}
	res, err := client.GetUser(context.TODO(), userreq)
	if err != nil {
		log.Fatalf("failed. error:%v\n", err)
	}
	log.Printf("result: %v \n", res)
	if *matchid == 0 {
		os.Exit(0)
	}
	log.Println("Starting event streaming...")
	streamreq := &pb.MatchEventRequest{
		Matchid: int32(*matchid), // MatchID should be int32,not int
	}
	stream, err := client.MatchEvent(context.Background(), streamreq)
	if err != nil {
	}
	for {
		// Wait for events...
		resp, err := stream.Recv()
		// Stops streaming if its over...
		if err == io.EOF {
			break
		}
		// If error occured...
		if err != nil {
			log.Fatalf("ERR : %v", err)
			break
		}
		// Real Data(if you dont want to handle by their event type)
		// log.Printf("Received : %v\n", resp)

		// Handle type assertion...
		switch e := resp.GetEvent().(type) {
		case *pb.MatchEventReply_Matchfinish:
			log.Printf("Match %d Match Finished. %v\n", *matchid, e.Matchfinish)
			break
		case *pb.MatchEventReply_Mapplayerupdate:
			log.Printf("Match %d Player stats updated. %v\n", *matchid, e.Mapplayerupdate)
		default:
			log.Printf("Match %d Received Event %v\n", *matchid, e)
		}
	}
	log.Println("Event streaming DONE.")
}
