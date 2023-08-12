package query

import (
	"fmt"
	"time"

	"github.com/rumblefrog/go-a2s"
)

type queryAdapter struct {
	timeout time.Duration
}

type QueryAdapter interface {
	QueryPlayer(ip string, port int, password string) (uint8, error)
}

func NewQueryAdapter(timeout time.Duration) QueryAdapter {
	return &queryAdapter{
		timeout: timeout,
	}
}

func (q *queryAdapter) QueryPlayer(ip string, port int, password string) (uint8, error) {
	ipport := fmt.Sprintf("%s:%d", ip, port)
	client, err := a2s.NewClient(
		ipport,
		a2s.TimeoutOption(q.timeout),
	)
	if err != nil {
		return 0, err
	}
	defer client.Close()

	player, err := client.QueryPlayer()
	if err != nil {
		return 0, err
	}
	return player.Count, nil
}
