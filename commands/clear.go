package commands

import (
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type ItemEnum string

func (ItemEnum) Type() string { return "item" }
func (ItemEnum) Options(source cmd.Source) []string {
	var names []string
	for _, it := range world.Items() {
		name, _ := it.EncodeItem()
		if parts := strings.Split(name, ":"); len(parts) > 1 {
			names = append(names, parts[1])
		} else {
			names = append(names, name)
		}
	}
	return names
}

type ClearCommand struct{}

func (ClearCommand) Allow(src cmd.Source) bool { return true }
func (ClearCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	if p, ok := src.(*player.Player); ok {
		p.Inventory().Clear()
		o.Printf("Cleared the inventory of %v", p.Name())
	}
}

type ClearTarget struct {
	Player []cmd.Target `cmd:"player"`
}

func (c ClearTarget) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	for _, target := range c.Player {
		if p, ok := target.(*player.Player); ok {
			p.Inventory().Clear()
			o.Printf("Cleared the inventory of %v", p.Name())
		}
	}
}

type ClearFull struct {
	Player   []cmd.Target      `cmd:"player"`
	ItemName ItemEnum          `cmd:"itemName"`
	Data     cmd.Optional[int] `cmd:"data"`
	MaxCount cmd.Optional[int] `cmd:"maxCount"`
}

func (c ClearFull) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	maxCount, hasMaxCount := c.MaxCount.Load()
	if !hasMaxCount {
		maxCount = -1
	}
	for _, target := range c.Player {
		if p, ok := target.(*player.Player); ok {
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
				continue
			}
			count := 0
			searchStack := item.NewStack(targetItem, 1)
			for _, slot := range p.Inventory().Slots() {
				if !slot.Empty() && slot.Comparable(searchStack) {
					if maxCount != -1 && count >= maxCount {
						break
					}
					err := p.Inventory().RemoveItem(slot)
					if err != nil {
						o.Errorf("unknown item %v", slot)
						return
					}
					count += slot.Count()
				}
			}
			o.Printf("Cleared %v items from %v", count, p.Name())
		}
	}
}
