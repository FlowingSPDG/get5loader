package api

import (
	"encoding/json"
	"fmt"
	"github.com/FlowingSPDG/get5-web-go/server/src/util"

	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/FlowingSPDG/get5-web-go/server/src/db"
)

const (
	// VERSION get5-web-go Version
	VERSION = "0.1.0"
)

// CheckLoggedInJSON Struct type for /api/v1/CheckLoggedIn API.
type CheckLoggedInJSON struct {
	IsLoggedIn bool   `json:"isLoggedIn"`
	IsAdmin    bool   `json:"isAdmin"`
	SteamID    string `json:"steamid"`
	UserID     int    `json:"userid"`
}

// GetVersionJSON Struct type for /api/v1/GetVersion API Response.
type GetVersionJSON struct {
	Version string `json:"version"`
}

// GetVersion handler for /api/v1/GetVersion API.
func GetVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CheckLoggedIn")
	response := GetVersionJSON{
		Version: VERSION,
	}
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

// CheckLoggedIn handler for /api/v1/CheckLoggedIn API.
func CheckLoggedIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CheckLoggedIn")
	response := CheckLoggedInJSON{
		IsLoggedIn: false,
	}
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		response.IsLoggedIn = s.Get("Loggedin").(bool)
		response.SteamID = s.Get("SteamID").(string)
		response.UserID = s.Get("UserID").(int)
		response.IsAdmin = s.Get("Admin").(bool)
	}
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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
		return
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
		return
	}
	fmt.Println(string(jsonbyte))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetUserInfo\n")
	userid := vars["userID"]
	response := APIUserData{}
	db.SQLAccess.Gorm.Where("id = ?", userid).First(&response)
	db.SQLAccess.Gorm.Where("user_id = ?", userid).Limit(20).Find(&response.Teams)
	db.SQLAccess.Gorm.Where("user_id = ?", userid).Limit(20).Find(&response.Servers)

	matches := []db.MatchData{}
	db.SQLAccess.Gorm.Where("user_id = ?", userid).Limit(20).Find(&matches)

	var m []APIMatchData
	for i := 0; i < len(matches); i++ {
		mapstats := []APIMapStatsData{}
		server := APIGameServerData{}
		team1 := APITeamData{}
		team2 := APITeamData{}
		user := APIUserData{}
		db.SQLAccess.Gorm.Where("match_id = ?", matches[i].ID).Limit(7).Find(&mapstats)
		db.SQLAccess.Gorm.Where("id = ?", matches[i].ServerID).First(&server)
		db.SQLAccess.Gorm.Where("id = ?", matches[i].Team1ID).First(&team1)
		db.SQLAccess.Gorm.Where("id = ?", matches[i].Team2ID).First(&team2)
		db.SQLAccess.Gorm.Where("id = ?", matches[i].UserID).First(&user)
		var winner int64
		if matches[i].Winner.Valid {
			winner_, err := matches[i].Winner.Value()
			if err != nil {
				winner = 0
			}
			winner = winner_.(int64)
		}
		starttime := matches[i].StartTime.Time
		endtime := matches[i].EndTime.Time
		status, err := matches[i].GetStatusString(true)
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
		matchdata := APIMatchData{
			ID:            matches[i].ID,
			UserID:        matches[i].UserID,
			Winner:        winner,
			Cancelled:     matches[i].Cancelled,
			StartTimeJSON: starttime,
			EndTimeJSON:   endtime,
			MaxMaps:       matches[i].MaxMaps,
			Title:         matches[i].Title,
			SkipVeto:      matches[i].SkipVeto,
			VetoMapPool:   strings.Split(matches[i].VetoMapPool, " "),
			Team1Score:    matches[i].Team1Score,
			Team2Score:    matches[i].Team2Score,
			Team1String:   matches[i].Team1String,
			Team2String:   matches[i].Team2String,
			Forfeit:       matches[i].Forfeit,
			Pending:       matches[i].Pending(),
			Live:          matches[i].Live(),
			Server:        server,
			MapStats:      mapstats,
			Team1:         team1,
			Team2:         team2,
			User:          user,
			Status:        status,
		}
		m = append(m, matchdata)
	}
	response.Matches = m

	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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
		return
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
		return
	}
	w.Write([]byte(status))
}

