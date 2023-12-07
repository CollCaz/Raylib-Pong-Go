package ball

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Position rl.Vector2
	Speed    rl.Vector2
	Radius   float32
	Color    rl.Color
}

func (b *Ball) Draw() {
	rl.DrawCircle(int32(b.Position.X), int32(b.Position.Y), b.Radius, b.Color)
}

func (b *Ball) Move(delta *float32) int {
	if (b.Position.X+b.Radius >= float32(rl.GetScreenWidth())+(b.Radius)*2.0) || b.Position.X <= -b.Radius*2 {

		b.Position.X = float32(rl.GetScreenWidth() / 2.0)
		b.Position.Y = float32(rl.GetScreenHeight() / 2.0)

		b.Speed = rl.Vector2Multiply(b.Speed, rl.Vector2Negate(rl.Vector2One()))
		return 1
	}

	if (b.Position.Y+b.Radius >= float32(rl.GetScreenHeight())) ||
		(b.Position.Y <= b.Radius) {
		b.Speed = rl.Vector2Reflect(b.Speed, rl.Vector2Normalize(
			rl.Vector2{X: 0, Y: -b.Speed.Y}))
	}

	b.Position.X += b.Speed.X * *delta
	b.Position.Y += b.Speed.Y * *delta
	return 0
}

func (b *Ball) Collide(playerRect, cpuRect rl.Rectangle) {
	ns := rl.Vector2Normalize(b.Speed)

	if ns.X < 0 {
		if rl.CheckCollisionCircleRec(b.Position, b.Radius, playerRect) {
			b.Speed = rl.Vector2Reflect(b.Speed, rl.Vector2Normalize(rl.Vector2{X: ns.X, Y: 0}))
		}
	}

	if ns.X > 0 {
		if rl.CheckCollisionCircleRec(b.Position, b.Radius, cpuRect) {
			b.Speed = rl.Vector2Reflect(b.Speed, rl.Vector2Normalize(rl.Vector2{X: ns.X, Y: 0}))
		}
	}
}
