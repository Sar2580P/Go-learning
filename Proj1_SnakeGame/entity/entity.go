package entity

import "github.com/hajimehoshi/ebiten/v2"

const(
	TagEnemy="enemy"
	TagPlayer="player"
	TagFood="food"
)

type Entity interface { // anything part of our game world like Food, Snake etc...
	Draw(screen *ebiten.Image)
	Update(world worldView) bool
	Tag() string   // to give tags/names to entities like snake or food
}