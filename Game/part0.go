package main

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type DefaultScene struct{}

func (*DefaultScene) Type() string { return "GameWorld" }

func (*DefaultScene) Preload() {}

func (*DefaultScene) Setup(w *ecs.World) {
  w.AddSystem(&common.RenderSystem{})

	common.SetBackground(color.White)
}

func main() {
	opts := engo.RunOptions{
		Title:  "Hello World",
		Width:  1024,
		Height: 640,
	}
	engo.Run(opts, &DefaultScene{})
}
