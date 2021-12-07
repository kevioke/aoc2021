package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type binaryNumber struct {
	bits []int
}

func (b binaryNumber) toNumber() int {
	num := 0
	pow := 1
	for i := len(b.bits) - 1; i >= 0; i-- {
		num += b.bits[i] * pow
		pow *= 2
	}
	return num
}

func main() {
	file, err := os.Open("inputs.txt")
	//file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	binaryNumbers, err := readInput(file)
	if err != nil {
		panic(err)
	}

	//fmt.Println(part1(binaryNumbers))
	fmt.Println(part2(binaryNumbers))
}

func part2(numbers []binaryNumber) int {
	oxygenNums := make([]binaryNumber, len(numbers))
	co2Nums := make([]binaryNumber, len(numbers))
	copy(oxygenNums, numbers)
	copy(co2Nums, numbers)

	// calculate oxygen
	for i := 0; i < len(numbers[0].bits) && len(oxygenNums) > 1; i++ {
		oxygenNums = filter(oxygenNums, i, true)
	}
	oxygen := oxygenNums[0].toNumber()

	// calculate co2
	for i := 0; i < len(numbers[0].bits) && len(co2Nums) > 1; i++ {
		co2Nums = filter(co2Nums, i, false)
	}
	co2 := co2Nums[0].toNumber()

	return oxygen * co2
}

func filter(numbers []binaryNumber, position int, isOxygen bool) []binaryNumber {
	filteredNums := make([]binaryNumber, 0, len(numbers))

	var include int
	if isOxygen {
		include = findTargetBit(numbers, position, isOxygen, 1)
	} else {
		include = findTargetBit(numbers, position, isOxygen, 0)
	}

	for _, number := range numbers {
		if number.bits[position] == include {
			filteredNums = append(filteredNums, number)
		}
	}
	return filteredNums
}

func findTargetBit(numbers []binaryNumber, position int, majority bool, equalNum int) int {
	m := 1
	n := 0
	if !majority {
		m = 0
		n = 1
	}
	sum := 0
	for _, num := range numbers {
		sum += num.bits[position]
	}

	if sum*2 == len(numbers) {
		return equalNum
	}
	if sum*2 > len(numbers) {
		return m
	}
	return n
}

func part1(numbers []binaryNumber) int {
	numBits := len(numbers[0].bits)
	sum := make([]int, numBits)

	for _, num := range numbers {
		for i, bitValue := range num.bits {
			sum[i] += bitValue
		}
	}

	gamma := binaryNumber{}
	epsilon := binaryNumber{}
	for _, s := range sum {
		if s > len(numbers)/2 {
			gamma.bits = append(gamma.bits, 1)
			epsilon.bits = append(epsilon.bits, 0)
		} else {
			gamma.bits = append(gamma.bits, 0)
			epsilon.bits = append(epsilon.bits, 1)
		}
	}

	return gamma.toNumber() * epsilon.toNumber()
}

func readInput(r io.Reader) ([]binaryNumber, error) {
	numbers := []binaryNumber{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		binNum := binaryNumber{}
		for _, bit := range line {
			switch bit {
			case '1':
				binNum.bits = append(binNum.bits, 1)
			case '0':
				binNum.bits = append(binNum.bits, 0)
			default:
				return nil, errors.New("malformed bit")
			}
		}

		numbers = append(numbers, binNum)
	}
	return numbers, scanner.Err()
}
