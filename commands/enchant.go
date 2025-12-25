package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type Enchantment string

func (Enchantment) Type() string { return "Enchantment" }
func (Enchantment) Options(src cmd.Source) []string {
	return []string{
		"protection", "fire_protection", "feather_falling", "blast_protection", "projectile_protection",
		"thorns", "respiration", "aqua_affinity", "depth_strider", "frost_walker", "binding_curse",
		"sharpness", "smite", "bane_of_arthropods", "knockback", "fire_aspect", "looting",
		"efficiency", "silk_touch", "unbreaking", "fortune", "power", "punch", "flame", "infinity",
		"luck_of_the_sea", "lure", "mending", "vanishing_curse", "impaling", "riptide", "loyalty",
		"channeling", "multishot", "piercing", "quick_charge", "swift_sneak",
	}
}

type EnchantCommand struct {
	Enchant Enchantment `cmd:"enchantment"`
	Level   int         `cmd:"level"`
}

func (EnchantCommand) Allow(src cmd.Source) bool {
	_, ok := src.(*player.Player)
	return ok
}

func (e EnchantCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	p := src.(*player.Player)
	main, off := p.HeldItems()

	if main.Empty() {
		o.Errorf("You must be holding an item to enchant it.")
		return
	}
	var id int
	switch string(e.Enchant) {
	case "protection":
		id = 0
	case "fire_protection":
		id = 1
	case "feather_falling":
		id = 2
	case "blast_protection":
		id = 3
	case "projectile_protection":
		id = 4
	case "thorns":
		id = 5
	case "respiration":
		id = 6
	case "aqua_affinity":
		id = 8
	case "depth_strider":
		id = 7
	case "frost_walker":
		id = 25
	case "binding_curse":
		id = 27
	case "sharpness":
		id = 9
	case "smite":
		id = 10
	case "bane_of_arthropods":
		id = 11
	case "knockback":
		id = 12
	case "fire_aspect":
		id = 13
	case "looting":
		id = 14
	case "efficiency":
		id = 15
	case "silk_touch":
		id = 16
	case "unbreaking":
		id = 17
	case "fortune":
		id = 18
	case "power":
		id = 19
	case "punch":
		id = 20
	case "flame":
		id = 21
	case "infinity":
		id = 22
	case "luck_of_the_sea":
		id = 23
	case "lure":
		id = 24
	case "mending":
		id = 26
	case "vanishing_curse":
		id = 28
	case "impaling":
		id = 29
	case "riptide":
		id = 30
	case "loyalty":
		id = 31
	case "channeling":
		id = 32
	case "multishot":
		id = 33
	case "piercing":
		id = 34
	case "quick_charge":
		id = 35
	case "swift_sneak":
		id = 37
	default:
		o.Errorf("Enchantment ID not mapped.")
		return
	}
	enchantType, ok := item.EnchantmentByID(id)
	if !ok {
		o.Errorf("Enchantment not found.")
		return
	}
	ench := item.NewEnchantment(enchantType, e.Level)
	newMain := main.WithEnchantments(ench)
	p.SetHeldItems(newMain, off)

	o.Printf("Applied %v level %v to your item.", e.Enchant, e.Level)
}
