package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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
	data, err := os.ReadFile("day12.txt")
	check(err)
	input := string(data)
	part := getPart()

	if part == "1" {
		fmt.Println(part1(input))
	} else if part == "2" {
		fmt.Println(part2(input))
	}
}

func getBoard(input string) [][]byte {
	rows := strings.Split(input, "\n")
	board := make([][]byte, len(rows))
	for i := range rows {
		board[i] = make([]byte, len(rows[i]))
		for j := range board[i] {
			board[i][j] = rows[i][j]
		}
	}

	return board
}

func onBoard(board [][]byte, y int, x int) bool {
	return y >= 0 && y < len(board) && x >= 0 && x < len(board[0])
}

func flood1(board [][]byte, y int, x int, plant byte) (int, int) {
	board[y][x] = board[y][x] + 32

	area := 1
	perimeter := 0

	if onBoard(board, y-1, x) {
		if board[y-1][x] == plant {
			a, p := flood1(board, y-1, x, plant)
			area += a
			perimeter += p
		}
		if board[y-1][x] != plant+32 {
			perimeter += 1
		}
	} else {
		perimeter += 1
	}
	if onBoard(board, y+1, x) {
		if board[y+1][x] == plant {
			a, p := flood1(board, y+1, x, plant)
			area += a
			perimeter += p
		}
		if board[y+1][x] != plant+32 {
			perimeter += 1
		}
	} else {
		perimeter += 1
	}
	if onBoard(board, y, x-1) {
		if board[y][x-1] == plant {
			a, p := flood1(board, y, x-1, plant)
			area += a
			perimeter += p
		}
		if board[y][x-1] != plant+32 {
			perimeter += 1
		}
	} else {
		perimeter += 1
	}
	if onBoard(board, y, x+1) {
		if board[y][x+1] == plant {
			a, p := flood1(board, y, x+1, plant)
			area += a
			perimeter += p
		}
		if board[y][x+1] != plant+32 {
			perimeter += 1
		}
	} else {
		perimeter += 1
	}

	return area, perimeter
}

func part1(input string) string {
	board := getBoard(input)
	total := 0

	for i := range board {
		for j := range board[i] {
			if unicode.IsUpper(rune(board[i][j])) {
				area, perimeter := flood1(board, i, j, board[i][j])
				total += area * perimeter
			}
		}
	}

	out := strconv.Itoa(total)
	return out
}

func flood2(board [][]byte, y int, x int, plant byte, left []int) int {
	board[y][x] = board[y][x] + 32

	area := 1

	if onBoard(board, y-1, x) {
		if board[y-1][x] == plant {
			a := flood2(board, y-1, x, plant, left)
			area += a
		}
		if board[y-1][x] != plant+32 {
			board[y-1][x] = '#'
		}
	}
	if onBoard(board, y+1, x) {
		if board[y+1][x] == plant {
			a := flood2(board, y+1, x, plant, left)
			area += a
		}
		if board[y+1][x] != plant+32 {
			board[y+1][x] = '#'
		}
	}
	if onBoard(board, y, x-1) {
		if board[y][x-1] == plant {
			a := flood2(board, y, x-1, plant, left)
			area += a
		}
		if board[y][x-1] != plant+32 {
			board[y][x-1] = '#'
			if x-1 < left[1] {
				left[0] = y
				left[1] = x - 1
			}
		}
	}
	if onBoard(board, y, x+1) {
		if board[y][x+1] == plant {
			a := flood2(board, y, x+1, plant, left)
			area += a
		}
		if board[y][x+1] != plant+32 {
			board[y][x+1] = '#'
		}
	}

	if onBoard(board, y-1, x-1) && board[y-1][x-1] != plant && board[y-1][x-1] != plant+32 {
		board[y-1][x-1] = '#'
	}
	if onBoard(board, y+1, x-1) && board[y+1][x-1] != plant && board[y+1][x-1] != plant+32 {
		board[y+1][x-1] = '#'
	}
	if onBoard(board, y-1, x+1) && board[y-1][x+1] != plant && board[y-1][x+1] != plant+32 {
		board[y-1][x+1] = '#'
	}
	if onBoard(board, y+1, x+1) && board[y+1][x+1] != plant && board[y+1][x+1] != plant+32 {
		board[y+1][x+1] = '#'
	}

	return area
}

