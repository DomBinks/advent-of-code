package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getPart() string {
	args := os.Args[1:]

	if len(args) == 0 || (args[0] != "1" && args[0] != "2") {
		panic(errors.New("no part specified"))
	} else {
		return args[0]
	}
}

func main() {
	data, err := os.ReadFile("day04.txt")
	check(err)
	input := string(data)
	part := getPart()

	if part == "1" {
		fmt.Println(part1(input))
	} else if part == "2" {
		fmt.Println(part2(input))
	}
}

func flip(board [][]string) [][]string {
	rows := len(board)
	cols := len(board[0])

	newBoard := make([][]string, rows)
	for i := range len(newBoard) {
		newBoard[i] = make([]string, cols)
	}

	for i := range rows {
		for j := range cols {
			newBoard[i][cols-1-j] = board[i][j]
		}
	}

	return newBoard
}

func transpose(board [][]string) [][]string {
	rows := len(board)
	cols := len(board[0])

	newBoard := make([][]string, cols)
	for i := range len(newBoard) {
		newBoard[i] = make([]string, rows)
	}

	for i := range rows {
		for j := range cols {
			newBoard[j][i] = board[i][j]
		}
	}

	return newBoard
}

func shiftUp(board [][]string) [][]string {
	rows := len(board)
	cols := len(board[0])

	newBoard := make([][]string, rows*2)
	for i := range len(newBoard) {
		newBoard[i] = make([]string, cols)
	}

	for i := range rows {
		for j := range cols {
			newI := i - j
			newBoard[newI+rows][j] = board[i][j]
		}
	}

	return newBoard
}

func shiftDown(board [][]string) [][]string {
	rows := len(board)
	cols := len(board[0])

	newBoard := make([][]string, rows*2)
	for i := range len(newBoard) {
		newBoard[i] = make([]string, cols)
	}

	for i := range rows {
		for j := range cols {
			newI := i + j
			newBoard[newI][j] = board[i][j]
		}
	}

	return newBoard
}

func search(board [][]string) int {
	total := 0

	for i := range len(board) {
		for j := range len(board[0]) - 3 {
			if board[i][j]+board[i][j+1]+board[i][j+2]+board[i][j+3] == "XMAS" {
				total += 1
			}
		}
	}

	return total
}

func part1(input string) string {
	total := 0

	rows := strings.Split(input, "\n")

	board := make([][]string, len(rows))
	transposeBoard := make([][]string, len(rows))
	shiftUpBoard := make([][]string, len(rows))
	shiftDownBoard := make([][]string, len(rows))
	for i := range len(rows) {
		board[i] = strings.Split(rows[i], "")
		transposeBoard[i] = strings.Split(rows[i], "")
		shiftUpBoard[i] = strings.Split(rows[i], "")
		shiftDownBoard[i] = strings.Split(rows[i], "")
	}

	total += search(board)
	board = flip(board)
	total += search(board)

	transposeBoard = transpose(transposeBoard)
	total += search(transposeBoard)
	transposeBoard = flip(transposeBoard)
	total += search(transposeBoard)

	shiftUpBoard = shiftUp(shiftUpBoard)
	total += search(shiftUpBoard)
	shiftUpBoard = flip(shiftUpBoard)
	total += search(shiftUpBoard)

	shiftDownBoard = shiftDown(shiftDownBoard)
	total += search(shiftDownBoard)
	shiftDownBoard = flip(shiftDownBoard)
	total += search(shiftDownBoard)

	out := strconv.Itoa(total)
	return out
}

func part2(input string) string {
	total := 0

	rows := strings.Split(input, "\n")

	board := make([][]string, len(rows))
	for i := range len(rows) {
		board[i] = strings.Split(rows[i], "")
	}

	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[0])-1; j++ {
			if board[i][j] == "A" {
				if (board[i-1][j-1] == "M" && board[i+1][j+1] == "S") ||
					(board[i-1][j-1] == "S" && board[i+1][j+1] == "M") {
					if (board[i+1][j-1] == "M" && board[i-1][j+1] == "S") ||
						(board[i+1][j-1] == "S" && board[i-1][j+1] == "M") {

						total += 1
					}
				}
			}
		}
	}

	out := strconv.Itoa(total)
	return out
}
