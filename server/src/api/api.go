package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/FlowingSPDG/get5-web-go/server/src/db"
)

type CheckLoggedInJSON struct {
	IsLoggedIn bool   `json:"isLoggedIn"`
	SteamID    string `json:"steamid"`
	UserID     int    `json:"userid"`
}

func CheckLoggedIn(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CheckLoggedIn\n")
	response := CheckLoggedInJSON{
		IsLoggedIn: false,
	}
	session, _ := db.SessionStore.Get(r, db.SessionData)
	if _, ok := session.Values["Loggedin"]; ok {
		response.IsLoggedIn = session.Values["Loggedin"].(bool)
		response.SteamID = session.Values["SteamID"].(string)
		response.UserID = session.Values["UserID"].(int)
	}
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetMatchInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetMatchInfo\n")
	matchid := vars["matchID"]
	match := db.MatchData{}
	mapstats := []APIMapStatsData{}
	server := APIGameServerData{}
	team1 := APITeamData{}
	team2 := APITeamData{}
	user := APIUserData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&match)
	db.SQLAccess.Gorm.Where("match_id = ?", matchid).Limit(7).Find(&mapstats)
	db.SQLAccess.Gorm.Where("id = ?", match.ServerID).First(&server)
	db.SQLAccess.Gorm.Where("id = ?", match.Team1ID).First(&team1)
	db.SQLAccess.Gorm.Where("id = ?", match.Team2ID).First(&team2)
	db.SQLAccess.Gorm.Where("id = ?", match.UserID).First(&user)
	var winner int64
	if match.Winner.Valid {
		winner_, err := match.Winner.Value()
		if err != nil {
			winner = 0
		}
		winner = winner_.(int64)
	}
	starttime := match.StartTime.Time
	endtime := match.EndTime.Time
	status, err := match.GetStatusString(true)
	if err != nil {
		status = ""
	}

	for i := 0; i < len(mapstats); i++ {
		if mapstats[i].StartTime.Valid {
			mapstats[i].StartTimeJSON = mapstats[i].StartTime.Time
		}
		if mapstats[i].EndTime.Valid {
			mapstats[i].EndTimeJSON = mapstats[i].EndTime.Time
		}
	}
	response := APIMatchData{
		ID:            match.ID,
		UserID:        match.UserID,
		Winner:        winner,
		Cancelled:     match.Cancelled,
		StartTimeJSON: starttime,
		EndTimeJSON:   endtime,
		MaxMaps:       match.MaxMaps,
		Title:         match.Title,
		SkipVeto:      match.SkipVeto,
		VetoMapPool:   strings.Split(match.VetoMapPool, " "),
		Team1Score:    match.Team1Score,
		Team2Score:    match.Team2Score,
		Team1String:   match.Team1String,
		Team2String:   match.Team2String,
		Forfeit:       match.Forfeit,
		Pending:       match.Pending(),
		Live:          match.Live(),
		Server:        server,
		MapStats:      mapstats,
		Team1:         team1,
		Team2:         team2,
		User:          user,
		Status:        status,
	}

	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(string(jsonbyte))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetPlayerStatInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetPlayerStatInfo\n")
	q := r.URL.Query()
	matchID := vars["matchID"]
	mapID := q.Get("mapID")
	response := []APIPlayerStatsData{}
	db.SQLAccess.Gorm.Where("match_id = ? AND map_id = ?", matchID, mapID).Limit(10).Find(&response)
	for i := 0; i < len(response); i++ { // Calculates by server-side for avoiding JavaScript's float restrcition
		response[i].Rating = response[i].GetRating()
		response[i].KDR = response[i].GetKDR()
		response[i].HSP = response[i].GetHSP()
		response[i].ADR = response[i].GetADR()
		response[i].FPR = response[i].GetFPR()
	}

	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(string(jsonbyte))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetUserInfo\n")
	userid := vars["userID"]
	response := db.UserData{}
	db.SQLAccess.Gorm.Where("id = ?", userid).First(&response)
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(string(jsonbyte))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetServerInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetServerInfo\n")
	serverID := vars["serverID"]
	response := db.GameServerData{}
	db.SQLAccess.Gorm.Where("id = ?", serverID).First(&response)
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(string(jsonbyte))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetStatusString(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetStatusString\n")
	matchid := vars["matchID"]
	response := db.MatchData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&response)
	w.Header().Set("Content-Type", "application/json")
	status, err := response.GetStatusString(true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write([]byte(status))
}

func GetMatches(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetMatches\n")
	q := r.URL.Query()
	userID := q.Get("userID")
	response := []db.MatchData{}
	if userID != "" {
		db.SQLAccess.Gorm.Limit(20).Where("user_id = ?", userID).Order("id DESC").Find(&response)
	} else {
		db.SQLAccess.Gorm.Limit(20).Order("id DESC").Find(&response)
	}
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(string(jsonbyte))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetTeamInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetMatchInfo\n")
	teamid := vars["teamID"]
	response := APITeamData{}
	db.SQLAccess.Gorm.Where("id = ?", teamid).First(&response)
	steamids, err := response.GetPlayers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	for i := 0; i < len(steamids); i++ {
		response.SteamIDs = append(response.SteamIDs, steamids[i])
	}
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetRecentMatches(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetRecentMatches\n vars:%v", vars)
	teamID := vars["teamID"]
	response := []db.MatchData{}
	if teamID == "" {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		db.SQLAccess.Gorm.Limit(20).Order("id DESC").Where("team1_id = ?", teamID).Or("team2_id = ?", teamID).Find(&response)
	}
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(string(jsonbyte))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func CheckUserCanEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CheckUserCanEdit\n")
	vars := mux.Vars(r)
	q := r.URL.Query()
	teamid := vars["teamID"]
	useridstr := q.Get("userID")
	team := db.TeamData{}
	res := SimpleJSONResponse{}
	db.SQLAccess.Gorm.Where("id = ?", teamid).First(&team)
	userid, err := strconv.Atoi(useridstr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	res.Response = strconv.FormatBool(team.CanEdit(userid))
	jsonbyte, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetMetrics\n")
	Metrics := db.GetMetrics()
	jsonbyte, err := json.Marshal(Metrics)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(string(jsonbyte))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetSteamName(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetSteamName\n")
	q := r.URL.Query()
	steamid := q.Get("steamID")
	steamid64, err := strconv.Atoi(steamid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	steamname, err := db.GetSteamName(uint64(steamid64))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write([]byte(steamname))
}
