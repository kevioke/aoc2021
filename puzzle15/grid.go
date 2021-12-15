package main

import "fmt"

func printGrid(grid [][]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Print(grid[y][x])
		}
		fmt.Println()
	}
}

func getAllAdjacent(x, y, maxX, maxY int) []pair {
	adjacent := []pair{}
	if x+1 < maxX {
		adjacent = append(adjacent, pair{x: x + 1, y: y})
	}
	if x-1 >= 0 {
		adjacent = append(adjacent, pair{x: x - 1, y: y})
	}
	if y+1 < maxY {
		adjacent = append(adjacent, pair{x: x, y: y + 1})
	}
	if y-1 >= 0 {
		adjacent = append(adjacent, pair{x: x, y: y - 1})
	}
	return adjacent
}
