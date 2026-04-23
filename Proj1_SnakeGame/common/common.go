package common

import (
	"time"
)

// first capital letter makes them visible outside package
const (
	GameSpeed=time.Second/6   // to control how fast game moves
	ScreenWidth=640
	ScreenHeight=480
	GridSize=20  // a 20x20 grid
)