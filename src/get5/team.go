package get5

import (
	"fmt"

	"github.com/FlowingSPDG/get5-web-go/src/db"
	"github.com/FlowingSPDG/get5-web-go/templates"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "github.com/solovev/steam_go"

	//"html/template"
	"net/http"
	"strconv"
	_ "time"
)

func TeamCreateHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := db.SessionStore.Get(r, db.SessionData)
	m := &db.TeamCreatePageData{
		Edit: false,
		//Content: tpl,
	}
	if _, ok := session.Values["Loggedin"]; ok {
		m.LoggedIn = true
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func TeamHandler(w http.ResponseWriter, r *http.Request) { // NEED Team player's info and recent matches info
	vars := mux.Vars(r)
	teamID := vars["teamID"]

	team, err := db.SQLAccess.MySQLGetTeamData(1, "id", teamID)
	if err != nil {
		panic(err)
	}

	session, _ := db.SessionStore.Get(r, db.SessionData)
	fmt.Printf("TeamHandler\nvars : %v", vars)

	loggedin := false
	IsYourTeam := false

	if _, ok := session.Values["Loggedin"]; ok {
		loggedin = session.Values["Loggedin"].(bool)
		if _, ok := session.Values["UserID"]; ok {
			fmt.Println("session.Values[UserID] : " + strconv.Itoa(session.Values["UserID"].(int)))
			IsYourTeam = session.Values["UserID"].(int) == team[0].UserID
		}
	}

	user, err := db.SQLAccess.MySQLGetUserData(1, "id", vars["userID"])
	if err != nil {
		panic(err)
	}

	u := &db.TeamPageData{
		LoggedIn:   loggedin,
		Team:       team[0],
		IsYourTeam: IsYourTeam,
		User:       user,
	}
	fmt.Println(team[0].Name)

	fmt.Fprintf(w, templates.Team(u)) // TODO
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
	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["userID"])
	if err != nil {
		panic(err)
	}
	var user db.UserData
	var team []db.TeamData
	user, _ = db.SQLAccess.MySQLGetUserData(1, "id", vars["userID"])
	team, _ = db.SQLAccess.MySQLGetTeamData(20, "user_id", vars["userID"])

	session, _ := db.SessionStore.Get(r, db.SessionData)
	fmt.Printf("TeamsHandler\nvars : %v", vars)

	loggedin := false
	IsYourTeam := false

	if _, ok := session.Values["Loggedin"]; ok {
		loggedin = session.Values["Loggedin"].(bool)
		if _, ok := session.Values["UserID"]; ok {
			fmt.Println("session.Values[UserID] : " + strconv.Itoa(session.Values["UserID"].(int)))
			IsYourTeam = userid == session.Values["UserID"].(int)
		}
	}

	PageData := &db.TeamsPageData{
		LoggedIn:   loggedin,
		User:       user,
		Teams:      team,
		IsYourTeam: IsYourTeam,
	}

	fmt.Fprintf(w, templates.Teams(PageData)) // TODO
}

func MyTeamsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := db.SessionStore.Get(r, db.SessionData)
	if _, ok := session.Values["Loggedin"]; ok {
		if session.Values["Loggedin"] == true {
			if _, ok := session.Values["UserID"]; ok {
				http.Redirect(w, r, "/teams/"+strconv.Itoa(session.Values["UserID"].(int)), 302)
			}
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	}
	w.WriteHeader(http.StatusOK)
}
