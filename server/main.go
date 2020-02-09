package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"github.com/FlowingSPDG/get5-web-go/server/src/api"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	"github.com/FlowingSPDG/get5-web-go/server/src/grpc"
	"github.com/gin-contrib/static"
	"github.com/go-ini/ini"
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
	APIONLY     bool
}

var (
	// StaticDir Directly where serves static files
	StaticDir = "./static"
	// HOST Server's domain name
	HOST string
	//Cnf Configration Data
	Cnf Config
	// SQLAccess SQL Access Object for MySQL and GORM things
	SQLAccess db.DBdatas

	GRPC_ADDR  string
	EnablegRPC bool
)

func init() {
	c, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}
	Cnf = Config{
		HOST:      c.Section("GET5").Key("HOST").MustString("localhost:8080"),
		SQLHost:   c.Section("sql").Key("host").MustString(""),
		SQLUser:   c.Section("sql").Key("user").MustString(""),
		SQLPass:   c.Section("sql").Key("pass").MustString(""),
		SQLPort:   c.Section("sql").Key("port").MustInt(3306),
		SQLDBName: c.Section("sql").Key("database").MustString(""),
		APIONLY:   c.Section("GET5").Key("API_ONLY").MustBool(false),
	}
	HOST = Cnf.HOST
	GRPC_ADDR = c.Section("GET5").Key("GPRC_ADDR").MustString(":50055")
	EnablegRPC = c.Section("GET5").Key("ENABLE_gRPC").MustBool(false)
}

func main() {

	defer SQLAccess.Gorm.Close()

	r := gin.Default()

	//s := r.Host(HOST).Subrouter() // in-case if we need vhost thing

	// misc
	r.GET("/login", db.LoginHandler)
	r.GET("/logout", db.LogoutHandler)

	v1 := r.Group("/api/v1")
	{
		// session handling
		v1.GET("/login", db.LoginHandler)
		v1.GET("/logout", db.LogoutHandler)

		v1.GET("/GetMatches", api.GetMatches)
		v1.GET("/GetMetrics", api.GetMetrics)
		v1.GET("/GetSteamName", api.GetSteamName)
		v1.GET("/GetTeamList", api.GetTeamList)
		v1.GET("/GetServerList", api.GetServerList)
		v1.GET("/GetVersion", api.GetVersion)
		v1.GET("/GetMapList", api.GetMapList)
		v1.GET("/CheckLoggedIn", api.CheckLoggedIn)

		// Match API for front(Vue)
		match := v1.Group("/match")
		{
			match.GET("/:matchID/GetMatchInfo", api.GetMatchInfo)
			match.GET("/:matchID/GetPlayerStatInfo", api.GetPlayerStatInfo)
			match.GET("/:matchID/GetStatusString", api.GetStatusString)
			match.POST("/:matchID", api.CreateMatch) // avoid conflicts...

			// GET5 API
			match.GET("/:matchID/config", api.MatchConfigHandler)
			match.POST("/:matchID/finish", api.MatchFinishHandler)
			match.POST("/:matchID/start", api.MatchMapStartHandler)
			match.POST("/:matchID/map/:mapNumber/update", api.MatchMapUpdateHandler)
			match.POST("/:matchID/map/:mapNumber/finish", api.MatchMapFinishHandler)
			match.POST("/:matchID/map/:mapNumber/player/:steamid64/update", api.MatchMapPlayerUpdateHandler)

			match.POST("/:matchID/cancel", api.MatchCancelHandler)
			match.POST("/:matchID/rcon", api.MatchRconHandler)
			match.POST("/:matchID/pause", api.MatchPauseHandler)
			match.POST("/:matchID/unpause", api.MatchUnpauseHandler)
			match.POST("/:matchID/adduser", api.MatchAddUserHandler)
			// // match.POST("/:matchID/sendconfig", api.MatchSendConfigHandler) // ? // I won't implement this
			match.GET("/:matchID/backup", api.MatchListBackupsHandler)
			match.POST("/:matchID/backup", api.MatchLoadBackupsHandler)

			// match.POST("/:matchID/vetoUpdate", api.MatchVetoUpdateHandler)
			// match.POST("/:matchID/map/:mapNumber/demo", api.MatchDemoUploadHandler)
		}

		team := v1.Group("/team")
		{
			team.GET("/:teamID/GetTeamInfo", api.GetTeamInfo)
			team.GET("/:teamID/GetRecentMatches", api.GetRecentMatches)
			team.GET("/:teamID/CheckUserCanEdit", api.CheckUserCanEdit)
			team.POST("/create", api.CreateTeam)
			team.PUT("/teamID/edit", api.EditTeam)
			team.DELETE("/:teamID/delete", api.DeleteTeam)
		}

		user := v1.Group("/user")
		{
			user.GET("/:userID/GetUserInfo", api.GetUserInfo)
		}

		server := v1.Group("/server")
		{
			server.GET("/:serverID/GetServerInfo", api.GetServerInfo)
			server.POST("/create", api.CreateServer)
			server.PUT("/:serverID/edit", api.EditServer)
			server.DELETE("/:serverID/delete", api.DeleteServer)
		}
	}

	if !Cnf.APIONLY {
		/*
			entrypoint := "./static/index.html"
			r.Path("/").HandlerFunc(ServeStaticFile(entrypoint))
			r.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
			r.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
			r.PathPrefix("/img").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
			r.PathPrefix("/fonts").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts"))))
		*/
	} else {
		log.Println("API ONLY MODE")
	}

	if EnablegRPC {
		log.Println("EnableGRPC option enabled. Starting gRPC server...")
		go func() {
			err := get5grpc.StartGrpc(GRPC_ADDR)
			if err != nil {
				panic(err)
			}
		}()
	}

	r.Use(static.Serve("/", static.LocalFile("/static", false)))

	log.Panicf("Failed to listen port %s : %v\n", HOST, r.Run(HOST))

}
