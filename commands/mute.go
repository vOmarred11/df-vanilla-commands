package commands

import (
	"fmt"
	"sync"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

var (
	muteMu       sync.RWMutex
	mutedPlayers = make(map[string]bool)
)

type MuteCommand struct {
	Target []cmd.Target           `cmd:"target"`
	State  cmd.Optional[BoolEnum] `cmd:"state"`
	Reason cmd.Optional[string]   `cmd:"reason"`
}

func (MuteCommand) Allow(src cmd.Source) bool { return true }

func (m MuteCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	muteMu.Lock()
	defer muteMu.Unlock()
	reason := "No reason provided"
	if r, ok := m.Reason.Load(); ok {
		reason = r
	}
	for _, t := range m.Target {
		if p, ok := t.(*player.Player); ok {
			name := p.Name()
			isCurrentlyMuted := mutedPlayers[name]
			var shouldMute bool
			if state, ok := m.State.Load(); ok {
				shouldMute = state == "true"
				if shouldMute && isCurrentlyMuted {
					o.Print(fmt.Sprintf("Player %s is already muted.", name))
					continue
				}
				if !shouldMute && !isCurrentlyMuted {
					o.Print(fmt.Sprintf("Player %s is not muted.", name))
					continue
				}
			} else {
				shouldMute = !isCurrentlyMuted
			}
			mutedPlayers[name] = shouldMute
			p.Handle(&EventHandler{p: p})
			if shouldMute {
				o.Print(fmt.Sprintf("Player %s has been muted. Reason: %s", name, reason))
				p.Message(fmt.Sprintf("You've been muted, reason: %s", reason))
			} else {
				o.Print(fmt.Sprintf("Player %s has been unmuted.", name))
				p.Message("You've been unmuted.")
			}
		}
	}
}
