package get5

import (
	"database/sql" //ここでパッケージをimport
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/solovev/steam_go"
	_ "html/template"
	"log"
	"net/http"
	_ "strconv"
	_ "time"
)

type HomeData struct {
	LoggedIn bool
	Content  interface{} // should be template
	UserName string
	UserID   string
}

type Config struct {
	SteamAPIKey string
	DefaultPage string
	SQLHost     string
	SQLPass     string
	SQLDBName   string
}

var (
	Cnf          Config
	UserDatas    = map[string]*UserData{}
	SteamAPIKey  = ""
	SessionStore = sessions.NewCookieStore([]byte("GET5_GO_SESSIONKEY"))
	SessionData  = "SessionData"
	DefaultPage  string
)

func init() {
	c, _ := ini.Load("config.ini")
	Cnf = Config{
		SteamAPIKey: c.Section("Steam").Key("APIKey").MustString(""),
		DefaultPage: c.Section("GET5").Key("DefaultPage").MustString(""),
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
		user, err := opID.ValidateAndGetUser(SteamAPIKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Printf("user : %v", *user)
		//steam_go.GetPlayerSummaries
		if err != nil {
			panic(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if user == nil {
			panic(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Printf("\nUserName : %s\n", user.PersonaName)
		UserDatas[user.SteamId] = &UserData{
			SteamID: user.SteamId,
			Name:    user.PersonaName,
		}
		session, _ := SessionStore.Get(r, SessionData)
		session.Options = &sessions.Options{MaxAge: 0}
		// Set some session values.
		session.Values["Loggedin"] = true
		session.Values["UserID"] = user.SteamId
		session.Values["Name"] = user.PersonaName
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

type MySQLConf struct {
	host  string
	user  string
	pass  string
	db    string
	port  int
	limit int
}

func MySQLGetUserData(conf MySQLConf) []SQLUserData {
	sqloption := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.user, conf.pass, conf.host, conf.port, conf.db)
	fmt.Println(sqloption)

	s, err := sql.Open("mysql", sqloption)
	log.Println("Connected to mysql.")

	//接続でエラーが発生した場合の処理
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	rows, err := s.Query("SELECT * FROM user")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	Users := make([]SQLUserData, 0)
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var User SQLUserData
		err := rows.Scan(&User.id, &User.steam_id, &User.name, &User.steam_id)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(User.id, User.steam_id, User.name, User.steam_id) //結果　1 yamada 2 suzuki
		Users = append(Users, User)
	}
	return Users
}

func MySQLGetTeamData(conf MySQLConf) []SQLTeamData {
	sqloption := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.user, conf.pass, conf.host, conf.port, conf.db)
	fmt.Println(sqloption)

	s, err := sql.Open("mysql", sqloption)
	log.Println("Connected to mysql.")

	//接続でエラーが発生した場合の処理
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	rows, err := s.Query("SELECT * FROM team")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	Teams := make([]SQLTeamData, 0)
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var Team SQLTeamData
		err := rows.Scan(&Team.id, &Team.user_id, &Team.name, &Team.flag, &Team.logo, &Team.auth, &Team.tag, &Team.public_team)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(Team.id, Team.user_id, Team.name, Team.flag, Team.logo, Team.auth, Team.tag, Team.public_team) //結果　1 yamada 2 suzuki
		Teams = append(Teams, Team)
	}
	return Teams
}

func MySQLGetMatchData(conf MySQLConf) []SQLMatchData {
	sqloption := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.user, conf.pass, conf.host, conf.port, conf.db)
	fmt.Println(sqloption)

	s, err := sql.Open("mysql", sqloption)
	log.Println("Connected to mysql.")

	//接続でエラーが発生した場合の処理
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	rows, err := s.Query("SELECT * FROM `match`")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	Matches := make([]SQLMatchData, 0)
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var Match SQLMatchData
		err := rows.Scan(&Match.id, &Match.user_id, &Match.server_id, &Match.team1_id, &Match.team2_id, &Match.winner, &Match.cancelled, &Match.start_time, &Match.end_time, &Match.max_maps, &Match.title, &Match.skip_veto, &Match.api_key, &Match.veto_mappool, &Match.team1_score, &Match.team2_score, &Match.team1_string, &Match.team2_string, &Match.forfeit, &Match.plugin_version)
		if err != nil {
			panic(err)
		}
		fmt.Println(Match.id, Match.user_id, Match.server_id, Match.team1_id, Match.team2_id, Match.winner, Match.cancelled, Match.start_time, Match.end_time, Match.max_maps, Match.title, Match.skip_veto, Match.api_key, Match.veto_mappool, Match.team1_score, Match.team2_score, Match.team1_string, Match.team2_string, Match.forfeit, Match.plugin_version)
		Matches = append(Matches, Match)
	}
	return Matches
}

func MySQLGetPlayerStatsData(conf MySQLConf) []SQLPlayerStatsData {
	sqloption := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.user, conf.pass, conf.host, conf.port, conf.db)
	fmt.Println(sqloption)

	s, err := sql.Open("mysql", sqloption)
	log.Println("Connected to mysql.")

	//接続でエラーが発生した場合の処理
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	rows, err := s.Query("SELECT * FROM `player_stats`")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	StatsDatas := make([]SQLPlayerStatsData, 0)
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var StatsData SQLPlayerStatsData
		err := rows.Scan(&StatsData.id, &StatsData.match_id, &StatsData.map_id, &StatsData.team_id, &StatsData.steam_id, &StatsData.name, &StatsData.kills, &StatsData.deaths, &StatsData.roundsplayed, &StatsData.assists, &StatsData.flashbang_assists, &StatsData.teamkills, &StatsData.suicides, &StatsData.headshot_kills, &StatsData.damage, &StatsData.bomb_plants, &StatsData.bomb_defuses, &StatsData.v1, &StatsData.v2, &StatsData.v3, &StatsData.v4, &StatsData.v5, &StatsData.k1, &StatsData.k2, &StatsData.k3, &StatsData.k4, &StatsData.k5, &StatsData.firstdeath_Ct, &StatsData.firstdeath_t, &StatsData.firstkill_ct, &StatsData.firstkill_t)
		if err != nil {
			panic(err)
		}
		fmt.Println(StatsData.id, StatsData.match_id, StatsData.map_id, StatsData.team_id, StatsData.steam_id, StatsData.name, StatsData.kills, StatsData.deaths, StatsData.roundsplayed, StatsData.assists, StatsData.flashbang_assists, StatsData.teamkills, StatsData.suicides, StatsData.headshot_kills, StatsData.damage, StatsData.bomb_plants, StatsData.bomb_defuses, StatsData.v1, StatsData.v2, StatsData.v3, StatsData.v4, StatsData.v5, StatsData.k1, StatsData.k2, StatsData.k3, StatsData.k4, StatsData.k5, StatsData.firstdeath_Ct, StatsData.firstdeath_t, StatsData.firstkill_ct, StatsData.firstkill_t)
		StatsDatas = append(StatsDatas, StatsData)
	}
	return StatsDatas
}

