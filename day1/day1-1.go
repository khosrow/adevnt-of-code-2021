package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.ReadFile("input2.txt")
	if err != nil {
		panic(err)
	}
	numbers := strings.Split(string(input), "\n")

	var measurements []int
	for _, number := range numbers {
		if number == "" {
			continue
		}
		measure, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		measurements = append(measurements, measure)
	}

	increased := 0
	prev := measurements[0]
	fmt.Printf("%d (N/A - no previous measurement)\n", prev)

	for _, current := range measurements[1:] {
		if current > prev {
			increased += 1
			fmt.Printf("%d (increased)\n", current)
		} else {
			fmt.Printf("%d (decreased)\n", current)
		}
		prev = current
	}

	fmt.Printf("There are %d measurement increases\n", increased)
}
