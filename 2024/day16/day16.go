package main

import (
	"errors"
	"fmt"
	"maps"
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
	data, err := os.ReadFile("day16.txt")
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
		for j := range rows[i] {
			board[i][j] = rows[i][j]
		}
	}

	return board
}

func find(board [][]byte, target byte) (int, int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == target {
				return i, j
			}
		}
	}

	return -1, -1
}

func copyBoard(board [][]byte) [][]byte {
	cb := make([][]byte, len(board))
	for i := range board {
		cb[i] = make([]byte, len(board[i]))
		for j := range board[i] {
			cb[i][j] = board[i][j]
		}
	}

	return cb
}

func candidate(board [][]byte, y int, x int) bool {
	return y > 0 && y < len(board)-1 && x > 0 && x < len(board[0])-1 &&
		board[y][x] != '#' && board[y][x] != '~'
}

type pos struct {
	y, x int
}

type ci struct {
	y, x int
	d    byte
}

func findPath(board [][]byte) map[pos]int {
	path := make(map[pos]int)

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '~' {
				path[pos{i, j}] = 1
			}
		}
	}

	return path
}

func flood(board [][]byte, cache map[ci]int, best int, dir byte, score int, y int, x int) (int, map[pos]int) {
	cached, in := cache[ci{y, x, dir}]
	if (in && score < cached) || !in {
		cache[ci{y, x, dir}] = score
	}

	if (score > cache[ci{y, x, dir}]) || score > best {
		return best + 1, make(map[pos]int)
	}

	if board[y][x] == 'E' {
		board[y][x] = '~'

		return score, findPath(board)
	}

	board[y][x] = '~'

	paths := make(map[pos]int)
	p := make(map[pos]int)

	if candidate(board, y-1, x) {
		n := score + 1
		if dir != 'N' {
			if dir == 'E' || dir == 'W' {
				n += 1000
			} else {
				n += 2000
			}
		}

		nb := copyBoard(board)
		n, p = flood(nb, cache, best, 'N', n, y-1, x)

		if n < 9999999999 && n == best {
			maps.Copy(paths, p)
		} else if n < best {
			paths = p
		}

		best = min(best, n)
	}
	if candidate(board, y+1, x) {
		s := score + 1
		if dir != 'S' {
			if dir == 'E' || dir == 'W' {
				s += 1000
			} else {
				s += 2000
			}
		}

		sb := copyBoard(board)
		s, p = flood(sb, cache, best, 'S', s, y+1, x)

		if s < 9999999999 && s == best {
			maps.Copy(paths, p)
		} else if s < best {
			paths = p
		}

		best = min(best, s)
	}
	if candidate(board, y, x-1) {
		w := score + 1
		if dir != 'W' {
			if dir == 'N' || dir == 'S' {
				w += 1000
			} else {
				w += 2000
			}
		}

		wb := copyBoard(board)
		w, p = flood(wb, cache, best, 'W', w, y, x-1)

		if w < 9999999999 && w == best {
			maps.Copy(paths, p)
		} else if w < best {
			paths = p
		}

		best = min(best, w)
	}
	if candidate(board, y, x+1) {
		e := score + 1
		if dir != 'E' {
			if dir == 'N' || dir == 'S' {
				e += 1000
			} else {
				e += 2000
			}
		}

		eb := copyBoard(board)
		e, p = flood(eb, cache, best, 'E', e, y, x+1)

		if e < 9999999999 && e == best {
			maps.Copy(paths, p)
		} else if e < best {
			paths = p
		}

		best = min(best, e)
	}

	return best, paths
}

func part1(input string) string {
	board := getBoard(input)
	startY, startX := find(board, 'S')
	score, _ := flood(board, make(map[ci]int), 9999999999, 'E', 0, startY, startX)
	return strconv.Itoa(score)
}

func part2(input string) string {
	board := getBoard(input)
	startY, startX := find(board, 'S')
	_, paths := flood(board, make(map[ci]int), 9999999999, 'E', 0, startY, startX)
	return strconv.Itoa(len(paths))
}
