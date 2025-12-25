package commands

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type SetWorldSpawnPos struct {
	SpawnPos mgl64.Vec3 `cmd:"spawnPos"`
}

func (SetWorldSpawnPos) Allow(src cmd.Source) bool {
	return true
}

func (s SetWorldSpawnPos) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	pos := cube.Pos{int(s.SpawnPos.X()), int(s.SpawnPos.Y()), int(s.SpawnPos.Z())}
	tx.World().SetSpawn(pos)
	o.Printf("Set the world spawn point to %v, %v, %v", pos.X(), pos.Y(), pos.Z())
}
