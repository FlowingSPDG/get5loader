package api

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	response := APIMatchData{
		ID:          match.ID,
		UserID:      match.UserID,
		Winner:      winner,
		Cancelled:   match.Cancelled,
		StartTime:   starttime,
		EndTime:     endtime,
		MaxMaps:     match.MaxMaps,
		Title:       match.Title,
		SkipVeto:    match.SkipVeto,
		VetoMapPool: strings.Split(match.VetoMapPool, " "),
		Team1Score:  match.Team1Score,
		Team2Score:  match.Team2Score,
		Team1String: match.Team1String,
		Team2String: match.Team2String,
		Forfeit:     match.Forfeit,
		Pending:     match.Pending(),
		Live:        match.Live(),
		Server:      server,
		MapStats:    mapstats,
		Team1:       team1,
		Team2:       team2,
		User:        user,
		Status:      status,
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
	//vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("GetMatches\n")
	response := []db.MatchData{}
	db.SQLAccess.Gorm.Limit(20).Order("id DESC").Find(&response)
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
	matchid := vars["teamID"]
	response := db.TeamData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&response)
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(string(jsonbyte))
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
