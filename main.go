package main

import (
	"sokoban_solution/bfs"
	"sokoban_solution/config"
)

func main() {
	config.InitFromFile("config.toml")
	bfs.BFS()
}