func GetMatches(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetMatches\n")
	q := r.URL.Query()
	offset := q.Get("offset")
	if offset == "" {
		offset = "0"
	}
	userID := q.Get("userID")
	response := []db.MatchData{}
	if userID != "" {
		db.SQLAccess.Gorm.Limit(20).Where("user_id = ?", userID).Order("id DESC").Offset(offset).Find(&response)
	} else {
		db.SQLAccess.Gorm.Limit(20).Order("id DESC").Offset(offset).Find(&response)
	}
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(jsonbyte))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetTeamInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetTeamInfo\n")
	teamid := vars["teamID"]
	response := APITeamData{}
	db.SQLAccess.Gorm.Where("id = ?", teamid).First(&response)
	steamids, err := response.GetPlayers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for i := 0; i < len(steamids); i++ {
		response.SteamIDs = append(response.SteamIDs, steamids[i])
	}
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetRecentMatches(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetRecentMatches\n")
	teamID := vars["teamID"]
	matches := []db.MatchData{}
	response := []APIMatchData{}
	db.SQLAccess.Gorm.Where("team1_id = ?", teamID).Or("team2_id = ?", teamID).Limit(20).Order("id DESC").Find(&matches)
	for i := 0; i < len(matches); i++ {
		mapstats := []APIMapStatsData{}
		server := APIGameServerData{}
		team1 := APITeamData{}
		team2 := APITeamData{}
		user := APIUserData{}
		db.SQLAccess.Gorm.Where("match_id = ?", matches[i].ID).Limit(7).Find(&mapstats)
		db.SQLAccess.Gorm.Where("id = ?", matches[i].ServerID).First(&server)
		db.SQLAccess.Gorm.Where("id = ?", matches[i].Team1ID).First(&team1)
		db.SQLAccess.Gorm.Where("id = ?", matches[i].Team2ID).First(&team2)
		db.SQLAccess.Gorm.Where("id = ?", matches[i].UserID).First(&user)
		var winner int64
		if matches[i].Winner.Valid {
			winner_, err := matches[i].Winner.Value()
			if err != nil {
				winner = 0
			}
			winner = winner_.(int64)
		}
		starttime := matches[i].StartTime.Time
		endtime := matches[i].EndTime.Time
		status, err := matches[i].GetStatusString(true)
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
		m := APIMatchData{
			ID:            matches[i].ID,
			UserID:        matches[i].UserID,
			Winner:        winner,
			Cancelled:     matches[i].Cancelled,
			StartTimeJSON: starttime,
			EndTimeJSON:   endtime,
			MaxMaps:       matches[i].MaxMaps,
			Title:         matches[i].Title,
			SkipVeto:      matches[i].SkipVeto,
			VetoMapPool:   strings.Split(matches[i].VetoMapPool, " "),
			Team1Score:    matches[i].Team1Score,
			Team2Score:    matches[i].Team2Score,
			Team1String:   matches[i].Team1String,
			Team2String:   matches[i].Team2String,
			Forfeit:       matches[i].Forfeit,
			Pending:       matches[i].Pending(),
			Live:          matches[i].Live(),
			Server:        server,
			MapStats:      mapstats,
			Team1:         team1,
			Team2:         team2,
			User:          user,
			Status:        status,
		}
		response = append(response, m)
	}

	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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
		return
	}
	res.Response = strconv.FormatBool(team.CanEdit(userid))
	jsonbyte, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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
		return
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
		return
	}
	steamname, err := db.GetSteamName(uint64(steamid64))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(steamname))
}

