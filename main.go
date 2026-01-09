package main

import (
	"ray-go/window"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LifeCycle struct{}

func (l LifeCycle) OnCleanup() {
	rl.CloseWindow()
}

func (l LifeCycle) OnStart() {

}

func (l LifeCycle) OnUpdate() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.SkyBlue)
	rl.DrawFPS(40, 40)
}

func main() {
	win := window.Window{
		Vsync:  true,
		Width:  400,
		Height: 400,
	}
	lifeCycleHandler := LifeCycle{}

	win.Init("Ray-Golang", lifeCycleHandler)
	win.Start()
}
