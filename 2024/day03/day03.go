package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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
	data, err := os.ReadFile("day03.txt")
	check(err)
	input := string(data)
	part := getPart()

	if part == "1" {
		fmt.Println(part1(input))
	} else if part == "2" {
		fmt.Println(part2(input))
	}
}

func parseByte(input string, index int, target byte) (bool, int) {
	if index < len(input) && input[index] == target {
		return true, index + 1
	} else {
		return false, index
	}
}

func parseStr(input string, index int, target string) (bool, int) {
	if index+len(target) <= len(input) && input[index:index+len(target)] == target {
		return true, index + len(target)
	} else {
		return false, index
	}
}

func parseNum(input string, index int) (string, int) {
	nums := map[byte]int{'0': 0, '1': 0, '2': 0, '3': 0, '4': 0, '5': 0, '6': 0, '7': 0, '8': 0, '9': 0}

	start := index
	var end int

	for index < len(input) {
		_, in := nums[input[index]]
		if in {
			index += 1
		} else {
			end = index
			break
		}
	}

	if start < end {
		return input[start:end], end
	} else {
		return "NaN", index
	}
}

func parseMul(input string, index int) (int, int) {
	out := -1
	var parsed bool

	parsed, index = parseStr(input, index, "mul(")
	if parsed {
		var a string
		a, index = parseNum(input, index)

		if a != "NaN" && index < len(input)-2 {
			parsed, index = parseByte(input, index, ',')
			if parsed {
				var b string
				b, index = parseNum(input, index)

				if b != "NaN" && index < len(input) {
					parsed, index = parseByte(input, index, ')')
					if parsed {
						aI, err := strconv.Atoi(a)
						check(err)
						bI, err := strconv.Atoi(b)
						check(err)

						out = aI * bI
					}
				}
			}
		}
	}

	return out, index
}

func part1(input string) string {
	total := 0
	index := 0

	for index < len(input)-7 {
		var mul int
		mul, index = parseMul(input, index)

		if mul != -1 {
			total += mul
		} else {
			index += 1
		}
	}

	out := strconv.Itoa(total)
	return out
}

func part2(input string) string {
	total := 0
	index := 0
	enabled := true

	for index < len(input)-7 {
		var parsed bool

		parsed, index = parseStr(input, index, "do()")
		if parsed {
			enabled = true
			continue
		}

		parsed, index = parseStr(input, index, "don't()")
		if parsed {
			enabled = false
			continue
		}

		if enabled {
			var mul int
			mul, index = parseMul(input, index)

			if mul != -1 {
				total += mul
				continue
			}
		}

		index += 1
	}

	out := strconv.Itoa(total)
	return out
}
