package mapdata

import (
	"errors"
	"sort"
)

type Map struct {
	Target PointSlice `toml:"target"`
	Map    [][]int    `toml:"map"`
	N      int
	M      int
}

// InitOnLoad 初始化
func (m *Map) InitOnLoad() error {
	sort.Sort(m.Target)
	if m.N = len(m.Map); m.N != 0 {
		m.M = len(m.Map[0])
	}
	if m.N == 0 || m.M == 0 {
		return errors.New("invalid map data")
	}
	return nil
}

// canEnter 是否可以进入
func (m *Map) canEnter(pos *Point) bool {
	if pos.X < 0 || pos.X >= m.N || pos.Y < 0 || pos.Y >= m.M {
		return false
	}
	return m.Map[pos.X][pos.Y] == 0
}

// Forward 向四个方向前进
func (m *Map) Forward(curStep *Step) (nextSteps []*Step) {
	for _, dir := range dirs {
		newPlayer := curStep.Player.forward(dir)
		if !m.canEnter(newPlayer) {
			continue
		}
		// 当前位置有箱子
		if curStep.posIsBox(newPlayer) {
			newBox := newPlayer.forward(dir)
			if !m.canEnter(newBox) {
				continue
			}
			if curStep.posIsBox(newBox) {
				continue
			}
		}
		nextStep := curStep.toNext(dir)
		nextSteps = append(nextSteps, nextStep)
	}
	return
}
