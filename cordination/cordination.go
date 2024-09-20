package cordination

type Position struct {
	dimensionSize int
	dimensions    []int
}

func NewPosition(dSize int) *Position {
	return &Position{
		dimensionSize: dSize,
		dimensions:    make([]int, dSize),
	}
}

func (p *Position) getX() int {
	if p.dimensionSize < 1 {
		return -1
	}
	return p.dimensions[0]
}

func (p *Position) getY() int {
	if p.dimensionSize < 2 {
		return -1
	}
	return p.dimensions[1]
}

func (p *Position) getAt(index int) int {
	return p.dimensions[index]
}

func (p *Position) set(index, value int) {
	p.dimensions[index] = value
}

func (p *Position) setBulk(value []int) {

	for i := range p.dimensions {
		p.dimensions[i] = value[i]
	}
}
