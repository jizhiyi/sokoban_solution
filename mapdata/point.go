package mapdata

import (
	"fmt"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p *Point) UniqCode() string {
	return fmt.Sprintf("%d-%d", p.X, p.Y)
}

func (p *Point) forward(dir Dir) *Point {
	return &Point{p.X + dir.X, p.Y + dir.Y}
}

type PointSlice []*Point

func (p PointSlice) UniqCode() string {
	builder := strings.Builder{}
	for _, v := range p {
		builder.WriteString("-" + v.UniqCode())
	}
	return builder.String()
}

func (ps PointSlice) Len() int {
	return len(ps)
}

func (ps PointSlice) Less(i, j int) bool {
	if ps[i].X == ps[j].X {
		return ps[i].Y < ps[j].Y
	}
	return ps[i].X < ps[j].X
}

func (ps PointSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
