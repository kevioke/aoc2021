package main

import (
	"bufio"
	"encoding/json"
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
	Version      int
	PacketType   int
	LiteralValue int
	SubPackets   []*packet
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

func parsePacket(bits []int) (*packet, int, error) {
	p := &packet{}
	p.Version = bitsToNumber(bits[0:3])
	p.PacketType = bitsToNumber(bits[3:6])
	var lastIndex int

	if p.PacketType == 4 {
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
		lastIndex = index
		p.LiteralValue = bitsToNumber(vals)
	} else {
		switch bits[6] {
		case 0:
			subPacketsStartIdx := 7 + 15
			subPacketsIdx := subPacketsStartIdx
			lengthSubPackets := bitsToNumber(bits[7:subPacketsStartIdx])

			for subPacketsIdx < subPacketsStartIdx+lengthSubPackets {
				subPacketBits := bits[subPacketsIdx:]
				subPacket, idx, err := parsePacket(subPacketBits)
				if err != nil {
					return nil, 0, err
				}
				subPacketsIdx += idx
				lastIndex = subPacketsIdx
				p.SubPackets = append(p.SubPackets, subPacket)
			}
		case 1:
			subPacketsIdx := 7 + 11
			numSubPackets := bitsToNumber(bits[7:subPacketsIdx])

			for i := 0; i < numSubPackets; i++ {
				subPacketBits := bits[subPacketsIdx:]
				subPacket, idx, err := parsePacket(subPacketBits)
				if err != nil {
					return nil, 0, err
				}
				subPacketsIdx += idx
				lastIndex = subPacketsIdx
				p.SubPackets = append(p.SubPackets, subPacket)
			}
		default:
			return nil, 0, errors.Errorf("invalid length type ID: %v", bits[6])
		}
	}

	return p, lastIndex, nil
}

func sumVersions(p *packet) int {
	sum := p.Version

	for _, subPacket := range p.SubPackets {
		sum += sumVersions(subPacket)
	}

	return sum
}

func eval(p *packet) int {
	switch p.PacketType {
	case 4:
		return p.LiteralValue
	case 0:
		sum := 0
		for _, subP := range p.SubPackets {
			sum += eval(subP)
		}
		return sum
	case 1:
		product := 1
		for _, subP := range p.SubPackets {
			product *= eval(subP)
		}
		return product
	case 2:
		min := eval(p.SubPackets[0])
		for _, subP := range p.SubPackets {
			val := eval(subP)
			if val < min {
				min = val
			}
		}
		return min
	case 3:
		max := eval(p.SubPackets[0])
		for _, subP := range p.SubPackets {
			val := eval(subP)
			if val > max {
				max = val
			}
		}
		return max
	case 5:
		first := eval(p.SubPackets[0])
		second := eval(p.SubPackets[1])
		if first > second {
			return 1
		}
		return 0
	case 6:
		first := eval(p.SubPackets[0])
		second := eval(p.SubPackets[1])
		if first < second {
			return 1
		}
		return 0
	case 7:
		first := eval(p.SubPackets[0])
		second := eval(p.SubPackets[1])
		if first == second {
			return 1
		}
		return 0
	}

	panic("invalid packet")
	return -1
}

func part1(input *puzzleInput) int {
	bits, err := getBits(input.hexInput)
	if err != nil {
		panic(err)
	}

	p, _, err := parsePacket(bits)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}

	//b, err := json.MarshalIndent(p, "", "  ")
	//fmt.Println(string(b))
	return sumVersions(p)
}

func part2(input *puzzleInput) int {
	bits, err := getBits(input.hexInput)
	if err != nil {
		panic(err)
	}

	p, _, err := parsePacket(bits)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}

	b, err := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(b))

	return eval(p)
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
