package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readInput(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func isMatching(a rune, b rune) bool {
	switch a {
	case '(':
		return b == ')'
	case '[':
		return b == ']'
	case '{':
		return b == '}'
	case '<':
		return b == '>'
	}
	return false
}

func errorValue(a rune) int {
	switch a {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	return 0
}

func isOpen(r rune) bool {
	switch r {
	case '(', '[', '{', '<':
		return true
	}
	return false
}

func errorPos(line string) int {
	s := stack{}
	for i, char := range line {
		if isOpen(char) {
			s.push(char)
			continue
		}

		potentialMatch, empty := s.pop()
		// missing data
		if empty {
			return -1
		}

		// mismatch
		if !isMatching(potentialMatch, char) {
			return i
		}
	}
	return -1
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		if pos := errorPos(line); pos >= 0 {
			sum += errorValue(rune(line[pos]))
		}
	}
	return sum
}

func main() {
	file, err := os.Open("inputs.txt")
	//file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	lines, err := readInput(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(lines))
}
