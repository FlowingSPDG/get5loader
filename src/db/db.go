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
	fmt.Println(sqloption)
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
		s.Gorm.Limit(limit).Find(&Matches)
	} else {
		s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&Matches)
	}
	return Matches, nil
	/*
		err := s.sql.Ping()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		var whr string = ""
		if len(where) > 0 {
			whr = "WHERE " + where
		}

		//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
		q := fmt.Sprintf("SELECT * FROM `match` %s ORDER BY ID DESC LIMIT %d ", whr, limit)
		fmt.Println(q)
		rows, err := s.sql.Query(q)
		if err != nil {
			return nil, err
			panic(err.Error())
		}
		defer rows.Close()

		Matches := make([]SQLMatchData, 0)
		//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
		for rows.Next() {
			var Match SQLMatchData
			err := rows.Scan(&Match.Id, &Match.User_id, &Match.Server_id, &Match.Team1_id, &Match.Team2_id, &Match.Winner, &Match.Cancelled, &Match.Start_time, &Match.End_time, &Match.Max_maps, &Match.Title, &Match.Skip_veto, &Match.Api_key, &Match.Veto_mappool, &Match.Team1_score, &Match.Team2_score, &Match.Team1_string, &Match.Team2_string, &Match.Forfeit, &Match.Plugin_version)
			if err != nil {
				panic(err)
				return nil, err
			}
			//fmt.Println(Match.id, Match.user_id, Match.server_id, Match.team1_id, Match.team2_id, Match.winner, Match.cancelled, Match.start_time, Match.end_time, Match.max_maps, Match.title, Match.skip_veto, Match.api_key, Match.veto_mappool, Match.team1_score, Match.team2_score, Match.team1_string, Match.team2_string, Match.forfeit, Match.plugin_version)
			Matches = append(Matches, Match)
		}
		return Matches, nil
	*/
}

func (s *DBdatas) MySQLGetPlayerStatsData(limit int, wherekey string, wherevalue string) ([]SQLPlayerStatsData, error) {
	PlayerStatsData := []SQLPlayerStatsData{}
	s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&PlayerStatsData)
	return PlayerStatsData, nil
	/*
		//接続でエラーが発生した場合の処理
		err := s.sql.Ping()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
		q := fmt.Sprintf("SELECT * FROM `player_stats` LIMIT %d WHERE %s", limit, where)
		fmt.Println(q)
		rows, err := s.sql.Query(q)
		if err != nil {
			panic(err.Error())
			return nil, err
		}
		defer rows.Close()

		StatsDatas := make([]SQLPlayerStatsData, 0)
		//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
		for rows.Next() {
			var StatsData SQLPlayerStatsData
			err := rows.Scan(&StatsData.Id, &StatsData.Match_id, &StatsData.Map_id, &StatsData.Team_id, &StatsData.Steam_id, &StatsData.Name, &StatsData.Kills, &StatsData.Deaths, &StatsData.Roundsplayed, &StatsData.Assists, &StatsData.Flashbang_assists, &StatsData.Teamkills, &StatsData.Suicides, &StatsData.Headshot_kills, &StatsData.Damage, &StatsData.Bomb_plants, &StatsData.Bomb_defuses, &StatsData.V1, &StatsData.V2, &StatsData.V3, &StatsData.V4, &StatsData.V5, &StatsData.K1, &StatsData.K2, &StatsData.K3, &StatsData.K4, &StatsData.K5, &StatsData.Firstdeath_Ct, &StatsData.Firstdeath_t, &StatsData.Firstkill_ct, &StatsData.Firstkill_t)
			if err != nil {
				return nil, err
				panic(err)
			}
			//fmt.Println(StatsData.id, StatsData.match_id, StatsData.map_id, StatsData.team_id, StatsData.steam_id, StatsData.name, StatsData.kills, StatsData.deaths, StatsData.roundsplayed, StatsData.assists, StatsData.flashbang_assists, StatsData.teamkills, StatsData.suicides, StatsData.headshot_kills, StatsData.damage, StatsData.bomb_plants, StatsData.bomb_defuses, StatsData.v1, StatsData.v2, StatsData.v3, StatsData.v4, StatsData.v5, StatsData.k1, StatsData.k2, StatsData.k3, StatsData.k4, StatsData.k5, StatsData.firstdeath_Ct, StatsData.firstdeath_t, StatsData.firstkill_ct, StatsData.firstkill_t)
			StatsDatas = append(StatsDatas, StatsData)
		}
		return StatsDatas, nil
	*/
}

