package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type DayLightCycleCommand struct {
	Value BoolEnum `cmd:"doLightCycle"`
}

func (DayLightCycleCommand) Allow(src cmd.Source) bool { return true }

func (d DayLightCycleCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	w := tx.World()
	if d.Value.Raw() {
		w.StartTime()
		o.Printf("DayLightCycle set to true.")
	} else {
		w.StopTime()
		o.Printf("DayLightCycle set to false.")
	}
}
