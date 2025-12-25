package commands

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

func RegisterAll() {
	cmd.Register(cmd.New("ping", "Show you current ping.", nil, PingCommand{}))
	cmd.Register(cmd.New("gamemode", "Switch game mode.", nil,
		GamemodeEnum{}, GameModeInteger{},
	))
	cmd.Register(cmd.New("effect", "Add or remove effects.", nil, EffectCommand{}))
	cmd.Register(cmd.New("time", "Set a time in the server.", nil, TimeCommand{}))
	cmd.Register(cmd.New("weather", "Change weather in the server", nil, WeatherCommand{}))
	cmd.Register(cmd.New("teleport", "Teleport to a position or a player.", []string{"tp"}, TpCommand{},
		TeleportToPos{},
		TeleportToPlayer{},
		TeleportPlayerToPos{},
		TeleportPlayerToPlayer{}),
	)
	cmd.Register(cmd.New("clear", "Clears items from player inventory.", nil,
		ClearCommand{},
		ClearTarget{},
		ClearFull{},
	))
	cmd.Register(cmd.New("difficulty", "Sets the game difficulty.", nil, DifficultyCommand{}))
	cmd.Register(cmd.New("give", "Gives an item to a player.", nil, GiveCommand{}))
	cmd.Register(cmd.New("setworldspawn", "Sets the world spawn.", nil, SetWorldSpawnPos{}))
	cmd.Register(cmd.New("coords", "Returns target coordinates.", nil, CoordsCommand{}))
	cmd.Register(cmd.New("kill", "Kills entities.", nil,
		KillSelf{},
		KillTarget{},
	))
	cmd.Register(cmd.New("list", "Lists online players.", nil, ListCommand{}))
	cmd.Register(cmd.New("title", "Send title with different action types.", nil, TitleCommand{}))
	cmd.Register(cmd.New("tell", "Send a private message to a target.", []string{"w", "msg", "whisper"}, TellCommand{}))
	cmd.Register(cmd.New("damage", "Apply damage target.", nil, DamageCommand{}))
	cmd.Register(cmd.New("enchant", "Enchant held item.", nil, EnchantCommand{}))
	cmd.Register(cmd.New("kick", "Kicks a player from the server.", nil, KickCommand{}))
	cmd.Register(cmd.New("fill", "Fill an area.", nil, FillCommand{}))
	cmd.Register(cmd.New("daylightcycle", "Manage day light cycle.", nil, DayLightCycleCommand{}))
	cmd.Register(cmd.New("freeze", "Freeze the target.", nil, FreezeCommand{}))
	cmd.Register(cmd.New("broadcast", "Send a message to everyone.", []string{"announce"}, BroadcastCommand{}))
	cmd.Register(cmd.New("size", "Change target scale.", nil, SizeCommand{}))
	cmd.Register(cmd.New("heal", "Give health to the target.", nil, HealCommand{}))
	cmd.Register(cmd.New("vanish", "Makes you invisibile and gives you the ability to fly", nil, VanishCommand{}))
	cmd.Register(cmd.New("xp", "Give or remove xp of the target.", nil, XPCommand{}))
	cmd.Register(cmd.New("rename", "Rename the item in your hand.", nil, RenameCommand{}))
	cmd.Register(cmd.New("top", "Teleport to the surface.", nil, TopCommand{}))
	cmd.Register(cmd.New("tpa", "Request to teleport to a player.", nil, TpaCommand{}))
	cmd.Register(cmd.New("tpaccept", "Accept a teleport request.", nil, TpAcceptCommand{}))
	cmd.Register(cmd.New("tpdeny", "Deny a teleport request.", nil, TpDenyCommand{}))
	cmd.Register(cmd.New("mute", "Mutes or unmutes a player.", nil, MuteCommand{}))
	cmd.Register(cmd.New("invsee", "View player inventory.", nil, InvSeeCommand{}))
	cmd.Register(cmd.New("hud", "Hide or reset HUD elements.", nil, HudCommand{}))
	cmd.Register(cmd.New("transfer", "Sposta i giocatori in un altro server", nil, TransferCommand{}))
	cmd.Register(cmd.New("particle", "Summons particles in the world.", nil, ParticleCommand{}))
	cmd.Register(cmd.New("playsound", "Play minecraft sounds.", nil, SoundCommand{}))

}

func requirePlayer(src interface{}) (*player.Player, error) {
	pl, ok := src.(*player.Player)
	if !ok {
		return nil, fmt.Errorf("only players can use this command")
	}
	return pl, nil
}
