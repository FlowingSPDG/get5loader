package get5

import (
	"fmt"
	"strconv"

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
	//vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchesHandler\n")

	name := ""
	userid := 0
	loggedin := false
	session, _ := db.SessionStore.Get(r, db.SessionData)

	matches, err := db.SQLAccess.MySQLGetMatchData(20, "", "")
	if err != nil {
		panic(err)
	}

	u := &db.MatchesPageData{
		LoggedIn: loggedin,
		UserName: name,
		UserID:   userid,
		Matches:  matches,
	}

	if _, ok := session.Values["Loggedin"]; ok { // FUCK.
		if _, ok := session.Values["Name"]; ok {
			if _, ok := session.Values["UserID"].(int); ok {
				if session.Values["Loggedin"].(bool) == true {
					u.LoggedIn = true
					u.UserName = session.Values["Name"].(string)
					u.UserID = session.Values["UserID"].(int)
				}
			}
		}
	}

	fmt.Fprintf(w, templates.Match(u)) // TODO
}

func MatchesWithIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchesWithIDHandler\n")

	name := ""
	userid := 0
	session, _ := db.SessionStore.Get(r, db.SessionData)

	u := &db.MatchesPageData{
		LoggedIn: false,
		UserName: name,
		UserID:   userid,
	}

	if _, ok := session.Values["Loggedin"]; ok { // FUCK.
		if _, ok := session.Values["Name"]; ok {
			if _, ok := session.Values["UserID"].(int); ok {
				if session.Values["Loggedin"].(bool) == true {
					u.LoggedIn = true
					u.UserName = session.Values["Name"].(string)
					u.UserID = session.Values["UserID"].(int)
					userid = session.Values["UserID"].(int)
				}
			}
		}
	}

	matches, err := db.SQLAccess.MySQLGetMatchData(20, "user_id", vars["userID"])

	u.Matches = matches

	fmt.Println(matches)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, templates.Match(u)) // TODO
}

func MyMatchesHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchesHandler\n")

	name := ""
	userid := 0
	loggedin := false
	session, _ := db.SessionStore.Get(r, db.SessionData)

	u := &db.MatchesPageData{
		LoggedIn: loggedin,
		UserName: name,
		UserID:   userid,
	}

	if _, ok := session.Values["Loggedin"]; ok { // FUCK.
		if _, ok := session.Values["Name"]; ok {
			if _, ok := session.Values["UserID"].(int); ok {
				if session.Values["Loggedin"].(bool) == true {
					u.LoggedIn = true
					loggedin = true
					u.UserName = session.Values["Name"].(string)
					u.UserID = session.Values["UserID"].(int)
					userid = session.Values["UserID"].(int)
				}
			}
		}
	}

	if !loggedin {
		http.Redirect(w, r, "/login", 302)
	}

	matches, err := db.SQLAccess.MySQLGetMatchData(20, "user_id", strconv.Itoa(userid))

	u.Matches = matches

	fmt.Println(matches)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, templates.Match(u)) // TODO
}
