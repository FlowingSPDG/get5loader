package db

import (
	"database/sql" //ここでパッケージをimport
	"fmt"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/solovev/steam_go"

	//_ "html/template"
	"log"
	"net/http"
	_ "strconv"
	"time"

	"github.com/FlowingSPDG/get5-web-go/src/models"
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
	//Limit int
	sql *sql.DB
}

func (s *DBdatas) Ping() error {
	err := s.sql.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (s *DBdatas) InitOrReconnect() error {
	sqloption := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", s.User, s.Pass, s.Host, s.Port, s.Db)
	if s.sql == nil {
		s.sql, _ = sql.Open("mysql", sqloption)
	}
	err := s.sql.Ping()
	if err != nil {
		return fmt.Errorf("sql ping failure.")
	}
	s.sql.SetConnMaxLifetime(time.Second * 30)
	return nil
}

func (s *DBdatas) close() error {
	err := s.sql.Close()
	if err != nil {
		return err
	}
	return nil
}

var (
	UserDatas    = map[string]*models.UserData{}
	SteamAPIKey  = ""
	SessionStore = sessions.NewCookieStore([]byte("GET5_GO_SESSIONKEY"))
	SessionData  = "SessionData"
	DefaultPage  string
	SQLAccess    DBdatas
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
		user, err := SQLAccess.MySQLGetUserData(1, "steam_id="+steamid)
		session, _ := SessionStore.Get(r, SessionData)
		session.Options = &sessions.Options{MaxAge: 0}
		// Set some session values.
		session.Values["Loggedin"] = true
		session.Values["UserID"] = user.Id // should be get5 id
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

func (s *DBdatas) GetUserData(limit int, where string) []models.SQLUserData {
	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	err := s.sql.Ping()
	if err != nil {
		panic(err.Error())
	}
	q := fmt.Sprintf("SELECT * FROM user LIMIT %d WHERE %s", limit, where)
	fmt.Println(q)
	rows, err := s.sql.Query(q)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	Users := make([]models.SQLUserData, 0)
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var User models.SQLUserData
		err := rows.Scan(&User.Id, &User.Steam_id, &User.Name, &User.Steam_id)

		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(User.id, User.steam_id, User.name, User.steam_id) //
		Users = append(Users, User)
	}
	return Users
}

func (s *DBdatas) MySQLGetTeamData(limit int, where string) ([]models.SQLTeamData, error) {
	if s == nil {
		return nil, fmt.Errorf("sql is nil")
	}
	//接続でエラーが発生した場合の処理
	err := s.sql.Ping()
	if err != nil {
		fmt.Println(s.sql)
		log.Fatal(err)
		return nil, err
	}
	//defer s.Close()
	var whr string = ""
	if len(where) > 0 {
		whr = "WHERE " + where
	}

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	q := fmt.Sprintf("SELECT * FROM `team` %s LIMIT %d ", whr, limit)
	//fmt.Println(q)
	rows, err := s.sql.Query(q)
	if err != nil {
		return nil, err
	}
	fmt.Println(q)
	defer rows.Close()

	Teams := make([]models.SQLTeamData, 0)
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var Team models.SQLTeamData
		err := rows.Scan(&Team.Id, &Team.User_id, &Team.Name, &Team.Flag, &Team.Logo, &Team.Auth, &Team.Tag, &Team.Public_team)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(Team.Id, Team.User_id, Team.Name, Team.Flag, Team.Logo, Team.Auth, Team.Tag, Team.Public_team) //結果　1 yamada 2 suzuki
		Teams = append(Teams, Team)
	}
	return Teams, nil
}

func (s *DBdatas) MySQLGetMatchData(limit int, where string) ([]models.SQLMatchData, error) {
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

	Matches := make([]models.SQLMatchData, 0)
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var Match models.SQLMatchData
		err := rows.Scan(&Match.Id, &Match.User_id, &Match.Server_id, &Match.Team1_id, &Match.Team2_id, &Match.Winner, &Match.Cancelled, &Match.Start_time, &Match.End_time, &Match.Max_maps, &Match.Title, &Match.Skip_veto, &Match.Api_key, &Match.Veto_mappool, &Match.Team1_score, &Match.Team2_score, &Match.Team1_string, &Match.Team2_string, &Match.Forfeit, &Match.Plugin_version)
		if err != nil {
			panic(err)
			return nil, err
		}
		//fmt.Println(Match.id, Match.user_id, Match.server_id, Match.team1_id, Match.team2_id, Match.winner, Match.cancelled, Match.start_time, Match.end_time, Match.max_maps, Match.title, Match.skip_veto, Match.api_key, Match.veto_mappool, Match.team1_score, Match.team2_score, Match.team1_string, Match.team2_string, Match.forfeit, Match.plugin_version)
		Matches = append(Matches, Match)
	}
	return Matches, nil
}

