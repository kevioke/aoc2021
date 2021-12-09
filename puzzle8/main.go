package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
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

		entry := logEntry{
			pattern: parseDigits(entryStr[0]),
			output:  parseDigits(entryStr[1]),
		}
		logEntries = append(logEntries, entry)
	}
	return logEntries, nil
}

func parseDigits(str string) []digit {
	digits := []digit{}
	tokens := strings.Fields(str)
	for _, token := range tokens {
		d := digit{}
		for _, segment := range token {
			d.segments = append(d.segments, segment)
		}
		digits = append(digits, d)
	}
	return digits
}

func part1(logEntries []logEntry) int {
	sum := 0
	for _, entry := range logEntries {
		for _, d := range entry.output {
			switch len(d.segments) {
			case 2, 4, 3, 7:
				sum++
			}
		}
	}
	return sum
}

func part2(logEntries []logEntry) int {
	sum := 0
	perms := permutations([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'})
	pMaps := permMaps(perms)

	for _, entry := range logEntries {
		for _, pMap := range pMaps {
			if mappingExists(entry, pMap) {
				sum += decodeSegment(entry.output, pMap)
				break
			}
		}
	}
	return sum
}

func decodeSegment(digits []digit, pMap map[rune]rune) int {
	sum := 0
	for _, digit := range digits {
		segmentMap := map[rune]bool{}
		for _, segment := range digit.segments {
			segmentMap[pMap[segment]] = true
		}
		_, val := validSegments(segmentMap)
		sum *= 10
		sum += val
	}
	return sum
}

func validSegments(segments map[rune]bool) (bool, int) {
	validNumbers := []map[rune]bool{
		// 0
		{
			'a': true,
			'b': true,
			'c': true,
			'e': true,
			'f': true,
			'g': true,
		},
		// 1
		{
			'c': true,
			'f': true,
		},
		// 2
		{
			'a': true,
			'c': true,
			'd': true,
			'e': true,
			'g': true,
		},
		// 3
		{
			'a': true,
			'c': true,
			'd': true,
			'f': true,
			'g': true,
		},
		// 4
		{
			'b': true,
			'c': true,
			'd': true,
			'f': true,
		},
		// 5
		{
			'a': true,
			'b': true,
			'd': true,
			'f': true,
			'g': true,
		},
		// 6
		{
			'a': true,
			'b': true,
			'd': true,
			'e': true,
			'f': true,
			'g': true,
		},
		// 7
		{
			'a': true,
			'c': true,
			'f': true,
		},
		// 8
		{
			'a': true,
			'b': true,
			'c': true,
			'd': true,
			'e': true,
			'f': true,
			'g': true,
		},
		// 9
		{
			'a': true,
			'b': true,
			'c': true,
			'd': true,
			'f': true,
			'g': true,
		},
	}

	for i, validNumber := range validNumbers {
		if reflect.DeepEqual(validNumber, segments) {
			return true, i
		}
	}
	return false, 0
}

func mappingExists(entry logEntry, mapping map[rune]rune) bool {
	digits := append(entry.pattern, entry.output...)
	for _, d := range digits {
		segmentsOn := map[rune]bool{}
		for _, rawSegment := range d.segments {
			segmentsOn[mapping[rawSegment]] = true
		}

		if valid, _ := validSegments(segmentsOn); !valid {
			return false
		}
	}
	return true
}

func remove(s []rune, i int) []rune {
	result := make([]rune, len(s))
	copy(result, s)
	result[i] = result[len(result)-1]
	return result[:len(result)-1]
}

func permMaps(perms [][]rune) []map[rune]rune {
	pMaps := []map[rune]rune{}
	for _, perm := range perms {
		pMap := map[rune]rune{}
		segment := 'a'
		for _, p := range perm {
			pMap[p] = segment
			segment++
		}
		pMaps = append(pMaps, pMap)
	}
	return pMaps
}

func permutations(segments []rune) [][]rune {
	if len(segments) == 0 {
		return nil
	}
	if len(segments) == 1 {
		return [][]rune{segments}
	}

	perms := [][]rune{}
	for i, segment := range segments {
		permRest := permutations(remove(segments, i))

		for _, p := range permRest {
			perm := []rune{segment}
			perms = append(perms, append(perm, p...))
		}
	}
	return perms
}

func main() {
	file, err := os.Open("inputs.txt")
	//file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	logEntries, err := readInput(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(logEntries))
	fmt.Println(part2(logEntries))
}
