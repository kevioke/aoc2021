package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type gridCount struct {
	counts [][]int
}

func (g *gridCount) initGrid(size int) {
	g.counts = make([][]int, size)
	for i := 0; i < size; i++ {
		g.counts[i] = make([]int, size)
	}
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}

	return b, a
}

func xyDir(segment lineSegment) (int, int) {
	xDir := 1
	if segment.start.x > segment.end.x {
		xDir = -1
	}

	yDir := 1
	if segment.start.y > segment.end.y {
		yDir = -1
	}
	return xDir, yDir
}

func (g *gridCount) updateWithDiagonals(segment lineSegment) {
	// vertical line
	if segment.start.x == segment.end.x {
		x := segment.start.x
		for currY, endY := minMax(segment.start.y, segment.end.y); currY <= endY; currY++ {
			g.counts[x][currY]++
		}
	} else if segment.start.y == segment.end.y { // horizontal
		y := segment.start.y
		for currX, endX := minMax(segment.start.x, segment.end.x); currX <= endX; currX++ {
			g.counts[currX][y]++
		}
	} else {
		xDir, yDir := xyDir(segment)
		x, y := segment.start.x, segment.start.y
		for x != segment.end.x+xDir {
			g.counts[x][y]++
			x += xDir
			y += yDir
		}
	}
}

func (g *gridCount) update(segment lineSegment) {
	// vertical line
	if segment.start.x == segment.end.x {
		x := segment.start.x
		for currY, endY := minMax(segment.start.y, segment.end.y); currY <= endY; currY++ {
			g.counts[x][currY]++
		}
	} else if segment.start.y == segment.end.y { // horizontal
		y := segment.start.y
		for currX, endX := minMax(segment.start.x, segment.end.x); currX <= endX; currX++ {
			g.counts[currX][y]++
		}
	}
}

func (g *gridCount) print() {
	for y := 0; y < len(g.counts); y++ {
		for x := 0; x < len(g.counts); x++ {
			fmt.Print(g.counts[x][y])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func (g *gridCount) numOverlaps() int {
	numOverlaps := 0
	for i := 0; i < len(g.counts); i++ {
		for j := 0; j < len(g.counts); j++ {
			if g.counts[i][j] > 1 {
				numOverlaps++
			}
		}
	}
	return numOverlaps
}

type segmentPoint struct {
	x int
	y int
}

func (s segmentPoint) findMax() int {
	if s.x > s.y {
		return s.x
	}
	return s.y
}

type lineSegment struct {
	start segmentPoint
	end   segmentPoint
}

func (l lineSegment) findMax() int {
	startMax := l.start.findMax()
	endMax := l.end.findMax()
	if startMax > endMax {
		return startMax
	}
	return endMax
}

func findMax(segments []lineSegment) int {
	max := 0
	for _, segment := range segments {
		segMax := segment.findMax()
		if segMax > max {
			max = segMax
		}
	}
	return max
}

func parseXY(s string) (int, int) {
	parts := strings.Split(s, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return x, y
}

func readInput(r io.Reader) ([]lineSegment, error) {
	scanner := bufio.NewScanner(r)
	segments := []lineSegment{}
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		segment := lineSegment{}
		segment.start.x, segment.start.y = parseXY(fields[0])
		segment.end.x, segment.end.y = parseXY(fields[2])
		segments = append(segments, segment)
	}
	return segments, nil
}

func part1(segments []lineSegment) int {
	// init grid
	grid := gridCount{}
	grid.initGrid(findMax(segments) + 1)

	// for each segment update board
	for _, segment := range segments {
		grid.update(segment)
	}

	// count overlaps and return
	return grid.numOverlaps()
}

func part2(segments []lineSegment) int {
	// init grid
	grid := gridCount{}
	grid.initGrid(findMax(segments) + 1)

	// for each segment update board
	for _, segment := range segments {
		grid.updateWithDiagonals(segment)
	}

	// count overlaps and return
	return grid.numOverlaps()
}

func main() {
	file, err := os.Open("inputs.txt")
	//file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	segments, err := readInput(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(segments))
	fmt.Println(part2(segments))
}
