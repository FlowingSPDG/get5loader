package services

import (
	"log"
	"strings"
	// "github.com/FlowingSPDG/get5-web-go/server/src/api"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	// "github.com/FlowingSPDG/get5-web-go/server/src/util"
	"context"
	// "google.golang.org/grpc"
	// "log"
)

func (s Server) RegisterMatch(ctx context.Context, req *pb.RegisterMatchRequest) (*pb.RegisterMatchReply, error) {
	log.Println("RegisterMatch")
	log.Printf("req : %v\n", req)
	match := &db.MatchData{}
	match, err := match.Create(
		int(req.GetUserid()),
		int(req.GetTeam1Id()),
		int(req.GetTeam2Id()),
		req.GetTeam1String(),
		req.GetTeam2String(),
		int(req.GetMaxmaps()),
		req.GetSkipveto(),
		req.GetTitle(),
		req.GetVetomappool(),
		int(req.GetServerid()),
		req.GetCvars(),
		req.GetSideType(),
		req.GetIsPug(),
	)

	log.Printf("match : %v\n", match)
	if err != nil {
		return &pb.RegisterMatchReply{
			Error:        true,
			Errormessage: "Could not register match",
		}, err
	}
	return &pb.RegisterMatchReply{
		Error:        false,
		Errormessage: "",
		Id:           int32(match.ID),
	}, nil
}

func (s Server) GetMatch(ctx context.Context, req *pb.GetMatchRequest) (*pb.GetMatchReply, error) {
	match := &db.MatchData{}
	match, err := match.Get(int(req.GetId()))
	if err != nil {
		return &pb.GetMatchReply{
			Error:        true,
			Errormessage: "Match not found",
			Match:        &pb.MatchData{},
		}, err
	}
	match.GetMapStat()
	mapstats := make([]*pb.MapStatsData, 0, len(match.MapStats))
	for i := 0; i < len(match.MapStats); i++ {
		mapstats = append(mapstats, &pb.MapStatsData{
			Id:        int32(match.MapStats[i].ID),
			Matchid:   int32(match.MapStats[i].MatchID),
			Mapnumber: int32(match.MapStats[i].MapNumber),
			Mapname:   match.MapStats[i].MapName,
			//Starttime:match.MapStats[i].StartTime,
			// Endtime:match.MapStats[i].EndTime,
			Winner:     int32(match.MapStats[i].Winner.Int32),
			Team1Score: int32(match.MapStats[i].Team1Score),
			Team2Score: int32(match.MapStats[i].Team2Score),
		})
	}

	return &pb.GetMatchReply{
		Error:        false,
		Errormessage: "",
		Match: &pb.MatchData{
			Id:        int32(match.ID),
			Userid:    int32(match.UserID),
			Serverid:  int32(match.ServerID),
			Winner:    match.Winner.Int32,
			Cancelled: match.Cancelled,
			// Starttime: match.StartTime, // ??
			// Endtime:match.EndTime,
			Maxmaps:       int32(match.MaxMaps),
			Title:         match.Title,
			Skipveto:      match.SkipVeto,
			Apikey:        match.APIKey,
			Vetomappool:   strings.Split(match.VetoMapPool, ","),
			Team1Score:    int32(match.Team1Score),
			Team2Score:    int32(match.Team2Score),
			Team1String:   match.Team1String,
			Team2String:   match.Team2String,
			Forfeit:       match.Forfeit,
			Pluginversion: match.PluginVersion,
			Mapstats:      mapstats,
		},
	}, nil
}

func (s Server) EditMatch(ctx context.Context, req *pb.EditMatchRequest) (*pb.EditMatchReply, error) {
	return nil, nil // TODO
}

func (s Server) DeleteMatch(ctx context.Context, req *pb.DeleteMatchRequest) (*pb.DeleteMatchReply, error) {
	return nil, nil // TODO
}
