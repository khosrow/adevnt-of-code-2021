package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Digit string

func atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

const (
	one   = "cf"
	four  = "bcdf"
	seven = "acf"
	eight = "abcdefg"
)

func parseInput(input string) (signals []string, outputs []string) {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		for _, sig := range strings.Split(parts[0], " ") {
			signals = append(signals, sig)
		}

		for _, sig := range strings.Split(parts[1], " ") {
			outputs = append(outputs, sig)
		}
	}

	return signals, outputs
}

func checkOutput(outputs []string) int {
	count := 0
	for _, o := range outputs {
		switch len(o) {
		case len(one):
			count++
		case len(four):
			count++
		case len(seven):
			count++
		case len(eight):
			count++
		}
	}
	return count
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	_, outputs := parseInput(string(input))
	// fmt.Printf("signals: %v\n", signals)
	// fmt.Printf("outputs: %v\n", outputs)
	count := checkOutput(outputs)
	fmt.Printf("Total number of 1, 4, 7, or 8: %d\n", count)
}
