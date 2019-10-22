package db

import (
	// "database/sql" //ここでパッケージをimport
	"fmt"
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/go-ini/ini"
	//"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/solovev/steam_go"

	//_ "html/template"
	_ "log"
	"net/http"
	_ "strconv"
	_ "time"
)

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

type DBdatas struct {
	Host string
	User string
	Pass string
	Db   string
	Port int
	Gorm *gorm.DB
}

var (
	UserDatas    = map[string]*UserData{}
	SteamAPIKey  = ""
	SessionStore = sessions.NewCookieStore([]byte("GET5_GO_SESSIONKEY"))
	SessionData  = "SessionData"
	DefaultPage  string
	SQLAccess    DBdatas
	Cnf          Config
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
	// gorm.Open("mysql", "gorm:password@/country?charset=utf8&parseTime=True&loc=Local")
	//fmt.Println(sqloption)
	DB, err := gorm.Open("mysql", sqloption)
	if err != nil {
		panic(err)
	}
	DB.SingularTable(true)
	SQLAccess.Gorm = DB
	SteamAPIKey = Cnf.SteamAPIKey
	DefaultPage = Cnf.DefaultPage
}

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

func (s *DBdatas) GetUserData(limit int, wherekey string, wherevalue string) ([]UserData, error) {
	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	UserData := []UserData{}
	s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&UserData)
	return UserData, nil
}

func (s *DBdatas) MySQLGetTeamData(limit int, wherekey string, wherevalue string) ([]TeamData, error) {
	TeamData := []TeamData{}
	s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&TeamData)
	return TeamData, nil
}

func (s *DBdatas) MySQLGetMatchData(limit int, wherekey string, wherevalue string) ([]MatchData, error) {
	Matches := []MatchData{}
	if wherekey == "" || wherevalue == "" {
		s.Gorm.Order("id DESC").Limit(limit).Find(&Matches)
	} else {
		s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&Matches)
	}
	return Matches, nil
}

func (s *DBdatas) MySQLGetPlayerStatsData(limit int, wherekey string, wherevalue string) ([]SQLPlayerStatsData, error) {
	PlayerStatsData := []SQLPlayerStatsData{}
	s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&PlayerStatsData)
	return PlayerStatsData, nil
}

func (s *DBdatas) MySQLGetMapStatsData(limit int, wherekey string, wherevalue string) ([]MapStatsData, error) {
	MapStatsData := []MapStatsData{}
	s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&MapStatsData)
	return MapStatsData, nil
}

func (s *DBdatas) MySQLGetGameServerData(limit int, wherekey string, wherevalue string) ([]GameServerData, error) {
	GameServer := []GameServerData{}
	s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&GameServer)
	return GameServer, nil
}

func (s *DBdatas) MySQLGetUserData(limit int, wherekey string, wherevalue string) (UserData, error) {
	UserData := UserData{}
	s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&UserData)
	return UserData, nil
}
