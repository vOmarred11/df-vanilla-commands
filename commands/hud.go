package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/hud"
	"github.com/df-mc/dragonfly/server/world"
)

type HudCommand struct {
	Target  []cmd.Target             `cmd:"target"`
	Action  hudAction                `cmd:"action"`
	Element cmd.Optional[hudElement] `cmd:"element"`
}

func (c HudCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	for _, t := range c.Target {
		if p, ok := t.(*player.Player); ok {
			elem, ok := c.Element.Load()
			if !ok {
				o.Errorf("Specify an element.")
				return
			}

			if elem == "all" {
				var status string
				for _, e := range hud.All() {
					if c.Action == "hide" {
						status = "hidden"
						p.HideHudElement(e)
					} else {
						status = "shown"
						p.ShowHudElement(e)
					}
				}
				p.Messagef("All hud elements %s.", status)
			} else {
				if c.Action == "hide" {
					p.HideHudElement(elem.HUDElement())
					p.Messagef("%s hidden.", elem)
				} else {
					p.ShowHudElement(elem.HUDElement())
					p.Messagef("%s resetted.", elem)
				}
			}
		}
	}
}

type hudElement string

func (hudElement) Type() string { return "element" }
func (hudElement) Options(src cmd.Source) []string {
	return []string{
		"air_bubbles", "all", "armor", "crosshair", "health",
		"horse_health", "hotbar", "hunger", "item_text",
		"paperdoll", "progress_bar", "status_effects",
		"tooltips", "touch_controls",
	}
}

func (h hudElement) HUDElement() hud.Element {
	switch h {
	case "air_bubbles":
		return hud.AirBubbles()
	case "armor":
		return hud.Armour()
	case "crosshair":
		return hud.Crosshair()
	case "health":
		return hud.Health()
	case "horse_health":
		return hud.HorseHealth()
	case "hotbar":
		return hud.HotBar()
	case "hunger":
		return hud.Hunger()
	case "item_text":
		return hud.ItemText()
	case "paperdoll":
		return hud.PaperDoll()
	case "progress_bar":
		return hud.ProgressBar()
	case "status_effects":
		return hud.StatusEffects()
	case "tooltips":
		return hud.ToolTips()
	case "touch_controls":
		return hud.TouchControls()
	}
	return hud.Element{}
}

type hudAction string

func (hudAction) Type() string { return "action" }
func (hudAction) Options(src cmd.Source) []string {
	return []string{"hide", "show"}
}
