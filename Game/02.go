
package main

import (
    "engo.io/engo"
    "engo.io/ecs"
		"engo.io/engo/common"
		"image/color"

    "Dokunuma/Game/dknm/mainGame"

    "fmt"
    //"log"
)

var (
  mgScene   *mainGame.MainGameScene
)

// Registering Scene
type mainScene struct {}
// Type uniquely defines your game type
func (*mainScene) Type() string { return "Dokunuma" }

// Resitering Control System
type controlEntity struct {
	*ecs.BasicEntity
}
type ControlSystem struct {
	entities []controlEntity
}

// Preload assetsのロード
func (*mainScene) Preload() {
  //engo.Files.Load("images/player.png")
}

// メインループの前に実行される関数
// シーンに対してエンティティやシステムを追加する処理などを書く
func (*mainScene) Setup(world *ecs.World) {
	world.AddSystem(&common.RenderSystem{})
  world.AddSystem(&ControlSystem{})

  engo.Input.RegisterButton("leftWays", engo.ArrowLeft)
  engo.Input.RegisterButton("rightWays", engo.ArrowRight)
  engo.Input.RegisterButton("upWays", engo.ArrowUp)
  engo.Input.RegisterButton("downWays", engo.ArrowDown)
  engo.Input.RegisterButton("escape", engo.Escape)

	//player.PlayerSet(world, "player")

  mgScene               = &mainGame.MainGameScene{PlayerName: "player"}
  engo.RegisterScene(mgScene)

	common.SetBackground(color.RGBA{120, 220, 120, 255})
}

func (ctrlSystem *ControlSystem) Add(basic *ecs.BasicEntity) {
  ctrlSystem.entities  = append(ctrlSystem.entities, controlEntity{basic})
}

func (ctrlSystem *ControlSystem) Remove(basic ecs.BasicEntity) {
  delete := -1

	for index, e := range ctrlSystem.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		ctrlSystem.entities = append(ctrlSystem.entities[:delete], ctrlSystem.entities[delete+1:]...)
	}
}

func (ctrlSystem *ControlSystem) Update(dt float32) {
  if btn := engo.Input.Button("leftWays"); btn.Down() {
		fmt.Println("left!")
    engo.SetSceneByName("GameScene", true)
	}
  if btn := engo.Input.Button("rightWays"); btn.Down() {
    fmt.Println("right")
	}
  if btn := engo.Input.Button("upWays"); btn.Down() {
		fmt.Println("up!")
	}
  if btn := engo.Input.Button("downWays"); btn.Down() {
		fmt.Println("down!")
	}

	if btn := engo.Input.Button("escape"); btn.JustPressed() {
		engo.Exit()
	}
}


func main() {
    opts := engo.RunOptions{
        Title		: "Dokunuma",
        Width		: 1200,
        Height	: 800,
    }
    engo.Run(opts, &mainScene{})
}
