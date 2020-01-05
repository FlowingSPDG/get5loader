package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FlowingSPDG/get5-web-go/server/src/api"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	"github.com/go-ini/ini"
	"github.com/gorilla/mux"
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

const (
	// VERSION get5-web-go Version
	VERSION = "v0.1.0"
)

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

// ServeStaticFile Host static files
func ServeStaticFile(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}

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

	defer SQLAccess.Gorm.Close()

	r := mux.NewRouter()
	entrypoint := "./static/index.html"
	r.Path("/").HandlerFunc(ServeStaticFile(entrypoint))
	r.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	r.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	r.PathPrefix("/img").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	r.PathPrefix("/fonts").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts"))))

	//s := r.Host(HOST).Subrouter() // in-case if we need vhost thing

	// misc
	r.HandleFunc("/api/v1/GetMatches", api.GetMatches).Methods("GET")
	r.HandleFunc("/api/v1/GetMetrics", api.GetMetrics).Methods("GET")
	r.HandleFunc("/api/v1/GetSteamName", api.GetSteamName).Methods("GET")
	r.HandleFunc("/api/v1/GetTeamList", api.GetTeamList).Methods("GET")
	r.HandleFunc("/api/v1/GetServerList", api.GetServerList).Methods("GET")

	// API for front(Vue)
	r.HandleFunc("/api/v1/CheckLoggedIn", api.CheckLoggedIn).Methods("GET")
	r.HandleFunc("/api/v1/match/{matchID}/GetMatchInfo", api.GetMatchInfo).Methods("GET")
	r.HandleFunc("/api/v1/match/{matchID}/GetPlayerStatInfo", api.GetPlayerStatInfo).Methods("GET")
	r.HandleFunc("/api/v1/match/{matchID}/GetStatusString", api.GetStatusString).Methods("GET")
	r.HandleFunc("/api/v1/match/create", api.CreateMatch).Methods("POST")
	r.HandleFunc("/api/v1/team/{teamID}/GetTeamInfo", api.GetTeamInfo).Methods("GET")
	r.HandleFunc("/api/v1/team/{teamID}/GetRecentMatches", api.GetRecentMatches).Methods("GET")
	r.HandleFunc("/api/v1/team/{teamID}/CheckUserCanEdit", api.CheckUserCanEdit).Methods("GET")
	r.HandleFunc("/api/v1/team/create", api.CreateTeam).Methods("POST")
	r.HandleFunc("/api/v1/team/{teamID}/edit", api.EditTeam).Methods("POST")
	r.HandleFunc("/api/v1/team/{teamID}/delete", api.DeleteTeam).Methods("DELETE")
	r.HandleFunc("/api/v1/user/{userID}/GetUserInfo", api.GetUserInfo).Methods("GET")
	r.HandleFunc("/api/v1/server/{serverID}/GetServerInfo", api.GetServerInfo).Methods("GET")
	r.HandleFunc("/api/v1/server/create", api.CreateServer).Methods("POST")
	r.HandleFunc("/api/v1/server/{serverID}/edit", api.EditServer).Methods("POST")
	r.HandleFunc("/api/v1/server/{serverID}/delete", api.DeleteServer).Methods("DELETE")

	// GET5 API
	r.HandleFunc("/api/v1/match/{matchID}/config", api.MatchConfigHandler)
	r.HandleFunc("/api/v1/match/{matchID}/finish", api.MatchFinishHandler).Methods("POST")
	r.HandleFunc("/api/v1/match/{matchID}/map/{mapNumber}/start", api.MatchMapStartHandler).Methods("POST")
	r.HandleFunc("/api/v1/match/{matchID}/map/{mapNumber}/update", api.MatchMapUpdateHandler).Methods("POST")
	r.HandleFunc("/api/v1/match/{matchID}/map/{mapNumber}/finish", api.MatchMapFinishHandler).Methods("POST")
	r.HandleFunc("/api/v1/match/{matchID}/map/{mapNumber}/player/{steamid64}/update", api.MatchMapPlayerUpdateHandler).Methods("POST")
	r.HandleFunc("/api/v1/match/{matchID}/cancel", api.MatchCancelHandler).Methods("POST")
	//r.HandleFunc("/api/v1/match/{matchID}/vetoUpdate", api.MatchVetoUpdateHandler).Methods("POST") // TODO
	//r.HandleFunc("/api/v1/match/{matchID}/map/{mapNumber}/demo", api.MatchDemoUploadHandler).Methods("POST") // TODO

	// session handling
	r.HandleFunc("/login", db.LoginHandler).Methods("GET")
	r.HandleFunc("/logout", db.LogoutHandler).Methods("GET")

	r.HandleFunc("/api/v1/login", db.LoginHandler).Methods("GET")
	r.HandleFunc("/api/v1/logout", db.LogoutHandler).Methods("GET")
	/*
		r.HandleFunc("/api/v1/match/{matchID}/rcon", api.MatchRconHandler)       // ?
		r.HandleFunc("/api/v1/match/{matchID}/pause", api.MatchPauseHandler)     // ?
		r.HandleFunc("/api/v1/match/{matchID}/unpause", api.MatchUnpauseHandler) // ?
		r.HandleFunc("/api/v1/match/{matchID}/adduser", api.MatchAddUserHandler) // ?
		//r.HandleFunc("/api/v1/match/{matchID}/sendconfig", api.MatchSendConfigHandler) // ?
		r.HandleFunc("/api/v1/match/{matchID}/backup", api.MatchBackupHandler).Methods("GET") // GET
	*/

	r.Methods("GET", "POST", "DELETE")
	http.Handle("/", r)
	fmt.Printf("RUNNING at %v\n", HOST)
	log.Fatal(http.ListenAndServe(HOST, nil))
}
