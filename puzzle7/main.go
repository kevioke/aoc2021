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

func part1(crabPositions []int) int {
	return 0
}

func main() {
	//file, err := os.Open("inputs.txt")
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	numbers, err := readInput(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(numbers)
}
