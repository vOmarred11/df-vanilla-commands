package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type DamageCommand struct {
	Target []cmd.Target `cmd:"target"`
	Amount float64      `cmd:"amount"`
}

func (DamageCommand) Allow(src cmd.Source) bool { return true }

func (d DamageCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	for _, target := range d.Target {
		if p, ok := target.(*player.Player); ok {
			p.Hurt(d.Amount, VoidDamageSource{})
			o.Printf("Applied %v damage to %v", d.Amount, p.Name())
		}
	}
}
