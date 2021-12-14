package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	x int
	y int
}

type dir int

const (
	xDir dir = iota
	yDir
)

type fold struct {
	dir dir
	val int
}

type puzzleInput struct {
	dots  []pair
	folds []fold
}

func readInput(r io.Reader) (*puzzleInput, error) {
	scanner := bufio.NewScanner(r)
	input := &puzzleInput{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		coord := strings.Split(line, ",")
		x, err := strconv.Atoi(coord[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(coord[1])
		if err != nil {
			return nil, err
		}

		input.dots = append(input.dots, pair{x: x, y: y})
	}

	for scanner.Scan() {
		line := scanner.Text()
		foldPart := strings.Split(line, "=")
		d := xDir
		if foldPart[0][len(foldPart[0])-1] == 'y' {
			d = yDir
		}

		val, err := strconv.Atoi(foldPart[1])
		if err != nil {
			return nil, err
		}
		input.folds = append(input.folds, fold{dir: d, val: val})
	}

	return input, nil
}

func generateGrid(pairs []pair) [][]bool {
	// determine dimension
	maxX, maxY := 0, 0
	for _, pair := range pairs {
		if pair.x > maxX {
			maxX = pair.x
		}
		if pair.y > maxY {
			maxY = pair.y
		}
	}

	// create empty grid
	grid := make([][]bool, maxY+1)
	for i := 0; i < maxY+1; i++ {
		grid[i] = make([]bool, maxX+1)
	}

	// set the values
	for _, pair := range pairs {
		grid[pair.y][pair.x] = true
	}

	return grid
}

func printGrid(grid [][]bool) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func copyGrid(grid [][]bool) [][]bool {
	c := make([][]bool, len(grid))
	for y := range grid {
		c[y] = make([]bool, len(grid[y]))
		copy(c[y], grid[y])
	}
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func foldY(grid [][]bool, yLine int) [][]bool {
	foldedGrid := copyGrid(grid)
	newHeight := max(yLine, len(grid)-yLine-1)
	foldedGrid = foldedGrid[:newHeight]

	for yOffset := 1; yLine+yOffset < len(grid); yOffset++ {
		for x := 0; x < len(grid[yLine+yOffset]); x++ {
			foldedGrid[yLine-yOffset][x] = grid[yLine+yOffset][x] || grid[yLine-yOffset][x]
		}
	}

	return foldedGrid
}

func foldX(grid [][]bool, xLine int) [][]bool {
	foldedGrid := copyGrid(grid)
	newWidth := max(xLine, len(grid[0])-xLine-1)
	for y := range foldedGrid {
		foldedGrid[y] = foldedGrid[y][:newWidth]
	}

	for xOffset := 1; xLine+xOffset < len(grid[0]); xOffset++ {
		for y := 0; y < len(grid); y++ {
			foldedGrid[y][xLine-xOffset] = grid[y][xLine+xOffset] || grid[y][xLine-xOffset]
		}
	}

	return foldedGrid
}

func countDots(grid [][]bool) int {
	sum := 0
	for _, row := range grid {
		for _, val := range row {
			if val {
				sum++
			}
		}
	}
	return sum
}

func part1(input *puzzleInput) int {
	grid := generateGrid(input.dots)
	for i, f := range input.folds {
		if f.dir == xDir {
			grid = foldX(grid, f.val)
		} else {
			grid = foldY(grid, f.val)
		}
		if i == 0 {
			break
		}
	}
	printGrid(grid)

	return countDots(grid)
}

func part2(input *puzzleInput) int {
	grid := generateGrid(input.dots)
	for _, f := range input.folds {
		if f.dir == xDir {
			grid = foldX(grid, f.val)
		} else {
			grid = foldY(grid, f.val)
		}
	}
	printGrid(grid)

	return countDots(grid)
}

func main() {
	//file, err := os.Open("example.txt")
	file, err := os.Open("inputs.txt")
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
