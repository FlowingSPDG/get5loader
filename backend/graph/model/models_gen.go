// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type GameServer struct {
	ID     string `json:"id"`
	IP     string `json:"Ip"`
	Port   int    `json:"port"`
	Name   string `json:"name"`
	Public bool   `json:"public"`
}

type NewGameServer struct {
	UserID       string `json:"userID"`
	IP           string `json:"Ip"`
	Port         int    `json:"port"`
	Name         string `json:"name"`
	RconPassword string `json:"RconPassword"`
	Public       bool   `json:"public"`
}

type User struct {
	ID          string        `json:"id"`
	SteamID     string        `json:"steamId"`
	Name        string        `json:"name"`
	Admin       bool          `json:"admin"`
	Gameservers []*GameServer `json:"gameservers"`
}
