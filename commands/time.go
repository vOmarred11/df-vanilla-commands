package commands

import (
	"fmt"
	"strconv"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type TimeAction string

func (TimeAction) Type() string { return "action" }
func (TimeAction) Options(source cmd.Source) []string {
	return []string{"set", "add"}
}

type TimeEnum string

func (TimeEnum) Type() string { return "time" }
func (TimeEnum) Options(source cmd.Source) []string {
	return []string{"day", "night", "noon", "midnight", "sunrise", "sunset"}
}

type TimeCommand struct {
	Action TimeAction `cmd:"action"`
	Time   TimeEnum   `cmd:"time"`
}

func (t TimeCommand) Allow(src cmd.Source) bool {
	return true
}

func (t TimeCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	currentWorld := tx.World()
	var value int

	switch t.Time {
	case "sunrise":
		value = 0
	case "day":
		value = 1000
	case "noon":
		value = 6000
	case "sunset":
		value = 12000
	case "night":
		value = 13000
	case "midnight":
		value = 18000
	default:
		v, err := strconv.Atoi(string(t.Time))
		if err != nil {
			o.Error(fmt.Errorf("unknown time '%v'", t.Time))
			return
		}
		value = v
	}

	action := string(t.Action)
	if action == "set" {
		currentWorld.SetTime(value)
		o.Printf("Time set to %v.", value)
	} else if action == "add" {
		newTime := currentWorld.Time() + value
		currentWorld.SetTime(newTime)
		o.Printf("Added %v ticks (Total: %v).", value, newTime)
	}
}