func nextD(dy int, dx int) int {
	if dy == -1 {
		return 0
	}
	if dy == 1 {
		return 1
	}
	if dx == -1 {
		return 2
	}
	if dx == 1 {
		return 3
	}

	return -1
}

func sides(board [][]byte, startY int, startX int) int {
	s := 0
	v := true
	y := startY
	x := startX
	d := 0

	dirs := [][][]int{{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}, // U
		{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}, // D
		{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}, // L
		{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}} // R

	for {
		//fmt.Println("d ", d, " y ", y, " x ", x)
		board[y][x] = '~'

		for i := range dirs {
			if (d == 2 || d == 3) && startY == y-1 && startX == x ||
				(d == 2 || d == 3) && startY == y+1 && startX == x ||
				(d == 0 || d == 1) && startY == y && startX == x-1 ||
				(d == 0 || d == 1) && startY == y && startX == x+1 {

				return s + 1
			}
			if (d == 0) && startY == y-1 && startX == x ||
				(d == 1) && startY == y+1 && startX == x ||
				(d == 2) && startY == y && startX == x-1 ||
				(d == 3) && startY == y && startX == x+1 {

				return s
			}

			dir := dirs[d][i]
			if onBoard(board, y+dir[0], x+dir[1]) && board[y+dir[0]][x+dir[1]] == '#' {
				if dir[0] != 0 && !v {
					s += 1
					v = true
				}
				if dir[1] != 0 && v {
					s += 1
					v = false
				}

				d = nextD(dir[0], dir[1])
				y += dir[0]
				x += dir[1]
				break
			}
		}
	}

	return -1
}

func part2(input string) string {
	boardInside := getBoard(input)
	board := make([][]byte, len(boardInside)+2)
	for i := range boardInside {
		board[i+1] = make([]byte, len(boardInside[i])+2)
		for j := range boardInside[i] {
			board[i+1][j+1] = boardInside[i][j]
		}
	}
	board[0] = make([]byte, len(boardInside)+2)
	board[len(board)-1] = make([]byte, len(boardInside)+2)
	for i := range board {
		board[i][0] = '0'
		board[i][len(board[i])-1] = '0'
	}
	for i := range board[0] {
		board[0][i] = '0'
		board[len(board)-1][i] = '0'
	}

	total := 0

	for i := range board {
		for j := range board[i] {
			fmt.Println("i ", i, " j ", j)
			if unicode.IsUpper(rune(board[i][j])) {
				boardC := make([][]byte, len(board))
				for k := range boardC {
					boardC[k] = make([]byte, len(board[i]))
					copy(boardC[k], board[k])
				}

				left := []int{i, j - 1}
				area := flood2(boardC, i, j, boardC[i][j], left)

				boardD := make([][]byte, len(boardC)*2)
				for k := range boardD {
					boardD[k] = make([]byte, len(boardC[i])*2)
				}

				for m := range boardC {
					for n := range boardC {
						boardD[m*2][n*2] = boardC[m][n]
						boardD[(m*2)+1][n*2] = boardC[m][n]
						boardD[(m*2)+1][(n*2)+1] = boardC[m][n]
						boardD[m*2][(n*2)+1] = boardC[m][n]
					}
				}

				/*
					for k := range boardD {
						for l := range boardD[k] {
							fmt.Print(string(boardD[k][l]))
						}
						fmt.Println("")
					}
				*/

				s := sides(boardD, (left[0]*2)+1, (left[1]*2)+1)

				//fmt.Println(string(board[i][j]), " a ", area, " s ", s)
				total += area * s

				flood1(board, i, j, board[i][j])
			}
		}
	}

	out := strconv.Itoa(total)
	return out
}
