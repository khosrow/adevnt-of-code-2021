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

type Basin []Coord

func parseInput(input string) [][]Coord {
	var board [][]Coord

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

func lowPoints(board [][]Coord) []Coord {
	var lowpoints []Coord

	for i := range board {
		for j := range board[i] {
			if lowestNeighbour(board, i, j) {
				lowpoints = append(lowpoints, board[i][j])
			}
		}
	}
	return lowpoints
}

func getBasins(lowpoints []Coord, board [][]Coord) []Basin {
	var basins []Basin

	for _, point := range lowpoints {
		row := point.row
		col := point.col
		basin := make([]Basin, 0)
	}

	return basins
}

func main() {
	input, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	board := parseInput(string(input))
	lowpoints := lowPoints(board)
	fmt.Printf("low points: %v\n", lowpoints)
	basins := getBasins(lowpoints, board)
	risk := 0
	for _, point := range lowpoints {
		risk += point.value + 1
	}
	fmt.Printf("risk: %d\n", risk)
}
