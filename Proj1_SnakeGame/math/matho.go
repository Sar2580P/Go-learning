package math

import (
	"main/common"
	"math/rand"
	"slices"
)

var (
	DirUp    = Point{X: 0, Y: -1}
	DirDown  = Point{X: 0, Y: 1}
	DirLeft  = Point{X: -1, Y: 0}
	DirRight = Point{X: 1, Y: 0}
)

type Point struct {
	X, Y int // we could construct getter-setter for these instead of capital letter
}

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

// pointer as we manipulate the point
func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}


func (p Point) Distance(q Point) int {
	return Abs(p.X-q.X)+Abs(p.Y-q.Y)
}

func Abs(num int) int {
	if num<0 {return -num}
	return num
}

func (p Point) IsCollision(points []Point) bool {

	// head going out of bound
	if p.X < 0 || p.Y < 0 || p.X >= common.ScreenWidth/common.GridSize || p.Y >= common.ScreenHeight/common.GridSize {
		return true
	}

	// head colliding with the body
	return slices.ContainsFunc(points, p.Equals)
}

func RandomPosition() Point { // for food spawning usecase
	return Point{
		X: rand.Intn(common.ScreenWidth / common.GridSize),
		Y: rand.Intn(common.ScreenHeight / common.GridSize),
	}
}

