package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type TpCommand struct{}

func (TpCommand) Allow(src cmd.Source) bool {
	return true
}

func (TpCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {}

type TeleportToPlayer struct {
	Destination []cmd.Target `cmd:"destination"`
}

func (t TeleportToPlayer) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	p, ok := src.(*player.Player)
	if !ok {
		return
	}
	if len(t.Destination) > 0 {
		if dest, ok := t.Destination[0].(*player.Player); ok {
			p.Teleport(dest.Position())
			o.Printf("Teleported %v to %v", p.Name(), dest.Name())
		}
	}
}

type TeleportPlayerToPlayer struct {
	Victim      []cmd.Target `cmd:"victim"`
	Destination []cmd.Target `cmd:"destination"`
}

func (t TeleportPlayerToPlayer) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	if len(t.Destination) == 0 {
		return
	}
	dest, ok := t.Destination[0].(*player.Player)
	if !ok {
		return
	}

	for _, victim := range t.Victim {
		if vic, ok := victim.(*player.Player); ok {
			vic.Teleport(dest.Position())
			o.Printf("Teleported %v to %v", vic.Name(), dest.Name())
		}
	}
}

type TeleportToPos struct {
	Destination mgl64.Vec3 `cmd:"destination"`
}

func (t TeleportToPos) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	if p, ok := src.(*player.Player); ok {
		p.Teleport(t.Destination)
		o.Printf("Teleported %v to %.2f, %.2f, %.2f", p.Name(), t.Destination.X(), t.Destination.Y(), t.Destination.Z())
	}
}

type TeleportPlayerToPos struct {
	Victim      []cmd.Target `cmd:"victim"`
	Destination mgl64.Vec3   `cmd:"destination"`
}

func (t TeleportPlayerToPos) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	for _, victim := range t.Victim {
		if vic, ok := victim.(*player.Player); ok {
			vic.Teleport(t.Destination)
			o.Printf("Teleported %v to %.2f, %.2f, %.2f", vic.Name(), t.Destination.X(), t.Destination.Y(), t.Destination.Z())
		}
	}
}
