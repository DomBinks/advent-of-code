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
	data, err := os.ReadFile("day14.txt")
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

func part1(input string) string {
	w := 101
	h := 103

	board := make([][][]int, h)
	for i := range board {
		board[i] = make([][]int, w)
	}

	lines := strings.Split(input, "\n")
	for i := range lines {
		p := strings.Split(lines[i], " ")[0]
		v := strings.Split(lines[i], " ")[1]

		pXY := strings.Split(p, "=")[1]
		vXY := strings.Split(v, "=")[1]

		pX, err := strconv.Atoi(strings.Split(pXY, ",")[0])
		check(err)
		pY, err := strconv.Atoi(strings.Split(pXY, ",")[1])
		check(err)
		vX, err := strconv.Atoi(strings.Split(vXY, ",")[0])
		check(err)
		vY, err := strconv.Atoi(strings.Split(vXY, ",")[1])
		check(err)

		eY := (pY + (100 * vY)) % h
		eX := (pX + (100 * vX)) % w
		for eY < 0 {
			eY += h
		}
		for eX < 0 {
			eX += w
		}

		board[eY][eX] = append(board[eY][eX], i)
	}

	total := 1
	quad := []pos{{0, 0}, {0, (w / 2) + 1}, {(h / 2) + 1, 0}, {(h / 2) + 1, (w / 2) + 1}}

	for q := range quad {
		iO := quad[q].y
		jO := quad[q].x

		local := 0
		for i := range h / 2 {
			for j := range w / 2 {
				local += len(board[i+iO][j+jO])
			}
		}

		total *= local
	}

	out := strconv.Itoa(total)
	return out
}

func part2(input string) string {
	w := 101
	h := 103

	board := make([][]int, h)
	for i := range board {
		board[i] = make([]int, w)
	}

	location := make(map[int]pos)
	velocity := make(map[int]pos)

	lines := strings.Split(input, "\n")
	for i := range lines {
		p := strings.Split(lines[i], " ")[0]
		v := strings.Split(lines[i], " ")[1]

		pXY := strings.Split(p, "=")[1]
		vXY := strings.Split(v, "=")[1]

		pX, err := strconv.Atoi(strings.Split(pXY, ",")[0])
		check(err)
		pY, err := strconv.Atoi(strings.Split(pXY, ",")[1])
		check(err)
		vX, err := strconv.Atoi(strings.Split(vXY, ",")[0])
		check(err)
		vY, err := strconv.Atoi(strings.Split(vXY, ",")[1])
		check(err)

		location[i] = pos{pY, pX}
		velocity[i] = pos{vY, vX}
		board[pY][pX] += 1
	}

	for it := range 10000 {
		for i, p := range location {
			nY := (p.y + velocity[i].y) % h
			nX := (p.x + velocity[i].x) % w

			for nY < 0 {
				nY += h
			}
			for nX < 0 {
				nX += w
			}

			board[p.y][p.x] -= 1
			board[nY][nX] += 1
			location[i] = pos{nY, nX}
		}

		for i := range board {
			brk := false
			for j := range len(board[i]) - 5 {
				if board[i][j] != 0 &&
					board[i][j] == board[i][j+1] &&
					board[i][j+1] == board[i][j+2] &&
					board[i][j+2] == board[i][j+3] &&
					board[i][j+3] == board[i][j+4] {

					fmt.Println(it)
					for m := range board {
						for n := range board[m] {
							if board[m][n] == 0 {
								fmt.Print(".")
							} else {
								fmt.Print(board[m][n])
							}
						}
						fmt.Print("\n")
					}

					brk = true
					break
				}
			}

			if brk {
				break
			}
		}
	}

	return ""
}
