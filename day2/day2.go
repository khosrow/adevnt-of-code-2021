package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type move struct {
	direction string
	amount    int
}

func parseMoves(input string) ([]move, error) {
	var movements []move
	for _, line := range strings.Split(string(input), "\n") {
		s := strings.Fields(line)
		if len(s) == 2 {
			// no error checking commands for now
			c := s[0]
			n, err := strconv.Atoi(s[1])
			if err != nil {
				return nil, err
			}
			movements = append(movements, move{direction: c, amount: n})
		}
	}

	return movements, nil
}

func calcMovements(movements []move) (horizontal int, depth int) {
	h := 0
	d := 0
	a := 0

	for _, c := range movements {
		switch c.direction {
		case "forward":
			h += c.amount
			d += a * c.amount
		case "down":
			a += c.amount
		case "up":
			a -= c.amount
		}
	}
	return h, d
}

func main() {

	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	movements, err := parseMoves(string(input))
	if err != nil {
		panic(err)
	}

	horizontal, depth := calcMovements(movements)
	// fmt.Printf("%v\n", movements)
	fmt.Printf("h=%d, d=%d, h*d=%d\n", horizontal, depth, horizontal*depth)
}
