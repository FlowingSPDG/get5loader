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
	"strconv"
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
	vars := mux.Vars(r)
	userid := vars["userID"]
	var user models.SQLUserData
	var team []models.SQLTeamData
	user, _ = db.SQLAccess.MySQLGetUserData(1, "id = "+userid)
	team, _ = db.SQLAccess.MySQLGetTeamData(20, "user_id = "+userid)

	session, _ := db.SessionStore.Get(r, db.SessionData)
	fmt.Printf("TeamsHandler\nvars : %v", vars)

	loggedin := false
	IsYourTeam := false

	if _, ok := session.Values["Loggedin"]; ok {
		loggedin = session.Values["Loggedin"].(bool)
		if _, ok := session.Values["UserID"]; ok {
			fmt.Println("session.Values[UserID] : " + strconv.Itoa(session.Values["UserID"].(int)))
			IsYourTeam = userid == strconv.Itoa(session.Values["UserID"].(int))
		}
	}

	PageData := &models.TeamsPageData{
		LoggedIn:   loggedin,
		User:       user,
		Teams:      team,
		IsYourTeam: IsYourTeam,
	}
	fmt.Println(team[0].Name)

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
