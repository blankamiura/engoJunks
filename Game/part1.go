package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image/color"

	"log"
)

// Create a Scene "GameScene"
type GameScene struct{}

// Create an Entity "Player"
type Player struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

// Setting GameScene's Type
func (*GameScene) Type() string { return "GameScene" }

// Setting GameScene's Preload()
func (*GameScene) Preload() {
	// Loading player texture
	engo.Files.Load("images/player.png")
}

// Setting GameScene's Setup()
// This is executed before entering main loop
func (*GameScene) Setup(world *ecs.World) {
  // Making RenderSystem registering RenderSystem
  // This System is necessary for drawing background here
  world.AddSystem(&common.RenderSystem{})

	// Setting player
	player 							 	 := Player{BasicEntity: ecs.NewBasic()}
	/// Setting SpaceComponent of player
	/// SpaceComponent defines initial position and size of Objects
	player.SpaceComponent		= common.SpaceComponent {
		Position:	engo.Point{200-40, 200-50}, // Initial position
		Width		:	80,													// Width of images
		Height	: 100,												// Height of images
	}
	/// Putting a texture on player
	playerTexture, err := common.LoadedSprite("images/player.png")
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}
	/// Setting status needed to render player
	player.RenderComponent = common.RenderComponent {
		Drawable	: playerTexture,
		Scale			: engo.Point{1, 1},
	}
	/// Making player registering RenderSystem
	for _, system := range world.Systems() {
		switch sys := system.(type){
		case *common.RenderSystem :
				sys.Add(&player.BasicEntity,
								&player.RenderComponent,
								&player.SpaceComponent)
		}
	}

  // Setting a listless green background
	common.SetBackground(color.RGBA{120, 255, 120, 255})
}

func main() {
  // Setting window options as RunOptions
	opts := engo.RunOptions{
		Title:  "Dokunuma",
		Width:  400,
		Height: 400,
	}
  // Executing the engo project using GameScene
	engo.Run(opts, &GameScene{})
}
