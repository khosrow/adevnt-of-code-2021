package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Move struct {
	start  Coord
	finish Coord
}

type CrabPositions []int

func atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func parseInput(input string) CrabPositions {
	var positions CrabPositions
	line := strings.Split(string(input), "\n")[0]
	for _, num := range strings.Split(line, ",") {
		p := atoi(num)
		positions = append(positions, p)
	}
	sort.Ints(positions)

	return positions
}

func calcCost(positions CrabPositions, median int) int {
	var cost float64
	for _, value := range positions {
		cost += math.Abs(float64(value - positions[median]))
	}

	return int(cost)
}

func main() {
	input, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	positions := parseInput(string(input))
	average := 0
	for _, num := range positions {
		average += num
	}
	average = average / len(positions)
	fmt.Printf("Average: %d\n", average)
	median := len(positions) / 2
	fmt.Printf("median value: %d\n", positions[median])
	cost := calcCost(positions, median)
	fmt.Printf("Median: %d, Cost: %d\n", positions[median], cost)
	// for i := median - 1; i <= median+1; i++ {
	// 	cost := calcCost(positions, i)
	// 	fmt.Printf("Median: %d, Cost: %d\n", positions[i], cost)
	// }
}
