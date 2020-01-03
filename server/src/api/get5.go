package api

import (
	"encoding/json"
	"fmt"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	"github.com/gorilla/mux"
	"net/http"
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

// MatchFinishHandler Handler for /api/v1/match/{matchID}/config API.
func MatchFinishHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchFinishHandler")
	fmt.Printf("winner : %s\n", r.FormValue("winner"))
	fmt.Printf("forfeit : %s\n", r.FormValue("forfeit"))

	vars := mux.Vars(r)
	matchid := vars["matchID"]
	var m = db.MatchData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&m)
	if m.APIKey != r.FormValue("key") {
		http.Error(w, "Wrong API Key", http.StatusBadRequest)
		return
	}
	if m.Finalized() {
		http.Error(w, "Match already finalized", http.StatusBadRequest)
		return
	}
	// Finish match here...
}

// MatchMapStartHandler Handler for /api/v1/match/{matchID}/map/{mapNumber}/start  API.
func MatchMapStartHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchMapStartHandler")
	fmt.Printf("mapname : %s\n", r.FormValue("mapname"))
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
	fmt.Printf("suicides : %s\n", r.FormValue("suicides"))
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
