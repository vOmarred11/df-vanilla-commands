package commands

import (
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type TransferCommand struct {
	Target  []cmd.Target `cmd:"target"`
	Address string       `cmd:"address"`
}

func (c TransferCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	for _, t := range c.Target {
		if p, ok := t.(*player.Player); ok {
			if !strings.Contains(c.Address, ":") {
				o.Error("Usage: ip:port")
				return
			}
			err := p.Transfer(c.Address)
			if err != nil {
				o.Error(err)
			}
		}
	}
}
