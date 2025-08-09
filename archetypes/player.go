package archetypes

import (
	"github.com/BrianAnakPintar/ducktape/assets"
	"github.com/BrianAnakPintar/ducktape/components"
	c "github.com/BrianAnakPintar/ducktape/constants"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

func NewPlayer(w donburi.World, pos math.Vec2) {
	playerEntity := w.Create(
		components.Player,
		components.Transform,
		components.Velocity,
		components.Sprite,
		components.Jump,
		components.Collider)

	entry := w.Entry(playerEntity)
	components.Player.SetValue(entry, components.PlayerData{Name: "Ducky", Health: c.PlayerHealth})
	components.Transform.SetValue(entry, components.TransformData{Pos: pos, Rot: 0})
	components.Sprite.SetValue(entry, components.SpriteData{Image: assets.PlayerAsset0})
	components.Jump.SetValue(entry, components.JumpData{MaxJumps: 2, JumpsLeft: 2})
	components.Collider.SetValue(entry, components.ColliderData{HalfWidth: 1, HalfHeight: 1, Static: false})
}
