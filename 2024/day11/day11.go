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
	data, err := os.ReadFile("day11.txt")
	check(err)
	input := string(data)
	part := getPart()

	if part == "1" {
		fmt.Println(part1(input))
	} else if part == "2" {
		fmt.Println(part2(input))
	}
}

func part1(input string) string {
	stonesStrs := strings.Split(input, " ")
	stones := make([]int, len(stonesStrs))
	for i := range stonesStrs {
		elem, err := strconv.Atoi(stonesStrs[i])
		check(err)
		stones[i] = elem
	}

	for _ = range 25 {
		i := 0
		for i < len(stones) {
			if stones[i] == 0 {
				stones[i] = 1
				i += 1
				continue
			}

			stone := strconv.Itoa(stones[i])
			if len(stone)%2 == 0 {
				l, err := strconv.Atoi(stone[:len(stone)/2])
				check(err)
				r, err := strconv.Atoi(stone[len(stone)/2:])
				check(err)

				var end []int
				if i+1 < len(stones) {
					end = stones[i+1:]
				}

				stones = slices.Concat(stones[:i], []int{l}, []int{r}, end)
				i += 2
				continue
			}

			stones[i] = stones[i] * 2024
			i += 1
		}
	}

	out := strconv.Itoa(len(stones))
	return out
}

func count(stone int, iters int, cache map[string]int) int {
	key := strconv.Itoa(stone) + "-" + strconv.Itoa(iters)
	val, in := cache[key]
	if in {
		return val
	}

	if iters == 0 {
		return 1
	}

	if stone == 0 {
		cache[key] = count(1, iters-1, cache)
		return cache[key]
	}

	stoneStr := strconv.Itoa(stone)
	if len(stoneStr)%2 == 0 {
		l, err := strconv.Atoi(stoneStr[:len(stoneStr)/2])
		check(err)
		r, err := strconv.Atoi(stoneStr[len(stoneStr)/2:])
		check(err)

		cache[key] = count(l, iters-1, cache) + count(r, iters-1, cache)
		return cache[key]
	}

	cache[key] = count(stone*2024, iters-1, cache)
	return cache[key]
}

func part2(input string) string {
	stonesStrs := strings.Split(input, " ")
	stones := make([]int, len(stonesStrs))
	for i := range stonesStrs {
		elem, err := strconv.Atoi(stonesStrs[i])
		check(err)
		stones[i] = elem
	}

	total := 0
	cache := make(map[string]int)

	for i := range stones {
		total += count(stones[i], 75, cache)
	}

	out := strconv.Itoa(total)
	return out
}
