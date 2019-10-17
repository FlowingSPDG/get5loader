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

func TeamCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("TeamCreateHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("TeamHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func TeamEditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("TeamEditHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func TeamDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("TeamDeleteHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func TeamsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("TeamsHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MyTeamsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MyTeamsHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}