func MySQLGetMapStatsData(conf MySQLConf) []SQLMapStatsData {
	sqloption := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.user, conf.pass, conf.host, conf.port, conf.db)
	fmt.Println(sqloption)

	s, err := sql.Open("mysql", sqloption)
	log.Println("Connected to mysql.")

	//接続でエラーが発生した場合の処理
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	rows, err := s.Query("SELECT * FROM `map_stats`")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	MapStatsDatas := make([]SQLMapStatsData, 0)
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var MapStatsData SQLMapStatsData
		err := rows.Scan(&MapStatsData.id, &MapStatsData.match_id, &MapStatsData.map_number, &MapStatsData.map_name, &MapStatsData.start_time, &MapStatsData.end_time, &MapStatsData.winner, &MapStatsData.team1_score, &MapStatsData.team2_score)
		if err != nil {
			panic(err)
		}
		fmt.Println(MapStatsData.id, MapStatsData.match_id, MapStatsData.map_number, MapStatsData.map_name, MapStatsData.start_time, MapStatsData.end_time, MapStatsData.winner, MapStatsData.team1_score, MapStatsData.team2_score)
		MapStatsDatas = append(MapStatsDatas, MapStatsData)
	}
	return MapStatsDatas
}

func MySQLGetGameServerData(conf MySQLConf) []GameServerData {
	sqloption := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.user, conf.pass, conf.host, conf.port, conf.db)
	fmt.Println(sqloption)

	s, err := sql.Open("mysql", sqloption)
	log.Println("Connected to mysql.")

	//接続でエラーが発生した場合の処理
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
	rows, err := s.Query("SELECT * FROM `game_server`")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	GameServerDatas := make([]GameServerData, 0)
	//レコード一件一件をあらかじめ用意しておいた構造体に当てはめていく。
	for rows.Next() {
		var serverdata GameServerData
		err := rows.Scan(&serverdata.id, &serverdata.user_id, &serverdata.in_use, &serverdata.ip_string, &serverdata.port, &serverdata.rcon_password, &serverdata.display_name, &serverdata.public_server)
		if err != nil {
			panic(err)
		}
		fmt.Println(serverdata.id, serverdata.user_id, serverdata.in_use, serverdata.ip_string, serverdata.port, serverdata.rcon_password, serverdata.display_name, serverdata.public_server)
		GameServerDatas = append(GameServerDatas, serverdata)
	}
	return GameServerDatas
}
