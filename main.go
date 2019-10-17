package main

import (
	"./get5"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	STATIC_DIR = "./static"
	HOST       = "localhost:8081"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//s := r.Host("get5.flowing.tokyo").Subrouter()
	r.HandleFunc("/", get5.HomeHandler).Methods("GET")
	r.HandleFunc("/login", get5.LoginHandler).Methods("GET")
	r.HandleFunc("/logout", get5.LogoutHandler).Methods("GET")

	r.HandleFunc("/match/create", get5.MatchCreateHandler)             // GET/POST
	r.HandleFunc("/match/{matchID}", get5.MatchHandler)                // ?
	r.HandleFunc("/match/{matchID}/config", get5.MatchConfigHandler)   // ?
	r.HandleFunc("/match/{matchID}/cancel", get5.MatchCancelHandler)   // ?
	r.HandleFunc("/match/{matchID}/rcon", get5.MatchRconHandler)       // ?
	r.HandleFunc("/match/{matchID}/pause", get5.MatchPauseHandler)     // ?
	r.HandleFunc("/match/{matchID}/unpause", get5.MatchUnpauseHandler) // ?
	r.HandleFunc("/match/{matchID}/adduser", get5.MatchAdduserHandler) // ?
	//r.HandleFunc("/match/{matchID}/sendconfig", get5.MatchSendConfigHandler) // ?
	r.HandleFunc("/match/{matchID}/backup", get5.MatchBackupHandler).Methods("GET")

	r.HandleFunc("/match/{matchID}/finish", get5.MatchFinishHandler).Methods("POST")                  // POST
	r.HandleFunc("/match/{matchID}/map/{mapNumber}/start", get5.MatchMapStartHandler).Methods("POST") // POST
	r.HandleFunc("/match/{matchID}/map/{mapNumber}/update", get5.MatchMapUpdateHandler).Methods("POST")
	r.HandleFunc("/match/{matchID}/map/{mapNumber}/finish", get5.MatchMapFinishHandler).Methods("POST")
	r.HandleFunc("/match/{matchID}/map/{mapNumber}/player/{steamid64}/update", get5.MatchMapPlayerUpdateHandler).Methods("POST")

	r.HandleFunc("/matches", get5.MatchesHandler)                // ?
	r.HandleFunc("/matches/{userID}", get5.MatchesWithIDHandler) // ?
	r.HandleFunc("/mymatches", get5.MyMatchesHandler)            // ?

	r.HandleFunc("/team/create", get5.TeamCreateHandler) // GET/POST
	r.HandleFunc("/team/{teamID}", get5.TeamHandler).Methods("GET")
	r.HandleFunc("/team/{teamID}/edit", get5.TeamEditHandler) // GET/POST
	r.HandleFunc("/team/{teamID}/delete", get5.TeamDeleteHandler)
	r.HandleFunc("/teams/{userID}", get5.TeamsHandler).Methods("GET")
	r.HandleFunc("/myteams", get5.MyTeamsHandler).Methods("GET")

	r.HandleFunc("/server/create", get5.ServerCreateHandler)                           // GET/POST
	r.HandleFunc("/server/{serverid}/edit", get5.ServerEditHandler)                    // GET/POST
	r.HandleFunc("/server/{serverid}/delete", get5.ServerDeleteHandler).Methods("GET") // GET
	r.HandleFunc("/myservers", get5.MyServersHandler)                                  // ?

	r.HandleFunc("/user/{userID}", get5.UserHandler)
	r.Methods("GET", "POST")
	http.Handle("/", r)
	fmt.Println("RUNNING")
	log.Fatal(http.ListenAndServe(HOST, nil))
}
