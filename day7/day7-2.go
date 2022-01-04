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

func calcCost(positions CrabPositions, average int) int {
	var cost float64
	for _, value := range positions {
		// cost += math.Abs(float64(value - average))
		diff := math.Abs(float64(value - average))
		for i := 1.0; i <= diff; i++ {
			cost += i
		}
	}

	return int(cost)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	positions := parseInput(string(input))
	average := 0
	for _, num := range positions {
		average += num
	}
	average = average / len(positions)
	var minCost, cost int
	cost = calcCost(positions, average)
	fmt.Printf("Average: %d, Cost: %d\n", average, cost)
	minCost = cost
	for i := average - 1; i <= average+1; i++ {
		cost = calcCost(positions, i)
		if cost < minCost {
			minCost = cost
		}
		fmt.Printf("Average: %d, Cost: %d\n", i, cost)
	}
	fmt.Printf("Minimum cost: %d\n", minCost)
}
