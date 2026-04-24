package entity

import (
	"image/color"
	"main/common"
	// "main/entity"
	"main/math"
	"slices"

	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

var _ Entity = (*Player)(nil)

type Player struct{
	body []math.Point
	direction math.Point
}

// constructor for creating Player instance
func NewPlayer(start math.Point, dir math.Point) *Player{
	return &Player{
		body: []math.Point{start},
		direction: dir,
	}
}

func (p *Player) Update(worldView worldView) bool {
	head := p.body[0]   // head of the snake
	dir := p.direction   // the compiler auto-dereference the pointer (*p)
	newHead := head.Add(dir)

	// check for collision for this new head
	if newHead.IsCollision((*p).body){
		return true 
	}

	grow:=false 
	// eat the food if present
	for _, entity := range worldView.GetEntities(TagFood){
		food:= entity.(*Food)  // type assertion --> I know this interface value secretly holds a *Food underneath, give me that concrete value
		
		if newHead.Equals(food.position){
			grow=true 
			food.Respawn()
			break 			
		}
	}
	

	for _, entity:= range worldView.GetEntities(TagEnemy){
		enemy:= entity.(*Enemy)
		if slices.Contains(enemy.body, newHead){return true}
	}

	if grow {
		p.body = append(    // add new head to the snake
			[]math.Point{newHead},
			p.body...,
		)
	}else{
		// just move the snake
		p.body=append(
			[]math.Point{newHead}, 
			p.body[:len(p.body)-1]...,
		)
	}
	return false 
}

func (p *Player) Draw(screen *ebiten.Image){
	for _, pt := range p.body{
		vector.FillRect(
			screen, 
			float32(pt.X * common.GridSize),
			float32(pt.Y * common.GridSize) ,
			common.GridSize,
			common.GridSize, 
			color.White, 
			true,
		)
	}
}

func (p Player) Tag()string{
	return TagPlayer
}

func (p *Player) SetDirection(dir math.Point){  // to set direction based on key-press
	p.direction=dir
}