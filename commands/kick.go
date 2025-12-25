package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type KickCommand struct {
	Target []cmd.Target         `cmd:"target"`
	Reason cmd.Optional[string] `cmd:"reason"`
}

func (KickCommand) Allow(src cmd.Source) bool { return true }

func (k KickCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	reason, ok := k.Reason.Load()
	if !ok {
		reason = "Kicked by console"
	}
	for _, target := range k.Target {
		if p, ok := target.(*player.Player); ok {
			p.Disconnect(reason)
			o.Printf("Kicked %s for: %s", p.Name(), reason)
		}
	}
}
