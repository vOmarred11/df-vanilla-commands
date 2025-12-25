package commands

import (
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
)

type SoundCommand struct {
	Type      soundEnum `cmd:"type"`
	Duration  int       `cmd:"duration"`
	Intensity float64   `cmd:"intensity"`
}

func (c SoundCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	p, ok := src.(*player.Player)
	if !ok {
		return
	}
	s := c.Type.Sound()
	if c.Intensity <= 0 {
		c.Intensity = 1.0
	}

	if c.Duration <= 0 {
		p.PlaySound(s)
		o.Printf("Playing %s once.", c.Type)
		return
	}

	o.Printf("Playing %s for %v seconds with intensity %v ms.", c.Type, c.Duration, c.Intensity)

	go func() {
		interval := time.Duration(800.0/c.Intensity) * time.Millisecond
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		endTime := time.Now().Add(time.Duration(c.Duration) * time.Second)
		for time.Now().Before(endTime) {
			p.PlaySound(s)
			<-ticker.C
		}
	}()
}

type soundEnum string

func (soundEnum) Type() string { return "sound" }

func (soundEnum) Options(src cmd.Source) []string {
	return []string{
		"attack", "burp", "click", "door_crash", "door_open",
		"explosion", "extinguish", "firework_launch", "fizz",
		"glass_break", "item_break", "item_throw", "level_up",
		"orb", "pop", "potion_throw", "totem", "thunder",
	}
}

func (s soundEnum) Sound() world.Sound {
	switch string(s) {
	case "attack":
		return sound.Attack{}
	case "burp":
		return sound.Burp{}
	case "click":
		return sound.Click{}
	case "door_crash":
		return sound.DoorCrash{}
	case "door_open":
		return sound.DoorOpen{}
	case "explosion":
		return sound.Explosion{}
	case "extinguish":
		return sound.FireExtinguish{}
	case "firework_launch":
		return sound.FireworkLaunch{}
	case "fizz":
		return sound.Fizz{}
	case "glass_break":
		return sound.GlassBreak{}
	case "item_break":
		return sound.ItemBreak{}
	case "item_throw":
		return sound.ItemThrow{}
	case "level_up":
		return sound.LevelUp{}
	case "orb":
		return sound.Experience{}
	case "pop":
		return sound.Pop{}
	case "potion_throw":
		return sound.PotionBrewed{}
	case "totem":
		return sound.Totem{}
	case "thunder":
		return sound.Thunder{}
	default:
		return sound.Pop{}
	}
}
