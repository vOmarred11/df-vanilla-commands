package commands

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type GamemodeEnum struct {
	Mode   GameMode                   `cmd:"GameMode"`
	Target cmd.Optional[[]cmd.Target] `cmd:"player"`
}

type GameModeInteger struct {
	Mode   int                        `cmd:"GameMode"`
	Target cmd.Optional[[]cmd.Target] `cmd:"player"`
}

func (g GameModeInteger) Allow(src cmd.Source) bool {
	return true
}

func (g GameModeInteger) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	pl, err := requirePlayer(src)
	if err != nil {
		o.Error(err)
		return
	}
	gm, ok := world.GameModeByID(g.Mode)
	if !ok {
		o.Error(fmt.Errorf("unknown gamemode %d", g.Mode))
		return
	}
	pl.SetGameMode(gm)
}

func (g GamemodeEnum) Allow(src cmd.Source) bool {
	return true
}

func (g GamemodeEnum) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	pl, err := requirePlayer(src)
	if err != nil {
		o.Error(err)
		return
	}
	switch g.Mode {
	case "a", "adventure":
		pl.SetGameMode(world.GameModeAdventure)
	case "c", "creative":
		pl.SetGameMode(world.GameModeCreative)
	case "d", "default":
		tx.World().DefaultGameMode()
	case "s", "survival":
		pl.SetGameMode(world.GameModeSurvival)
	case "spectator":
		pl.SetGameMode(world.GameModeSpectator)
	}
}

type GameMode string

func (g GameMode) Type() string {
	return "GameMode"
}

func (g GameMode) Options(source cmd.Source) []string {
	return []string{"a", "adventure", "c", "creative", "d", "default", "s", "spectator", "survival"}
}
