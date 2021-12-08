package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type digit struct {
	segments []rune
}

type logEntry struct {
	pattern []digit
	output  []digit
}

func readInput(r io.Reader) ([]logEntry, error) {
	scanner := bufio.NewScanner(r)
	logEntries := []logEntry{}
	for scanner.Scan() {
		entryStr := strings.Split(scanner.Text(), "|")
		strings.Fields(entryStr[0])
	}
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

func main() {
	//file, err := os.Open("inputs.txt")
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	logEntries, err := readInput(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(logEntries)
}
