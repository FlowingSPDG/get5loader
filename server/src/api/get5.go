package api

import (
	"encoding/json"
	"fmt"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

// MatchAPICheck Checks if API is available or not
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
		} else { // added,this is not exist in original get5-web
			MatchUpdate.Cancelled = true
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
	matchid, err := strconv.Atoi(vars["matchID"])
	if err != nil {
		http.Error(w, "matchid should be int", http.StatusBadRequest)
		return
	}
	mapnumber, err := strconv.Atoi(vars["mapNumber"])
	if err != nil {
		http.Error(w, "mapnumber should be int", http.StatusBadRequest)
		return
	}
	mapname := r.FormValue("mapname")
	fmt.Printf("mapname : %s\n", mapname)
	var m = db.MatchData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&m)
	err = MatchAPICheck(&m, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mUpdate := m
	db.SQLAccess.Gorm.First(&mUpdate)
	MapStats := &db.MapStatsData{}
	MapStats, err = MapStats.GetOrCreate(matchid, mapnumber, mapname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mUpdate.StartTime.Scan(time.Now())
	db.SQLAccess.Gorm.Model(&m).Update(&mUpdate)
	db.SQLAccess.Gorm.Save(&mUpdate)
	fmt.Println(MapStats)
}

// MatchMapUpdateHandler Handler for /api/v1/match/{matchID}/map/{mapNumber}/update API. // TODO
func MatchMapUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchMapUpdateHandler")
	fmt.Printf("team1score : %s\n", r.FormValue("team1score"))
	fmt.Printf("team2score : %s\n", r.FormValue("team2score"))
}

// MatchMapFinishHandler Handler for /api/v1/match/{matchID}/map/{mapNumber}/finish API. // TODO
func MatchMapFinishHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchMapFinishHandler")
	fmt.Printf("winner : %s\n", r.FormValue("winner"))
}

