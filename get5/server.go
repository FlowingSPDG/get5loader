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

func ServerCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("ServerCreateHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func ServerEditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("ServerEditHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func ServerDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("ServerDeleteHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MyServersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MyServersHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}
