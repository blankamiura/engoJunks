package gameScene

import (
  "engo.io/engo"
  "engo.io/ecs"
  "engo.io/engo/common"
  "image/color"

  "Dokunuma/Game/Scenes/GameMode/Player"
  "Dokunuma/Game/Scenes/GameMode/PlayerBullet"

  "log"
  //"fmt"
)

var (
  world   *ecs.World
  vx      float32
  vy      float32
)

//Defining main game \
type GameScene struct{
  PlayerName    string
}
func (*GameScene) Type() string { return "GameScene" }

func (*GameScene) Preload() {}

func (gScn *GameScene) Setup(world *ecs.World) {
  // Registoring RenderSystem
  world.AddSystem(&common.RenderSystem{})
  // Getting player's texture
  engo.Files.Load("images/" + gScn.PlayerName + ".png")

  // Setting Background
  common.SetBackground(color.RGBA{255, 255, 255, 255})

  // Setting player's texture
  character   := player.PlayerSet("player")

  for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&character.BasicEntity,
              &character.RenderComponent,
              &character.SpaceComponent)
		}
	}

  world   = world

  world.AddSystem(&GameSystem{character : &character})
}

type GameSystem struct {
  character    *player.Player
  plyBullets   *playerBullet.PlayerBullet
}

func (gSys *GameSystem) Add(player *player.Player) {
  gSys.character  = player
}

func (gSys *GameSystem) PlayerBulletAdd(world *ecs.World) {
  plyBullet       := playerBullet.PlayerFire("player",
                                             gSys.character.SpaceComponent.Position.X,
                                             gSys.character.SpaceComponent.Position.Y)

  for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem :
			sys.Add(&plyBullet.BasicEntity,
              &plyBullet.RenderComponent,
              &plyBullet.SpaceComponent)
		}
	}
  log.Println("fire!")
  gSys.plyBullets  = &plyBullet
}

func (gSys *GameSystem) Remove(basic ecs.BasicEntity) {
}

func (gSys *GameSystem) BulletsRunning(dt float32) {
  gSys.plyBullets.SpaceComponent.Position.X += 10
}

func (gSys *GameSystem) Update(dt float32) {
  if btn := engo.Input.Button("leftWays"); btn.Down() {
    if gSys.character.SpaceComponent.Position.X > 0 {
      gSys.character.PlayerRun("left")
    }
  }
  if btn := engo.Input.Button("rightWays"); btn.Down() {
    if gSys.character.SpaceComponent.Position.X < 1200 - gSys.character.SpaceComponent.Width {
      gSys.character.PlayerRun("right")
    }
  }
  if btn := engo.Input.Button("upWays"); btn.Down() {
	  if gSys.character.SpaceComponent.Position.Y > 0 {
      gSys.character.PlayerRun("up")
    }
  }
  if btn := engo.Input.Button("downWays"); btn.Down() {
  	if gSys.character.SpaceComponent.Position.Y < 800 - gSys.character.SpaceComponent.Height {
      gSys.character.PlayerRun("down")
    }
  }


  if btn := engo.Input.Button("fire"); btn.JustPressed() {
    gSys.PlayerBulletAdd(world)
  }

  gSys.BulletsRunning(dt)
}
