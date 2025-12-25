package commands

import (
	"image/color"
	"math/rand"

	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/particle"
	"github.com/df-mc/dragonfly/server/world/sound"
)

type ParticleCommand struct {
	Type   particleEnum `cmd:"type"`
	Amount int          `cmd:"amount"`
}

func (c ParticleCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	p, ok := src.(*player.Player)
	if !ok {
		return
	}
	pos := p.Position()
	part := c.Type.Particle()
	for i := 0; i < c.Amount; i++ {
		spread := 1.5
		spawnPos := pos.Add(cube.Pos{
			int(rand.Float64()*spread*2 - spread),
			int(rand.Float64() * spread),
			int(rand.Float64()*spread*2 - spread),
		}.Vec3())
		tx.AddParticle(spawnPos, part)
	}
}

type particleEnum string

func (particleEnum) Type() string { return "type" }

func (particleEnum) Options(src cmd.Source) []string {
	return []string{
		"flame", "lava", "note", "endermanteleport", "blockbreak",
		"bonemeal", "dragoneggteleport", "dust", "dustplume",
		"eggsmash", "entityflame", "evaporate", "hugeexplosion",
		"lavadrip", "punchblock", "snowballpoof", "splash", "waterdrip",
	}
}

func (p particleEnum) Particle() world.Particle {
	switch string(p) {
	case "flame":
		return particle.Flame{}
	case "lava":
		return particle.Lava{}
	case "note":
		return particle.Note{Instrument: sound.Instrument{}, Pitch: 0}
	case "endermanteleport":
		return particle.EndermanTeleport{}
	case "blockbreak":
		return particle.BlockBreak{Block: block.Air{}}
	case "bonemeal":
		return particle.BoneMeal{}
	case "dragoneggteleport":
		return particle.DragonEggTeleport{}
	case "dust":
		return particle.Dust{Colour: color.RGBA{R: 255, G: 50, B: 50, A: 255}}
	case "dustplume":
		return particle.DustPlume{}
	case "eggsmash":
		return particle.EggSmash{}
	case "entityflame":
		return particle.EntityFlame{}
	case "evaporate":
		return particle.Evaporate{}
	case "hugeexplosion":
		return particle.HugeExplosion{}
	case "lavadrip":
		return particle.LavaDrip{}
	case "punchblock":
		return particle.PunchBlock{Block: block.Air{}}
	case "snowballpoof":
		return particle.SnowballPoof{}
	case "splash":
		return particle.Splash{}
	case "waterdrip":
		return particle.WaterDrip{}
	default:
		return particle.Flame{}
	}
}
