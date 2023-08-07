// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package matches_gen

import (
	"database/sql"
)

type Match struct {
	ID         string
	UserID     string
	ServerID   string
	Team1ID    string
	Team2ID    string
	Winner     sql.NullString
	Cancelled  bool
	StartTime  sql.NullTime
	EndTime    sql.NullTime
	MaxMaps    int32
	Title      string
	SkipVeto   bool
	ApiKey     string
	Team1Score uint32
	Team2Score uint32
	Forfeit    sql.NullBool
	Status     int32
}