func (s *DBdatas) MySQLGetMapStatsData(limit int, wherekey string, wherevalue string) ([]SQLMapStatsData, error) {
	MapStatsData := []SQLMapStatsData{}
	s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&MapStatsData)
	return MapStatsData, nil
	/*
		//接続でエラーが発生した場合の処理
		err := s.sql.Ping()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
		q := fmt.Sprintf("SELECT * FROM `map_stats` LIMIT %d WHERE %s", limit, where)
		fmt.Println(q)
		rows, err := s.sql.Query(q)
		if err != nil {
			panic(err.Error())
			return nil, err
		}
		defer rows.Close()

		MapStatsDatas := make([]SQLMapStatsData, 0)
		//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
		for rows.Next() {
			var MapStatsData SQLMapStatsData
			err := rows.Scan(&MapStatsData.Id, &MapStatsData.Match_id, &MapStatsData.Map_number, &MapStatsData.Map_name, &MapStatsData.Start_time, &MapStatsData.End_time, &MapStatsData.Winner, &MapStatsData.Team1_score, &MapStatsData.Team2_score)
			if err != nil {
				panic(err)
				return nil, err
			}
			//fmt.Println(MapStatsData.id, MapStatsData.match_id, MapStatsData.map_number, MapStatsData.map_name, MapStatsData.start_time, MapStatsData.end_time, MapStatsData.winner, MapStatsData.team1_score, MapStatsData.team2_score)
			MapStatsDatas = append(MapStatsDatas, MapStatsData)
		}
		return MapStatsDatas, nil
	*/
}

func (s *DBdatas) MySQLGetGameServerData(limit int, wherekey string, wherevalue string) ([]GameServerData, error) {
	GameServer := []GameServerData{}
	s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&GameServer)
	return GameServer, nil
	/*
		//接続でエラーが発生した場合の処理
		err := s.sql.Ping()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		var whr string = ""
		if len(where) > 0 {
			whr = "WHERE " + where
		}

		//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
		q := fmt.Sprintf("SELECT * FROM `game_server` %s LIMIT %d ", whr, limit)
		fmt.Println(q)
		rows, err := s.sql.Query(q)
		if err != nil {
			panic(err.Error())
			return nil, err
		}
		defer rows.Close()

		GameServerDatas := make([]GameServerData, 0)
		//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
		for rows.Next() {
			var serverdata GameServerData
			err := rows.Scan(&serverdata.Id, &serverdata.User_id, &serverdata.In_use, &serverdata.Ip_string, &serverdata.Port, &serverdata.Rcon_password, &serverdata.Display_name, &serverdata.Public_server)
			if err != nil {
				panic(err)
				return nil, err
			}
			//fmt.Println(serverdata.id, serverdata.user_id, serverdata.in_use, serverdata.ip_string, serverdata.port, serverdata.rcon_password, serverdata.display_name, serverdata.public_server)
			GameServerDatas = append(GameServerDatas, serverdata)
		}
		return GameServerDatas, nil
	*/
}

func (s *DBdatas) MySQLGetUserData(limit int, wherekey string, wherevalue string) (UserData, error) {
	UserData := UserData{}
	s.Gorm.Limit(limit).Where(wherekey+" = ?", wherevalue).Find(&UserData)
	return UserData, nil
	/*
		//接続でエラーが発生した場合の処理
		var User = SQLUserData{}

		err := s.sql.Ping()
		if err != nil {
			log.Fatal(err)
			return User, err
		}

		var whr string = ""
		if len(where) > 0 {
			whr = "WHERE " + where
		}

		//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
		q := fmt.Sprintf("SELECT * FROM `user` %s LIMIT %d ", whr, limit)
		fmt.Println(q)
		rows, err := s.sql.Query(q)
		if err != nil {
			panic(err.Error())
			return User, err
		}
		defer rows.Close()

		//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
		for rows.Next() {
			err := rows.Scan(&User.Id, &User.Steam_id, &User.Name, &User.Admin)
			if err != nil {
				panic(err)
				return User, err
			}
		}
		fmt.Println(User)
		return User, nil
	*/
}
