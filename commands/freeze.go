package commands

import (
	"sync"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

var (
	frozenMu      sync.RWMutex
	frozenPlayers = make(map[string]bool)
)

type FreezeCommand struct {
	Target []cmd.Target `cmd:"target"`
}

func (FreezeCommand) Allow(src cmd.Source) bool { return true }

func (f FreezeCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	frozenMu.Lock()
	defer frozenMu.Unlock()
	for _, t := range f.Target {
		if p, ok := t.(*player.Player); ok {
			name := p.Name()
			isFrozen := !frozenPlayers[name]
			frozenPlayers[name] = isFrozen
			p.Handle(&EventHandler{p: p})
			if isFrozen {
				o.Printf("%v frozed.", name)
				p.Message("You've been frozen.")
			} else {
				o.Printf("%v unfrozed.", name)
				p.Message("You've been unfrozen.")
			}
		}
	}
}
