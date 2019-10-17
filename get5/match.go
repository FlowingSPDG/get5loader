package get5

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "github.com/solovev/steam_go"
	_ "html/template"
	"net/http"
	_ "strconv"
	_ "time"
)

func MatchCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchCreateHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchConfigHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchCancelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchCancelHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchRconHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchRconHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchPauseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchPauseHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchUnpauseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchUnpauseHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchAdduserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchAdduserHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

/*
func MatchSendConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchSendConfigHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}
*/

func MatchBackupHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchBackupHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchesHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchesWithIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchesWithIDHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MyMatchesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MyMatchesHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}