func (s *DBdatas) MySQLGetPlayerStatsData(limit int, where string) ([]models.SQLPlayerStatsData, error) {
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

	StatsDatas := make([]models.SQLPlayerStatsData, 0)
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var StatsData models.SQLPlayerStatsData
		err := rows.Scan(&StatsData.Id, &StatsData.Match_id, &StatsData.Map_id, &StatsData.Team_id, &StatsData.Steam_id, &StatsData.Name, &StatsData.Kills, &StatsData.Deaths, &StatsData.Roundsplayed, &StatsData.Assists, &StatsData.Flashbang_assists, &StatsData.Teamkills, &StatsData.Suicides, &StatsData.Headshot_kills, &StatsData.Damage, &StatsData.Bomb_plants, &StatsData.Bomb_defuses, &StatsData.V1, &StatsData.V2, &StatsData.V3, &StatsData.V4, &StatsData.V5, &StatsData.K1, &StatsData.K2, &StatsData.K3, &StatsData.K4, &StatsData.K5, &StatsData.Firstdeath_Ct, &StatsData.Firstdeath_t, &StatsData.Firstkill_ct, &StatsData.Firstkill_t)
		if err != nil {
			return nil, err
			panic(err)
		}
		//fmt.Println(StatsData.id, StatsData.match_id, StatsData.map_id, StatsData.team_id, StatsData.steam_id, StatsData.name, StatsData.kills, StatsData.deaths, StatsData.roundsplayed, StatsData.assists, StatsData.flashbang_assists, StatsData.teamkills, StatsData.suicides, StatsData.headshot_kills, StatsData.damage, StatsData.bomb_plants, StatsData.bomb_defuses, StatsData.v1, StatsData.v2, StatsData.v3, StatsData.v4, StatsData.v5, StatsData.k1, StatsData.k2, StatsData.k3, StatsData.k4, StatsData.k5, StatsData.firstdeath_Ct, StatsData.firstdeath_t, StatsData.firstkill_ct, StatsData.firstkill_t)
		StatsDatas = append(StatsDatas, StatsData)
	}
	return StatsDatas, nil
}

func (s *DBdatas) MySQLGetMapStatsData(limit int, where string) ([]models.SQLMapStatsData, error) {
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

	MapStatsDatas := make([]models.SQLMapStatsData, 0)
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var MapStatsData models.SQLMapStatsData
		err := rows.Scan(&MapStatsData.Id, &MapStatsData.Match_id, &MapStatsData.Map_number, &MapStatsData.Map_name, &MapStatsData.Start_time, &MapStatsData.End_time, &MapStatsData.Winner, &MapStatsData.Team1_score, &MapStatsData.Team2_score)
		if err != nil {
			panic(err)
			return nil, err
		}
		//fmt.Println(MapStatsData.id, MapStatsData.match_id, MapStatsData.map_number, MapStatsData.map_name, MapStatsData.start_time, MapStatsData.end_time, MapStatsData.winner, MapStatsData.team1_score, MapStatsData.team2_score)
		MapStatsDatas = append(MapStatsDatas, MapStatsData)
	}
	return MapStatsDatas, nil
}

func (s *DBdatas) MySQLGetGameServerData(limit int, where string) ([]models.GameServerData, error) {
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

	GameServerDatas := make([]models.GameServerData, 0)
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var serverdata models.GameServerData
		err := rows.Scan(&serverdata.Id, &serverdata.User_id, &serverdata.In_use, &serverdata.Ip_string, &serverdata.Port, &serverdata.Rcon_password, &serverdata.Display_name, &serverdata.Public_server)
		if err != nil {
			panic(err)
			return nil, err
		}
		//fmt.Println(serverdata.id, serverdata.user_id, serverdata.in_use, serverdata.ip_string, serverdata.port, serverdata.rcon_password, serverdata.display_name, serverdata.public_server)
		GameServerDatas = append(GameServerDatas, serverdata)
	}
	return GameServerDatas, nil
}

func (s *DBdatas) MySQLGetUserData(limit int, where string) (models.SQLUserData, error) {
	//接続でエラーが発生した場合の処理
	var User = models.SQLUserData{}

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
}
