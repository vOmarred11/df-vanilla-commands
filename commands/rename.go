package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type RenameCommand struct {
	Name string `cmd:"name"`
}

func (r RenameCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	p, err := requirePlayer(src)
	if err != nil {
		return
	}

	held, _ := p.HeldItems()
	if held.Empty() {
		o.Errorf("You must hold an item.")
		return
	}

	p.SetHeldItems(held.WithCustomName(r.Name), held)
	p.Messagef("Item renamed to: %s", r.Name)
}
