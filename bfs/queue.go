package bfs

import "sokoban_solution/mapdata"

type queue struct {
	data       []*mapdata.Step
	head, tail int
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) push(step *mapdata.Step) {
	q.data = append(q.data, step)
	q.tail++
}

func (q *queue) pop() *mapdata.Step {
	if q.head == q.tail {
		return nil
	}
	step := q.data[q.head]
	q.head++
	return step
}

func (q *queue) empty() bool {
	return q.head == q.tail
}
