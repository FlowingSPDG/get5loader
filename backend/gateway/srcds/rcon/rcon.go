package rcon

import (
	"fmt"

	"github.com/FlowingSPDG/go-steam"
)

type RCONAdapter interface {
	Execute(ip string, port int, password string, cmd string) (string, error)
}

type rconAdapter struct {
}

func NewRCONAdapter() RCONAdapter {
	return &rconAdapter{}
}

func (r *rconAdapter) Execute(ip string, port int, password string, cmd string) (string, error) {
	ipport := fmt.Sprintf("%s:%d", ip, port)
	connectOptions := &steam.ConnectOptions{RCONPassword: password}
	rcon, err := steam.Connect(ipport, connectOptions)
	if err != nil {
		return "", err
	}
	defer rcon.Close()

	resp, err := rcon.Send(cmd)
	if err != nil {
		return "", err
	}
	return resp, nil
}
