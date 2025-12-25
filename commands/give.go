package commands

import (
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type GiveCommand struct {
	Player   []cmd.Target      `cmd:"player"`
	ItemName ItemEnum          `cmd:"itemName"`
	Amount   cmd.Optional[int] `cmd:"amount"`
	Data     cmd.Optional[int] `cmd:"data"`
}

func (GiveCommand) Allow(src cmd.Source) bool {
	return true
}

func (c GiveCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	amount, ok := c.Amount.Load()
	if !ok {
		amount = 1
	}

	var targetItem world.Item
	for _, it := range world.Items() {
		name, _ := it.EncodeItem()
		cleanName := name
		if parts := strings.Split(name, ":"); len(parts) > 1 {
			cleanName = parts[1]
		}

		if cleanName == string(c.ItemName) {
			targetItem = it
			break
		}
	}

	if targetItem == nil {
		return
	}

	stack := item.NewStack(targetItem, amount)
	for _, target := range c.Player {
		if p, ok := target.(*player.Player); ok {
			_, err := p.Inventory().AddItem(stack)
			if err != nil {
				o.Errorf("unknown item: %s", stack.String())
			}
			o.Printf("Gave %v * %v to %v", amount, string(c.ItemName), p.Name())
		}
	}
}
