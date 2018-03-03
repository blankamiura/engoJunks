package main

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// Create a Scene "DefaultScene"
type DefaultScene struct{}

// Setting DefaultScene's Type
func (*DefaultScene) Type() string { return "GameWorld" }

// Setting DefaultScene's Preload()
func (*DefaultScene) Preload() {}

// Setting DefaultScene's Setup()
// This is executed before entering main loop
func (*DefaultScene) Setup(w *ecs.World) {
  // Making RenderSystem registering RenderSystem
  // This System is neccesary for drawing background here
  w.AddSystem(&common.RenderSystem{})

  // Setting a white background
	common.SetBackground(color.White)
}

func main() {
  // Setting window options as RunOptions
	opts := engo.RunOptions{
		Title:  "Hello World",
		Width:  1024,
		Height: 640,
	}
  // Executing the engo project using DefaultScene
	engo.Run(opts, &DefaultScene{})
}
