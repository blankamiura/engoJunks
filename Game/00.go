package main

import (
    "engo.io/engo"
    "engo.io/ecs"
		"engo.io/engo/common"
		"image/color"

		"log"
)

var (
	player	Player
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

// メインループの前に実行される関数
// シーンに対してエンティティやシステムを追加する処理などを書く
func (*myScene) Setup(world *ecs.World) {
	world.AddSystem(&common.RenderSystem{})

	player 							 	 := Player{BasicEntity: ecs.NewBasic()}
	player.SpaceComponent		= common.SpaceComponent {
		Position:	engo.Point{100-20, 100-25},
		Width		:	40,
		Height	: 50,
	}

	playerTexture, err := common.LoadedSprite("images/player.png")
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}

	player.RenderComponent = common.RenderComponent {
		Drawable	: playerTexture,
		Scale			: engo.Point{1, 1},
	}

	for _, system := range world.Systems() {
		switch sys := system.(type){
		case *common.RenderSystem :
				sys.Add(&player.BasicEntity,
								&player.RenderComponent,
								&player.SpaceComponent)
		}
	}

	common.SetBackground(color.RGBA{120, 255, 120, 255})
}


func main() {
    opts := engo.RunOptions{
        Title		: "go",
        Width		: 400,
        Height	: 400,
    }
    engo.Run(opts, &myScene{})
}
