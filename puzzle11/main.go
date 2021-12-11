package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readInput(r io.Reader) ([][]int, error) {
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
	return grid, nil
}

func flashStep(grid [][]int, flashState [][]bool) int {
	numFlashes := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 9 && !flashState[i][j] {
				flashState[i][j] = true
				numFlashes++
				adjacents := getAllAdjacent(i, j, len(grid), len(grid[i]))
				for _, adj := range adjacents {
					grid[adj[0]][adj[1]]++
				}
			}
		}
	}
	return numFlashes
}

func resetIfFlash(grid [][]int) int {
	numFlash := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 9 {
				grid[i][j] = 0
				numFlash++
			}
		}
	}

	return numFlash
}

func incrementGrid(grid [][]int, count int) {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] += count
		}
	}
}

func newFlashState(rows, cols int) [][]bool {
	flashState := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		flashState[i] = make([]bool, cols)
	}

	return flashState
}

func part1(grid [][]int) int {
	flashCount := 0
	for i := 0; i < 100; i++ {
		incrementGrid(grid, 1)
		flashState := newFlashState(len(grid), len(grid[0]))
		for numFlashes := flashStep(grid, flashState); numFlashes > 0; numFlashes = flashStep(grid, flashState) {
			flashCount += numFlashes
		}
		resetIfFlash(grid)
	}
	return flashCount
}

func part2(grid [][]int) int {
	for step := 1; ; step++ {
		incrementGrid(grid, 1)
		flashState := newFlashState(len(grid), len(grid[0]))
		for numFlashes := flashStep(grid, flashState); numFlashes > 0; numFlashes = flashStep(grid, flashState) {
		}
		if resetIfFlash(grid) == len(grid)*len(grid[0]) {
			return step
		}
	}
	return 0
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
