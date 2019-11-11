package db

import (
	// "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	// "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/go-ini/ini"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/solovev/steam_go"

	"net/http"
)

// Config Configration Struct for config.ini
type Config struct {
	SteamAPIKey string
	DefaultPage string
	SQLHost     string
	SQLUser     string
	SQLPass     string
	SQLPort     int
	SQLDBName   string
	HOST        string
}

// DBdatas Struct for MySQL configration and Gorm
type DBdatas struct {
	Host string
	User string
	Pass string
	Db   string
	Port int
	Gorm *gorm.DB
}

var (
	// SteamAPIKey Steam Web API Key for accessing Steam API.
	SteamAPIKey = ""
	// SessionStore Session cookie store data.
	SessionStore = sessions.NewCookieStore([]byte("GET5_GO_SESSIONKEY"))
	// SessionData Contains Session data.
	SessionData = "SessionData"
	// DefaultPage Default page where player access root directly.
	DefaultPage string
	// SQLAccess SQL Access Object for MySQL and GORM things
	SQLAccess DBdatas
	// Cnf Configration Data
	Cnf Config
)

func init() {
	c, _ := ini.Load("config.ini")
	Cnf := Config{
		SteamAPIKey: c.Section("Steam").Key("APIKey").MustString(""),
		DefaultPage: c.Section("GET5").Key("DefaultPage").MustString(""),
		HOST:        c.Section("GET5").Key("HOST").MustString(""),
		SQLHost:     c.Section("sql").Key("host").MustString(""),
		SQLUser:     c.Section("sql").Key("user").MustString(""),
		SQLPass:     c.Section("sql").Key("pass").MustString(""),
		SQLPort:     c.Section("sql").Key("port").MustInt(3306),
		SQLDBName:   c.Section("sql").Key("database").MustString(""),
	}
	SQLAccess = DBdatas{
		Host: Cnf.SQLHost,
		User: Cnf.SQLUser,
		Pass: Cnf.SQLPass,
		Db:   Cnf.SQLDBName,
		Port: Cnf.SQLPort,
	}
	sqloption := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", SQLAccess.User, SQLAccess.Pass, SQLAccess.Host, SQLAccess.Port, SQLAccess.Db)
	//fmt.Println(sqloption)
	DB, err := gorm.Open("mysql", sqloption)
	if err != nil {
		panic(err)
	}
	DB.LogMode(true)
	SQLAccess.Gorm = DB
	SteamAPIKey = Cnf.SteamAPIKey
	DefaultPage = Cnf.DefaultPage
}

// HomeHandler HTTP Handler for default page redirecting.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, DefaultPage, 302)
}

// LoginHandler HTTP Handler for /login page.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	opID := steam_go.NewOpenId(r)
	switch opID.Mode() {
	case "":
		http.Redirect(w, r, opID.AuthUrl(), 302)
	case "cancel":
		http.Redirect(w, r, DefaultPage, 302)
	default:
		steamid, err := opID.ValidateAndGetId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
		fmt.Println("SteamID : " + steamid)

		user := UserData{}
		user.GetOrCreate(SQLAccess.Gorm, steamid)
		session, _ := SessionStore.Get(r, SessionData)
		session.Options = &sessions.Options{MaxAge: 0}
		// Set some session values.
		session.Values["Loggedin"] = true
		session.Values["UserID"] = user.ID // should be get5 id
		session.Values["Name"] = user.Name
		session.Values["SteamID"] = steamid
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

// LogoutHandler HTTP Handler for /logout
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

// GetUserData Gets UserData array via MySQL(GORM).
func (s *DBdatas) GetUserData(limit int, wherekey string, wherevalue string) ([]UserData, error) {
	UserData := []UserData{}
	s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&UserData)
	return UserData, nil
}
