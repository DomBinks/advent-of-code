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
	data, err := os.ReadFile("day01.txt")
	check(err)
	input := string(data)
	part := getPart()

	if part == "1" {
		fmt.Println(part1(input))
	} else if part == "2" {
		fmt.Println(part2(input))
	}
}

func parseLeftRightLength(input string) ([]string, []string, int) {
	rows := strings.Split(input, "\n")
	length := len(rows)

	left := make([]string, length)
	right := make([]string, length)

	for i := range length {
		cols := strings.Split(rows[i], "   ")

		left[i] = cols[0]
		right[i] = cols[1]
	}

	return left, right, length
}

func part1(input string) string {
	left, right, length := parseLeftRightLength(input)

	slices.Sort(left)
	slices.Sort(right)

	sum := 0

	for i := range length {
		leftI, err := strconv.Atoi(left[i])
		check(err)

		rightI, err := strconv.Atoi(right[i])
		check(err)

		diff := leftI - rightI

		if diff < 0 {
			diff = -diff
		}

		sum += diff
	}

	out := strconv.Itoa(sum)
	return out
}

func part2(input string) string {
	left, right, length := parseLeftRightLength(input)

	freq := make(map[string]int)

	for i := range length {
		_, inFreq := freq[left[i]]

		if !inFreq {
			freq[left[i]] = 0
		}
	}

	for i := range length {
		freq[right[i]] += 1
	}

	total := 0

	for i := range length {
		key, err := strconv.Atoi(left[i])
		check(err)
		val := freq[left[i]]

		total += key * val
	}

	out := strconv.Itoa(total)
	return out
}
