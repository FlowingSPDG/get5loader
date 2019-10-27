package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FlowingSPDG/get5-web-go/server/src/db"
)

type CheckLoggedInJSON struct {
	IsLoggedIn bool `json:"isLoggedIn"`
}

func CheckLoggedIn(w http.ResponseWriter, r *http.Request) {
	response := CheckLoggedInJSON{
		IsLoggedIn: false,
	}
	session, _ := db.SessionStore.Get(r, db.SessionData)
	if _, ok := session.Values["Loggedin"]; ok {
		response.IsLoggedIn = session.Values["Loggedin"].(bool)
	}
	jsonbyte, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(string(jsonbyte))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}
