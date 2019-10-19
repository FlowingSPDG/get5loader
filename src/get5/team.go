package get5

import (
	"fmt"
	"github.com/FlowingSPDG/get5-web-go/src/db"
	"github.com/FlowingSPDG/get5-web-go/src/models"
	"github.com/FlowingSPDG/get5-web-go/templates"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "github.com/solovev/steam_go"

	//"html/template"
	"net/http"
	_ "strconv"
	_ "time"
)

type TeamCreatePageData struct {
	LoggedIn bool
	Edit     bool
	Content  interface{} // should be template
}

func TeamCreateHandler(w http.ResponseWriter, r *http.Request) {
	//tpl := template.Must(template.ParseFiles("get5/templates/layout.html", "get5/templates/team_create.html")) // template
	session, _ := db.SessionStore.Get(r, db.SessionData)
	m := &TeamCreatePageData{
		Edit: false,
		//Content: tpl,
	}
	if _, ok := session.Values["Loggedin"]; ok {
		m.LoggedIn = true
	} else {
		http.Redirect(w, r, "/login", 302)
	}

	// テンプレートを描画
	//tpl.Execute(w, m)
	//p := &templates.MainPage{}
	//templates.WritePageTemplate(w, p)
}

type TeamPageData struct {
	LoggedIn   bool
	IsYourTeam bool
	team       models.SQLTeamData
	tp         string
	test       string
	Content    interface{} // should be template
}

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := vars["teamID"]
	//tpl := template.Must(template.ParseFiles("get5/templates/layout.html", "get5/templates/team.html")) // template
	session, _ := db.SessionStore.Get(r, db.SessionData)
	team, err := db.SQLAccess.MySQLGetTeamData(1, "id = "+t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("TeamHandler\nvars : %v", vars)

	loggedin := false

	if _, ok := session.Values["Loggedin"]; ok {
		loggedin = session.Values["Loggedin"].(bool)
	}

	PageData := &models.TeamPageData{
		LoggedIn:   loggedin,
		Teams:      team,
		IsYourTeam: false, // currently steamid
	}
	fmt.Println(team[0].Name)

	fmt.Fprintf(w, templates.Team(PageData)) // TODO
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
	session, _ := db.SessionStore.Get(r, db.SessionData)
	if _, ok := session.Values["Loggedin"]; ok {
		if session.Values["Loggedin"] == true {

		}
	}
	w.WriteHeader(http.StatusOK)
}
