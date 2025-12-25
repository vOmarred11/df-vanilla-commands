package commands

import (
	"sync"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/hud"
	"github.com/df-mc/dragonfly/server/world"
)

var (
	vanishMu        sync.RWMutex
	vanishedPlayers = make(map[string]bool)
)

type VanishCommand struct {
	Target cmd.Optional[[]cmd.Target] `cmd:"target"`
}

func (VanishCommand) Allow(src cmd.Source) bool { return true }

func (v VanishCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	targets, ok := v.Target.Load()

	if !ok {
		p, isPlayer := src.(*player.Player)
		if !isPlayer {
			o.Errorf("This command can only be executed by a player.")
			return
		}
		targets = []cmd.Target{p}
	}

	vanishMu.Lock()
	defer vanishMu.Unlock()

	for _, t := range targets {
		if p, ok := t.(*player.Player); ok {
			name := p.Name()
			isVanished := !vanishedPlayers[name]
			vanishedPlayers[name] = isVanished
			for ent := range tx.Players() {
				if other, isOtherPlayer := ent.(*player.Player); isOtherPlayer {
					if other.Name() == name {
						continue
					}
					if isVanished {
						other.HideEntity(p)
					} else {
						other.ShowEntity(p)
					}
				}
			}

			if isVanished {
				p.SetGameMode(world.GameModeCreative)
				for _, elem := range hud.All() {
					p.HideHudElement(elem)
				}
				o.Printf("%v is now vanished.", name)
				p.Message("Vanish enabled.")
			} else {
				p.StopFlying()
				p.SetGameMode(world.GameModeSurvival)
				for _, elem := range hud.All() {
					p.ShowHudElement(elem)
				}
				o.Printf("%v is no longer vanished.", name)
				p.Message("Vanish disabled.")
			}
		}
	}
}
