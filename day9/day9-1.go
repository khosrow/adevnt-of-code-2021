package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	row   int
	col   int
	value int
}

type Move struct {
	start  Coord
	finish Coord
}

func atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func parseInput(input string) [][]Coord {
	var board [][]Coord

	// board := make([][]Coord)
	// for j := range board {
	// 	board[j] = make([]int, xmax)
	// }

	// get all the moves
	for row, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		fullrow := make([]Coord, 0)
		for col, s := range line {
			num, _ := strconv.Atoi(string(s))
			fullrow = append(
				fullrow,
				Coord{
					col:   col,
					row:   row,
					value: num,
				})
		}
		board = append(board, fullrow)
	}

	return board
}

func lowestNeighbour(board [][]Coord, row int, col int) bool {
	current := board[row][col].value
	rowMax := len(board) - 1
	colMax := len(board[rowMax]) - 1

	if row == 0 && col == 0 {
		if current < board[row+1][col].value &&
			current < board[row][col+1].value {
			return true
		}
	} else if row == 0 && col == colMax {
		if current < board[row+1][col].value &&
			current < board[row][col-1].value {
			return true
		}
	} else if row == rowMax && col == 0 {
		if current < board[row-1][col].value &&
			current < board[row][col+1].value {
			return true
		}
	} else if row == rowMax && col == colMax {
		if current < board[row-1][col].value &&
			current < board[row][col-1].value {
			return true
		}
	} else if row == 0 {
		if current < board[row][col-1].value &&
			current < board[row][col+1].value &&
			current < board[row+1][col].value {
			return true
		}
	} else if col == 0 {
		if current < board[row-1][col].value &&
			current < board[row+1][col].value &&
			current < board[row][col+1].value {
			return true
		}
	} else if row == rowMax {
		if current < board[row][col-1].value &&
			current < board[row][col+1].value &&
			current < board[row-1][col].value {
			return true
		}
	} else if col == colMax {
		if current < board[row-1][col].value &&
			current < board[row+1][col].value &&
			current < board[row][col-1].value {
			return true
		}
	} else {
		if current < board[row-1][col].value &&
			current < board[row+1][col].value &&
			current < board[row][col-1].value &&
			current < board[row][col+1].value {
			return true
		}
	}
	return false
}

func lowPoints(board [][]Coord) []int {
	var lowpoints []int

	for i := range board {
		for j := range board[i] {
			if lowestNeighbour(board, i, j) {
				lowpoints = append(lowpoints, board[i][j].value)
			}
		}
	}
	return lowpoints
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	board := parseInput(string(input))
	lowpoints := lowPoints(board)
	fmt.Printf("low points: %v\n", lowpoints)
	risk := 0
	for _, num := range lowpoints {
		risk += num + 1
	}
	fmt.Printf("risk: %d\n", risk)
}
