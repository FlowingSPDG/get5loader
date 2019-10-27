package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

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
	response := db.MatchData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&response)
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
	fmt.Printf("GetUserInfo\n")
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