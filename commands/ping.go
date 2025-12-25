package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type PingCommand struct {
	Target cmd.Optional[[]cmd.Target] `cmd:"player"`
}

func (PingCommand) Allow(cmd.Source) bool { return true }

func (p PingCommand) Run(src cmd.Source, o *cmd.Output, _ *world.Tx) {
	target, ok := p.Target.Load()
	pl, _ := src.(*player.Player)
	finalTarget := pl
	if ok && len(target) > 0 {
		if t, ok := target[0].(*player.Player); ok {
			finalTarget = t
		}
	}
	latency := finalTarget.Latency().Milliseconds()
	o.Printf("%s's Ping: %v", finalTarget.Name(), latency)
}
