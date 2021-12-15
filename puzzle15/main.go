package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

type pair struct {
	x, y int
}

type puzzleInput struct {
	grid [][]int
}

func readInput(r io.Reader) (*puzzleInput, error) {
	scanner := bufio.NewScanner(r)
	grid := [][]int{}
	for scanner.Scan() {
		row := []int{}
		for _, cell := range scanner.Text() {
			num, err := strconv.Atoi(string(cell))
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	return &puzzleInput{grid: grid}, nil
}

func visitMin(distance [][]int, unvisited []pair) (pair, []pair) {
	min := -1
	var minIndex int
	for index, node := range unvisited {
		node := node
		if distance[node.y][node.x] < min || min == -1 {
			minIndex = index
			min = distance[node.y][node.x]
		}
	}

	minNode := unvisited[minIndex]
	unvisited = append(unvisited[:minIndex], unvisited[minIndex+1:]...)

	return minNode, unvisited
}

func containsPair(pairs []pair, target pair) bool {
	for _, p := range pairs {
		if target == p {
			return true
		}
	}
	return false
}

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func dijkstra(grid [][]int) [][]int {
	distance := make([][]int, len(grid))
	unvisited := []pair{}

	for y := 0; y < len(grid); y++ {
		distance[y] = make([]int, len(grid[y]))
		for x := 0; x < len(grid[y]); x++ {
			unvisited = append(unvisited, pair{x: x, y: y})
			distance[y][x] = math.MaxInt64
		}
	}
	distance[0][0] = 0

	for len(unvisited) != 0 {
		var currNode pair
		currNode, unvisited = visitMin(distance, unvisited)
		neighbors := getAllAdjacent(currNode.x, currNode.y, len(grid[0]), len(grid))
		for _, neighbor := range neighbors {
			if containsPair(unvisited, neighbor) {
				// update distance
				if distance[neighbor.y][neighbor.x] > distance[currNode.y][currNode.x]+grid[neighbor.y][neighbor.x] {
					distance[neighbor.y][neighbor.x] = distance[currNode.y][currNode.x] + grid[neighbor.y][neighbor.x]
				}
			}
		}
	}

	return distance
}

// TODO: still pretty slow. Use better data structures for getting min node
// and checking for node existence in set
func dijkstra2(grid [][]int, dst pair) int {
	distance := make([][]int, len(grid))
	unvisited := []pair{}

	for y := 0; y < len(grid); y++ {
		distance[y] = make([]int, len(grid[y]))
		for x := 0; x < len(grid[y]); x++ {
			unvisited = append(unvisited, pair{x: x, y: y})
			distance[y][x] = math.MaxInt64
		}
	}
	distance[0][0] = 0

	for len(unvisited) != 0 {
		var currNode pair
		currNode, unvisited = visitMin(distance, unvisited)
		if currNode == dst {
			return distance[currNode.y][currNode.x]
		}
		neighbors := getAllAdjacent(currNode.x, currNode.y, len(grid[0]), len(grid))
		for _, neighbor := range neighbors {
			if containsPair(unvisited, neighbor) {
				// update distance
				if distance[neighbor.y][neighbor.x] > distance[currNode.y][currNode.x]+grid[neighbor.y][neighbor.x] {
					distance[neighbor.y][neighbor.x] = distance[currNode.y][currNode.x] + grid[neighbor.y][neighbor.x]
				}
			}
		}
	}

	return 0
}

func generateBiggerGrid(grid [][]int, scale int) [][]int {
	yLen := len(grid)
	xLen := len(grid[0])
	bigGrid := make([][]int, len(grid)*scale)

	for y := range bigGrid {
		bigGrid[y] = make([]int, len(grid[0])*scale)
	}

	for yScale := 0; yScale < scale; yScale++ {
		for xScale := 0; xScale < scale; xScale++ {
			yOffset := yScale * yLen
			xOffset := xScale * xLen
			for y := 0; y < yLen; y++ {
				for x := 0; x < xLen; x++ {
					newNum := grid[y][x] + xScale + yScale
					if newNum > 9 {
						newNum -= 9
					}
					bigGrid[y+yOffset][x+xOffset] = newNum
				}
			}
		}
	}

	return bigGrid
}

func part1(input *puzzleInput) int {
	distance := dijkstra(input.grid)
	return distance[len(input.grid)-1][len(input.grid[0])-1]
}

func part2(input *puzzleInput) int {
	biggerGrid := generateBiggerGrid(input.grid, 5)
	return dijkstra2(biggerGrid, pair{x: len(biggerGrid[0]) - 1, y: len(biggerGrid) - 1})
}

func main() {
	file, err := os.Open("inputs.txt")
	//file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	input, err := readInput(file)
	if err != nil {
		panic(err)
	}

	//fmt.Println(part1(input))
	fmt.Println(part2(input))
}
