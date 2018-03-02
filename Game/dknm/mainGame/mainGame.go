package mainGame

import (
  "engo.io/engo"
  "engo.io/ecs"
  "engo.io/engo/common"
  "image/color"

  "log"
  //"fmt"
)

var (
  vx      float32
  vy      float32
  slowVx  float32
  slowVy  float32
)

//Defining main game \
type MainGameScene struct{
  PlayerName    string
}
func (*MainGameScene) Type() string { return "GameScene" }

func (*MainGameScene) Preload() {}

func (mgs *MainGameScene) Setup(world *ecs.World) {
  // Registoring RenderSystem
  world.AddSystem(&common.RenderSystem{})
  // Getting player's texture
  engo.Files.Load("images/" + mgs.PlayerName + ".png")

  // Setting Background
  common.SetBackground(color.RGBA{255, 255, 255, 255})

  // Setting player's status
  if mgs.PlayerName == "player" {
    vx    = 10.0
    vy    = 10.0
  }

  // Setting player's texture
  player                := Player{BasicEntity: ecs.NewBasic()}
  player.SpaceComponent  = common.SpaceComponent {
    Position  : engo.Point{600-7.5, 600-12},
    Width     : 15,
    Height    : 24,
  }
  playerTexture, err    := common.LoadedSprite("images/" + mgs.PlayerName + ".png")
  if err != nil {
    log.Println("Unable to load texture: " + err.Error())
  }
  player.RenderComponent = common.RenderComponent {
    Drawable  : playerTexture,
    Scale     : engo.Point{1, 1},
  }

  //mgs.player  = player

  for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&player.BasicEntity,
              &player.RenderComponent,
              &player.SpaceComponent)
		}
	}

  world.AddSystem(&MainGameSystem{&player})

}

type Player struct {
  ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}
type playerBullets struct {
  *ecs.BasicEntity
}
type MainGameSystem struct {
  player    *Player
  pBullets  []playerBullets
}

func (mgs *MainGameSystem) Add(player *Player) {
  mgs.player  = player
}

func (mgs *MainGameSystem) Remove(basic ecs.BasicEntity) {
}

func (mgs *MainGameSystem) Update(dt float32) {
  if btn := engo.Input.Button("leftWays"); btn.Down() {
    if mgs.player.SpaceComponent.Position.X > 0 {
      mgs.player.SpaceComponent.Position.X  -= vx
    }
  }
  if btn := engo.Input.Button("rightWays"); btn.Down() {
    if mgs.player.SpaceComponent.Position.X < 1165 {
      mgs.player.SpaceComponent.Position.X  += vx
    }
  }
  if btn := engo.Input.Button("upWays"); btn.Down() {
	  if mgs.player.SpaceComponent.Position.Y > 0 {
      mgs.player.SpaceComponent.Position.Y  -= vy
    }
  }
  if btn := engo.Input.Button("downWays"); btn.Down() {
  	if mgs.player.SpaceComponent.Position.Y < 750 {
      mgs.player.SpaceComponent.Position.Y  += vy
    }
  }
}
