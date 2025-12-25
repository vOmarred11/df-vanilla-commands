package commands

import (
	"fmt"
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/world"
)

type invSeeMenu struct {
	target *player.Player
	items  []item.Stack
}

func (m invSeeMenu) Submit(submitter form.Submitter, pressed form.Button, tx *world.Tx) {
	p, ok := submitter.(*player.Player)
	if !ok {
		return
	}

	if pressed.Text == "CLOSE" {
		return
	}

	shouldReopen := false

	if pressed.Text == "REMOVE ALL" {
		for _, stack := range m.target.Inventory().Items() {
			if !stack.Empty() {
				m.target.Inventory().RemoveItem(stack)
			}
		}
		p.Message(fmt.Sprintf("Cleared all items from %s's inventory.", m.target.Name()))
		shouldReopen = true
	} else {
		for _, stack := range m.items {
			rawName, _ := stack.Item().EncodeItem()
			cleanName := strings.ReplaceAll(rawName, "minecraft:", "")
			buttonText := fmt.Sprintf("%dx %s", stack.Count(), cleanName)

			if pressed.Text == buttonText {
				m.target.Inventory().RemoveItem(stack)
				p.Message(fmt.Sprintf("Removed %dx %s from %s's inventory.", stack.Count(), cleanName, m.target.Name()))
				shouldReopen = true
				break
			}
		}
	}

	if shouldReopen {
		sendInvForm(p, m.target)
	}
}

type InvSeeCommand struct {
	Target []cmd.Target `cmd:"target"`
}

func (c InvSeeCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	p, ok := src.(*player.Player)
	if !ok {
		o.Errorf("This command can only be used by a player.")
		return
	}

	for _, t := range c.Target {
		if target, ok := t.(*player.Player); ok {
			sendInvForm(p, target)
			return
		}
	}
}

func sendInvForm(p *player.Player, target *player.Player) {
	var buttons []form.Button
	var itemsShown []item.Stack

	buttons = append(buttons, form.NewButton("CLOSE", ""))
	buttons = append(buttons, form.NewButton("REMOVE ALL", ""))

	for _, stack := range target.Inventory().Items() {
		if !stack.Empty() {
			rawName, _ := stack.Item().EncodeItem()
			cleanName := strings.ReplaceAll(rawName, "minecraft:", "")
			buttonText := fmt.Sprintf("%dx %s", stack.Count(), cleanName)

			buttons = append(buttons, form.NewButton(buttonText, ""))
			itemsShown = append(itemsShown, stack)
		}
	}

	menu := form.NewMenu(invSeeMenu{target: target, items: itemsShown}, "Inventory: "+target.Name()).
		WithBody("Options:").
		WithButtons(buttons...)

	p.SendForm(menu)
}
