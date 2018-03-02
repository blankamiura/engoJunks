package player

import (
  "engo.io/engo"
  "engo.io/ecs"
  "engo.io/engo/common"

  "log"
  //"fmt"
)

var (
  vx      float32
  vy      float32
  slowVx  float32
  slowVy  float32
)

type playerScene struct {}

type Player struct {
  ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}
type PlayerSystem struct {
  player *Player
}

func (*playerScene) Preload() {
  //engo.Files.Load("images/player.png")
}

func PlayerSet(world *ecs.World, playerName string) {
  // making PlayerSystem regitering world
  //world.AddSystem(&PlayerSystem{})
  //engo.Input.RegisterAxis("rightway", engo.AxisKeyPair{engo.A, engo.D})

  // loading texture of player
  engo.Files.Load("images/" + playerName + ".png")

  if playerName == "player" {
    vx    = 10.0
    vy    = 10.0
  }

  player                := Player{BasicEntity: ecs.NewBasic()}
  player.SpaceComponent  = common.SpaceComponent {
    Position  : engo.Point{600-7.5, 600-12},
    Width     : 15,
    Height    : 24,
  }

  playerTexture, err    := common.LoadedSprite("images/" + playerName + ".png")
  if err != nil {
    log.Println("Unable to load texture: " + err.Error())
  }

  player.RenderComponent = common.RenderComponent {
    Drawable  : playerTexture,
    Scale     : engo.Point{1, 1},
  }

  //plySystem.player  = player

  for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&player.BasicEntity,
              &player.RenderComponent,
              &player.SpaceComponent)
		}
	}

  world.AddSystem(&PlayerSystem{&player})
}

func (plySystem *PlayerSystem) Add(player *Player) {
  plySystem.player  = player
}

func (plySystem *PlayerSystem) Remove(basic ecs.BasicEntity) {
}

func (plySystem *PlayerSystem) Update(dt float32) {
  if btn := engo.Input.Button("leftWays"); btn.Down() {
    if plySystem.player.SpaceComponent.Position.X > 0 {
      plySystem.player.SpaceComponent.Position.X  -= vx
    }
  }
  if btn := engo.Input.Button("rightWays"); btn.Down() {
    if plySystem.player.SpaceComponent.Position.X < 1165 {
      plySystem.player.SpaceComponent.Position.X  += vx
    }
  }
  if btn := engo.Input.Button("upWays"); btn.Down() {
	  if plySystem.player.SpaceComponent.Position.Y > 0 {
      plySystem.player.SpaceComponent.Position.Y  -= vy
    }
  }
  if btn := engo.Input.Button("downWays"); btn.Down() {
  	if plySystem.player.SpaceComponent.Position.Y < 750 {
      plySystem.player.SpaceComponent.Position.Y  += vy
    }
  }
}
