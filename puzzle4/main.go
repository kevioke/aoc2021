package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	boardLen = 5
)

type bingoBoard struct {
	array  [][]int
	marked [][]bool
	hasWon bool
}

func (b *bingoBoard) String() string {
	return fmt.Sprintf("%v %v", b.array, b.marked)
}

func (b *bingoBoard) hasBingo() bool {
	// check horizontal
	for i := 0; i < boardLen; i++ {
		allMarked := true
		for j := 0; j < boardLen; j++ {
			if !b.marked[i][j] {
				allMarked = false
			}
		}
		if allMarked {
			return true
		}
	}

	// check vertical
	for j := 0; j < boardLen; j++ {
		allMarked := true
		for i := 0; i < boardLen; i++ {
			if !b.marked[i][j] {
				allMarked = false
			}
		}
		if allMarked {
			return true
		}
	}

	return false
}

func (b *bingoBoard) initBoardState() {
	for i := 0; i < boardLen; i++ {
		markedRow := make([]bool, boardLen)
		b.marked = append(b.marked, markedRow)
	}
}

func (b *bingoBoard) updateState(bingoNum int) {
	for i := 0; i < boardLen; i++ {
		for j := 0; j < boardLen; j++ {
			if b.array[i][j] == bingoNum {
				b.marked[i][j] = true
			}
		}
	}
}

func (b *bingoBoard) sumOfUnmarked() int {
	sum := 0
	for i := 0; i < boardLen; i++ {
		for j := 0; j < boardLen; j++ {
			if !b.marked[i][j] {
				sum += b.array[i][j]
			}
		}
	}
	return sum
}

func readInput(r io.Reader) ([]int, []*bingoBoard, error) {
	scanner := bufio.NewScanner(r)
	seq := []int{}
	scanner.Scan()
	seqString := scanner.Text()
	for _, numString := range strings.Split(seqString, ",") {
		num, err := strconv.Atoi(numString)
		if err != nil {
			return nil, nil, err
		}
		seq = append(seq, num)
	}

	scanner.Scan()
	boards := []*bingoBoard{}
	currentBoard := &bingoBoard{}
	row := 0
	for scanner.Scan() {
		if row == boardLen {
			row = 0
			boards = append(boards, currentBoard)
			currentBoard = &bingoBoard{}
			continue
		}

		currentRow := []int{}
		for _, numString := range strings.Fields(scanner.Text()) {
			num, err := strconv.Atoi(numString)
			if err != nil {
				return nil, nil, err
			}
			currentRow = append(currentRow, num)
		}
		currentBoard.array = append(currentBoard.array, currentRow)
		row++
	}
	boards = append(boards, currentBoard)

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return seq, boards, scanner.Err()
}

func part1(boards []*bingoBoard, seq []int) int {
	for _, bingoNum := range seq {
		for _, board := range boards {
			board.updateState(bingoNum)
			if board.hasBingo() {
				return board.sumOfUnmarked() * bingoNum
			}
		}
	}
	return 0
}

func part2(boards []*bingoBoard, seq []int) int {
	numberOfBingos := 0
	for _, bingoNum := range seq {
		for _, board := range boards {
			if board.hasWon {
				continue
			}

			board.updateState(bingoNum)
			if board.hasBingo() {
				board.hasWon = true
				numberOfBingos++
			}
			if numberOfBingos == len(boards) {
				return board.sumOfUnmarked() * bingoNum
			}
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

	seq, boards, err := readInput(file)
	if err != nil {
		panic(err)
	}

	for _, board := range boards {
		board.initBoardState()
	}

	fmt.Println(part1(boards, seq))
	fmt.Println(part2(boards, seq))
}
