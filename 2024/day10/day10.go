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
	data, err := os.ReadFile("day10.txt")
	check(err)
	input := string(data)
	part := getPart()

	if part == "1" {
		fmt.Println(part1(input))
	} else if part == "2" {
		fmt.Println(part2(input))
	}
}

func getBoard(input string) [][]int {
	lines := strings.Split(input, "\n")
	board := make([][]int, len(lines))
	for i := range lines {
		board[i] = make([]int, len(lines[i]))
		for j := range lines[i] {
			var ij int
			if string(lines[i][j]) == "." {
				ij = -1
			} else {
				var err error
				ij, err = strconv.Atoi(string(lines[i][j]))
				check(err)
			}

			board[i][j] = ij
		}
	}

	return board
}

func onBoard(board [][]int, y int, x int) bool {
	return y >= 0 && y < len(board) && x >= 0 && x < len(board[0])
}

func search1(board [][]int, reached map[string]bool, y int, x int, num int) int {
	if num == 9 && !reached[strconv.Itoa(y)+","+strconv.Itoa(x)] {
		reached[strconv.Itoa(y)+","+strconv.Itoa(x)] = true
		return 1
	}

	total := 0
	if onBoard(board, y-1, x) && board[y-1][x] == num+1 {
		total += search1(board, reached, y-1, x, num+1)
	}
	if onBoard(board, y+1, x) && board[y+1][x] == num+1 {
		total += search1(board, reached, y+1, x, num+1)
	}
	if onBoard(board, y, x-1) && board[y][x-1] == num+1 {
		total += search1(board, reached, y, x-1, num+1)
	}
	if onBoard(board, y, x+1) && board[y][x+1] == num+1 {
		total += search1(board, reached, y, x+1, num+1)
	}

	return total
}

func part1(input string) string {
	board := getBoard(input)
	total := 0

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				reached := make(map[string]bool)
				total += search1(board, reached, i, j, 0)
			}
		}
	}

	out := strconv.Itoa(total)
	return out
}

func search2(board [][]int, y int, x int, num int) int {
	if num == 9 {
		return 1
	}

	total := 0
	if onBoard(board, y-1, x) && board[y-1][x] == num+1 {
		total += search2(board, y-1, x, num+1)
	}
	if onBoard(board, y+1, x) && board[y+1][x] == num+1 {
		total += search2(board, y+1, x, num+1)
	}
	if onBoard(board, y, x-1) && board[y][x-1] == num+1 {
		total += search2(board, y, x-1, num+1)
	}
	if onBoard(board, y, x+1) && board[y][x+1] == num+1 {
		total += search2(board, y, x+1, num+1)
	}

	return total
}

func part2(input string) string {
	board := getBoard(input)
	total := 0

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				total += search2(board, i, j, 0)
			}
		}
	}

	out := strconv.Itoa(total)
	return out
}
