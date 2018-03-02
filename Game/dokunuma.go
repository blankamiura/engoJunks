package main

import (
    "engo.io/engo"
    "engo.io/ecs"
		"engo.io/engo/common"
		"image/color"

    "Dokunuma/Game/dknm/player"

    "log"
)

var (
  player  Player
)

type myScene struct {}

type Player struct {
  ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Type uniquely defines your game type
func (*myScene) Type() string { return "myGame" }


// Preload assetsのロード
func (*myScene) Preload() {
  engo.Files.Load("images/player.png")
}

func PlayerSet(world *ecs.World) {
  player                := Player{BasicEntity: ecs.NewBasic()}
  player.SpaceComponent  = common.SpaceComponent {
    Position  : engo.Point{100, 100},
    Width     : 150,
    Height    : 238,
  }

  playerTexture, err    := common.LoadedSprite("images/player.png")
  if err != nil {
    log.Println("Unable to load texture: " + err.Error())
  }

  player.RenderComponent = common.RenderComponent {
    Drawable  : playerTexture,
    Scale     : engo.Point{1, 1},
  }

  for _, system := range world.Systems() {
		switch sys := system.(type){
		case *common.RenderSystem :
				sys.Add(&player.BasicEntity,
								&player.RenderComponent,
								&player.SpaceComponent)
		}
	}
}

// メインループの前に実行される関数
// シーンに対してエンティティやシステムを追加する処理などを書く
func (*myScene) Setup(world *ecs.World) {
	world.AddSystem(&common.RenderSystem{})

	PlayerSet(world)

	common.SetBackground(color.RGBA{120, 220, 120, 255})
}


func main() {
    opts := engo.RunOptions{
        Title		: "Dokunuma",
        Width		: 1200,
        Height	: 800,
    }
    engo.Run(opts, &myScene{})
}
