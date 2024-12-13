package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
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
	data, err := os.ReadFile("example.txt")
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
	x, y int
}

func part1(input string) string {
	total := 0

	lines := strings.Split(input, "\n")
	l := 0
	for l < len(lines) {
		a := strings.Split(lines[l], " ")
		b := strings.Split(lines[l+1], " ")
		p := strings.Split(lines[l+2], " ")

		xaS := strings.Split(a[2], "+")[1]
		xa, err := strconv.Atoi(xaS[:len(xaS)-1])
		check(err)
		ya, err := strconv.Atoi(strings.Split(a[3], "+")[1])
		check(err)
		xbS := strings.Split(b[2], "+")[1]
		xb, err := strconv.Atoi(xbS[:len(xbS)-1])
		check(err)
		yb, err := strconv.Atoi(strings.Split(b[3], "+")[1])
		check(err)
		xpS := strings.Split(p[1], "=")[1]
		xp, err := strconv.Atoi(xpS[:len(xpS)-1])
		check(err)
		yp, err := strconv.Atoi(strings.Split(p[2], "=")[1])
		check(err)

		table := make(map[pos]int)
		table[pos{xb, yb}] = 1

		xaC := xa
		yaC := ya
		pC := 3
		for xaC <= xp && yaC <= yp {
			pts, in := table[pos{xaC, yaC}]
			if !in {
				pts = 99999999
			}
			table[pos{xaC, yaC}] = min(pC, pts)

			pC += 3
			xaC += xa
			yaC += ya
		}

		xbC := xb
		ybC := yb
		pC = 1
		for xbC <= xp && ybC <= yp {
			pts, in := table[pos{xbC, ybC}]
			if !in {
				pts = 99999999
			}
			table[pos{xbC, ybC}] = min(pC, pts)

			pC += 1
			xbC += xb
			ybC += yb
		}

		xyS := make([]pos, 0)
		for k := range table {
			xyS = append(xyS, k)
		}
		sort.Slice(xyS, func(i, j int) bool {
			if xyS[i].x != xyS[j].x {
				return xyS[i].x < xyS[j].x
			} else {
				return xyS[i].y < xyS[j].y
			}
		})

		for i := range xyS {
			for j := range xyS {
				xy := xyS[i]
				xyO := xyS[j]
				xyP := table[xy]
				xyOP := table[xyO]

				xN := xy.x + xyO.x
				yN := xy.y + xyO.y

				if xN > xp || yN > yp {
					break
				}

				pts, in := table[pos{xN, yN}]
				if !in {
					pts = 999999999
				}
				table[pos{xN, yN}] = min(xyP+xyOP, pts)
			}
		}

		pts, in := table[pos{xp, yp}]
		if in {
			total += pts
		} else {
		}

		l += 4
	}

	out := strconv.Itoa(total)
	return out
}

func part2(input string) string {
	return input
}
