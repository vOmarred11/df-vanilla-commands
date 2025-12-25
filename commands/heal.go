package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type HealCommand struct {
	Target cmd.Optional[[]cmd.Target] `cmd:"target"`
}

func (HealCommand) Allow(src cmd.Source) bool { return true }

func (h HealCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	targets, ok := h.Target.Load()
	if !ok {
		p, isPlayer := src.(*player.Player)
		if !isPlayer {
			o.Errorf("Specify a player.")
			return
		}
		targets = []cmd.Target{p}
	}

	for _, t := range targets {
		if p, ok := t.(*player.Player); ok {
			p.Heal(p.MaxHealth(), nil)
			p.SetFood(20)
			p.Message("You have been healed.")
			o.Printf("%s healed.", p.Name())
		}
	}
}
