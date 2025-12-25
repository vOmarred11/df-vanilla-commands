package commands

import (
	"fmt"
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type EffectEnum string

func (e EffectEnum) Type() string {
	return "effect"
}

func (e EffectEnum) Options(source cmd.Source) []string {
	return []string{
		"speed", "slowness", "haste", "mining_fatigue", "strength",
		"instant_health", "instant_damage", "jump_boost", "regeneration",
		"resistance", "fire_resistance", "water_breathing", "invisibility",
		"blindness", "night_vision", "hunger", "weakness", "poison",
		"wither", "health_boost", "absorption", "levitation",
	}
}

type EffectAction string

func (e EffectAction) Type() string {
	return "action"
}

func (e EffectAction) Options(source cmd.Source) []string {
	return []string{
		"add", "remove",
	}
}

type EffectCommand struct {
	Action    EffectAction               `cmd:"action"`
	Effect    EffectEnum                 `cmd:"effect"`
	Target    cmd.Optional[[]cmd.Target] `cmd:"player"`
	Duration  cmd.Optional[int]          `cmd:"duration"`
	Intensity cmd.Optional[int]          `cmd:"intensity"`
	Infinite  cmd.Optional[BoolEnum]     `cmd:"infinite"`
	Particles cmd.Optional[BoolEnum]     `cmd:"particles"`
}

func (e EffectCommand) Allow(src cmd.Source) bool {
	return true
}

func (e EffectCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	var targets []cmd.Target
	if t, ok := e.Target.Load(); ok {
		targets = t
	} else if p, ok := src.(*player.Player); ok {
		targets = []cmd.Target{p}
	}

	if len(targets) == 0 {
		return
	}

	eff, ok := mapEffect(e.Effect)
	if !ok {
		o.Error(fmt.Errorf("unknown effect type: %s", e.Effect))
		return
	}

	action := string(e.Action)

	for _, t := range targets {
		p, ok := t.(*player.Player)
		if !ok {
			continue
		}

		if action == "remove" {
			found := false
			for _, activeEff := range p.Effects() {
				if activeEff.Type() == eff {
					found = true
					break
				}
			}
			if !found {
				o.Error(fmt.Errorf("%s does not have the effect %s.", p.Name(), e.Effect))
				break
			}
			p.RemoveEffect(eff)
			o.Printf("Effect %s removed from %s.", e.Effect, p.Name())
			continue
		}

		if action == "add" {
			intensity := e.Intensity.LoadOr(1)

			var inf bool
			if v, ok := e.Infinite.Load(); ok {
				inf = v.Raw()
			} else {
				inf = false
			}

			if inf {
				p.AddEffect(effect.NewInfinite(eff, intensity))
				o.Printf("Effect %s added to %s for infinite with intensity %v.", e.Effect, p.Name(), intensity)
			} else {
				duration := time.Second * time.Duration(e.Duration.LoadOr(30))
				p.AddEffect(effect.New(eff, intensity, duration))
				o.Printf("Effect %s added to %s for %v with intensity %v.", e.Effect, p.Name(), duration, intensity)
			}
		}
	}
}

func mapEffect(name EffectEnum) (effect.LastingType, bool) {
	switch name {
	case "speed":
		return effect.Speed, true
	case "slowness":
		return effect.Slowness, true
	case "haste":
		return effect.Haste, true
	case "mining_fatigue":
		return effect.MiningFatigue, true
	case "strength":
		return effect.Strength, true
	case "instant_health":
		return effect.InstantHealth, true
	case "instant_damage":
		return effect.InstantDamage, true
	case "jump_boost":
		return effect.JumpBoost, true
	case "regeneration":
		return effect.Regeneration, true
	case "resistance":
		return effect.Resistance, true
	case "fire_resistance":
		return effect.FireResistance, true
	case "water_breathing":
		return effect.WaterBreathing, true
	case "invisibility":
		return effect.Invisibility, true
	case "blindness":
		return effect.Blindness, true
	case "night_vision":
		return effect.NightVision, true
	case "hunger":
		return effect.Hunger, true
	case "weakness":
		return effect.Weakness, true
	case "poison":
		return effect.Poison, true
	case "wither":
		return effect.Wither, true
	case "health_boost":
		return effect.HealthBoost, true
	case "absorption":
		return effect.Absorption, true
	case "levitation":
		return effect.Levitation, true
	}
	return nil, false
}
