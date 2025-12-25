package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type WeatherEnum string

func (WeatherEnum) Type() string { return "weather" }
func (WeatherEnum) Options(source cmd.Source) []string {
	return []string{"clear", "rain", "thunder"}
}

type WeatherCommand struct {
	Weather  WeatherEnum          `cmd:"weather"`
	Duration cmd.Optional[string] `cmd:"duration"`
}

func (w WeatherCommand) Allow(src cmd.Source) bool {
	return true
}

func (w WeatherCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	currentWorld := tx.World()
	duration := 60
	if d, ok := w.Duration.Load(); ok {
		v, err := strconv.Atoi(d)
		if err == nil && v > 0 {
			duration = v
		}
	}

	dur := time.Second * time.Duration(duration)

	switch w.Weather {
	case "clear":
		currentWorld.StopRaining()
		o.Printf("Weather clear.")
	case "rain":
		currentWorld.StartRaining(dur)
		o.Printf("Weather set to rain with duration %v s.", duration)
	case "thunder":
		currentWorld.StartThundering(dur)
		o.Printf("Weather set to rain thunder duration %v s.", duration)
	default:
		o.Error(fmt.Errorf("unknown weather type %v", w.Weather))
	}
}
