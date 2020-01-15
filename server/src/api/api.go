package api

import (
	"encoding/json"
	"fmt"
	"github.com/FlowingSPDG/get5-web-go/server/src/util"
	"github.com/go-ini/ini"

	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/FlowingSPDG/get5-web-go/server/src/db"
)

const (
	// VERSION get5-web-go Version
	VERSION = "0.1.1"
)

var (
	ActiveMapPool  []string
	ReserveMapPool []string
)

func init() {
	c, _ := ini.Load("config.ini")
	active := c.Section("MAPLIST").Key("Active").MustString("de_dust2,de_mirage,de_inferno,de_overpass,de_train,de_nuke,de_vertigo")
	reserve := c.Section("MAPLIST").Key("Reserve").MustString("de_cache,de_season")
	ActiveMapPool = strings.Split(strings.ToLower(strings.TrimSpace(active)), ",")
	ReserveMapPool = strings.Split(strings.ToLower(strings.TrimSpace(reserve)), ",")
}

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
		http.Error(w, "JSON Format invalid", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

// GetMapList Struct type for /api/v1/GetMapList API Response.
type GetMapListJSON struct {
	Active  []string `json:"active"`
	Reserve []string `json:"reserve"`
}

// GetMapList handler for /api/v1/GetMapList API.
func GetMapList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CheckLoggedIn")
	response := GetMapListJSON{
		Active:  ActiveMapPool,
		Reserve: ReserveMapPool,
	}
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "JSON Format invalid", http.StatusBadRequest)
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
		http.Error(w, "JSON Format invalid", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

// GetMatchInfo Gets match info by ID
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
	db.SQLAccess.Gorm.First(&match, matchid)
	db.SQLAccess.Gorm.Where("match_id = ?", matchid).Limit(7).Find(&mapstats)
	db.SQLAccess.Gorm.First(&server, match.ServerID)
	db.SQLAccess.Gorm.First(&team1, match.Team1ID)
	db.SQLAccess.Gorm.First(&team2, match.Team2ID)
	db.SQLAccess.Gorm.First(&user, match.UserID)
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
		http.Error(w, "JSON Format invalid", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
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
		http.Error(w, "JSON Format invalid", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetUserInfo\n")
	userid := vars["userID"]
	response := APIUserData{}
	db.SQLAccess.Gorm.First(&response, userid)
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
		db.SQLAccess.Gorm.First(&server, matches[i].ServerID)
		db.SQLAccess.Gorm.First(&user, matches[i].UserID)
		db.SQLAccess.Gorm.First(&team1, matches[i].Team1ID)
		db.SQLAccess.Gorm.First(&team2, matches[i].Team2ID)
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
		http.Error(w, "JSON Format invalid", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetServerInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetServerInfo\n")
	serverID := vars["serverID"]
	response := db.GameServerData{}
	db.SQLAccess.Gorm.First(&response, serverID)
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "JSON Format invalid", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetStatusString(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetStatusString\n")
	matchid := vars["matchID"]
	response := db.MatchData{}
	db.SQLAccess.Gorm.First(&response, matchid)
	status, err := response.GetStatusString(true)
	if err != nil {
		http.Error(w, "Failed to get status", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
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
		http.Error(w, "JSON Format invalid", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetTeamInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetTeamInfo\n")
	teamid, err := strconv.Atoi(vars["teamID"])
	if err != nil {
		http.Error(w, "teamID should be int.", http.StatusBadRequest)
		return
	}
	response := APITeamData{}
	db.SQLAccess.Gorm.First(&response, teamid)
	var steamids = make([]string, 5)
	steamids, err = response.GetPlayers()
	if err != nil {
		http.Error(w, "Failed to get players.", http.StatusInternalServerError)
		return
	}
	for i := 0; i < len(steamids); i++ {
		response.SteamIDs = append(response.SteamIDs, steamids[i])
	}
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "JSON Format invalid", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func GetRecentMatches(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("GetRecentMatches\n")
	teamID := vars["teamID"]
	matches := []db.MatchData{}
	response := make([]APIMatchData, 0, len(matches))
	db.SQLAccess.Gorm.Where("team1_id = ?", teamID).Or("team2_id = ?", teamID).Limit(20).Order("id DESC").Find(&matches)
	for i := 0; i < len(matches); i++ {
		mapstats := []APIMapStatsData{}
		server := APIGameServerData{}
		team1 := APITeamData{}
		team2 := APITeamData{}
		user := APIUserData{}
		db.SQLAccess.Gorm.Where("match_id = ?", matches[i].ID).Limit(7).Find(&mapstats)
		db.SQLAccess.Gorm.First(&server, matches[i].ServerID)
		db.SQLAccess.Gorm.First(&team1, matches[i].Team1ID)
		db.SQLAccess.Gorm.First(&team2, matches[i].Team2ID)
		db.SQLAccess.Gorm.First(&user, matches[i].UserID)
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
		http.Error(w, "JSON Format invalid", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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
	db.SQLAccess.Gorm.First(&team, teamid)
	userid, err := strconv.Atoi(useridstr)
	if err != nil {
		http.Error(w, "userid should be int", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(strconv.FormatBool(team.CanEdit(userid))))
}

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetMetrics\n")
	Metrics := db.GetMetrics()
	jsonbyte, err := json.Marshal(Metrics)
	if err != nil {
		http.Error(w, "JSON Format Invalid", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

// GetSteamName Get Steam Profile name by SteamWebAPI
func GetSteamName(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetSteamName\n")
	q := r.URL.Query()
	steamid := q.Get("steamID")
	steamid64, err := strconv.Atoi(steamid)
	if err != nil {
		http.Error(w, "steamID should be int", http.StatusBadRequest)
		return
	}
	steamname, err := db.GetSteamName(uint64(steamid64))
	if err != nil {
		http.Error(w, "failed to get steamname", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
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
		http.Error(w, "JSON Format invalid", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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
		http.Error(w, "JSON Format invalid", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
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
			http.Error(w, "JSON Format invalid", http.StatusBadRequest)
			return
		}
		_, err = Team.Create(userid, TeamTemp.Name, TeamTemp.Tag, TeamTemp.Flag, TeamTemp.Logo, TeamTemp.Auths, TeamTemp.PublicTeam)
		if err != nil {
			http.Error(w, "Failed to create team", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
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
			http.Error(w, "teamid should be int.", http.StatusBadRequest)
			return
		}
		Team := db.TeamData{ID: teamid}
		db.SQLAccess.Gorm.First(&Team)
		if !Team.CanEdit(userid) {
			http.Error(w, "You do not have permission to edit this team.", http.StatusUnauthorized)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&Team)
		if err != nil {
			http.Error(w, "JSON Format Invalid", http.StatusBadRequest)
			return
		}
		Team.ID = teamid
		Team.UserID = userid
		fmt.Printf("Team : %v\n", Team)

		_, err = Team.Edit()
		if err != nil {
			http.Error(w, "Failed to edit team", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
	}
}

// DeleteTeam Deletes team
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
			http.Error(w, "teamID should be int.", http.StatusBadRequest)
			return
		}
		Team := db.TeamData{ID: teamID}

		if !Team.CanDelete(userid) {
			http.Error(w, "You dont have permission to delete this team.", http.StatusUnauthorized)
			return
		}

		err = Team.Delete()
		if err != nil {
			http.Error(w, "Failed to delete team", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
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
			http.Error(w, "JSON Format invalid", http.StatusBadRequest)
			return
		}
		_, err = Server.Create(userid, ServerTemp.DisplayName, ServerTemp.IPString, ServerTemp.Port, ServerTemp.RconPassword, ServerTemp.PublicServer)
		if err != nil {
			http.Error(w, "Failed to create server", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))

	} else {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
	}
}

// EditServer Edits Server information
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
			http.Error(w, "serverID should be int", http.StatusBadRequest)
			return
		}
		Server := db.GameServerData{ID: serverID}
		if !Server.CanEdit(userid) {
			http.Error(w, "You do not have permission to edit this server.", http.StatusForbidden)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&Server)
		if err != nil {
			http.Error(w, "JSON Format invalid", http.StatusBadRequest)
			return
		}
		Server.ID = serverID
		Server.UserID = userid

		_, err = Server.Edit()
		if err != nil {
			http.Error(w, "Failed to edit server", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
	}
}

// DeleteServer Deletes Server information
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
			http.Error(w, "serverID should be int", http.StatusBadRequest)
			return
		}
		Server := db.GameServerData{ID: serverID}

		if !Server.CanDelete(userid) {
			http.Error(w, "You dont have permission to delete this server.", http.StatusUnauthorized)
			return
		}

		err = Server.Delete()
		if err != nil {
			http.Error(w, "Failed to delete server", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
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
		err := json.NewDecoder(r.Body).Decode(&MatchTemp)
		if err != nil {
			http.Error(w, "JSON Format invalid", http.StatusBadRequest)
			return
		}
		_, err = Match.Create(userid, MatchTemp.Team1ID, MatchTemp.Team2ID, MatchTemp.Team1String, MatchTemp.Team2String, MatchTemp.MaxMaps, MatchTemp.SkipVeto, MatchTemp.Title, MatchTemp.VetoMapPoolJSON, MatchTemp.ServerID)
		if err != nil {
			http.Error(w, "Failed to create match", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))

	} else {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
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
		rec := db.SQLAccess.Gorm.First(&Match, matchid)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find match", http.StatusNotFound)
		}
		MatchUpdate := Match
		db.SQLAccess.Gorm.First(&MatchUpdate)
		MatchUpdate.Cancelled = true
		db.SQLAccess.Gorm.Model(&Match).Update(&MatchUpdate)
		db.SQLAccess.Gorm.Save(&MatchUpdate)

		Server := db.GameServerData{}
		db.SQLAccess.Gorm.First(&Server, Match.ServerID)
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
		http.Error(w, "Forbidden", http.StatusUnauthorized)
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
		rec := db.SQLAccess.Gorm.First(&Match, matchid)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find match", http.StatusNotFound)
			return
		}

		Server := db.GameServerData{}
		rec = db.SQLAccess.Gorm.First(&Server, Match.ServerID)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find server", http.StatusNotFound)
			return
		}
		RconRes := "No output"
		res, err := Server.SendRcon(command)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send command : %s", err), http.StatusInternalServerError)
			return
		}
		if res != "" {
			RconRes = res
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(RconRes))
	} else {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
		return
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
		rec := db.SQLAccess.Gorm.First(&Match, matchid)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find match", http.StatusNotFound)
		}

		Server := db.GameServerData{}
		rec = db.SQLAccess.Gorm.First(&Server, Match.ServerID)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find server", http.StatusNotFound)
		}
		_, err = Server.SendRcon("sm_pause")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send pause command : %s", err), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
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
		rec := db.SQLAccess.Gorm.First(&Match, matchid)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find match", http.StatusNotFound)
		}

		Server := db.GameServerData{}
		rec = db.SQLAccess.Gorm.First(&Server, Match.ServerID)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find server", http.StatusNotFound)
		}
		_, err = Server.SendRcon("sm_unpause")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send unpause command : %s", err), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
	}
}

// MatchAddUserHandler Handler for /api/v1/match/{matchID}/adduser API.
func MatchAddUserHandler(w http.ResponseWriter, r *http.Request) {
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
		rec := db.SQLAccess.Gorm.First(&Match, matchid)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find match", http.StatusNotFound)
			return
		}

		Server := db.GameServerData{}
		rec = db.SQLAccess.Gorm.First(&Server, Match.ServerID)
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
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(res))
	} else {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
	}
}

// BackupListJSON Struct type for /api/v1/match/{matchID}/backup
type BackupListJSON struct {
	Files []string `json:"files"`
}

// MatchListBackupsHandler Handler for /api/v1/match/{matchID}/backup API(GET).
func MatchListBackupsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchListBackupsHandler")
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
			http.Error(w, "matchID should be int", http.StatusBadRequest)
			return
		}
		Match := db.MatchData{}
		rec := db.SQLAccess.Gorm.First(&Match, matchid)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find match", http.StatusNotFound)
			return
		}

		Server := db.GameServerData{}
		rec = db.SQLAccess.Gorm.First(&Server, Match.ServerID)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find server", http.StatusNotFound)
			return
		}
		res, err := Server.SendRcon(fmt.Sprintf("get5_listbackups %d", matchid))
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send command : %s", err), http.StatusInternalServerError)
			return
		}
		resJSON := BackupListJSON{
			Files: strings.Split(strings.TrimSpace(res), "\n"),
		}
		fmt.Printf("BackupFiles : %v", resJSON)
		jsonbyte, err := json.Marshal(resJSON)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	} else {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
	}
}

// MatchLoadBackupsHandler Handler for /api/v1/match/{matchID}/backup API(POST).
func MatchLoadBackupsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MatchLoadBackupsHandler")
	var IsLoggedIn = false
	var IsAdmin = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
		IsAdmin = s.Get("Admin").(bool)
	}
	if IsLoggedIn && IsAdmin {
		q := r.URL.Query()
		file := q.Get("file")
		if file == "" {
			http.Error(w, "No file specified", http.StatusBadRequest)
			return
		}
		vars := mux.Vars(r)
		matchid, err := strconv.Atoi(vars["matchID"])
		if err != nil {
			http.Error(w, "matchID should be int", http.StatusBadRequest)
			return
		}
		Match := db.MatchData{}
		rec := db.SQLAccess.Gorm.First(&Match, matchid)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find match", http.StatusNotFound)
			return
		}

		Server := db.GameServerData{}
		rec = db.SQLAccess.Gorm.First(&Server, Match.ServerID)
		if rec.RecordNotFound() {
			http.Error(w, "Failed to find server", http.StatusNotFound)
			return
		}
		res, err := Server.SendRcon(fmt.Sprintf("get5_loadbackup %s", file))
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to load backup : %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(res))
	} else {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
	}
}
