package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FlowingSPDG/get5-web-go/server/src/api"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	"github.com/FlowingSPDG/get5-web-go/server/src/get5"
	"github.com/go-ini/ini"
	"github.com/gorilla/mux"
	_ "github.com/hydrogen18/stalecucumber"
	_ "github.com/sipin/gorazor/gorazor"
	_ "github.com/solovev/steam_go"
)

// Config Configration Struct for config.ini
type Config struct {
	SteamAPIKey string
	DefaultPage string
	SQLHost     string
	SQLUser     string
	SQLPass     string
	SQLPort     int
	SQLDBName   string
	HOST        string
}

var (
	// StaticDir Directly where serves static files
	StaticDir = "./static"
	// HOST Server's domain name
	HOST string
	//Cnf Configration Data
	Cnf Config
	// SQLAccess SQL Access Object for MySQL and GORM things
	SQLAccess db.DBdatas
)

func init() {
	c, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}
	Cnf = Config{
		HOST:      c.Section("GET5").Key("HOST").MustString("localhost:8080"),
		SQLHost:   c.Section("sql").Key("host").MustString(""),
		SQLUser:   c.Section("sql").Key("user").MustString(""),
		SQLPass:   c.Section("sql").Key("pass").MustString(""),
		SQLPort:   c.Section("sql").Key("port").MustInt(3306),
		SQLDBName: c.Section("sql").Key("database").MustString(""),
	}
	HOST = Cnf.HOST
	fmt.Println(db.SQLAccess)
}

func main() {

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//s := r.Host(HOST).Subrouter()

	r.HandleFunc("/api/v1/CheckLoggedIn", api.CheckLoggedIn).Methods("GET")
	r.HandleFunc("/api/v1/match/{matchID}/GetMatchInfo", api.GetMatchInfo).Methods("GET")
	r.HandleFunc("/api/v1/match/{matchID}/GetStatusString", api.GetStatusString).Methods("GET")
	r.HandleFunc("/api/v1/GetMatches", api.GetMatches).Methods("GET")
	r.HandleFunc("/api/v1/GetMetrics", api.GetMetrics).Methods("GET")
	r.HandleFunc("/api/v1/GetSteamName", api.GetSteamName).Methods("GET")
	r.HandleFunc("/api/v1/team/{teamID}/GetTeamInfo", api.GetTeamInfo).Methods("GET")
	r.HandleFunc("/api/v1/team/{teamID}/CheckUserCanEdit", api.CheckUserCanEdit).Methods("GET")
	r.HandleFunc("/api/v1/user/{userID}/GetUserInfo", api.GetTeamInfo).Methods("GET")
	r.HandleFunc("/api/v1/server/{serverID}/GetServerInfo", api.GetServerInfo).Methods("GET")

	r.HandleFunc("/", db.HomeHandler).Methods("GET")
	r.HandleFunc("/login", db.LoginHandler).Methods("GET")
	r.HandleFunc("/logout", db.LogoutHandler).Methods("GET")

	r.HandleFunc("/match/create", get5.MatchCreateHandler)             // GET/POST
	r.HandleFunc("/match/{matchID}", get5.MatchHandler)                // ?
	r.HandleFunc("/match/{matchID}/config", get5.MatchConfigHandler)   // ?
	r.HandleFunc("/match/{matchID}/cancel", get5.MatchCancelHandler)   // ?
	r.HandleFunc("/match/{matchID}/rcon", get5.MatchRconHandler)       // ?
	r.HandleFunc("/match/{matchID}/pause", get5.MatchPauseHandler)     // ?
	r.HandleFunc("/match/{matchID}/unpause", get5.MatchUnpauseHandler) // ?
	r.HandleFunc("/match/{matchID}/adduser", get5.MatchAddUserHandler) // ?
	//r.HandleFunc("/match/{matchID}/sendconfig", get5.MatchSendConfigHandler) // ?
	r.HandleFunc("/match/{matchID}/backup", get5.MatchBackupHandler).Methods("GET") // GET

	r.HandleFunc("/match/{matchID}/finish", get5.MatchFinishHandler).Methods("POST")                                             // POST
	r.HandleFunc("/match/{matchID}/map/{mapNumber}/start", get5.MatchMapStartHandler).Methods("POST")                            // POST
	r.HandleFunc("/match/{matchID}/map/{mapNumber}/update", get5.MatchMapUpdateHandler).Methods("POST")                          // POST
	r.HandleFunc("/match/{matchID}/map/{mapNumber}/finish", get5.MatchMapFinishHandler).Methods("POST")                          // POST
	r.HandleFunc("/match/{matchID}/map/{mapNumber}/player/{steamid64}/update", get5.MatchMapPlayerUpdateHandler).Methods("POST") // POST

	r.HandleFunc("/matches", get5.MatchesHandler)                // ?
	r.HandleFunc("/matches/{userID}", get5.MatchesWithIDHandler) // ?
	r.HandleFunc("/mymatches", get5.MyMatchesHandler)            // ?

	r.HandleFunc("/team/create", get5.TeamCreateHandler)              // GET/POST
	r.HandleFunc("/team/{teamID}", get5.TeamHandler).Methods("GET")   // GET
	r.HandleFunc("/team/{teamID}/edit", get5.TeamEditHandler)         // GET/POST
	r.HandleFunc("/team/{teamID}/delete", get5.TeamDeleteHandler)     // ?
	r.HandleFunc("/teams/{userID}", get5.TeamsHandler).Methods("GET") // GET
	r.HandleFunc("/myteams", get5.MyTeamsHandler).Methods("GET")      // GET

	r.HandleFunc("/server/create", get5.ServerCreateHandler)                           // GET/POST
	r.HandleFunc("/server/{serverid}/edit", get5.ServerEditHandler)                    // GET/POST
	r.HandleFunc("/server/{serverid}/delete", get5.ServerDeleteHandler).Methods("GET") // GET
	r.HandleFunc("/myservers", get5.MyServersHandler)                                  // ?

	r.HandleFunc("/user/{userID}", get5.UserHandler)

	r.HandleFunc("/metrics", get5.MetricsHandler)

	r.Methods("GET", "POST")
	http.Handle("/", r)
	fmt.Println("RUNNING")
	log.Fatal(http.ListenAndServe(HOST, nil))
}
