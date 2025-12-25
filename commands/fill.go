package commands

import (
	"math"
	"strings"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type BlockEnum string

func (BlockEnum) Type() string { return "block" }
func (BlockEnum) Options(source cmd.Source) []string {
	var names []string
	for _, b := range world.Blocks() {
		name, _ := b.EncodeBlock()
		if parts := strings.Split(name, ":"); len(parts) > 1 {
			names = append(names, parts[1])
		} else {
			names = append(names, name)
		}
	}
	return names
}

type FillCommand struct {
	From  mgl64.Vec3 `cmd:"from"`
	To    mgl64.Vec3 `cmd:"to"`
	Block BlockEnum  `cmd:"block"`
}

func (FillCommand) Allow(src cmd.Source) bool { return true }

func (f FillCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	x1, y1, z1 := int(math.Floor(f.From.X())), int(math.Floor(f.From.Y())), int(math.Floor(f.From.Z()))
	x2, y2, z2 := int(math.Floor(f.To.X())), int(math.Floor(f.To.Y())), int(math.Floor(f.To.Z()))
	minX, maxX := min(x1, x2), max(x1, x2)
	minY, maxY := min(y1, y2), max(y1, y2)
	minZ, maxZ := min(z1, z2), max(z1, z2)
	var targetBlock world.Block
	for _, b := range world.Blocks() {
		name, _ := b.EncodeBlock()
		cleanName := name
		if parts := strings.Split(name, ":"); len(parts) > 1 {
			cleanName = parts[1]
		}
		if cleanName == string(f.Block) {
			targetBlock = b
			break
		}
	}
	if targetBlock == nil {
		o.Errorf("Block %v not found", f.Block)
		return
	}
	count := 0
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				tx.SetBlock(cube.Pos{x, y, z}, targetBlock, nil)
				count++
			}
		}
	}
	o.Printf("Successfully filled %v blocks", count)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
