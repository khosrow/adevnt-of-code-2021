package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Fish map[int]int

func atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func parseInput(input string) Fish {
	fish := make(Fish)

	line := strings.Split(string(input), "\n")[0]
	for _, f := range strings.Split(line, ",") {
		num := atoi(f)
		fish[num]++
	}
	return fish
}

func countFish(fish Fish) int {
	var count int
	for day := 1; day <= 256; day++ {
		oldFish := fish
		fish = make(Fish)
		// fmt.Printf("oldFish: %v\n", oldFish)
		for k, v := range oldFish {
			// fmt.Printf("key: %d, value: %d\n", k, v)
			if k == 0 {
				fish[6] += v
				fish[8] += v
			} else {
				fish[k-1] += v
			}
		}
	}
	for _, v := range fish {
		count += v
	}
	return count
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fishList := parseInput(string(input))
	fmt.Printf("%v\n", fishList)
	fishCount := countFish(fishList)
	fmt.Printf("Fish count: %d\n", fishCount)
}
