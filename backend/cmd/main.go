package main

import (
	"fmt"
	"log"

	config "github.com/FlowingSPDG/get5-web-go/backend/cfg"
	"github.com/FlowingSPDG/get5-web-go/backend/cmd/di"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg := config.GetConfig()

	// misc
	// r.GET("/login", db.LoginHandler)
	// r.GET("/logout", db.LogoutHandler)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/version", di.InitializeGetVersionController().Handle)
		v1.GET("/maps", di.InitializeGetMaplistController().Handle)
		/*
			// session handling
			v1.GET("/login", db.LoginHandler)
			v1.GET("/logout", db.LogoutHandler)

			v1.GET("/GetMatches", api.GetMatches)
			v1.GET("/GetMetrics", api.GetMetrics)
			v1.GET("/GetSteamName", api.GetSteamName)
			v1.GET("/GetTeamList", api.GetTeamList)
			v1.GET("/GetServerList", api.GetServerList)
			v1.GET("/CheckLoggedIn", api.CheckLoggedIn)
		*/

		// Match API for front(Vue)
		match := v1.Group("/match")
		{
			match.GET("/:matchID", di.InitializeGetMatchController(cfg).Handle)
			/*
				match.GET("/:matchID/GetPlayerStatInfo", api.GetPlayerStatInfo)

					match.GET("/:matchID/GetStatusString", api.GetStatusString)
					match.POST("/:matchID", api.CreateMatch) // avoid conflicts...

					// GET5 API
					match.GET("/:matchID/config", api.MatchConfigHandler)
					match.POST("/:matchID/finish", api.MatchFinishHandler)
					match.POST("/:matchID/map/:mapNumber/start", api.MatchMapStartHandler)
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
			*/
		}

		/*
			team := v1.Group("/team")
			{
				team.GET("/:teamID/GetTeamInfo", api.GetTeamInfo)
				team.GET("/:teamID/GetRecentMatches", api.GetRecentMatches)
				team.GET("/:teamID/CheckUserCanEdit", api.CheckUserCanEdit)
				team.POST("/create", api.CreateTeam)
				team.PUT("/:teamID/edit", api.EditTeam)
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
		*/
	}
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Panicf("Failed to listen port %v", r.Run(addr))
}
