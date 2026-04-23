package entity

import (
	"main/common"
	"main/math"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

var _ Entity = (*Food)(nil)  // Food struct must satisfy Entity interface

type Food struct{
	position math.Point
}

// constructor function to create instance
func NewFood() *Food{
	return &Food{position: math.RandomPosition()}
}

// Entity interface func
func (f *Food) Update(world worldView) bool {
	return false
}

// Entity interface func
func (f *Food) Draw(screen *ebiten.Image){
	
	// drawing food	
	vector.FillRect(
		screen, 
		float32(f.position.X * common.GridSize),
		float32(f.position.Y * common.GridSize) ,
		common.GridSize,
		common.GridSize, 
		color.RGBA{255,0,0,255}, 
		true,
	)
}

// Entity interface func
func (f Food)Tag() string {
	return "food"
}

func (f Food) Respawn(){    // create a new random food
	f.position=math.RandomPosition()
}