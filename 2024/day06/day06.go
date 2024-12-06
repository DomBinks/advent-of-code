package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
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
	data, err := os.ReadFile("day06.txt")
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
		board[i] = []byte(rows[i])
	}

	return board
}

func findStart(board [][]byte) (int, int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '^' {
				return i, j
			}
		}
	}

	return -1, -1
}

func onBoard(board [][]byte, y int, x int) bool {
	return y >= 0 && y < len(board) && x >= 0 && x < len(board[0])
}

func getNewYX(y int, x int, direction byte) (int, int) {
	newY := y
	newX := x

	switch direction {
	case 'U':
		newY = y - 1
	case 'D':
		newY = y + 1
	case 'L':
		newX = x - 1
	case 'R':
		newX = x + 1
	}

	return newY, newX
}

func changeDirection(direction byte) byte {
	switch direction {
	case 'U':
		return 'R'
	case 'R':
		return 'D'
	case 'D':
		return 'L'
	case 'L':
		return 'U'
	}

	return direction
}

func blocked(board [][]byte, y int, x int) bool {
	return board[y][x] == '#'
}

func marked(board [][]byte, y int, x int) bool {
	return board[y][x] == 'X'
}

func mark(board [][]byte, y int, x int) [][]byte {
	board[y][x] = 'X'
	return board
}

func part1(input string) string {
	board := getBoard(input)
	direction := byte('U')
	total := 0

	y, x := findStart(board)
	mark(board, y, x)
	total += 1

	for onBoard(board, y, x) {
		newY, newX := getNewYX(y, x, direction)

		if onBoard(board, newY, newX) {
			if blocked(board, newY, newX) {
				direction = changeDirection(direction)
				continue
			} else if !marked(board, newY, newX) {
				mark(board, newY, newX)
				total += 1
			}
		}

		y = newY
		x = newX
	}

	out := strconv.Itoa(total)
	return out
}

func addCollision(collisions map[string][]byte, y int, x int, direction byte) {
	key := strconv.Itoa(y) + "-" + strconv.Itoa(x)
	_, in := collisions[key]
	if !in {
		collisions[key] = make([]byte, 0)
	}
	collisions[key] = append(collisions[key], direction)
}

func inCollisions(collisions map[string][]byte, y int, x int, direction byte) bool {
	key := strconv.Itoa(y) + "-" + strconv.Itoa(x)
	dirs, in := collisions[key]
	if in && slices.Contains(dirs, direction) {
		return true
	} else {
		return false
	}
}

func loops(board [][]byte, direction byte, collisions map[string][]byte, y int, x int) bool {
	for onBoard(board, y, x) {
		newY, newX := getNewYX(y, x, direction)

		if onBoard(board, newY, newX) {
			if blocked(board, newY, newX) {
				if inCollisions(collisions, newY, newX, direction) {
					return true
				} else {
					addCollision(collisions, newY, newX, direction)
					direction = changeDirection(direction)
					continue
				}
			}
		}

		y = newY
		x = newX
	}

	return false
}

func part2(input string) string {
	board := getBoard(input)
	direction := byte('U')
	collisions := make(map[string][]byte)
	total := 0

	startY, startX := findStart(board)
	y := startY
	x := startX

	for onBoard(board, y, x) {
		mark(board, y, x)
		newY, newX := getNewYX(y, x, direction)

		if onBoard(board, newY, newX) {
			if blocked(board, newY, newX) {
				addCollision(collisions, newY, newX, direction)
				direction = changeDirection(direction)
				continue
			} else if !marked(board, newY, newX) {
				boardC := make([][]byte, len(board))
				for i := range boardC {
					boardC[i] = make([]byte, len(board[i]))
					for j := range boardC[i] {
						boardC[i][j] = board[i][j]
					}
				}
				boardC[newY][newX] = '#'

				if loops(boardC, 'U', make(map[string][]byte), startY, startX) {
					total += 1
				}
			}
		}

		y = newY
		x = newX
	}

	out := strconv.Itoa(total)
	return out
}
