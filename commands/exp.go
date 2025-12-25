package commands

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type XPActionEnum string

func (XPActionEnum) Type() string { return "action" }
func (XPActionEnum) Options(src cmd.Source) []string {
	return []string{"add", "remove", "set"}
}

type XPCommand struct {
	Action XPActionEnum               `cmd:"action"`
	Amount int                        `cmd:"amount"`
	Target cmd.Optional[[]cmd.Target] `cmd:"target"`
}

func (XPCommand) Allow(src cmd.Source) bool { return true }

func (x XPCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	targets, ok := x.Target.Load()
	if !ok {
		p, isPlayer := src.(*player.Player)
		if !isPlayer {
			o.Errorf("You must specify a player from console.")
			return
		}
		targets = []cmd.Target{p}
	}

	for _, t := range targets {
		if p, ok := t.(*player.Player); ok {
			switch x.Action {
			case "add":
				p.AddExperience(x.Amount * 10)
				p.Message(fmt.Sprintf("Added %v levels to %s.", x.Amount, p.Name()))
			case "remove":
				p.AddExperience(-x.Amount * 10)
				p.Message(fmt.Sprintf("Removed %v levels to %s.", x.Amount, p.Name()))
			case "set":
				current := p.Experience() * 10
				p.AddExperience(x.Amount*10 - current)
				p.Message(fmt.Sprintf("%v levels set to %s.", x.Amount, p.Name()))
			}
		}
	}
}
