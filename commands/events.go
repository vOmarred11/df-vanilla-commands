package commands

import (
	"net"
	"time"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type EventHandler struct {
	p *player.Player
}

func (h *EventHandler) HandleMove(ctx *player.Context, newPos mgl64.Vec3, newRot cube.Rotation) {
	frozenMu.RLock()
	frozen := frozenPlayers[h.p.Name()]
	frozenMu.RUnlock()

	if frozen {
		ctx.Cancel()
	}
}

func (h *EventHandler) HandleTeleport(ctx *player.Context, pos mgl64.Vec3) {
	frozenMu.RLock()
	frozen := frozenPlayers[h.p.Name()]
	frozenMu.RUnlock()

	if frozen {
		ctx.Cancel()
	}
}
func (h *EventHandler) HandleJump(p *player.Player) {
	frozenMu.RLock()
	frozen := frozenPlayers[h.p.Name()]
	frozenMu.RUnlock()

	if frozen {
		return
	}
}
func (h *EventHandler) HandleChat(ctx *player.Context, message *string) {

}
func (h *EventHandler) HandleChangeWorld(p *player.Player, before, after *world.World) {}

func (h *EventHandler) HandleToggleSprint(ctx *player.Context, after bool) {}

func (h *EventHandler) HandleToggleSneak(ctx *player.Context, after bool) {}

func (h *EventHandler) HandleFoodLoss(ctx *player.Context, from int, to *int) {}

func (h *EventHandler) HandleHeal(ctx *player.Context, health *float64, src world.HealingSource) {}

func (h *EventHandler) HandleHurt(ctx *player.Context, damage *float64, immune bool, attackImmunity *time.Duration, src world.DamageSource) {

}

func (h *EventHandler) HandleDeath(p *player.Player, src world.DamageSource, keepInv *bool) {}

func (h *EventHandler) HandleRespawn(p *player.Player, pos *mgl64.Vec3, w **world.World) {}

func (h *EventHandler) HandleSkinChange(ctx *player.Context, skin *skin.Skin) {}

func (h *EventHandler) HandleFireExtinguish(ctx *player.Context, pos cube.Pos) {}

func (h *EventHandler) HandleStartBreak(ctx *player.Context, pos cube.Pos) {}

func (h *EventHandler) HandleBlockPlace(ctx *player.Context, pos cube.Pos, b world.Block) {}

func (h *EventHandler) HandleBlockPick(ctx *player.Context, pos cube.Pos, b world.Block) {}

func (h *EventHandler) HandleItemUse(ctx *player.Context) {}

func (h *EventHandler) HandleItemUseOnBlock(ctx *player.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
}

func (h *EventHandler) HandleItemUseOnEntity(ctx *player.Context, e world.Entity) {}

func (h *EventHandler) HandleItemRelease(ctx *player.Context, item item.Stack, dur time.Duration) {}

func (h *EventHandler) HandleItemConsume(ctx *player.Context, item item.Stack) {}

func (h *EventHandler) HandleAttackEntity(ctx *player.Context, e world.Entity, force, height *float64, critical *bool) {
}

func (h *EventHandler) HandleExperienceGain(ctx *player.Context, amount *int) {}

func (h *EventHandler) HandlePunchAir(ctx *player.Context) {}

func (h *EventHandler) HandleSignEdit(ctx *player.Context, pos cube.Pos, frontSide bool, oldText, newText string) {
}

func (h *EventHandler) HandleLecternPageTurn(ctx *player.Context, pos cube.Pos, oldPage int, newPage *int) {
}

func (h *EventHandler) HandleItemDamage(ctx *player.Context, i item.Stack, damage *int) {}

func (h *EventHandler) HandleItemPickup(ctx *player.Context, i *item.Stack) {}

func (h *EventHandler) HandleHeldSlotChange(ctx *player.Context, from, to int) {}

func (h *EventHandler) HandleItemDrop(ctx *player.Context, s item.Stack) {}

func (h *EventHandler) HandleTransfer(ctx *player.Context, addr *net.UDPAddr) {}

func (h *EventHandler) HandleCommandExecution(ctx *player.Context, command cmd.Command, args []string) {
}

func (h *EventHandler) HandleQuit(p *player.Player) {}

func (h *EventHandler) HandleDiagnostics(p *player.Player, d session.Diagnostics) {}
func (h *EventHandler) HandleBlockBreak(ctx *player.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {

}
