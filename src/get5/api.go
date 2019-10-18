package get5

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "github.com/solovev/steam_go"
	//_ "html/template"
	_ "github.com/valyala/quicktemplate/examples/basicserver/templates"
	"net/http"
	_ "strconv"
	_ "time"
)

func MatchFinishHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchFinishHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchMapStartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchMapStartHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchMapUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchMapUpdateHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchMapFinishHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchMapFinishHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchMapPlayerUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchMapPlayerUpdateHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("UserHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}
