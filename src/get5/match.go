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

// MatchCreateHandler HTTP Handler for /match/create . not implemented yet
func MatchCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchCreateHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// MatchHandler HTTP Handler for /match/{matchID} page.
func MatchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchHandler\nvars : %v", vars)

	loggedin := false
	session, _ := db.SessionStore.Get(r, db.SessionData)

	match, err := db.SQLAccess.MySQLGetMatchData(1, "id", vars["matchID"])
	if err != nil {
		panic(err)
	}

	u := &db.MatchPageData{
		LoggedIn: loggedin,
		Match:    match[0],
	}

	if _, ok := session.Values["Loggedin"]; ok { // FUCK.
		u.LoggedIn = session.Values["Loggedin"].(bool)
	}

	fmt.Fprintf(w, templates.Match(u)) // TODO
}

// MatchConfigHandler HTTP Handler for /match/{matchID}/config page.
func MatchConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchConfigHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// MatchCancelHandler HTTP Handler for /match/{matchID}/cancel page.
func MatchCancelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchCancelHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// MatchRconHandler HTTP Handler for /match/{matchID}/rcon
func MatchRconHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchRconHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// MatchPauseHandler HTTP Handler for /match/{matchID}/pause
func MatchPauseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchPauseHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// MatchUnpauseHandler HTTP Handler for /match/{matchID}/unpause
func MatchUnpauseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchUnpauseHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// MatchAddUserHandler HTTP Handler for /match/{matchID}/adduser
func MatchAddUserHandler(w http.ResponseWriter, r *http.Request) {
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

// MatchBackupHandler HTTP Handler for /match/{matchID}/backup
func MatchBackupHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchBackupHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

// MatchesHandler HTTP Handler for /matches/
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
		LoggedIn:   loggedin,
		UserName:   name,
		UserID:     userid,
		Matches:    matches,
		AllMatches: true,
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

	fmt.Fprintf(w, templates.Matches(u)) // TODO
}

// MatchesWithIDHandler HTTP Handler for /matches/{userID}
func MatchesWithIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchesWithIDHandler\n")

	name := ""
	userid, err := strconv.Atoi(vars["userID"])
	if err != nil {
		panic(err)
	}
	session, _ := db.SessionStore.Get(r, db.SessionData)

	user, err := db.SQLAccess.MySQLGetUserData(1, "id", vars["userID"])
	if err != nil {
		panic(err)
	}

	u := &db.MatchesPageData{
		LoggedIn:   false,
		UserName:   name,
		UserID:     userid,
		MyMatches:  false,
		AllMatches: false,
		Owner:      user,
	}

	if _, ok := session.Values["Loggedin"]; ok { // FUCK.
		if _, ok := session.Values["Name"]; ok {
			if _, ok := session.Values["UserID"].(int); ok {
				if session.Values["Loggedin"].(bool) == true {
					u.LoggedIn = true
					u.UserName = session.Values["Name"].(string)
					if session.Values["UserID"].(int) == userid {
						u.MyMatches = true
					}
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

	fmt.Fprintf(w, templates.Matches(u)) // TODO
}

// MyMatchesHandler HTTP Handler for /mymatches
func MyMatchesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchesHandler\n")

	name := ""
	userid := 0
	loggedin := false
	session, _ := db.SessionStore.Get(r, db.SessionData)

	user, err := db.SQLAccess.MySQLGetUserData(1, "id", vars["userID"])

	u := &db.MatchesPageData{
		LoggedIn:   loggedin,
		UserName:   name,
		UserID:     userid,
		MyMatches:  true,
		AllMatches: false,
		Owner:      user,
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

	fmt.Fprintf(w, templates.Matches(u)) // TODO
}
