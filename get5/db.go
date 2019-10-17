package get5

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/solovev/steam_go"
	_ "html/template"
	"net/http"
	_ "strconv"
	_ "time"
)

type HomeData struct {
	LoggedIn bool
	Content  interface{} // should be template
	UserName string
	UserID   string
}

var (
	UserDatas    = map[string]*UserData{}
	SteamAPIKey  = "7A9C505B9AA359CC5DF2AF43448B33B7"
	SessionStore = sessions.NewCookieStore([]byte("GET5_GO_SESSIONKEY"))
	SessionData  = "SessionData"
	DefaultPage  = "/matches"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, DefaultPage, 302)
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
		fmt.Printf("\nUserName : %s\n", user.PersonaName)
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
