package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type TopCommand struct{}

func (t TopCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	p, err := requirePlayer(src)
	if err != nil {
		o.Errorf(err.Error())
		return
	}
	pos := p.Position()
	highestY := tx.HighestBlock(int(pos.X()), int(pos.Z()))
	p.Teleport(mgl64.Vec3{pos.X(), float64(highestY) + 1, pos.Z()})
	p.Message("Teleported to the surface.")
}
