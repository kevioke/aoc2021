package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readInput(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	numbers := []int{}
	scanner.Scan()
	for _, numStr := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func fuelCost(crabPosCount map[int]int, candidatePos int) int {
	cost := 0
	for pos, numCrabs := range crabPosCount {
		cost += absDiff(pos, candidatePos) * numCrabs
	}
	return cost
}

func part1(crabPositions []int) int {
	// generate map of crab position counts
	crabPosCount := map[int]int{}
	for _, crabPos := range crabPositions {
		if _, found := crabPosCount[crabPos]; !found {
			crabPosCount[crabPos] = 1
		} else {
			crabPosCount[crabPos]++
		}
	}

	// for each crab position calculate difference and sum them keeping track of min
	minCost := -1
	for pos := range crabPosCount {
		// first cost
		if minCost == -1 {
			minCost = fuelCost(crabPosCount, pos)
			continue
		}

		cost := fuelCost(crabPosCount, pos)
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func main() {
	file, err := os.Open("inputs.txt")
	//file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	numbers, err := readInput(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(numbers))
}
