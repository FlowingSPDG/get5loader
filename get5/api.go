package get5

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/solovev/steam_go"
	"html/template"
	"net/http"
	_ "strconv"
	"time"
)

type HomeData struct {
	LoggedIn bool
	Content  interface{} // should be template
	UserName string
	UserID   string
}

type UserData struct {
	ID      int
	SteamID string
	Name    string
	Admin   bool
	Servers []GameServerData
	Teams   []TeamData
	Matches []MatchData
}

type GameServerData struct {
	ID           int
	UserID       int
	DisplayName  string
	IPstring     string
	port         int
	RconPassword string
	InUse        bool
	PublicServer bool
}

type TeamData struct {
	ID         int
	UserID     int
	Name       string
	Tag        string
	Flag       string
	Logo       string
	Auths      []string
	PublicTeam bool
}

type MatchData struct {
	ID            int64
	UserID        int64
	ServerID      int64
	Team1ID       int64
	Team1Score    int
	Team1String   string
	Team2ID       int64
	Team2Score    int
	Team2String   string
	winner        int64
	PluginVersion string
	forfeit       bool
	cancelled     bool
	StartTime     time.Time
	EndTime       time.Time
	MaxMaps       int
	title         string
	SkipVeto      bool
	APIKey        string

	VetoMapPool []string
	MapStats    []MapStatsData
}

type MapStatsData struct {
	ID         int
	MatchID    int
	MapNumber  int
	MapName    string
	StartTime  time.Time
	EndTIme    time.Time
	Winner     int
	Team1Score int
	Team2Score int
}

var (
	UserDatas    = map[string]*UserData{}
	SteamAPIKey  = "7A9C505B9AA359CC5DF2AF43448B33B7"
	SessionStore = sessions.NewCookieStore([]byte("GET5_GO_SESSIONKEY"))
	SessionData  = "SessionData"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("get5/templates/layout.html", "get5/templates/matches.html")) // template
	vars := mux.Vars(r)                                                                                    //パスパラメータ取得
	fmt.Printf("HomeHandler\nvars : %v\n", vars)

	name := ""
	userid := ""
	loggedin := false
	session, _ := SessionStore.Get(r, SessionData)

	m := &HomeData{
		LoggedIn: false,
		Content:  tpl,
		UserName: name,
		UserID:   userid,
	}

	if _, ok := session.Values["Name"]; ok {
		if ok {
			name, _ = session.Values["Name"].(string)
			loggedin = true
		}
	}

	if _, ok := session.Values["UserID"]; ok {
		if ok {
			userid, _ = session.Values["UserID"].(string)
			loggedin = true
		}
	}

	m.UserID = userid
	m.LoggedIn = loggedin

	fmt.Println(m)

	// テンプレートを描画
	tpl.Execute(w, m)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	opID := steam_go.NewOpenId(r)
	switch opID.Mode() {
	case "":
		http.Redirect(w, r, opID.AuthUrl(), 302)
	case "cancel":
		//w.Write([]byte("Authorization cancelled"))
		http.Redirect(w, r, "/", 302)
	default:
		user, err := opID.ValidateAndGetUser(SteamAPIKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Printf("user : %v", *user)
		//steam_go.GetPlayerSummaries
		if err != nil {
			panic(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if user == nil {
			panic(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		UserDatas[user.SteamId] = &UserData{
			SteamID: user.SteamId,
			Name:    user.PersonaName,
		}
		session, _ := SessionStore.Get(r, SessionData)
		session.Options = &sessions.Options{MaxAge: 0}
		// Set some session values.
		session.Values["UserID"] = user.SteamId
		session.Values["Name"] = user.PersonaName
		// Save it before we write to the response/return from the handler.
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Register to DB if its new player
		http.Redirect(w, r, "/", 302)
	}
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("LoginHandler\nvars : %v\n", vars)
	w.WriteHeader(http.StatusOK)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("LogoutHandler\nvars : %v", vars)
	session, _ := SessionStore.Get(r, SessionData)
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 302)
}

func MatchesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchesHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchesWithIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchesWithIDHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("TeamHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("UserHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}
