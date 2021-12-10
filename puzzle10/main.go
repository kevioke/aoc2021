package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func completionValue(a rune) int {
	switch a {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
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

		openChar, empty := s.pop()
		if empty {
			panic("not enough open characters")
		}

		// mismatch
		if !isMatching(openChar, char) {
			return i
		}
	}
	return -1
}

func pairClose(r rune) rune {
	switch r {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	}
	return 0
}

func getClosingChars(line string) string {
	var closingChars string
	s := stack{}
	for _, char := range line {
		if isOpen(char) {
			s.push(char)
			continue
		}

		openChar, empty := s.pop()
		if empty {
			panic("not enough open characters")
		}

		if !isMatching(openChar, char) {
			return ""
		}
	}

	for {
		char, empty := s.pop()
		if empty {
			break
		}

		closingChars = closingChars + string(pairClose(char))
	}

	return closingChars
}

func part2(lines []string) int {
	points := []int{}
	for _, line := range lines {
		closingChars := getClosingChars(line)
		if len(closingChars) > 0 {
			sum := 0
			for _, char := range closingChars {
				sum *= 5
				sum += completionValue(char)
			}
			points = append(points, sum)
		}
	}

	sort.Ints(points)
	return points[len(points)/2]
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

	//fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
