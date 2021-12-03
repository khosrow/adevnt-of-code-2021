package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMeasurements(input string) ([]int, error) {
	numbers := strings.Split(string(input), "\n")

	var measurements []int
	for _, number := range numbers {
		if number == "" {
			continue
		}
		measure, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		measurements = append(measurements, measure)
	}

	return measurements, nil
}

func depthIncreaseCount(measurements []int) int {
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
	return increased
}

func sumArray(array []int) int {
	res := 0
	for _, n := range array {
		res += n
	}

	return res
}

func main() {

	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	measurements, err := getMeasurements(string(input))
	if err != nil {
		panic(err)
	}

	var slidingWindow []int
	for i := 0; i < len(measurements)-2; i++ {
		// fmt.Printf("%v\n", measurements[i:i+3])
		num := sumArray(measurements[i : i+3])
		slidingWindow = append(slidingWindow, num)
	}

	increased := depthIncreaseCount(slidingWindow)
	fmt.Printf("There are %d measurement increases\n", increased)
}
