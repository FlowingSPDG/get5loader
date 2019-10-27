package db

// MatchesPageData Struct for /matches/ page.
type MatchesPageData struct {
	LoggedIn   bool
	UserName   string
	UserID     int
	Matches    []MatchData
	AllMatches bool
	MyMatches  bool
	Owner      UserData
}

// MatchPageData Struct for /match/{matchID} page.
type MatchPageData struct {
	LoggedIn    bool
	AdminAccess bool
	Match       MatchData
}

// TeamsPageData Struct for /teams/{userID} page.
type TeamsPageData struct {
	LoggedIn   bool
	User       UserData
	IsYourTeam bool
	Teams      []TeamData
}

// TeamPageData Struct for /team/{teamID} page.
type TeamPageData struct {
	LoggedIn   bool
	IsYourTeam bool
	User       UserData
	Team       TeamData
}

// UserPageData Struct for /user/{userID} page.
type UserPageData struct {
	LoggedIn bool
	User     UserData
}

// MyserversPageData Struct for /myservers page.
type MyserversPageData struct {
	Servers  []GameServerData
	LoggedIn bool
}

// TeamCreatePageData Struct for /team/create page.
type TeamCreatePageData struct {
	LoggedIn bool
	Edit     bool
	Content  interface{} // should be template
}

// MetricsDataPage Struct for /metrics page.
type MetricsDataPage struct {
	LoggedIn bool
	Data     MetricsData
}
