package get5

import (
	"fmt"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	"github.com/FlowingSPDG/get5-web-go/server/templates"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "github.com/solovev/steam_go"

	//_ "html/template"
	"net/http"
	"strconv"
	_ "time"
)

// ServerCreateHandler HTTP Handler for /server/create
func ServerCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("ServerCreateHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// ServerEditHandler HTTP Handler for /server/{serverID}/edit
func ServerEditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("ServerEditHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// ServerDeleteHandler HTTP Handler for /server/{serverID}/delete
func ServerDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("ServerDeleteHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// MyServersHandler HTTP Handler for /myservers
func MyServersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	var servers []db.GameServerData

	session, _ := db.SessionStore.Get(r, db.SessionData)
	fmt.Printf("TeamsHandler\nvars : %v", vars)

	loggedin := false

	userid := ""

	if _, ok := session.Values["Loggedin"]; ok {
		loggedin = session.Values["Loggedin"].(bool)
		if _, ok := session.Values["UserID"]; ok {
			userid = strconv.Itoa(session.Values["UserID"].(int))
		}
	}

	if !loggedin {
		http.Redirect(w, r, "/login", 302)
	}

	servers, _ = db.SQLAccess.MySQLGetGameServerData(20, "user_id", userid)

	PageData := &db.MyserversPageData{
		Servers:  servers,
		LoggedIn: loggedin,
	}

	fmt.Fprintf(w, templates.Myservers(PageData)) // TODO
}
