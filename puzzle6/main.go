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

func tick(fishTimers []int) []int {
	newFishTimers := make([]int, len(fishTimers))
	for i, fishTimer := range fishTimers {
		if fishTimer == 0 {
			newFishTimers[i] = 6
			newFishTimers = append(newFishTimers, 8)
		} else {
			newFishTimers[i] = fishTimer - 1
		}
	}
	return newFishTimers
}

func part1(fishTimers []int) int {
	for i := 0; i < 80; i++ {
		fishTimers = tick(fishTimers)
	}
	return len(fishTimers)
}

func optimizedTick(timerCounts []int) []int {
	newTimerCounts := make([]int, len(timerCounts))
	spawningFish := timerCounts[0]
	for i := 0; i < len(timerCounts)-1; i++ {
		newTimerCounts[i] = timerCounts[i+1]
	}

	newTimerCounts[len(timerCounts)-1] = spawningFish
	newTimerCounts[6] += spawningFish
	return newTimerCounts
}

func part2(fishTimers []int) int {
	timerCounts := make([]int, 9)
	for _, timer := range fishTimers {
		timerCounts[timer]++
	}
	fmt.Println(timerCounts)

	for i := 0; i < 256; i++ {
		timerCounts = optimizedTick(timerCounts)
	}

	numFishes := 0
	for _, timerCount := range timerCounts {
		numFishes += timerCount
	}

	return numFishes
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

	fmt.Println(part2(numbers))
}