// GetTeamList Returns registered team list in JSON.
func GetTeamList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetTeamList\n")
	Teams := []db.TeamData{}
	s := db.Sess.Start(w, r)
	var IsLoggedIn bool
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
	}
	if IsLoggedIn {
		userid := s.Get("UserID").(int)
		db.SQLAccess.Gorm.Where("public_team = 1").Or("user_id = ?", userid).Find(&Teams)
	} else {
		db.SQLAccess.Gorm.Where("public_team = 1").Find(&Teams)
	}
	for i := 0; i < len(Teams); i++ {
		Teams[i].GetPlayers()
	}
	jsonbyte, err := json.Marshal(Teams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

// GetServerList Returns registered public server and owned list in JSON
func GetServerList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetServerList\n")
	s := db.Sess.Start(w, r)
	var IsLoggedIn bool
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
	}
	Servers := []APIGameServerData{}
	if IsLoggedIn {
		userid := s.Get("UserID").(int)
		db.SQLAccess.Gorm.Where("public_server = true AND in_use = false").Or("user_id = ? AND in_use = false", userid).Find(&Servers)
	} else {
		db.SQLAccess.Gorm.Where("public_server = true AND in_use = false").Find(&Servers)
	}
	jsonbyte, err := json.Marshal(Servers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

// CreateTeam Registers team info to DB
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CreateTeam\n")
	var IsLoggedIn = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
	}
	if IsLoggedIn {
		userid := s.Get("UserID").(int)
		Team := db.TeamData{}
		TeamTemp := db.TeamData{}
		err := json.NewDecoder(r.Body).Decode(&TeamTemp)
		if err != nil {
			fmt.Println("Failed to decode JSON")
			http.Error(w, "JSON Format invalid", http.StatusBadRequest)
			return
		}
		_, err = Team.Create(userid, TeamTemp.Name, TeamTemp.Tag, TeamTemp.Flag, TeamTemp.Logo, TeamTemp.Auths, TeamTemp.PublicTeam)
		if err != nil {
			fmt.Println("Failed to create team")
			http.Error(w, "Failed to create team", http.StatusInternalServerError)
			return
		}
		res := SimpleJSONResponse{
			Response: "OK",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	} else {
		res := SimpleJSONResponse{
			Errorcode:    http.StatusUnauthorized,
			Errormessage: "Forbidden",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	}
}

// EditTeam Edits team information
func EditTeam(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("EditTeam\n")
	var IsLoggedIn = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
	}
	if IsLoggedIn {
		userid := s.Get("UserID").(int)
		vars := mux.Vars(r)
		teamid, err := strconv.Atoi(vars["teamID"])
		if err != nil {
			fmt.Println(err)
			fmt.Println("teamid should be int")
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusBadRequest,
				Errormessage: "teamid should be int.",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonbyte)
			return
		}
		Team := db.TeamData{ID: teamid}
		db.SQLAccess.Gorm.First(&Team)
		if !Team.CanEdit(userid) {
			fmt.Println("You do not have permission to edit this team.")
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusUnauthorized,
				Errormessage: "You dont have permission to edit this team.",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&Team)
		if err != nil {
			fmt.Println("Failed to decode JSON")
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusBadRequest,
				Errormessage: "JSON Format invalid",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}
		Team.ID = teamid
		Team.UserID = userid

		_, err = Team.Edit()
		if err != nil {
			fmt.Printf("Failed to edit team %v\n", teamid)
			http.Error(w, "", http.StatusInternalServerError)
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusInternalServerError,
				Errormessage: "Failed to edit team",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}
		w.WriteHeader(http.StatusOK)
		res := SimpleJSONResponse{
			Response: "OK",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		res := SimpleJSONResponse{
			Errorcode:    http.StatusUnauthorized,
			Errormessage: "Forbidden",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	}
}

// DeleteTeam Deletes team // TODO
func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DeleteTeam\n")
	var IsLoggedIn = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
	}
	if IsLoggedIn {
		userid := s.Get("UserID").(int)
		vars := mux.Vars(r)
		teamID, err := strconv.Atoi(vars["teamID"])
		if err != nil {
			fmt.Println("teamID should be int")
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusBadRequest,
				Errormessage: "teamID should be int.",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}
		Team := db.TeamData{ID: teamID}

		if !Team.CanDelete(userid) {
			fmt.Println("You dont have permission to delete this team.")
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusUnauthorized,
				Errormessage: "You dont have permission to delete this team.",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}

		err = Team.Delete()
		if err != nil {
			fmt.Printf("Failed to delete team %v\n", teamID)
			http.Error(w, "", http.StatusInternalServerError)
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusInternalServerError,
				Errormessage: "Failed to delete team",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}
		w.WriteHeader(http.StatusOK)
		res := SimpleJSONResponse{
			Response: "OK",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		res := SimpleJSONResponse{
			Errorcode:    http.StatusUnauthorized,
			Errormessage: "Forbidden",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	}
}

// CreateServer Register server to DB
func CreateServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CreateServer\n")
	var IsLoggedIn = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
	}
	if IsLoggedIn {
		userid := s.Get("UserID").(int)
		Server := db.GameServerData{}
		ServerTemp := db.GameServerData{}
		err := json.NewDecoder(r.Body).Decode(&ServerTemp)
		if err != nil {
			fmt.Println("Failed to decode JSON")
			http.Error(w, "JSON Format invalid", http.StatusBadRequest)
			return
		}
		_, err = Server.Create(userid, ServerTemp.DisplayName, ServerTemp.IPString, ServerTemp.Port, ServerTemp.RconPassword, ServerTemp.PublicServer)
		if err != nil {
			http.Error(w, "Failed to create server", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		res := SimpleJSONResponse{
			Response: "OK",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		res := SimpleJSONResponse{
			Errorcode:    http.StatusUnauthorized,
			Errormessage: "Forbidden",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	}
}

// EditServer Edits Server information // TODO
func EditServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("EditServer\n")
	var IsLoggedIn = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
	}
	if IsLoggedIn {
		userid := s.Get("UserID").(int)
		vars := mux.Vars(r)
		serverID, err := strconv.Atoi(vars["serverID"])
		if err != nil {
			fmt.Println(err)
			fmt.Println("serverID should be int")
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusBadRequest,
				Errormessage: "serverID should be int.",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonbyte)
			return
		}
		Server := db.GameServerData{ID: serverID}
		if !Server.CanEdit(userid) {
			fmt.Println("You do not have permission to edit this server.")
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusUnauthorized,
				Errormessage: "You dont have permission to edit this server.",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&Server)
		if err != nil {
			fmt.Println("Failed to decode JSON")
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusBadRequest,
				Errormessage: "JSON Format invalid",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}
		Server.ID = serverID
		Server.UserID = userid

		_, err = Server.Edit()
		if err != nil {
			fmt.Printf("Failed to edit server %v\n", serverID)
			http.Error(w, "", http.StatusInternalServerError)
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusInternalServerError,
				Errormessage: "Failed to edit server",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}
		w.WriteHeader(http.StatusOK)
		res := SimpleJSONResponse{
			Response: "OK",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		res := SimpleJSONResponse{
			Errorcode:    http.StatusUnauthorized,
			Errormessage: "Forbidden",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	}
}

// DeleteServer Deletes Server information // TODO
func DeleteServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DeleteTeam\n")
	var IsLoggedIn = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
	}
	if IsLoggedIn {
		userid := s.Get("UserID").(int)
		vars := mux.Vars(r)
		serverID, err := strconv.Atoi(vars["serverID"])
		if err != nil {
			fmt.Println("serverID should be int")
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusBadRequest,
				Errormessage: "serverID should be int.",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}
		Server := db.GameServerData{ID: serverID}

		if !Server.CanDelete(userid) {
			fmt.Println("You dont have permission to delete this server.")
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusUnauthorized,
				Errormessage: "You dont have permission to delete this server.",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}

		err = Server.Delete()
		if err != nil {
			fmt.Printf("Failed to delete server %v\n", serverID)
			http.Error(w, "", http.StatusInternalServerError)
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusInternalServerError,
				Errormessage: "Failed to delete server",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}
		w.WriteHeader(http.StatusOK)
		res := SimpleJSONResponse{
			Response: "OK",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		res := SimpleJSONResponse{
			Errorcode:    http.StatusUnauthorized,
			Errormessage: "Forbidden",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	}
}

// CreateMatch Registers match info
func CreateMatch(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CreateMatch\n")
	var IsLoggedIn = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
	}
	if IsLoggedIn {
		userid := s.Get("UserID").(int)
		var MatchTemp = db.MatchData{}
		var Match = db.MatchData{}
		// Returns error if JSON is not valid...
		err := json.NewDecoder(r.Body).Decode(&MatchTemp)
		if err != nil {
			fmt.Println("failed to decode Match, Format incorrect")
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    http.StatusBadRequest,
				Errormessage: "JSON Format invalid",
			}
			jsonbyte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonbyte)
			return
		}
		_, err = Match.Create(userid, MatchTemp.Team1ID, MatchTemp.Team2ID, MatchTemp.Team1String, MatchTemp.Team2String, MatchTemp.MaxMaps, MatchTemp.SkipVeto, MatchTemp.Title, MatchTemp.VetoMapPoolJSON, MatchTemp.ServerID)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		res := SimpleJSONResponse{
			Response: "ok",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		res := SimpleJSONResponse{
			Errorcode:    http.StatusUnauthorized,
			Errormessage: "Forbidden",
		}
		jsonbyte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	}
}

