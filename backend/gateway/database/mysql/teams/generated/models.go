// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package teams_gen

import (
	"database/sql"
)

type Team struct {
	ID         int64
	UserID     int64
	Name       string
	Flag       string
	Logo       string
	Tag        string
	PublicTeam sql.NullBool
}
