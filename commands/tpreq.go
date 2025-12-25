package commands

import (
	"sync"
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

var (
	tpaRequests = make(map[string]string)
	tpaMu       sync.Mutex
)

type TpaCommand struct {
	Target []cmd.Target `cmd:"target"`
}

func (t TpaCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	p, err := requirePlayer(src)
	if err != nil {
		return
	}
	if len(t.Target) == 0 {
		return
	}

	target, ok := t.Target[0].(*player.Player)
	if !ok || target.Name() == p.Name() {
		o.Errorf("Invalid target.")
		return
	}

	tpaMu.Lock()
	tpaRequests[target.Name()] = p.Name()
	tpaMu.Unlock()

	p.Messagef("Teleport request sent to %s.", target.Name())
	target.Messagef("%s has requested to teleport to you.", p.Name())
	time.AfterFunc(45*time.Second, func() {
		tpaMu.Lock()
		if sender, ok := tpaRequests[target.Name()]; ok && sender == p.Name() {
			delete(tpaRequests, target.Name())
			p.Message("Teleport request expired.")
		}
		tpaMu.Unlock()
	})
}

type TpAcceptCommand struct{}

func (t TpAcceptCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	p, err := requirePlayer(src)
	if err != nil {
		return
	}

	tpaMu.Lock()
	senderName, ok := tpaRequests[p.Name()]
	if !ok {
		tpaMu.Unlock()
		o.Errorf("You have no pending requests.")
		return
	}
	delete(tpaRequests, p.Name())
	tpaMu.Unlock()
	for entity := range tx.Players() {
		if sender, ok := entity.(*player.Player); ok && sender.Name() == senderName {
			sender.Teleport(p.Position())
			sender.Messagef("%s accepted your request.", p.Name())
			p.Message("Request accepted.")
			return
		}
	}
	o.Errorf("The player is no longer online.")
}

type TpDenyCommand struct{}

func (t TpDenyCommand) Run(src cmd.Source, o *cmd.Output, tx *world.Tx) {
	p, err := requirePlayer(src)
	if err != nil {
		return
	}

	tpaMu.Lock()
	senderName, ok := tpaRequests[p.Name()]
	if !ok {
		tpaMu.Unlock()
		o.Errorf("No requests to deny.")
		return
	}
	delete(tpaRequests, p.Name())
	tpaMu.Unlock()
	p.Message("Request denied.")
	for entity := range tx.Players() {
		if sender, ok := entity.(*player.Player); ok && sender.Name() == senderName {
			sender.Messagef("%s denied your teleport request.", p.Name())
		}
	}
}
