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

	"github.com/FlowingSPDG/get5-web-go/src/db"
	"github.com/FlowingSPDG/get5-web-go/templates"
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
