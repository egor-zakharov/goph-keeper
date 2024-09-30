package cli

import (
	"github.com/egor-zakharov/goph-keeper/internal/client"
	"strings"
)

type Executor struct {
	client client.Client
}

func NewExecutor(client client.Client) *Executor {
	return &Executor{
		client: client,
	}
}

func (e *Executor) parseRawData(raw string) []string {
	raw = strings.TrimSpace(raw)
	return strings.Split(raw, ";")
}
