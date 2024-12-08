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
	data, err := os.ReadFile("day08.txt")
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

func search(board [][]byte, freq byte, startY int, startX int) ([]int, []int) {
	ys := make([]int, 0)
	xs := make([]int, 0)

	for y := startY + 1; y < len(board); y++ {
		if board[y][startX] == freq {
			ys = append(ys, y)
			xs = append(xs, startX)
		}
	}

	for x := startX + 1; x < len(board[0]); x++ {
		for y := 0; y < len(board); y++ {
			if board[y][x] == freq {
				ys = append(ys, y)
				xs = append(xs, x)
			}
		}
	}

	return ys, xs
}

func onBoard(board [][]byte, y int, x int) bool {
	return y >= 0 && y < len(board) && x >= 0 && x < len(board[0])
}

func part1(input string) string {
	board := getBoard(input)
	antinodes := make([][]bool, len(board))
	for i := range antinodes {
		antinodes[i] = make([]bool, len(board[i]))
	}
	total := 0

	for j := range board[0] {
		for i := range board {
			if board[i][j] != '.' {
				ys, xs := search(board, board[i][j], i, j)
				for k := range ys {
					y := ys[k]
					x := xs[k]
					dY := max(y, i) - min(y, i)
					dX := max(x, j) - min(x, j)

					var rY, lY int
					if y < i {
						lY = i + dY
						rY = y - dY
					} else {
						lY = i - dY
						rY = y + dY
					}
					lX := j - dX
					rX := x + dX

					if onBoard(board, lY, lX) {
						if !antinodes[lY][lX] {
							total += 1
							antinodes[lY][lX] = true
						}
					}
					if onBoard(board, rY, rX) {
						if !antinodes[rY][rX] {
							total += 1
							antinodes[rY][rX] = true
						}
					}
				}
			}
		}
	}

	out := strconv.Itoa(total)
	return out
}

func part2(input string) string {
	board := getBoard(input)
	antinodes := make([][]bool, len(board))
	for i := range antinodes {
		antinodes[i] = make([]bool, len(board[i]))
	}
	total := 0

	for j := range board[0] {
		for i := range board {
			if board[i][j] != '.' {
				if !antinodes[i][j] {
					total += 1
					antinodes[i][j] = true
				}
				ys, xs := search(board, board[i][j], i, j)
				for k := 0; k < len(ys); k++ {
					y := ys[k]
					x := xs[k]
					dY := max(y, i) - min(y, i)
					dX := max(x, j) - min(x, j)

					var rY, lY int
					if y < i {
						lY = i + dY
						rY = y - dY
					} else {
						lY = i - dY
						rY = y + dY
					}
					lX := j - dX
					rX := x + dX

					for onBoard(board, lY, lX) || onBoard(board, rY, rX) {
						if onBoard(board, lY, lX) {
							if !antinodes[lY][lX] {
								total += 1
								antinodes[lY][lX] = true
							}
						}
						if onBoard(board, rY, rX) {
							if !antinodes[rY][rX] {
								total += 1
								antinodes[rY][rX] = true
							}
						}

						if y < i {
							lY = lY + dY
							rY = rY - dY
						} else {
							lY = lY - dY
							rY = rY + dY
						}
						lX = lX - dX
						rX = rX + dX
					}
				}
			}
		}
	}

	out := strconv.Itoa(total)
	return out
}
