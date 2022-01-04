package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x      int
	y      int
	marked bool
}

type BingoBoard struct {
	numbers    map[int]Coord
	x          [5]int
	y          [5]int
	winningNum int
}

func atoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func parseInput(input string) ([]int, map[int][]int, []BingoBoard) {
	var moves []int
	var bingo []BingoBoard
	dict := make(map[int][]int)

	// get all the moves
	lines := strings.Split(string(input), "\n")
	for _, i := range strings.Split(lines[0], ",") {
		moves = append(moves, atoi(i))
	}

	for i := 2; i < len(lines); i += 6 {

		bingo = append(bingo,
			BingoBoard{
				numbers:    make(map[int]Coord),
				x:          [5]int{0, 0, 0, 0, 0},
				y:          [5]int{0, 0, 0, 0, 0},
				winningNum: -1,
			})
		for j := 0; j < 5; j++ {
			for k, numStr := range strings.Fields(lines[i+j]) {
				num := atoi(numStr)

				bingo[len(bingo)-1].numbers[num] = Coord{x: k, y: j, marked: false}
				dict[num] = append(dict[num], len(bingo)-1)
			}
		}
	}

	return moves, dict, bingo
}

// func getFirstWin(moves []int, dict map[int][]int, bingo []BingoBoard) int {
// 	for _, num := range moves {
// 		boards := dict[num]
// 		for _, board := range boards {
// 			// check if the number is in this board
// 			if coord, ok := bingo[board].numbers[num]; ok {
// 				bingo[board].x[coord.x] += 1
// 				bingo[board].y[coord.y] += 1
// 				bingo[board].numbers[num] = Coord{
// 					x:      coord.x,
// 					y:      coord.y,
// 					marked: true,
// 				}
// 				if bingo[board].x[coord.x] == 5 || bingo[board].y[coord.y] == 5 {
// 					fmt.Println("====Bingo=====")
// 					fmt.Printf("Winning number: %d\n", num)
// 					return getBoardScore(bingo[board], num)
// 				}
// 			}
// 		}
// 	}
// 	return -1
// }

func getLastWin(moves []int, dict map[int][]int, bingo []BingoBoard) int {
	var winners []int
	for _, num := range moves {
		boards := dict[num]
		for _, board := range boards {
			// check if board is already won
			if bingo[board].winningNum == -1 {
				// check if the number is in this board
				if coord, ok := bingo[board].numbers[num]; ok {
					bingo[board].x[coord.x] += 1
					bingo[board].y[coord.y] += 1
					bingo[board].numbers[num] = Coord{
						x:      coord.x,
						y:      coord.y,
						marked: true,
					}
					if bingo[board].x[coord.x] == 5 || bingo[board].y[coord.y] == 5 {
						fmt.Println("==== Last Bingo =====")
						fmt.Printf("Winning number: %d\n", num)
						bingo[board].winningNum = num
						winners = append(winners, board)
					}
				}
			}
		}
	}
	lastWinner := winners[len(winners)-1]
	return getBoardScore(bingo[lastWinner],
		bingo[lastWinner].winningNum)
}

func getBoardScore(board BingoBoard, num int) int {
	sum := 0
	for key, value := range board.numbers {
		if value.marked == false {
			sum += key
		}
	}

	fmt.Printf("sum: %d\n", sum)

	return sum * num
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	moves, dict, boards := parseInput(string(input))
	// color.Blue("hello")
	fmt.Println(moves)
	fmt.Print(dict)
	fmt.Println(boards)

	// score := getFirstWin(moves, dict, boards)
	// fmt.Printf("First Win Score: %d\n", score)

	score := getLastWin(moves, dict, boards)
	fmt.Printf("Last Win score: %d\n", score)

}
