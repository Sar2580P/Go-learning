package main

import (
	"bytes"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	gameSpeed=time.Second/6   // to control how fast game moves
	screenWidth=640
	screenHeight=480
	gridSize=20  // a 20x20 grid
)

// keyboard control
var (
	dirUp=Point{x:0, y:-1}
	dirDown=Point{x:0, y:1}
	dirLeft=Point{x:-1, y:0}
	dirRight=Point{x:1, y:0}
	mplusFaceSource *text.GoTextFaceSource
)

type Point struct{
	x, y int
}

type Game struct {
	snake []Point
	direction Point   // how the snake moves (left, right, up, down)
	lastUpdate time.Time
	food Point
	gameOver bool
}

func (g *Game) Update() error {   // the method of the Game class/struct
	if g.gameOver { return nil }

	// get keyboard signals
	if ebiten.IsKeyPressed((ebiten.KeyW)){g.direction=dirUp}
	if ebiten.IsKeyPressed((ebiten.KeyS)){g.direction=dirDown}
	if ebiten.IsKeyPressed((ebiten.KeyA)){g.direction=dirLeft}
	if ebiten.IsKeyPressed((ebiten.KeyD)){g.direction=dirRight}


	if time.Since(g.lastUpdate)<gameSpeed{
		// don't update
		return nil
	}

	g.lastUpdate=time.Now()
	g.updateSnake(&g.snake, g.direction)  // changing snake points in-place 
	return nil 
	}    



func (g *Game) updateSnake(
	snake *[]Point, 
	direction Point,
){
	head := (*snake)[0]   // head of the snake
	newHead := Point{
		x: head.x + direction.x,
		y: head.y + direction.y,
	}

	// check for collision for this new head
	if g.isCollision(newHead, *snake){
		g.gameOver=true
		return
	}

	// food eating using collision detection
	if newHead==g.food {	
		// grow the snake
		*snake=append([]Point{newHead}, *snake...)

		// create new food
		g.spawnFood()

	}else{
		// simple snake movement
		*snake=append(
			[]Point{newHead}, 
			(*snake)[:len(*snake)-1]...,    // remove the tail and put new head
		)
	}

}

func (g Game) isCollision(
	p Point, 
	snake []Point, 
)bool{

	// head going out of bound
	if p.x<0 || p.y<0 || p.x>= screenWidth/gridSize || p.y>= screenHeight/gridSize{
		return true
	}

	// head colliding with the body
	for _, body_pt :=range snake{
		if p==body_pt {return true}
	}
	return false
}



func (g *Game) Draw(screen *ebiten.Image){

	// iterate over the snake points
	for _, p:= range g.snake{
		vector.FillRect(
			screen, 
			float32(p.x * gridSize),
			float32(p.y * gridSize) ,
			gridSize,
			gridSize, 
			color.White, 
			true,
		)
	}

	// drawing food	
	vector.FillRect(
		screen, 
		float32(g.food.x * gridSize),
		float32(g.food.y * gridSize) ,
		gridSize,
		gridSize, 
		color.RGBA{255,0,0,255}, 
		true,
	)

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
		op.GeoM.Translate(screenWidth/2-w/2, screenHeight/2 - h/2)  // places the text in centre
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
	return screenWidth, screenHeight
}

func (g *Game) spawnFood(){
	g.food=Point{
		rand.Intn(screenWidth/gridSize),
		rand.Intn(screenHeight/gridSize),
	}
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

	game:= &Game{
		snake: []Point{{
			x: (screenWidth/gridSize)/2 ,    // game initialized with snake in centre cell of the screen
			y: (screenHeight/gridSize)/2,
		}},
		direction: Point{x:1, y:0},
		gameOver: false,
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Vasuki naag")

	// Game needs to 'implement' ebiten.Game interface
	if err:= ebiten.RunGame(game) ; err != nil{
		log.Fatal(err)
	}
}