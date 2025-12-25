package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

// BoolEnum gestisce la selezione testuale tra true e false.
type BoolEnum string

func (BoolEnum) Type() string { return "bool" }
func (BoolEnum) Options(src cmd.Source) []string {
	return []string{"true", "false"}
}
func (b BoolEnum) Raw() bool {
	switch b {
	case "true":
		return true
	case "false":
		return false
	default:
		return false
	}
}
