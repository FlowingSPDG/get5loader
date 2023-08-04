package db

import (
	"fmt"
	"log"
	"time"

	config "github.com/FlowingSPDG/get5-web-go/server/src/cfg"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/go-sessions/v3"
	"github.com/solovev/steam_go"
)

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
	// DefaultPage Default page where player access root directly.
	DefaultPage string
	// SQLAccess SQL Access Object for MySQL and GORM things
	SQLAccess DBdatas
	// Sess Session
	Sess *sessions.Sessions
)

func init() {
	SQLAccess = DBdatas{
		Host: config.Cnf.SQLHost,
		User: config.Cnf.SQLUser,
		Pass: config.Cnf.SQLPass,
		Db:   config.Cnf.SQLDBName,
		Port: config.Cnf.SQLPort,
	}
	sqloption := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", SQLAccess.User, SQLAccess.Pass, SQLAccess.Host, SQLAccess.Port, SQLAccess.Db)
	//log.Println(sqloption)
	DB, err := gorm.Open("mysql", sqloption)
	if err != nil {
		panic(err)
	}
	if config.Cnf.SQLDebugMode {
		log.Println("SQL Debug mode Enabled. Transaction logs active")
	}
	DB.LogMode(config.Cnf.SQLDebugMode)
	SQLAccess.Gorm = DB
	SteamAPIKey = config.Cnf.SteamAPIKey
	DefaultPage = config.Cnf.DefaultPage

	Sess = sessions.New(sessions.Config{
		// Cookie string, the session's client cookie name, for example: "mysessionid"
		//
		// Defaults to "gosessionid"
		Cookie: config.Cnf.Cookie,
		// it's time.Duration, from the time cookie is created, how long it can be alive?
		// 0 means no expire.
		// -1 means expire when browser closes
		// or set a value, like 2 hours:
		Expires: time.Hour * 2,
		// if you want to invalid cookies on different subdomains
		// of the same host, then enable it
		DisableSubdomainPersistence: false,
		// want to be crazy safe? Take a look at the "securecookie" example folder.
	})
}

// LoginHandler HTTP Handler for /login page.
func LoginHandler(c *gin.Context) {
	log.Printf("LoginHandler\n")
	opID := steam_go.NewOpenId(c.Request)
	switch opID.Mode() {
	case "":
		http.Redirect(c.Writer, c.Request, opID.AuthUrl(), 302)
	case "cancel":
		http.Redirect(c.Writer, c.Request, DefaultPage, 302)
	default:
		steamid, err := opID.ValidateAndGetId()
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
		log.Println("SteamID : " + steamid)

		user := UserData{SteamID: steamid}
		user.GetOrCreate()
		s := Sess.Start(c.Writer, c.Request)
		// Set some session values.
		s.Set("Loggedin", true)
		s.Set("UserID", user.ID) // should be get5 id
		s.Set("Name", user.Name)
		s.Set("SteamID", steamid)
		s.Set("Loggedin", true)
		s.Set("Admin", false)
		if user.Admin {
			s.Set("Admin", true)
		}

		// Register to DB if its new player
		http.Redirect(c.Writer, c.Request, "/", 302)
	}
}

// LogoutHandler HTTP Handler for /logout
func LogoutHandler(c *gin.Context) {
	log.Printf("LogoutHandler\n")
	s := Sess.Start(c.Writer, c.Request)
	s.Destroy()
	http.Redirect(c.Writer, c.Request, "/", 302)
}

// GetUserData Gets UserData array via MySQL(GORM).
func (s *DBdatas) GetUserData(limit int, wherekey string, wherevalue string) ([]UserData, error) {
	UserData := []UserData{}
	s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&UserData)
	return UserData, nil
}
