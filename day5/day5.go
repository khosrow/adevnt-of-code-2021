package main

import (
	"fmt"
	"os"
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

func atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func parseInput(input string) (moves []Move, xMax int, yMax int) {
	xMax = 0
	yMax = 0

	// get all the moves
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		s := strings.Split(line, " -> ")

		begin := strings.Split(s[0], ",")
		x1 := atoi(begin[0])
		y1 := atoi(begin[1])

		end := strings.Split(s[1], ",")
		x2 := atoi(end[0])
		y2 := atoi(end[1])

		if x1 > xMax {
			xMax = x1
		} else if x2 > xMax {
			xMax = x2
		}

		if y1 > yMax {
			yMax = y1
		} else if y2 > yMax {
			yMax = y2
		}

		// only keep valid values
		if x1 == x2 || y1 == y2 {
			moves = append(moves, Move{
				start: Coord{
					x: x1,
					y: y1,
				},
				finish: Coord{
					x: x2,
					y: y2,
				},
			})
		}
	}
	return moves, xMax + 1, yMax + 1
}

func processMoves(moves []Move, xmax int, ymax int) [][]int {
	var begin, end int
	board := make([][]int, ymax)
	for j := range board {
		board[j] = make([]int, xmax)
	}
	for _, move := range moves {

		// if same X
		if move.start.x == move.finish.x {
			// fmt.Println(move)
			if move.start.y <= move.finish.y {
				begin = move.start.y
				end = move.finish.y
			} else {
				begin = move.finish.y
				end = move.start.y
			}
			for i := begin; i <= end; i++ {
				board[i][move.start.x]++
			}
		}

		// if same Y
		if move.start.y == move.finish.y {
			// fmt.Println(move)
			if move.start.x <= move.finish.x {
				begin = move.start.x
				end = move.finish.x
			} else {
				begin = move.finish.x
				end = move.start.x
			}
			for i := begin; i <= end; i++ {
				// fmt.Printf("x(i)=%d,y=%d", i, move.start.y)
				board[move.start.y][i]++
			}
		}
	}
	return board
}

func countOverlaps(board [][]int) int {
	overlaps := 0
	for _, b := range board {
		// fmt.Printf("%v\n", board[i])
		for _, j := range b {
			if j > 1 {
				overlaps++
			}
		}
	}

	return overlaps
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	moves, xmax, ymax := parseInput(string(input))
	fmt.Printf("Board size: %dx%d \n", xmax, ymax)
	board := processMoves(moves, xmax, ymax)
	overlaps := countOverlaps(board)

	// fmt.Println(moves)
	fmt.Printf("Overlaps > 2 : %d\n", overlaps)
}
