package commands

import (
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/bossbar"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/df-mc/dragonfly/server/world"
)

type TitleType string

func (TitleType) Type() string { return "TitleType" }
func (TitleType) Options(src cmd.Source) []string {
	return []string{"title", "subtitle", "tip", "toast", "bossbar", "jukeboxbossbar", "clear"}
}

type TitleCommand struct {
	Target  []cmd.Target         `cmd:"target"`
	Type    TitleType            `cmd:"mode"`
	Message cmd.Optional[string] `cmd:"message"`
	FadeIn  cmd.Optional[int]    `cmd:"fadeIn"`
	Stay    cmd.Optional[int]    `cmd:"stay"`
	FadeOut cmd.Optional[int]    `cmd:"fadeOut"`
}

func (TitleCommand) Allow(src cmd.Source) bool { return true }

func (t TitleCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	mode := string(t.Type)
	msg, _ := t.Message.Load()

	fInVal, _ := t.FadeIn.Load()
	stayVal, ok := t.Stay.Load()
	if !ok {
		stayVal = 2
	}
	fOutVal, _ := t.FadeOut.Load()

	fin := time.Duration(fInVal) * time.Second
	stay := time.Duration(stayVal) * time.Second
	fout := time.Duration(fOutVal) * time.Second
	var toastTitle string
	if p, ok := src.(*player.Player); ok {
		toastTitle = p.Name()
	}
	for _, target := range t.Target {
		if p, ok := target.(*player.Player); ok {
			switch mode {
			case "title":
				p.SendTitle(title.New(msg).WithFadeInDuration(fin).WithDuration(stay).WithFadeOutDuration(fout))
			case "subtitle":
				p.SendTitle(title.New(" ").WithSubtitle(msg).WithFadeInDuration(fin).WithDuration(stay).WithFadeOutDuration(fout))
			case "tip":
				p.SendTip(msg)
			case "toast":
				if len(t.Target) > 1 {
					toastTitle = "Server"
				}
				p.SendToast(toastTitle, msg)
			case "bossbar":
				p.SendBossBar(bossbar.New(msg))
			case "jukeboxbossbar":
				p.SendJukeboxPopup(msg)
			case "clear":
				p.SendTitle(title.New("").WithFadeOutDuration(0))
				p.RemoveBossBar()
			}
			if mode == "toast" || len(t.Target) > 1 {
				o.Printf("Sent toast to %v people", CurrentPlayers)
			}
			if mode == "clear" {
				o.Printf("Cleared titles for %s", p.Name())
				return
			}
			o.Printf("Sent %s to %s", mode, p.Name())
		}
	}
}