// MatchCancelHandler Handler for /api/v1/match/{matchID}/cancel API.
func MatchCancelHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchCancelHandler")
	var IsLoggedIn = false
	var IsAdmin = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
		IsAdmin = s.Get("Admin").(bool)
	}
	if IsLoggedIn && IsAdmin {
		vars := mux.Vars(r)
		matchid, err := strconv.Atoi(vars["matchID"])
		if err != nil {
			http.Error(w, "matchid should be int", http.StatusBadRequest)
			return
		}
		Match := db.MatchData{}
		rec := db.SQLAccess.Gorm.Where("id = ?", matchid).First(&Match)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find match", http.StatusNotFound)
		}
		MatchUpdate := Match
		db.SQLAccess.Gorm.First(&MatchUpdate)
		MatchUpdate.Cancelled = true
		db.SQLAccess.Gorm.Model(&Match).Update(&MatchUpdate)
		db.SQLAccess.Gorm.Save(&MatchUpdate)

		Server := db.GameServerData{}
		db.SQLAccess.Gorm.Where("id = ?", Match.ServerID).First(&Server)
		ServerUpdate := Server
		db.SQLAccess.Gorm.First(&ServerUpdate)
		ServerUpdate.InUse = false
		db.SQLAccess.Gorm.Model(&Server).Update(&ServerUpdate)
		db.SQLAccess.Gorm.Save(&ServerUpdate)

		_, err = Server.SendRcon("get5_endmatch")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to cancel match: %v", err), http.StatusNotFound)
		}
	} else {
		http.Error(w, "Please log in", http.StatusUnauthorized)
	}
}

