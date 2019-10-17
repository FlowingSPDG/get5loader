package get5

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

type HomeData struct {
	Date    string
	Time    string
	Content interface{} // should be template
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("get5/templates/layout.html", "get5/templates/matches.html")) // template
	vars := mux.Vars(r)                                                                                    //パスパラメータ取得
	fmt.Printf("HomeHandler\nvars : %v", vars)

	m := &HomeData{
		Date:    time.Now().Format("2006-01-02"),
		Time:    time.Now().Format("15:04:05"),
		Content: tpl,
	}

	// テンプレートを描画
	tpl.Execute(w, m)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("HomeHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("HomeHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchesWithIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("HomeHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("HomeHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("HomeHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("HomeHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}
