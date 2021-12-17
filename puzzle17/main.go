package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type pair struct {
	x, y int
}

type puzzleInput struct {
	minX, maxX, minY, maxY int
}

func readInput(r io.Reader) (*puzzleInput, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line := scanner.Text()
	re := regexp.MustCompile("^target area: x=(.*)\\.\\.(.*), y=(.*)\\.\\.(.*)$")
	matches := re.FindStringSubmatch(line)

	minX, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, err
	}
	maxX, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}
	minY, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil, err
	}
	maxY, err := strconv.Atoi(matches[4])
	if err != nil {
		return nil, err
	}

	return &puzzleInput{
		minX: minX,
		maxX: maxX,
		minY: minY,
		maxY: maxY,
	}, nil
}

func next(pos, vel pair) (pair, pair) {
	newVel := pair{}
	if vel.x > 0 {
		newVel.x = vel.x - 1
	} else if vel.x < 0 {
		newVel.x = vel.x + 1
	} else {
		newVel.x = 0
	}
	newVel.y = vel.y - 1

	return pair{
		x: pos.x + vel.x,
		y: pos.y + vel.y,
	}, newVel
}

func part1(input *puzzleInput) int {
	minX, maxX, minY, maxY := input.minX, input.maxX, input.minY, input.maxY
	maxHeight := 0

	for xVel := 0; xVel < maxX+1; xVel++ {
		for yVel := 0; yVel < -minY+1; yVel++ {
			landed := false
			pos := pair{}
			vel := pair{x: xVel, y: yVel}
			//fmt.Println("trying velocity", vel)
			currMaxY := 0

			for {
				//fmt.Println("evaluating", pos)
				if pos.x >= minX && pos.x <= maxX && pos.y >= minY && pos.y <= maxY {
					//fmt.Println("landed", pos)
					landed = true
					break
				}
				if pos.x > maxX || pos.y < minY {
					//fmt.Println("overreach", pos)
					break
				}

				if pos.y > currMaxY {
					currMaxY = pos.y
				}
				pos, vel = next(pos, vel)
			}

			if landed && currMaxY > maxHeight {
				maxHeight = currMaxY
			}
		}
	}

	return maxHeight
}

func part2(input *puzzleInput) int {
	minX, maxX, minY, maxY := input.minX, input.maxX, input.minY, input.maxY
	numLanded := 0

	for xVel := 0; xVel < maxX+1; xVel++ {
		for yVel := minY; yVel < -minY+1; yVel++ {
			pos := pair{}
			vel := pair{x: xVel, y: yVel}
			//fmt.Println("trying velocity", vel)

			for {
				//fmt.Println("evaluating", pos)
				if pos.x >= minX && pos.x <= maxX && pos.y >= minY && pos.y <= maxY {
					//fmt.Println("landed", pos)
					numLanded++
					break
				}
				if pos.x > maxX || pos.y < minY {
					//fmt.Println("overreach", pos)
					break
				}

				pos, vel = next(pos, vel)
			}
		}
	}

	return numLanded
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
