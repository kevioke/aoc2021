package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	up direction = iota
	down
	forward
)

type command struct {
	dir    direction
	scaler int
}

func main() {
	file, err := os.Open("inputs.txt")
	//file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	commands, err := readCommands(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(commands))
	fmt.Println(part2(commands))
}

func part1(commands []command) int {
	horizontalPos := 0
	depthPos := 0

	for _, cmd := range commands {
		switch cmd.dir {
		case up:
			depthPos -= cmd.scaler
		case down:
			depthPos += cmd.scaler
		case forward:
			horizontalPos += cmd.scaler
		default:
			panic("bad direction")
		}
	}

	return horizontalPos * depthPos
}

func part2(commands []command) int {
	aim := 0
	horizontalPos := 0
	depthPos := 0

	for _, cmd := range commands {
		switch cmd.dir {
		case up:
			aim -= cmd.scaler
		case down:
			aim += cmd.scaler
		case forward:
			horizontalPos += cmd.scaler
			depthPos += aim * cmd.scaler
		default:
			panic("bad direction")
		}
	}

	return horizontalPos * depthPos
}

func readCommands(r io.Reader) ([]command, error) {
	commands := []command{}
	scanner := bufio.NewScanner(r)
	//scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		scaler, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		cmd := command{
			scaler: scaler,
		}

		switch parts[0] {
		case "up":
			cmd.dir = up
		case "down":
			cmd.dir = down
		case "forward":
			cmd.dir = forward
		default:
			panic("bad dir")
		}

		commands = append(commands, cmd)
	}
	return commands, scanner.Err()
}
