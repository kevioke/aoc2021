package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"

	"kevinguy.com/aoc2021/stack"
)

type pair struct {
	First, Second *pair
	Value         *int
}

type puzzleInput struct {
	pairs []*pair
}

func parseLine(line string) *pair {
	s := stack.Stack[*pair]{}
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case ',':
		case '[':
			s.Push(&pair{})
		case ']':
			second := s.Pop()
			first := s.Pop()
			currentPair := s.Peek()
			currentPair.First = first
			currentPair.Second = second
		default:
			if !unicode.IsDigit(rune(line[i])) {
				panic(fmt.Sprintf("unexpected char %s at %d", string(line[i]), i))
			}
			numberStr := ""
			for unicode.IsDigit(rune(line[i])) {
				numberStr += string(line[i])
				i++
			}
			i--

			number, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}

			s.Push(&pair{Value: &number})
		}
	}
	return s.Pop()
}

func readInput(r io.Reader) (*puzzleInput, error) {
	scanner := bufio.NewScanner(r)
	input := &puzzleInput{}
	for scanner.Scan() {
		line := scanner.Text()
		input.pairs = append(input.pairs, parseLine(line))
	}

	return input, nil
}

func intPtr(i int) *int {
	return &i
}

func pairToStr(p *pair) string {
	if p.Value != nil {
		return fmt.Sprintf("%d", *p.Value)
	}

	return fmt.Sprintf("[%s,%s]", pairToStr(p.First), pairToStr(p.Second))
}

type state struct {
	replaced bool
	leftVal  int
	rightVal int
	leftAdd  bool
	rightAdd bool
}

func incrRightMost(p *pair, num int) {
	if p.Value != nil {
		*p.Value += num
		return
	}

	incrRightMost(p.Second, num)
}

func incrLeftMost(p *pair, num int) {
	if p.Value != nil {
		*p.Value += num
		return
	}

	incrLeftMost(p.First, num)
}

func add(p1, p2 *pair) *pair {
	return &pair{First: p1, Second: p2}
}

func split(p *pair) bool {
	if p.Value != nil && *p.Value >= 10 {
		first := *p.Value / 2
		second := (*p.Value + 1) / 2
		p.Value = nil
		p.First = &pair{Value: &first}
		p.Second = &pair{Value: &second}
		return true
	}
	if p.Value != nil {
		return false
	}

	if split(p.First) {
		return true
	}

	return split(p.Second)
}

func explode(p *pair, s *state, depth int) {
	// TODO: !s.replaced check might not be necessary
	if !s.replaced && depth >= 4 && p.First != nil && p.Second != nil && p.First.Value != nil && p.Second.Value != nil {
		s.leftVal = *p.First.Value
		s.rightVal = *p.Second.Value
		s.replaced = true

		p.Value = intPtr(0)
		p.First, p.Second = nil, nil
		return
	}

	if p.Value != nil {
		return
	}

	explode(p.First, s, depth+1)
	if s.replaced && !s.rightAdd {
		// find the leftmost value and increment by right most replaced value
		s.rightAdd = true
		incrLeftMost(p.Second, s.rightVal)
		return
	}

	if s.replaced {
		return
	}

	explode(p.Second, s, depth+1)
	if s.replaced && !s.leftAdd {
		// find the rightMost value and increment by right most replaced value
		s.leftAdd = true
		incrRightMost(p.First, s.leftVal)
		return
	}
}

func reduce(p *pair) {
	for {
		s := state{}
		explode(p, &s, 0)
		if s.replaced {
			fmt.Println("explode", pairToStr(p))
			continue
		}

		if split(p) {
			fmt.Println("split", pairToStr(p))
			continue
		}

		break
	}
}

func eval(pairs []*pair) *pair {
	p := pairs[0]
	reduce(p)
	for i := 1; i < len(pairs); i++ {
		p = add(p, pairs[i])
		fmt.Println("added", pairToStr(p))
		reduce(p)
	}

	return p
}

func magnitude(p *pair) int {
	if p.Value != nil {
		return *p.Value
	}

	return 3*magnitude(p.First) + 2*magnitude(p.Second)
}

func part1(input *puzzleInput) int {
	p := eval(input.pairs)
	fmt.Println(pairToStr(p))
	return magnitude(p)
}

func copyPair(p *pair) *pair {
	cp := &pair{}
	if p.Value != nil {
		cp.Value = intPtr(*p.Value)
		return cp
	}

	cp.First = copyPair(p.First)
	cp.Second = copyPair(p.Second)

	return cp
}

func copyPairs(pairs []*pair) []*pair {
	cp := []*pair{}
	for _, p := range pairs {
		cp = append(cp, copyPair(p))
	}

	return cp
}

func part2(input *puzzleInput) int {
	max := 0
	for i := range input.pairs {
		for j := range input.pairs {
			if i == j {
				continue
			}
			pairs := copyPairs(input.pairs)
			addRes := add(pairs[i], pairs[j])
			reduce(addRes)
			m := magnitude(addRes)
			if m > max {
				max = m
			}
		}
	}
	return max
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