// MatchRconHandler Handler for /api/v1/match/{matchID}/rcon API.
func MatchRconHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchRconHandler")
	var IsLoggedIn = false
	var IsAdmin = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
		IsAdmin = s.Get("Admin").(bool)
	}
	if IsLoggedIn && IsAdmin {
		vars := mux.Vars(r)
		matchid, err := strconv.Atoi(vars["matchID"])
		if err != nil {
			http.Error(w, "matchid should be int", http.StatusBadRequest)
			return
		}
		q := r.URL.Query()
		command := q.Get("command")
		Match := db.MatchData{}
		rec := db.SQLAccess.Gorm.Where("id = ?", matchid).First(&Match)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find match", http.StatusNotFound)
		}

		Server := db.GameServerData{}
		rec = db.SQLAccess.Gorm.Where("id = ?", Match.ServerID).First(&Server)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find server", http.StatusNotFound)
		}
		RconRes := "No output"
		res, err := Server.SendRcon(command)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send command : %s", err), http.StatusInternalServerError)
		}
		if res != "" {
			RconRes = res
		}
		JSONres := SimpleJSONResponse{
			Response: RconRes,
		}
		jsonbyte, err := json.Marshal(JSONres)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	} else {
		http.Error(w, "Please log in", http.StatusUnauthorized)
	}
}