// MatchMapPlayerUpdateHandler Handler for /api/v1/match/{matchID}/map/{mapNumber}/player/{steamid64}/update API.
func MatchMapPlayerUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchMapPlayerUpdateHandler")
	vars := mux.Vars(r)
	matchid, err := strconv.Atoi(vars["matchID"])
	if err != nil {
		http.Error(w, "matchid should be int", http.StatusBadRequest)
		return
	}
	mapnumber, err := strconv.Atoi(vars["mapNumber"])
	if err != nil {
		http.Error(w, "mapnumber should be int", http.StatusBadRequest)
		return
	}
	steamid64 := vars["steamid64"]

	FormName := r.FormValue("name")
	FormTeam := r.FormValue("team")
	FormKills, err := strconv.Atoi(r.FormValue("kills"))
	if err != nil {
		http.Error(w, "kills should be int", http.StatusBadRequest)
		return
	}
	FormAssists, err := strconv.Atoi(r.FormValue("assists"))
	if err != nil {
		http.Error(w, "assists should be int", http.StatusBadRequest)
		return
	}
	FormDeaths, err := strconv.Atoi(r.FormValue("deaths"))
	if err != nil {
		http.Error(w, "deaths should be int", http.StatusBadRequest)
		return
	}
	FormFlashbangAssists, err := strconv.Atoi(r.FormValue("flashbang_assists"))
	if err != nil {
		http.Error(w, "flashbang_assists should be int", http.StatusBadRequest)
		return
	}
	FormTeamKills, err := strconv.Atoi(r.FormValue("teamkills"))
	if err != nil {
		http.Error(w, "teamkills should be int", http.StatusBadRequest)
		return
	}
	FormSuicides, err := strconv.Atoi(r.FormValue("suicides"))
	if err != nil {
		http.Error(w, "suicides should be int", http.StatusBadRequest)
		return
	}
	FormDamage, err := strconv.Atoi(r.FormValue("damage"))
	if err != nil {
		http.Error(w, "damage should be int", http.StatusBadRequest)
		return
	}
	FormHeadShotKills, err := strconv.Atoi(r.FormValue("headshot_kills"))
	if err != nil {
		http.Error(w, "headshot_kills should be int", http.StatusBadRequest)
		return
	}
	FormRoundsPlayed, err := strconv.Atoi(r.FormValue("roundsplayed"))
	if err != nil {
		http.Error(w, "roundsplayed should be int", http.StatusBadRequest)
		return
	}
	FormBombPlants, err := strconv.Atoi(r.FormValue("bomb_plants"))
	if err != nil {
		http.Error(w, "bomb_plants should be int", http.StatusBadRequest)
		return
	}
	FormBombDefuses, err := strconv.Atoi(r.FormValue("bomb_defuses"))
	if err != nil {
		http.Error(w, "bomb_defuses should be int", http.StatusBadRequest)
		return
	}

	Form1KillRounds, err := strconv.Atoi(r.FormValue("1kill_rounds"))
	if err != nil {
		http.Error(w, "1kill_rounds should be int", http.StatusBadRequest)
		return
	}
	Form2KillRounds, err := strconv.Atoi(r.FormValue("2kill_rounds"))
	if err != nil {
		http.Error(w, "3kill_rounds should be int", http.StatusBadRequest)
		return
	}
	Form3KillRounds, err := strconv.Atoi(r.FormValue("3kill_rounds"))
	if err != nil {
		http.Error(w, "3kill_rounds should be int", http.StatusBadRequest)
		return
	}
	Form4KillRounds, err := strconv.Atoi(r.FormValue("4kill_rounds"))
	if err != nil {
		http.Error(w, "4kill_rounds should be int", http.StatusBadRequest)
		return
	}
	Form5KillRounds, err := strconv.Atoi(r.FormValue("5kill_rounds"))
	if err != nil {
		http.Error(w, "5kill_rounds should be int", http.StatusBadRequest)
		return
	}
	FormV1, err := strconv.Atoi(r.FormValue("v1"))
	if err != nil {
		http.Error(w, "v1 should be int", http.StatusBadRequest)
		return
	}
	FormV2, err := strconv.Atoi(r.FormValue("v2"))
	if err != nil {
		http.Error(w, "v2 should be int", http.StatusBadRequest)
		return
	}
	FormV3, err := strconv.Atoi(r.FormValue("v3"))
	if err != nil {
		http.Error(w, "v3 should be int", http.StatusBadRequest)
		return
	}
	FormV4, err := strconv.Atoi(r.FormValue("v4"))
	if err != nil {
		http.Error(w, "v4 should be int", http.StatusBadRequest)
		return
	}
	FormV5, err := strconv.Atoi(r.FormValue("v5"))
	if err != nil {
		http.Error(w, "v5 should be int", http.StatusBadRequest)
		return
	}
	FormFirstKillT, err := strconv.Atoi(r.FormValue("firstkill_t"))
	if err != nil {
		http.Error(w, "firstkill_t should be int", http.StatusBadRequest)
		return
	}
	FormFirstKillCT, err := strconv.Atoi(r.FormValue("firstkill_ct"))
	if err != nil {
		http.Error(w, "firstkill_ct should be int", http.StatusBadRequest)
		return
	}
	FormFirstDeathT, err := strconv.Atoi(r.FormValue("firstdeath_t"))
	if err != nil {
		http.Error(w, "firstdeath_t should be int", http.StatusBadRequest)
		return
	}
	FormFirstDeathCT, err := strconv.Atoi(r.FormValue("firstdeath_ct"))
	if err != nil {
		http.Error(w, "firstdeath_ct should be int", http.StatusBadRequest)
		return
	}
	/*
		Form, err := strconv.Atoi(r.FormValue("tradekill")) // https://github.com/FlowingSPDG/get5-webapi/blob/e41ac0ab3c698ed67dbadcd667e55feef403e074/scripting/get5_apistats.sp#L429
		if err != nil {
			http.Error(w, "tradekill should be int", http.StatusBadRequest)
			return
		}
	*/

	fmt.Printf("matchid : %d\n", matchid)
	fmt.Printf("mapnumber : %d\n", mapnumber)
	fmt.Printf("key : %s\n", r.FormValue("key"))
	fmt.Printf("name : %s\n", FormName)
	fmt.Printf("team : %s\n", FormTeam)
	fmt.Printf("kills : %d\n", FormKills)
	fmt.Printf("assists : %d\n", FormAssists)
	fmt.Printf("deaths : %d\n", FormDeaths)
	fmt.Printf("flashbang_assists : %d\n", FormFlashbangAssists)
	fmt.Printf("teamkills : %d\n", FormTeamKills)
	fmt.Printf("suicides : %d\n", FormSuicides) // not working?
	fmt.Printf("damage : %d\n", FormDamage)
	fmt.Printf("headshot_kills : %d\n", FormHeadShotKills)
	fmt.Printf("roundsplayed : %d\n", FormRoundsPlayed)
	fmt.Printf("bomb_plants : %d\n", FormBombPlants)
	fmt.Printf("bomb_defuses : %d\n", FormBombDefuses)
	fmt.Printf("1kill_rounds : %d\n", Form1KillRounds)
	fmt.Printf("2kill_rounds : %d\n", Form2KillRounds)
	fmt.Printf("3kill_rounds : %d\n", Form3KillRounds)
	fmt.Printf("4kill_rounds : %d\n", Form4KillRounds)
	fmt.Printf("5kill_rounds : %d\n", Form5KillRounds)
	fmt.Printf("v1 : %d\n", FormV1)
	fmt.Printf("v2 : %d\n", FormV2)
	fmt.Printf("v3 : %d\n", FormV3)
	fmt.Printf("v4 : %d\n", FormV4)
	fmt.Printf("v5 : %d\n", FormV5)
	fmt.Printf("firstkill_t : %d\n", FormFirstKillT)
	fmt.Printf("firstkill_ct : %d\n", FormFirstKillCT)
	fmt.Printf("firstdeath_t : %d\n", FormFirstDeathT)
	fmt.Printf("firstdeath_ct : %d\n", FormFirstDeathCT)

	var m = db.MatchData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&m)
	if m.APIKey != r.FormValue("key") {
		http.Error(w, "Wrong API Key", http.StatusBadRequest)
		return
	}

	MapStats := db.MapStatsData{}
	MapStatsRecord := db.SQLAccess.Gorm.Where("match_id = ? AND map_number = ? ", matchid, mapnumber).First(&MapStats)
	if !MapStatsRecord.RecordNotFound() {
		p := &db.PlayerStatsData{}
		p, err := p.GetOrCreate(matchid, mapnumber, steamid64)
		if err != nil {
			http.Error(w, "Failed to solve player stats object", http.StatusNotFound)
			return
		}
		pUpdate := p
		db.SQLAccess.Gorm.First(&pUpdate)

		if FormTeam == "team1" {
			p.TeamID = m.Team1ID
		} else if FormTeam == "team2" {
			p.TeamID = m.Team2ID
		}
		pUpdate.Name = FormName
		pUpdate.Kills = FormKills
		pUpdate.Assists = FormAssists
		pUpdate.Deaths = FormDeaths
		pUpdate.FlashbangAssists = FormFlashbangAssists
		pUpdate.Teamkills = FormTeamKills
		pUpdate.Suicides = FormSuicides
		pUpdate.Damage = FormDamage
		pUpdate.HeadshotKills = FormHeadShotKills
		pUpdate.Roundsplayed = FormRoundsPlayed
		pUpdate.BombPlants = FormBombPlants
		pUpdate.BombDefuses = FormBombDefuses
		pUpdate.K1 = Form1KillRounds
		pUpdate.K2 = Form2KillRounds
		pUpdate.K3 = Form3KillRounds
		pUpdate.K4 = Form4KillRounds
		pUpdate.K5 = Form5KillRounds
		pUpdate.V1 = FormV1
		pUpdate.V2 = FormV2
		pUpdate.V3 = FormV3
		pUpdate.V4 = FormV4
		pUpdate.V5 = FormV5
		pUpdate.FirstkillT = FormFirstKillT
		pUpdate.FirstkillCT = FormFirstKillCT
		pUpdate.FirstdeathT = FormFirstDeathT
		pUpdate.FirstdeathCT = FormFirstDeathCT
		db.SQLAccess.Gorm.Model(&p).Update(&pUpdate)
		db.SQLAccess.Gorm.Save(&pUpdate)
	} else {
		http.Error(w, "Failed to find map stats object", http.StatusNotFound)
		return
	}

}

