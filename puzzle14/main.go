package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type puzzleInput struct {
	template string
	rules    map[string]byte
}

func readInput(r io.Reader) (*puzzleInput, error) {
	scanner := bufio.NewScanner(r)
	input := &puzzleInput{
		rules: map[string]byte{},
	}
	scanner.Scan()
	input.template = scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		ruleStr := strings.Split(line, " -> ")
		input.rules[ruleStr[0]] = byte(ruleStr[1][0])
	}
	return input, nil
}

func getNewPolymer(polymer, insertions []byte) []byte {
	b := bytes.Buffer{}
	for i := 0; i < len(polymer)-1; i++ {
		b.WriteByte(polymer[i])
		if insertions[i] != 0 {
			b.WriteByte(insertions[i])
		}
	}
	b.WriteByte(polymer[len(polymer)-1])
	return b.Bytes()
}

func getInsertions(polymer []byte, rules map[string]byte) []byte {
	insertions := []byte{}
	for i := 0; i < len(polymer)-1; i++ {
		if insertion, found := rules[string(polymer[i:i+2])]; found {
			insertions = append(insertions, insertion)
		}
	}

	return insertions
}

func getElementCount(polymer []byte) map[byte]int {
	count := map[byte]int{}
	for _, element := range polymer {
		count[element]++
	}
	return count
}

func part1(input *puzzleInput) int {
	polymer := []byte(input.template)

	for i := 0; i < 10; i++ {
		insertions := getInsertions(polymer, input.rules)
		polymer = getNewPolymer(polymer, insertions)
	}

	elementCount := getElementCount(polymer)
	min, max := -1, -1

	for _, count := range elementCount {
		if min == -1 {
			min = count
		}
		if max == -1 {
			max = count
		}

		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}

	return max - min
}

func getNewPolymerMap(polymer map[string]int, rules map[string]byte) (map[string]int, map[byte]int) {
	counts := map[byte]int{}
	newPolymer := map[string]int{}
	for pair, count := range polymer {
		newElement := rules[pair]
		counts[newElement] += count

		newPair1 := string(pair[0]) + string(newElement)
		newPair2 := string(newElement) + string(pair[1])
		newPolymer[newPair1] += count
		newPolymer[newPair2] += count
	}
	return newPolymer, counts
}

func part2(input *puzzleInput) int {
	polymer := map[string]int{}
	for i := 0; i < len(input.template)-1; i++ {
		polymer[input.template[i:i+2]]++
	}

	counts := map[byte]int{}
	for i := 0; i < len(input.template); i++ {
		counts[input.template[i]]++
	}

	for i := 0; i < 40; i++ {
		var newCounts map[byte]int
		polymer, newCounts = getNewPolymerMap(polymer, input.rules)
		for el, count := range newCounts {
			counts[el] += count
		}
	}

	min, max := -1, 1

	for _, count := range counts {
		if min == -1 {
			min = count
		}
		if max == -1 {
			max = count
		}

		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}

	return max - min
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
