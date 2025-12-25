package commands

import (
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type DifficultyEnum string

func (DifficultyEnum) Type() string { return "difficulty" }
func (DifficultyEnum) Options(source cmd.Source) []string {
	return []string{"peaceful", "easy", "normal", "hard"}
}

type DifficultyCommand struct {
	Difficulty DifficultyEnum `cmd:"difficulty"`
}

func (DifficultyCommand) Allow(src cmd.Source) bool {
	return true
}

func (c DifficultyCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	value := strings.ToLower(string(c.Difficulty))
	var diff world.Difficulty

	switch value {
	case "peaceful":
		diff = world.DifficultyPeaceful
	case "easy":
		diff = world.DifficultyEasy
	case "normal":
		diff = world.DifficultyNormal
	case "hard":
		diff = world.DifficultyHard
	default:
		return
	}

	tx.World().SetDifficulty(diff)

	o.Printf("Set game difficulty to %v", strings.Title(value))
}
