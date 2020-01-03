package api

import (
	"encoding/json"
	"fmt"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	"github.com/gorilla/mux"
	"net/http"
)

// MatchConfigHandler Handler for /api/v1/match/{matchID}/config API.
func MatchConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("MatchConfigHandler\n")
	matchid := vars["matchID"]
	match := db.MatchData{}
	db.SQLAccess.Gorm.Where("id = ?", matchid).First(&match)
	res, err := match.BuildMatchDict()
	if err != nil {
		http.Error(w, "Internal ERROR", http.StatusInternalServerError)
		return
	}
	jsonbyte, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Internal ERROR", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonbyte)
}
