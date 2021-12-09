package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func readInput(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)
	grid := [][]int{}
	for scanner.Scan() {
		row := []int{}
		for _, cell := range scanner.Text() {
			row = append(row, int(cell-'0'))
		}
		grid = append(grid, row)
	}
	return grid, nil
}

func riskLevel(row, col int, grid [][]int) (bool, int) {
	num := grid[row][col]
	if row+1 < len(grid) && num >= grid[row+1][col] {
		return false, 0
	}
	if row-1 >= 0 && num >= grid[row-1][col] {
		return false, 0
	}
	if col+1 < len(grid[0]) && num >= grid[row][col+1] {
		return false, 0
	}
	if col-1 >= 0 && num >= grid[row][col-1] {
		return false, 0
	}

	return true, num + 1
}

func part1(grid [][]int) int {
	sum := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if isLow, riskLevel := riskLevel(row, col, grid); isLow {
				sum += riskLevel
			}
		}
	}
	return sum
}

func dfs(row, col int, grid [][]int, visited [][]bool) int {
	size := 1
	visited[row][col] = true

	if row+1 < len(grid) && !visited[row+1][col] && grid[row+1][col] != 9 {
		size += dfs(row+1, col, grid, visited)
	}
	if row-1 >= 0 && !visited[row-1][col] && grid[row-1][col] != 9 {
		size += dfs(row-1, col, grid, visited)
	}
	if col+1 < len(grid[0]) && !visited[row][col+1] && grid[row][col+1] != 9 {
		size += dfs(row, col+1, grid, visited)
	}
	if col-1 >= 0 && !visited[row][col-1] && grid[row][col-1] != 9 {
		size += dfs(row, col-1, grid, visited)
	}

	return size
}

func newVisited(numRows, numCols int) [][]bool {
	visited := make([][]bool, numRows)
	for row := range visited {
		visited[row] = make([]bool, numCols)
	}
	return visited
}

func part2(grid [][]int) int {
	// for each low point
	//		depth first search and store size in slice
	// sum 3 lowest sizes

	sizes := []int{}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if isLow, _ := riskLevel(row, col, grid); isLow {
				size := dfs(row, col, grid, newVisited(len(grid), len(grid[0])))
				sizes = append(sizes, size)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}

func main() {
	file, err := os.Open("inputs.txt")
	//file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	grid, err := readInput(file)
	if err != nil {
		panic(err)
	}

	//fmt.Println(part1(grid))
	fmt.Println(part2(grid))
}
