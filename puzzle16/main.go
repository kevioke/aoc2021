package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type puzzleInput struct {
	hexInput string
}

type packet struct {
	version    int
	packetType int

	literalValue int
	subPackets   []packet
}

func readInput(r io.Reader) (*puzzleInput, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return &puzzleInput{hexInput: scanner.Text()}, nil
}

func getBits(hexString string) ([]int, error) {
	bits := []int{}
	for _, hex := range hexString {
		num, err := strconv.ParseInt(string(hex), 16, 64)
		if err != nil {
			return nil, err
		}

		for i := 3; i >= 0; i-- {
			bit := num >> i & 1
			bits = append(bits, int(bit))
		}
	}
	return bits, nil
}

func printBits(bits []int) {
	for _, bit := range bits {
		fmt.Print(bit)
	}
	fmt.Println()
}

func bitsToNumber(bits []int) int {
	num := 0
	for i := 0; i < len(bits); i++ {
		num = num<<1 | bits[i]
	}

	return num
}

func parsePacket(bits []int) (*packet, error) {
	p := &packet{}
	p.version = bitsToNumber(bits[0:3])
	p.packetType = bitsToNumber(bits[3:6])

	if p.packetType == 4 {
		stop := false
		index := 6
		vals := []int{}

		for !stop {
			vals = append(vals, bits[index+1:index+5]...)
			if bits[index] == 0 {
				stop = true
			}
			index += 5
		}
		p.literalValue = bitsToNumber(vals)
	} else {
		var lengthBits int

		switch bits[6] {
		case 0:
			lengthBits = 15
		case 1:
			lengthBits = 11
		default:
			return nil, errors.Errorf("invalid length type ID: %v", bits[6])
		}

		length := bitsToNumber(bits[7 : 7+lengthBits])
	}

	return p, nil
}

func part1(input *puzzleInput) int {
	fmt.Println(input)
	bits, err := getBits(input.hexInput)
	if err != nil {
		panic(err)
	}
	printBits(bits)

	p, err := parsePacket(bits)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(p)
	return 0
}

func part2(input *puzzleInput) int {
	return 0
}

func main() {
	//file, err := os.Open("inputs.txt")
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	input, err := readInput(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(input))
	//fmt.Println(part2(input))
}
