package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type VoidDamageSource struct{}

func (VoidDamageSource) IgnoreTotem() bool         { return true }
func (VoidDamageSource) ReducedByArmour() bool     { return false }
func (VoidDamageSource) ReducedByResistance() bool { return false }
func (VoidDamageSource) Fire() bool                { return false }

type KillSelf struct{}

func (KillSelf) Allow(src cmd.Source) bool { return true }

func (k KillSelf) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	if p, ok := src.(*player.Player); ok {
		if p.GameMode() == world.GameModeCreative {
			p.SetGameMode(world.GameModeSurvival)
		}
		p.Hurt(1000, VoidDamageSource{})
		o.Printf("Killed %v", p.Name())
	}
}

type KillTarget struct {
	Target    []cmd.Target           `cmd:"target"`
	Anonymous cmd.Optional[BoolEnum] `cmd:"anonymous"`
}

func (KillTarget) Allow(src cmd.Source) bool { return true }

func (k KillTarget) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	var anonymous bool
	if v, ok := k.Anonymous.Load(); ok {
		anonymous = v.Raw()
	} else {
		anonymous = true
	}

	for _, target := range k.Target {
		if p, ok := target.(*player.Player); ok {
			if p.GameMode() == world.GameModeCreative {
				p.SetGameMode(world.GameModeSurvival)
			}

			p.Hurt(1000, VoidDamageSource{})

			if anonymous {
				p.Messagef("Killed by magic")
			} else {
				if i, ok := src.(*player.Player); ok {
					p.Messagef("Killed by %v", i.Name())
				} else {
					p.Messagef("Killed by the server")
				}
			}
			o.Printf("Killed %v", p.Name())
			continue
		}

		if living, ok := target.(interface {
			Hurt(damage float64, src world.DamageSource)
		}); ok {
			living.Hurt(1000, VoidDamageSource{})
			o.Printf("Killed entity")
		}
	}
}
