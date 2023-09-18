package mapdata

import "sort"

type Step struct {
	Player *Point     `toml:"player"` // Player 玩家位置
	Box    PointSlice `toml:"box"`    // Box   箱子位置
	prev   *Step      // prev 上一步
	Next   *Step      // Next 下一步, 最后求路径才有用
}

func (self *Step) InitOnLoad() error {
	sort.Sort(self.Box)
	return nil
}

// UniqCode 唯一码
func (self *Step) UniqCode() string {
	return self.Player.UniqCode() + self.Box.UniqCode()
}

// posIsBox 当前位置是否有箱子
func (self *Step) posIsBox(pos *Point) bool {
	for _, v := range self.Box {
		if v.X == pos.X && v.Y == pos.Y {
			return true
		}
	}
	return false
}

// toNext 下一步
func (s *Step) toNext(dir Dir) *Step {
	newStep := &Step{}
	newStep.Player = s.Player.forward(dir)
	newStep.Box = append(newStep.Box, s.Box...)
	for i := range newStep.Box {
		if newStep.Box[i].X == newStep.Player.X && newStep.Box[i].Y == newStep.Player.Y {
			newStep.Box[i] = newStep.Box[i].forward(dir)
			sort.Sort(newStep.Box)
			break
		}
	}
	newStep.prev = s
	return newStep
}

func (s *Step) BackSetNext() {
	prev := s.prev
	cur := s
	for prev != nil {
		prev.Next = cur
		prev = prev.prev
		cur = cur.prev
	}
}