// MatchPauseHandler Handler for /api/v1/match/{matchID}/pause API.
func MatchPauseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchPauseHandler")
	var IsLoggedIn = false
	var IsAdmin = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
		IsAdmin = s.Get("Admin").(bool)
	}
	if IsLoggedIn && IsAdmin {
		vars := mux.Vars(r)
		matchid, err := strconv.Atoi(vars["matchID"])
		if err != nil {
			http.Error(w, "matchid should be int", http.StatusBadRequest)
			return
		}
		Match := db.MatchData{}
		rec := db.SQLAccess.Gorm.Where("id = ?", matchid).First(&Match)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find match", http.StatusNotFound)
		}

		Server := db.GameServerData{}
		rec = db.SQLAccess.Gorm.Where("id = ?", Match.ServerID).First(&Server)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find server", http.StatusNotFound)
		}
		_, err = Server.SendRcon("sm_pause")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send pause command : %s", err), http.StatusInternalServerError)
		}
		JSONres := SimpleJSONResponse{
			Response: "ok",
		}
		jsonbyte, err := json.Marshal(JSONres)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	} else {
		http.Error(w, "Please log in", http.StatusUnauthorized)
	}
}

// MatchUnpauseHandler Handler for /api/v1/match/{matchID}/unpause API.
func MatchUnpauseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchUnpauseHandler")
	var IsLoggedIn = false
	var IsAdmin = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
		IsAdmin = s.Get("Admin").(bool)
	}
	if IsLoggedIn && IsAdmin {
		vars := mux.Vars(r)
		matchid, err := strconv.Atoi(vars["matchID"])
		if err != nil {
			http.Error(w, "matchid should be int", http.StatusBadRequest)
			return
		}
		Match := db.MatchData{}
		rec := db.SQLAccess.Gorm.Where("id = ?", matchid).First(&Match)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find match", http.StatusNotFound)
		}

		Server := db.GameServerData{}
		rec = db.SQLAccess.Gorm.Where("id = ?", Match.ServerID).First(&Server)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find server", http.StatusNotFound)
		}
		_, err = Server.SendRcon("sm_unpause")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send unpause command : %s", err), http.StatusInternalServerError)
		}
		JSONres := SimpleJSONResponse{
			Response: "ok",
		}
		jsonbyte, err := json.Marshal(JSONres)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	} else {
		http.Error(w, "Please log in", http.StatusUnauthorized)
	}
}

// MatchAddUserHandler Handler for /api/v1/match/{matchID}/adduser API.
func MatchAddUserHandler(w http.ResponseWriter, r *http.Request) { // TODO
	fmt.Println("MatchAddUserHandler")
	var IsLoggedIn = false
	var IsAdmin = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
		IsAdmin = s.Get("Admin").(bool)
	}
	if IsLoggedIn && IsAdmin {
		q := r.URL.Query()
		team := q.Get("team")
		if team != "team1" && team != "team2" && team != "spec" {
			http.Error(w, "No team specified", http.StatusBadRequest)
			return
		}
		auth := q.Get("auth")
		newauth, err := util.AuthToSteamID64(auth)
		if err != nil {
			http.Error(w, "Auth format invalid.", http.StatusBadRequest)
			return
		}
		fmt.Printf("auth : %s", newauth)
		vars := mux.Vars(r)
		matchid, err := strconv.Atoi(vars["matchID"])
		if err != nil {
			http.Error(w, "matchID should be int", http.StatusBadRequest)
			return
		}
		Match := db.MatchData{}
		rec := db.SQLAccess.Gorm.Where("id = ?", matchid).First(&Match)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find match", http.StatusNotFound)
			return
		}

		Server := db.GameServerData{}
		rec = db.SQLAccess.Gorm.Where("id = ?", Match.ServerID).First(&Server)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find server", http.StatusNotFound)
			return
		}
		res, err := Server.SendRcon(fmt.Sprintf("get5_addplayer %s %s", newauth, team))
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send command : %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, res)
	} else {
		http.Error(w, "Please log in", http.StatusUnauthorized)
	}
}
