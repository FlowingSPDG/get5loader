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
	fmt.Println("Gorilla test")
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//s := r.Host("get5.flowing.tokyo").Subrouter()
	r.HandleFunc("/", get5.HomeHandler).Methods("GET")
	r.HandleFunc("/login", get5.LoginHandler).Methods("GET")
	r.HandleFunc("/logout", get5.LogoutHandler).Methods("GET")
	r.HandleFunc("/matches", get5.MatchesHandler).Methods("GET")
	r.HandleFunc("/matches/{userID}", get5.MatchesWithIDHandler).Methods("GET")
	r.HandleFunc("/match/{matchID}", get5.MatchHandler)
	r.HandleFunc("/team/{teamID}", get5.TeamHandler)
	r.HandleFunc("/user/{userID}", get5.UserHandler)
	r.Methods("GET", "POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(HOST, nil))
}
