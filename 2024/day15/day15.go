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
	data, err := os.ReadFile("day15.txt")
	check(err)
	input := string(data)
	part := getPart()

	if part == "1" {
		fmt.Println(part1(input))
	} else if part == "2" {
		fmt.Println(part2(input))
	}
}

type pos struct {
	y, x int
}

func onBoard(board [][]byte, y int, x int) bool {
	return y > 0 && y < len(board)-1 && x > 0 && x < len(board[0])-1
}

func move1(board [][]byte, y int, x int, dy int, dx int) bool {
	ny := y + dy
	nx := x + dx

	if board[ny][nx] == '#' {
		return false
	} else if board[ny][nx] == '.' {
		board[ny][nx] = board[y][x]
		return true
	} else if move1(board, ny, nx, dy, dx) {
		board[ny][nx] = board[y][x]
		return true
	} else {
		return false
	}
}

func part1(input string) string {
	bc := strings.Split(input, "X")
	b := strings.Split(bc[0], "\n")
	cs := strings.Join(strings.Split(bc[1], "\n"), "")

	board := make([][]byte, len(b))
	robot := pos{-1, -1}

	for i := range b {
		board[i] = make([]byte, len(b[i]))
		for j := range b[i] {
			board[i][j] = b[i][j]

			if b[i][j] == '@' {
				robot.y = i
				robot.x = j
			}
		}
	}

	for k := range cs {
		c := cs[k]
		dy := 0
		dx := 0

		switch c {
		case '^':
			dy = -1
		case 'v':
			dy = 1
		case '<':
			dx = -1
		case '>':
			dx = 1
		}

		if onBoard(board, robot.y+dy, robot.x+dx) {
			if board[robot.y+dy][robot.x+dx] != '#' &&
				move1(board, robot.y, robot.x, dy, dx) {

				board[robot.y][robot.x] = '.'
				robot.y = robot.y + dy
				robot.x = robot.x + dx
			}
		}
	}

	total := 0

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				total += (100 * i) + j
			}
		}
	}

	out := strconv.Itoa(total)
	return out
}

func canMove2(board [][]byte, y int, x int, dy int, dx int) bool {
	if board[y][x] == '.' {
		return true
	} else if board[y][x] == '[' {
		return canMove2(board, y+dy, x+dx, dy, dx) && canMove2(board, y+dy, x+1+dx, dy, dx)
	} else if board[y][x] == ']' {
		return canMove2(board, y+dy, x+dx, dy, dx) && canMove2(board, y+dy, x-1+dx, dy, dx)
	} else {
		return false
	}
}

func move2(board [][]byte, y int, x int, dy int, dx int) bool {
	ny := y + dy
	nx := x + dx

	if board[ny][nx] == '#' {
		return false
	} else if board[ny][nx] == '.' {
		board[ny][nx] = board[y][x]
		return true
	} else if dy == 0 && move2(board, ny, nx, dy, dx) {
		board[ny][nx] = board[y][x]
		return true
	} else if board[ny][nx] == '[' || board[ny][nx] == ']' {
		o := 0
		if board[ny][nx] == '[' {
			o = 1
		} else {
			o = -1
		}

		if canMove2(board, ny, nx, dy, dx) && canMove2(board, ny, nx+o, dy, dx) {
			move2(board, ny, nx+o, dy, dx)
			move2(board, ny, nx, dy, dx)
			board[ny][nx] = board[y][x]
			board[ny][nx+o] = '.'

			return true
		}
	} else {
		return false
	}

	return false
}

func part2(input string) string {
	bc := strings.Split(input, "X")
	b := strings.Split(bc[0], "\n")
	cs := strings.Join(strings.Split(bc[1], "\n"), "")

	board := make([][]byte, len(b))
	robot := pos{-1, -1}

	for i := range b {
		board[i] = make([]byte, len(b[i])*2)
		for j := range b[i] {
			switch b[i][j] {
			case '@':
				board[i][j*2] = '@'
				board[i][(j*2)+1] = '.'
				robot.y = i
				robot.x = j * 2
			case '#':
				board[i][j*2] = '#'
				board[i][(j*2)+1] = '#'
			case 'O':
				board[i][j*2] = '['
				board[i][(j*2)+1] = ']'
			case '.':
				board[i][j*2] = '.'
				board[i][(j*2)+1] = '.'
			}
		}
	}

	for k := range cs {
		c := cs[k]
		dy := 0
		dx := 0

		switch c {
		case '^':
			dy = -1
		case 'v':
			dy = 1
		case '<':
			dx = -1
		case '>':
			dx = 1
		}

		if onBoard(board, robot.y+dy, robot.x+dx) {
			if board[robot.y+dy][robot.x+dx] != '#' &&
				move2(board, robot.y, robot.x, dy, dx) {
				board[robot.y][robot.x] = '.'
				robot.y = robot.y + dy
				robot.x = robot.x + dx
			}
		}
	}

	total := 0

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '[' {
				total += (100 * i) + j
			}
		}
	}

	out := strconv.Itoa(total)
	return out
}
