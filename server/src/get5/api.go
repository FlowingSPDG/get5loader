package get5

import (
	"fmt"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "github.com/solovev/steam_go"

	//_ "html/template"
	"net/http"
	_ "strconv"
	_ "time"

	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	"github.com/FlowingSPDG/get5-web-go/server/templates"
)

// MatchFinishHandler HTTP Handler for manage match finish. not implemented yet
func MatchFinishHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchFinishHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// MatchMapStartHandler HTTP Handler for Map start. not implemented yet
func MatchMapStartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchMapStartHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// MatchMapUpdateHandler HTTP Handler for Match map information update. not implemented yet
func MatchMapUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchMapUpdateHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// MatchMapFinishHandler HTTP Handler for Match map finish. not implemented yet
func MatchMapFinishHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchMapFinishHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// MatchMapPlayerUpdateHandler HTTP Handler for Match map player stats update. not implemented yet
func MatchMapPlayerUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchMapPlayerUpdateHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// UserHandler HTTP Handler for /user/{userID}
func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("UserHandler\nvars : %v", vars)

	loggedin := false
	session, _ := db.SessionStore.Get(r, db.SessionData)

	User, err := db.SQLAccess.GetUserData(1, "id", vars["userID"])
	if err != nil {
		panic(err)
	}

	u := &db.UserPageData{
		LoggedIn: loggedin,
		User:     User[0],
	}

	if _, ok := session.Values["Loggedin"]; ok { // FUCK.
		if session.Values["Loggedin"].(bool) == true {
			u.LoggedIn = true
		}
	}

	fmt.Fprintf(w, templates.User(u)) // TODO
}

// MetricsHandler HTTP Handler for /metrics
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MetricsHandler")

	loggedin := false
	session, _ := db.SessionStore.Get(r, db.SessionData)

	data := db.GetMetrics()
	fmt.Println(data)

	u := &db.MetricsDataPage{
		LoggedIn: loggedin,
		Data:     data,
	}

	if _, ok := session.Values["Loggedin"]; ok { //
		u.LoggedIn = session.Values["Loggedin"].(bool)
	}

	fmt.Fprintf(w, templates.Metrics(u)) // TODO
}
