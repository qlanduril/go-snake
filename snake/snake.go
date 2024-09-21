package snake

import (
	//"fmt"
	cord "github.com/qlanduril/go-snake/cordination"
)

const (
	HEAD_DIRECTION_NULL  = 0
	HEAD_DIRECTION_UP    = 1
	HEAD_DIRECTION_DOWN  = 2
	HEAD_DIRECTION_LEFT  = 3
	HEAD_DIRECTION_RIGHT = 4
)

type Snake struct {
	length        int
	headDirection int
	body          []*cord.Position
}

func NewSnake(dSize int, headPos *cord.Position) *Snake {

	tmp := make([]*cord.Position, 1)

	if headPos == nil {
		tmp[0] = cord.NewPosition(dSize)
	} else {
		tmp[0] = headPos
	}

	return &Snake{
		length:        1,
		headDirection: HEAD_DIRECTION_NULL,
		body:          tmp,
	}
}
