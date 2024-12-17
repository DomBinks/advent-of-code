package utils

import "strings"

type Pos struct {
	Y, X int
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Extract[T any](v T, e error) T {
	Check(e)
	return v
}

func GetSections(input string) []string {
	return strings.Split(input, "X")
}

func GetBoard(input string) [][]byte {
	rows := strings.Split(input, "\n")
	board := make([][]byte, len(rows))
	for i := range rows {
		board[i] = make([]byte, len(rows[i]))
		for j := range rows[i] {
			board[i][j] = rows[i][j]
		}
	}

	return board
}

func OnBoard(board [][]byte, x int, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func FindOnBoard(board [][]byte, target byte) Pos {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == target {
				return Pos{i, j}
			}
		}
	}

	return Pos{-1, -1}
}

func CountOnBoard(board [][]byte, target byte) int {
	found := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] == target {
				found += 1
			}
		}
	}

	return found
}

func BoardVFlip(board [][]string) [][]string {
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

func BoardHFlip(board [][]string) [][]string {
	rows := len(board)
	cols := len(board[0])

	newBoard := make([][]string, rows)
	for i := range len(newBoard) {
		newBoard[i] = make([]string, cols)
		copy(newBoard[i], board[len(board)-1-i])
	}

	return newBoard
}

func BoardTranspose(board [][]string) [][]string {
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

func BoardShiftUp(board [][]string) [][]string {
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

func BoardShiftDown(board [][]string) [][]string {
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
