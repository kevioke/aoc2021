package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func readInput(r io.Reader) (map[string][]string, error) {
	scanner := bufio.NewScanner(r)
	nodes := map[string][]string{}

	for scanner.Scan() {
		edge := strings.Split(scanner.Text(), "-")
		a := edge[0]
		b := edge[1]
		nodes[a] = append(nodes[a], b)
		nodes[b] = append(nodes[b], a)
	}
	return nodes, nil
}

func copyMap(m map[string]bool) map[string]bool {
	cp := map[string]bool{}
	for key, val := range m {
		cp[key] = val
	}

	return cp
}

func numPaths2(nodes map[string][]string, src, dst string, exclude map[string]bool, twiceVisit bool) [][]string {
	if src == dst {
		return [][]string{[]string{dst}}
	}

	paths := [][]string{}
	newExclude := copyMap(exclude)
	if unicode.IsLower(rune(src[0])) {
		newExclude[src] = true
	}

	for _, neighbor := range nodes[src] {
		if neighbor == "start" || (exclude[neighbor] && twiceVisit) {
			continue
		}

		if exclude[neighbor] {
			for _, path := range numPaths2(nodes, neighbor, dst, newExclude, true) {
				paths = append(paths, append([]string{src}, path...))
			}
			continue
		}

		for _, path := range numPaths2(nodes, neighbor, dst, newExclude, twiceVisit) {
			paths = append(paths, append([]string{src}, path...))
		}
	}
	return paths
}

func numPaths(nodes map[string][]string, src, dst string, exclude map[string]bool) int {
	totalPaths := 0
	if src == dst {
		return 1
	}

	for _, neighbor := range nodes[src] {
		if exclude[neighbor] {
			continue
		}

		newExclude := copyMap(exclude)
		if unicode.IsLower(rune(src[0])) {
			newExclude[src] = true
		}
		totalPaths += numPaths(nodes, neighbor, dst, newExclude)
	}
	return totalPaths
}

func part1(nodes map[string][]string) int {
	return numPaths(nodes, "start", "end", map[string]bool{})
}

func part2(nodes map[string][]string) int {
	paths := numPaths2(nodes, "start", "end", map[string]bool{}, false)
	//fmt.Println(paths)
	return len(paths)
}

func main() {
	file, err := os.Open("inputs.txt")
	//file, err := os.Open("example.txt")
	//file, err := os.Open("example2.txt")
	//file, err := os.Open("example3.txt")
	if err != nil {
		panic(err)
	}

	input, err := readInput(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
