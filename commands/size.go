package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type SizeCommand struct {
	Scale  float64                    `cmd:"scale"`
	Target cmd.Optional[[]cmd.Target] `cmd:"target"`
}

func (SizeCommand) Allow(src cmd.Source) bool { return true }

func (s SizeCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	targets, ok := s.Target.Load()
	if !ok {
		p, isPlayer := src.(*player.Player)
		if !isPlayer {
			return
		}
		targets = []cmd.Target{p}
	}

	for _, t := range targets {
		if p, ok := t.(*player.Player); ok {
			p.SetScale(s.Scale)
			p.Messagef("Sized modified to %v.", s.Scale)
			o.Printf("%s sized modified to %v.", p.Name(), s.Scale)
		}
	}
}