// MatchVetoUpdateHandler Handler for /api/v1/match/{matchID}/vetoUpdate API. // TODO
func MatchVetoUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchVetoUpdateHandler")
	vars := mux.Vars(r)
	matchid, err := strconv.Atoi(vars["matchID"])
	if err != nil {
		http.Error(w, "matchid should be int", http.StatusBadRequest)
		return
	}

	FormMap := r.FormValue("map")
	FormTeamString := r.FormValue("teamString")
	FormPickOrVeto := r.FormValue("pick_or_veto")

	fmt.Printf("matchid : %d\n", matchid)
	fmt.Printf("matchid : %s\n", FormMap)
	fmt.Printf("matchid : %s\n", FormTeamString)
	fmt.Printf("matchid : %s\n", FormPickOrVeto)

	m := &db.MatchData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&m)
	err = MatchAPICheck(m, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var TeamName string
	if FormTeamString == "team1" {
		team := &db.TeamData{}
		team.ID = m.Team1ID
		db.SQLAccess.Gorm.First(&team)
		TeamName = team.Name
	} else if FormTeamString == "team2" {
		team := &db.TeamData{}
		team.ID = m.Team2ID
		db.SQLAccess.Gorm.First(&team)
		TeamName = team.Name
	} else {
		TeamName = "Decider"
	}
	// veto = Veto.create(matchid, teamName, request.values.get('map'), request.values.get('pick_or_veto'))
	// TODO : Add Veto struct in db/models.go
	// Register to DB
	fmt.Printf("Confirmed Map Veto For %s on map %s\n", TeamName, FormMap)
}

// MatchDemoUploadHandler Handler for /api/v1/match/{matchID}/map/{mapNumber}/demo API. // TODO
func MatchDemoUploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchDemoUploadHandler")
	vars := mux.Vars(r)
	matchid, err := strconv.Atoi(vars["matchID"])
	if err != nil {
		http.Error(w, "matchid should be int", http.StatusBadRequest)
		return
	}
	mapNumber, err := strconv.Atoi(vars["mapNumber"])
	if err != nil {
		http.Error(w, "mapNumber should be int", http.StatusBadRequest)
		return
	}

	DemoFile := r.FormValue("demoFile")

	m := &db.MatchData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&m)
	err = MatchAPICheck(m, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mapstats := &db.MapStatsData{}
	mapstatsRecord := db.SQLAccess.Gorm.Where("match_id = ? AND map_number", matchid, mapNumber).First(&mapstats)
	if !mapstatsRecord.RecordNotFound() {
		mapstatsRecord.Update("demoFile", DemoFile)
		return
	}
	http.Error(w, "Failed to find map stats object", http.StatusBadRequest)
	fmt.Println("Made it through the demo post.")
}
