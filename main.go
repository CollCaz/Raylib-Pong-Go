package main

import (
	b "pong/Ball"
	p "pong/Paddle"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  int = 800
	ScreenHeight int = 450
)

var ball = b.Ball{
	Position: rl.Vector2{
		X: float32(ScreenWidth / 2.0), Y: float32(ScreenHeight) / 2.0,
	},
	Speed: rl.Vector2{
		X: float32(500.0), Y: float32(500.0),
	},
	Radius: float32(20),
	Color:  rl.White,
}

var player = p.Paddle{
	Position: rl.Vector2{X: 10, Y: float32(ScreenHeight)/2.0 - 50},
	Size:     rl.Vector2{X: 20, Y: 120},
	Speed:    400,
	Color:    rl.White,
}

var ai = p.Paddle{
	Position: rl.Vector2{X: float32(ScreenWidth) - 35, Y: float32(ScreenHeight)/2.0 - 50},
	Size:     rl.Vector2{X: 20, Y: 120},
	Speed:    400,
	Color:    rl.White,
}

func main() {
	rl.InitWindow(int32(ScreenWidth), int32(ScreenHeight), "PonGo")

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawFPS(20, 20)
		// Draw Game Objects
		rl.DrawLine(int32(ScreenWidth)/2, 0, int32(ScreenWidth)/2, int32(ScreenHeight), rl.White)
		ball.Draw()
		ai.Draw()
		player.Draw()

		// Game Logic
		player.Control(&dt)
		ball.Move(&dt)
		ai.Ai(ball.Position, &dt)
		ball.Collide(player.GetRect(), ai.GetRect())
		rl.EndDrawing()

	}
}
