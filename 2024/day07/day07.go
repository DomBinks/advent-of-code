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
	data, err := os.ReadFile("day07.txt")
	check(err)
	input := string(data)
	part := getPart()

	if part == "1" {
		fmt.Println(part1(input))
	} else if part == "2" {
		fmt.Println(part2(input))
	}
}

func applyOperator(target int, current int, operator string, operands []int, concat bool) bool {
	if current > target {
		return false
	}
	if len(operands) == 0 {
		return target == current
	}

	if operator == "+" {
		current += operands[0]
	}
	if operator == "*" {
		current *= operands[0]
	}
	if concat && operator == "||" {
		var err error
		current, err = strconv.Atoi(strconv.Itoa(current) + strconv.Itoa(operands[0]))
		check(err)
	}

	applyConcat := false
	if concat {
		applyConcat = applyOperator(target, current, "||", operands[1:], concat)
	}

	return applyOperator(target, current, "+", operands[1:], concat) ||
		applyOperator(target, current, "*", operands[1:], concat) ||
		applyConcat
}

func solveEquations(input string, concat bool) string {
	equations := strings.Split(input, "\n")
	total := 0

	for i := range equations {
		parts := strings.Split(equations[i], ":")
		target, err := strconv.Atoi(parts[0])
		check(err)
		operandsStrs := strings.Split(parts[1], " ")[1:]
		operands := make([]int, len(operandsStrs))
		for j := range operandsStrs {
			operands[j], err = strconv.Atoi(operandsStrs[j])
			check(err)
		}

		applyConcat := false
		if concat {
			applyConcat = applyOperator(target, operands[0], "||", operands[1:], concat)
		}

		if applyOperator(target, operands[0], "+", operands[1:], concat) ||
			applyOperator(target, operands[0], "*", operands[1:], concat) ||
			applyConcat {

			total += target
		}
	}

	out := strconv.Itoa(total)
	return out
}

func part1(input string) string {
	return solveEquations(input, false)
}

func part2(input string) string {
	return solveEquations(input, true)
}
