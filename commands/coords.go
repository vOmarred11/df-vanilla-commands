package commands

import (
	"math"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type CoordsCommand struct {
	Target  cmd.Optional[[]cmd.Target] `cmd:"player"`
	Rounded cmd.Optional[BoolEnum]     `cmd:"rounded"`
}

func (CoordsCommand) Allow(cmd.Source) bool { return true }

func (c CoordsCommand) Run(src cmd.Source, o *cmd.Output, _ *world.Tx) {
	var targets []cmd.Target

	if t, ok := c.Target.Load(); ok {
		targets = t
	} else if p, ok := src.(*player.Player); ok {
		targets = append(targets, p)
	}

	if len(targets) == 0 {
		return
	}
	var rounded bool
	if val, ok := c.Rounded.Load(); ok {
		rounded = val.Raw()
	} else {
		rounded = false
	}

	for _, target := range targets {
		if p, ok := target.(*player.Player); ok {
			pos := p.Position()
			x, y, z := pos.X(), pos.Y(), pos.Z()

			if rounded {
				o.Printf("%v is located at X: %.0f, Y: %.0f, Z: %.0f", p.Name(), math.Round(x), math.Round(y), math.Round(z))
			} else {
				o.Printf("%v is located at X: %.2f, Y: %.2f, Z: %.2f", p.Name(), x, y, z)
			}
		}
	}
}
