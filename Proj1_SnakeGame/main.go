package main

import (
	"bytes"
	"image/color"
	"log"
	"main/common"
	"main/entity"
	"main/game"
	"main/math"
	"time"
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// keyboard control
var (

	mplusFaceSource *text.GoTextFaceSource
)


type Game struct {
	world *game.World
	lastUpdate time.Time
	gameOver bool
}

func (g *Game) Update() error {   // the method of the Game class/struct
	if g.gameOver { return nil }

	playerRaw, ok := g.world.GetFirstEntity(entity.TagPlayer) // get the first snake

	if !ok{
		return errors.New("entity player was not found")
	}
	player:=playerRaw.(*entity.Player)  // type assert 


	// get keyboard signals
	if ebiten.IsKeyPressed((ebiten.KeyW)){player.SetDirection(math.DirUp)}
	if ebiten.IsKeyPressed((ebiten.KeyS)){player.SetDirection(math.DirDown)}
	if ebiten.IsKeyPressed((ebiten.KeyA)){player.SetDirection(math.DirLeft)}
	if ebiten.IsKeyPressed((ebiten.KeyD)){player.SetDirection(math.DirRight)}


	if time.Since(g.lastUpdate)<common.GameSpeed{
		// don't update
		return nil
	}

	g.lastUpdate=time.Now()

	// update all the entities
	for _, ent := range g.world.Entities(){
		if ent.Update(g.world){
			g.gameOver=true
			return nil
		}
	}
	return nil 
}    


func (g *Game) Draw(screen *ebiten.Image){

	for _, ent := range g.world.Entities(){
		ent.Draw(screen)
	}


	// game over scenario
	if g.gameOver{
		face:= &text.GoTextFace{
			Source: mplusFaceSource,
			Size: 48,
		}
		t:= "Khel Khatam!"
		w,h := text.Measure(
			t,
			face,
			face.Size,
		)

		op:= &text.DrawOptions{}
		op.GeoM.Translate(common.ScreenWidth/2-w/2, common.ScreenHeight/2 - h/2)  // places the text in centre
		op.ColorScale.ScaleWithColor(color.RGBA{159, 104, 20 ,255})
		text.Draw(
			screen,
			t,
			face,
			op,
		)
	}

}

func (g *Game) Layout(
	outsideWidth int, 
	outsideHeight int,
)(int, int){
	return common.ScreenWidth, common.ScreenHeight
}

func main(){
	s, err := text.NewGoTextFaceSource(
		bytes.NewReader(
			fonts.MPlus1pRegular_ttf,
		),
	)
	if err!=nil{
		log.Fatal(err)
	}
	mplusFaceSource=s

	// build the world
	world:= game.NewWorld()
	world.AddEntity(entity.NewPlayer(
		math.Point{
			X: (common.ScreenWidth/common.GridSize)/2 ,    // game initialized with snake in centre cell of the screen
			Y: (common.ScreenHeight/common.GridSize)/2,
		}, 
		math.DirUp,
	))

	world.AddEntity(entity.NewEnemy(
		math.Point{  // far from player, somewhere top-left
			X: (common.ScreenWidth/common.GridSize)/3 ,    
			Y: (common.ScreenHeight/common.GridSize)/3,
		}, 
		math.DirRight,
	))
	world.AddEntity(entity.NewFood())
	world.AddEntity(entity.NewFood())
	world.AddEntity(entity.NewFood())
	world.AddEntity(entity.NewFood())
	
	

	game:= &Game{
		world: world, 
		gameOver: false,
	}
	ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
	ebiten.SetWindowTitle("Vasuki naag")

	// Game needs to 'implement' ebiten.Game interface
	if err:= ebiten.RunGame(game) ; err != nil{
		log.Fatal(err)
	}
}