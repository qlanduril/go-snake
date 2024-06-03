package util

type DiscreteCordinate2D interface {
	GetX()
	GetY()
}

type Position struct {
	x, y int
}

func (p *Position) NewPostion(newx, newy int) *Position {

	s := Position{
		x: newx,
		y: newy,
	}

	return &s
}

func (p *Position) GetX() int { return p.x }

func (p *Position) GetY() int { return p.y }
