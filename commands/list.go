package commands

import (
	"strings"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

var MaxPlayers int
var CurrentPlayers int

type PlayerSource string

func (PlayerSource) Type() string { return "PlayerSource" }
func (PlayerSource) Options(src cmd.Source) []string {
	return []string{"server", "world"}
}

type ListCommand struct {
	Source PlayerSource `cmd:"playerSource"`
}

func (l ListCommand) ServerConfig(cfg server.Config) *server.Config {
	return &cfg
}
func (ListCommand) Allow(src cmd.Source) bool { return true }

func (l ListCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	source := string(l.Source)
	var onlinePlayers []*player.Player
	for e := range tx.Entities() {
		if p, ok := e.(*player.Player); ok {
			onlinePlayers = append(onlinePlayers, p)
		}
	}
	count := len(onlinePlayers)
	if source == "server" {
		o.Printf("There are %v/%v players online:", CurrentPlayers, MaxPlayers)
	} else {
		o.Printf("There are %v players online in %s world:", count, tx.World().Name())
	}
	var names []string
	for _, p := range onlinePlayers {
		names = append(names, p.Name())
	}
	if count > 0 {
		o.Printf(strings.Join(names, ", "))
	}
}
