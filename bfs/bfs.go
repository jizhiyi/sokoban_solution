package bfs

import (
	"fmt"
	"sokoban_solution/config"
)

func BFS() {
	trie := newTrie()
	trie.insert(config.GConfig.InitStep.UniqCode())
	queue := newQueue()
	queue.push(config.GConfig.InitStep)
	for !queue.empty() {
		step := queue.pop()
		if step == nil {
			break
		}
		if step.Box.UniqCode() == config.GConfig.MapData.Target.UniqCode() {
			step.BackSetNext()
			break
		}
		nextSteps := config.GConfig.MapData.Forward(step)
		for _, nextStep := range nextSteps {
			if trie.insert(nextStep.UniqCode()) {
				queue.push(nextStep)
			}
		}
	}
	count := 0
	for cur := config.GConfig.InitStep; cur != nil; cur = cur.Next {
		fmt.Println(count, " ----- ", cur.Player.X, cur.Player.Y)
	}
}
