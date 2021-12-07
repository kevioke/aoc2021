package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("inputs.txt")
	//file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	inputs, err := readInts(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(numIncreases(sumGroup(inputs, 3)))
}

func sumGroup(nums []int, size int) []int {
	groups := []int{}

	for i := range nums {
		if i+size > len(nums) {
			break
		}

		sum := 0
		for _, num := range nums[i : i+size] {
			sum += num
		}

		groups = append(groups, sum)
	}

	return groups
}

func numIncreases(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numIncreases := 0
	prevNum := nums[0]
	for _, num := range nums {
		if num > prevNum {
			numIncreases++
		}

		prevNum = num
	}
	return numIncreases
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}
