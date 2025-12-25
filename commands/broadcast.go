package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type BroadcastCommand struct {
	Message string `cmd:"message"`
}

func (BroadcastCommand) Allow(src cmd.Source) bool { return true }

func (b BroadcastCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {

	if len(b.Message) == 0 {
		o.Errorf("You must provide a message.")
		return
	}
	fullMessage := "§d[SERVER]§r " + b.Message
	for ent := range tx.Players() {
		if p, ok := ent.(*player.Player); ok {
			p.Message(fullMessage)
		}
	}
}
