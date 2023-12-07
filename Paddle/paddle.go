package paddle

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Paddle struct {
	Position rl.Vector2
	Size     rl.Vector2
	Speed    float32
	Color    rl.Color
}

func (p *Paddle) Draw() {
	rl.DrawRectangle(int32(p.Position.X), int32(p.Position.Y), int32(p.Size.X), int32(p.Size.Y), p.Color)
}

func (p *Paddle) Control(delta *float32) {
	var up int32 = rl.KeyW
	var down int32 = rl.KeyS

	if rl.IsKeyDown(up) {
		if p.Position.Y >= 0 {
			p.Position.Y -= p.Speed * *delta
		}
	}

	if rl.IsKeyDown(down) {
		if p.Position.Y+p.Size.Y <= float32(rl.GetScreenHeight()) {
			p.Position.Y += p.Speed * *delta
		}
	}
}

func (p *Paddle) Ai(ballPos rl.Vector2, delta *float32) {
	target := rl.Vector2{X: p.Position.X, Y: ballPos.Y + float32(rl.GetRandomValue(-20, 20))}

	if ballPos.X >= float32(rl.GetScreenWidth()/2.0) {
		p.Position = rl.Vector2MoveTowards(p.Position, target, (p.Speed * *delta))
	}

	if p.Position.Y+p.Size.Y >= float32(rl.GetScreenHeight()) {
		p.Position.Y = float32(float32(rl.GetScreenHeight()) - p.Size.Y)
	}

	if (p.Position.Y + p.Size.Y) <= 0 {
		p.Position.Y = 1
	}
}

func (p *Paddle) GetRect() rl.Rectangle {
	pr := rl.Rectangle{
		X:      p.Position.X,
		Y:      p.Position.Y,
		Width:  p.Size.X,
		Height: p.Size.Y,
	}
	return pr
}
