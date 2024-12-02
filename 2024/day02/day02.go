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
	data, err := os.ReadFile("day02.txt")
	check(err)
	input := string(data)
	part := getPart()

	if part == "1" {
		fmt.Println(part1(input))
	} else if part == "2" {
		fmt.Println(part2(input))
	}
}

func isIncreasing(a int, b int) bool {
	return a < b
}

func diff(a int, b int, increasing bool) bool {
	if increasing {
		return b-a >= 1 && b-a <= 3
	} else {
		return a-b >= 1 && a-b <= 3
	}
}

func isSafe(report string) bool {
	levels := strings.Split(report, " ")

	prev, err := strconv.Atoi(levels[0])
	check(err)
	next, err := strconv.Atoi(levels[1])
	check(err)
	increasing := isIncreasing(prev, next)

	for j := 1; j < len(levels); j++ {
		current, err := strconv.Atoi(levels[j])
		check(err)

		if !diff(prev, current, increasing) || (increasing != isIncreasing(prev, current)) {
			return false
		} else {
			prev = current
		}
	}

	return true
}

func part1(input string) string {
	reports := strings.Split(input, "\n")
	safe := 0

	for i := range len(reports) {
		if isSafe(reports[i]) {
			safe += 1
		}
	}

	return strconv.Itoa(safe)
}

func part2(input string) string {
	reports := strings.Split(input, "\n")
	safe := 0

	for i := range len(reports) {
		report := strings.Split(reports[i], " ")

		for j := range len(report) {
			start := report[0:j]
			end := report[j+1:]
			subReport := strings.Join(slices.Concat(start, end), " ")

			if isSafe(subReport) {
				safe += 1
				break
			}
		}
	}

	return strconv.Itoa(safe)
}
