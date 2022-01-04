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

		var slope float32
		slope = float32((y2 - y1)) / float32((x2 - x1))

		// only keep valid values)
		if x1 == x2 || y1 == y2 || slope == 1 || slope == -1 {
			// if x1 == x2 || y1 == y2 {
			fmt.Printf("%s : %f\n", line, slope)
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
	board := make([][]int, ymax)
	for j := range board {
		board[j] = make([]int, xmax)
	}
	for _, move := range moves {
		if move.start.x == move.finish.x {
			calcVertical(board, move)
		} else if move.start.y == move.finish.y {
			calcHorizontal(board, move)
		} else {
			fmt.Println("calculating diagonal")
			calcDiagonal(board, move)
		}
	}
	return board
}

func calcHorizontal(board [][]int, move Move) {
	var begin, end int
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

func calcVertical(board [][]int, move Move) {
	var begin, end int
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

func calcDiagonal(board [][]int, move Move) {
	if move.start.x <= move.finish.x {
		if move.start.y <= move.finish.y {
			i := move.start.y
			for j := move.start.x; j <= move.finish.x; j++ {
				board[i][j]++
				i++
			}
		} else {
			i := move.start.y
			for j := move.start.x; j <= move.finish.x; j++ {
				board[i][j]++
				i--
			}
		}
	} else {
		if move.start.y <= move.finish.y {
			i := move.start.y
			for j := move.start.x; j >= move.finish.x; j-- {
				board[i][j]++
				i++
			}
		} else {
			i := move.start.y
			for j := move.start.x; j >= move.finish.x; j-- {
				board[i][j]++
				i--
			}
		}
	}
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

	fmt.Println("hello world")
	moves, xmax, ymax := parseInput(string(input))
	fmt.Printf("Board size: %dx%d \n", xmax, ymax)
	board := processMoves(moves, xmax, ymax)
	overlaps := countOverlaps(board)

	// fmt.Println(moves)
	fmt.Printf("Overlaps > 2 : %d\n", overlaps)
}
