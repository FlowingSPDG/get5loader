package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	"github.com/FlowingSPDG/get5-web-go/server/src/util"
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
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		response.IsLoggedIn = s.Get("Loggedin").(bool)
		response.SteamID = s.Get("SteamID").(string)
		response.UserID = s.Get("UserID").(int)
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

// GetTeamList Returns registered team list in JSON.
func GetTeamList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetTeamList\n")
	Teams := []db.TeamData{}
	db.SQLAccess.Gorm.Where("public_team = 1").Find(&Teams)
	for i := 0; i < len(Teams); i++ {
		Teams[i].GetPlayers()
	}
	jsonbyte, err := json.Marshal(Teams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

// GetServerList Returns registered public server list in JSON
func GetServerList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetServerList\n")
	Servers := []APIGameServerData{}
	db.SQLAccess.Gorm.Where("public_server = 1").Find(&Servers)
	jsonbyte, err := json.Marshal(Servers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CreateTeam\n")
	var IsLoggedIn = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
	}
	if IsLoggedIn {
		Team := db.TeamData{}
		TeamTemp := db.TeamData{}
		err := json.NewDecoder(r.Body).Decode(&TeamTemp)
		if err != nil {
			fmt.Println("failed to decode JSON")
			http.Error(w, "JSON Format invalid", http.StatusBadRequest)
			return
		}
		if TeamTemp.Name == "" {
			fmt.Println("failed to decode Team Name")
			http.Error(w, "JSON Format invalid", http.StatusBadRequest)
			return
		}
		Team.UserID = s.Get("UserID").(int)
		Team.Name = TeamTemp.Name
		Team.Tag = TeamTemp.Tag
		Team.Flag = TeamTemp.Flag
		Team.Logo = TeamTemp.Logo
		Team.AuthsPickle, err = util.SteamID64sToPickle(TeamTemp.Auths)
		if err != nil {
			http.Error(w, "Internal ERROR", http.StatusInternalServerError)
			return
		}
		Team.PublicTeam = TeamTemp.PublicTeam
		db.SQLAccess.Gorm.Create(&Team)
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
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonbyte)
	}
}

func CreateServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CreateServer\n")
	var IsLoggedIn = false
	s := db.Sess.Start(w, r)
	if s.Get("Loggedin") != nil {
		IsLoggedIn = s.Get("Loggedin").(bool)
	}
	if IsLoggedIn {
		Server := db.GameServerData{}
		ServerTemp := db.GameServerData{}
		err := json.NewDecoder(r.Body).Decode(&ServerTemp)
		if err != nil {
			fmt.Println("failed to decode JSON")
			http.Error(w, "JSON Format invalid", http.StatusBadRequest)
			return
		}
		if ServerTemp.DisplayName == "" || ServerTemp.IPString == "" || ServerTemp.RconPassword == "" {
			fmt.Println("failed to decode Server Name")
			http.Error(w, "JSON Format invalid", http.StatusBadRequest)
			return
		}
		Server.UserID = s.Get("UserID").(int)
		Server.DisplayName = ServerTemp.DisplayName
		Server.IPString = ServerTemp.IPString
		Server.Port = ServerTemp.Port
		Server.RconPassword = ServerTemp.RconPassword
		Server.PublicServer = ServerTemp.PublicServer
		_, err = util.CheckServerAvailability(Server.IPString, Server.Port, Server.RconPassword)
		if err != nil {
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    500,
				Errormessage: err.Error(),
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
			db.SQLAccess.Gorm.Create(&Server)
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
		}

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
		User := db.UserData{
			ID: userid,
		}
		db.SQLAccess.Gorm.First(&User)
		Match := db.MatchData{}
		MatchTemp := db.MatchData{}

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
		if MatchTemp.Team1ID == 0 || MatchTemp.Team2ID == 0 || MatchTemp.ServerID == 0 {
			fmt.Println("failed to decode Match,lack of informations")
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
		server := db.GameServerData{}
		server.ID = MatchTemp.ServerID
		db.SQLAccess.Gorm.First(&server)
		// returns error if user wasnt owned server,or not an admin.
		if userid != MatchTemp.ServerID && !User.Admin && !server.PublicServer {
			http.Error(w, "This is not your server!", http.StatusUnauthorized)
			return
		}
		get5res, err := util.CheckServerAvailability(server.IPString, server.Port, server.RconPassword) // Returns error if SRCDS is not available
		if err != nil {
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    500,
				Errormessage: err.Error(),
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

		MatchOnServer := db.MatchData{
			ServerID:  MatchTemp.ServerID,
			Cancelled: false,
		}
		db.SQLAccess.Gorm.Where("EndTime = ?", "NULL").First(&MatchOnServer)
		if MatchOnServer.ID != 0 {
			res := SimpleJSONResponse{
				Response:     "error",
				Errorcode:    500,
				Errormessage: fmt.Sprintf("Match %v is already using this server", MatchOnServer.ID),
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

		Match.UserID = userid
		Match.ServerID = MatchTemp.ServerID
		Match.GetServer()
		Match.Team1ID = MatchTemp.Team1ID
		Match.Team2ID = MatchTemp.Team2ID
		Match.MaxMaps = MatchTemp.MaxMaps
		Match.Title = MatchTemp.Title
		Match.SkipVeto = MatchTemp.SkipVeto
		Match.VetoMapPool = strings.Join(MatchTemp.VetoMapPoolJSON, " ")
		Match.Team1String = MatchTemp.Team1String
		Match.Team2String = MatchTemp.Team2String
		if get5res.PluginVersion == "" {
			get5res.PluginVersion = "unknown"
		}
		Match.PluginVersion = get5res.PluginVersion
		err = Match.SendToServer()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal SRCDS ERROR", http.StatusInternalServerError)
			return
		}
		db.SQLAccess.Gorm.Model(&Match.Server).Update("in_use", true)
		db.SQLAccess.Gorm.Create(&Match)

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
