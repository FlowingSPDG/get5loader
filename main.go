package main

import (
	"./get5"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Gorilla test")
	r := mux.NewRouter()
	//s := r.Host("get5.flowing.tokyo").Subrouter()
	r.HandleFunc("/", get5.HomeHandler)
	r.HandleFunc("/login", get5.LoginHandler)
	r.HandleFunc("/matches", get5.MatchesHandler).Methods("GET")
	r.HandleFunc("/matches/{userID}", get5.MatchesWithIDHandler).Methods("GET")
	r.HandleFunc("/match/{matchID}", get5.MatchHandler)
	r.HandleFunc("/team/{teamID}", get5.TeamHandler)
	r.HandleFunc("/user/{userID}", get5.UserHandler)
	r.Methods("GET", "POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
