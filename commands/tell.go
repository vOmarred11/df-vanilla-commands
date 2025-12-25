package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type TellCommand struct {
	Target  []cmd.Target `cmd:"target"`
	Message string       `cmd:"message"`
}

func (TellCommand) Allow(src cmd.Source) bool { return true }

func (t TellCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	senderName := "Unknown"
	if p, ok := src.(*player.Player); ok {
		senderName = p.Name()
	}
	for _, target := range t.Target {
		if p, ok := target.(*player.Player); ok {
			if p.Name() == senderName {
				o.Error("You cannot whisper yourself.")
				return
			}
			p.Messagef("§7%v: %v", senderName, t.Message)
			o.Printf("§7%v: %v", p.Name(), t.Message)
		} else {
			o.Errorf("The target must be a player.")
		}
	}
}
