package api

import (
	"encoding/json"
	"fmt"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// MatchConfigHandler Handler for /api/v1/match/{matchID}/config API.
func MatchConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("MatchConfigHandler\n")
	matchid := vars["matchID"]
	match := db.MatchData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&match)
	res, err := match.BuildMatchDict()
	if err != nil {
		http.Error(w, "Internal ERROR", http.StatusInternalServerError)
		return
	}
	jsonbyte, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Internal ERROR", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

// MatchAPICheck Checks API is available or not
func MatchAPICheck(m *db.MatchData, r *http.Request) error {
	if m.APIKey != r.FormValue("key") {
		return fmt.Errorf("Wrong API Key")
	}
	if m.Finalized() {
		return fmt.Errorf("Match already finalized")
	}
	return nil
}

// MatchFinishHandler Handler for /api/v1/match/{matchID}/finish API.
func MatchFinishHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchFinishHandler")
	vars := mux.Vars(r)
	matchid := vars["matchID"]
	winner := r.FormValue("winner")
	forfeit := r.FormValue("forfeit")
	fmt.Printf("matchid : %s\n", matchid)
	fmt.Printf("winner : %s\n", winner)
	fmt.Printf("forfeit : %s\n", forfeit)
	var Match = db.MatchData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&Match)
	fmt.Printf("Requested API key : %s\n", r.FormValue("key"))
	fmt.Printf("Real API key : %s\n", Match.APIKey)
	var MatchUpdate = Match
	db.SQLAccess.Gorm.First(&MatchUpdate)
	err := MatchAPICheck(&Match, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if winner == "team1" {
		MatchUpdate.Winner.Scan(MatchUpdate.Team1ID)
	} else if winner == "team2" {
		MatchUpdate.Winner.Scan(MatchUpdate.Team2ID)
	} else {
		MatchUpdate.Winner.Scan(nil)
	}
	if forfeit == "1" {
		MatchUpdate.Forfeit = true
		if winner == "team1" {
			MatchUpdate.Team1Score = 1
			MatchUpdate.Team2Score = 0
		} else if winner == "team2" {
			MatchUpdate.Team1Score = 0
			MatchUpdate.Team2Score = 1
		}
	}
	MatchUpdate.EndTime.Scan(time.Now())
	server := MatchUpdate.GetServer()
	db.SQLAccess.Gorm.Model(&server).Update("in_use = true")
	db.SQLAccess.Gorm.Model(&Match).Update(&MatchUpdate)
	db.SQLAccess.Gorm.Save(&MatchUpdate)
	fmt.Printf("Finished match %v, winner = %v\n", MatchUpdate.ID, winner)
}

// MatchMapStartHandler Handler for /api/v1/match/{matchID}/map/{mapNumber}/start  API.
func MatchMapStartHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchMapStartHandler")
	vars := mux.Vars(r)
	matchid := vars["matchID"]
	//mapnumber := vars["mapNumber"]
	mapname := r.FormValue("mapname")
	fmt.Printf("mapname : %s\n", mapname)
	var m = db.MatchData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&m)
	err := MatchAPICheck(&m, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if m.StartTime.Value == nil {
		m.StartTime.Scan(time.Now())
	}

}

// MatchMapUpdateHandler Handler for /api/v1/match/{matchID}/map/{mapNumber}/update API.
func MatchMapUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchMapUpdateHandler")
	fmt.Printf("team1score : %s\n", r.FormValue("team1score"))
	fmt.Printf("team2score : %s\n", r.FormValue("team2score"))
}

// MatchMapFinishHandler Handler for /api/v1/match/{matchID}/map/{mapNumber}/finish API.
func MatchMapFinishHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchMapFinishHandler")
	fmt.Printf("winner : %s\n", r.FormValue("winner"))
}

// MatchMapPlayerUpdateHandler Handler for /api/v1/match/{matchID}/map/{mapNumber}/player/{steamid64}/update API.
func MatchMapPlayerUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchMapPlayerUpdateHandler")
	vars := mux.Vars(r)
	matchid := vars["matchID"]
	fmt.Printf("key : %s\n", r.FormValue("key"))
	fmt.Printf("name : %s\n", r.FormValue("name"))
	fmt.Printf("team : %s\n", r.FormValue("team"))
	fmt.Printf("kills : %s\n", r.FormValue("kills"))
	fmt.Printf("assists : %s\n", r.FormValue("assists"))
	fmt.Printf("deaths : %s\n", r.FormValue("deaths"))
	fmt.Printf("flashbang_assists : %s\n", r.FormValue("flashbang_assists"))
	fmt.Printf("teamkills : %s\n", r.FormValue("teamkills"))
	fmt.Printf("suicides : %s\n", r.FormValue("suicides")) // not working?
	fmt.Printf("damage : %s\n", r.FormValue("damage"))
	fmt.Printf("headshot_kills : %s\n", r.FormValue("headshot_kills"))
	fmt.Printf("roundsplayed : %s\n", r.FormValue("roundsplayed"))
	fmt.Printf("bomb_plants : %s\n", r.FormValue("bomb_plants"))
	fmt.Printf("bomb_defuses : %s\n", r.FormValue("bomb_defuses"))
	fmt.Printf("1kill_rounds : %s\n", r.FormValue("1kill_rounds"))
	fmt.Printf("2kill_rounds : %s\n", r.FormValue("2kill_rounds"))
	fmt.Printf("3kill_rounds : %s\n", r.FormValue("3kill_rounds"))
	fmt.Printf("4kill_rounds : %s\n", r.FormValue("4kill_rounds"))
	fmt.Printf("5kill_rounds : %s\n", r.FormValue("5kill_rounds"))
	fmt.Printf("v1 : %s\n", r.FormValue("v1"))
	fmt.Printf("v2 : %s\n", r.FormValue("v2"))
	fmt.Printf("v3 : %s\n", r.FormValue("v3"))
	fmt.Printf("v4 : %s\n", r.FormValue("v4"))
	fmt.Printf("v5 : %s\n", r.FormValue("v5"))
	fmt.Printf("firstkill_t : %s\n", r.FormValue("firstkill_t"))
	fmt.Printf("firstkill_ct : %s\n", r.FormValue("firstkill_ct"))
	fmt.Printf("firstdeath_t : %s\n", r.FormValue("firstdeath_t"))
	fmt.Printf("firstdeath_ct : %s\n", r.FormValue("firstdeath_ct"))

	var m = db.MatchData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&m)
	if m.APIKey != r.FormValue("key") {
		http.Error(w, "Wrong API Key", http.StatusBadRequest)
		return
	}
}
