package main

import (
	"advent-of-code/utils"
	"errors"
	"fmt"
	"math"
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
	data, err := os.ReadFile("day17.txt")
	check(err)
	input := string(data)
	part := getPart()

	if part == "1" {
		fmt.Println(part1(input))
	} else if part == "2" {
		fmt.Println(part2(input))
	}
}

func comboOperand(operand int, a int, b int, c int) int {
	switch operand {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	case 7:
		return -1
	}

	panic("invalid operand")
}

type state struct {
	A, B, C, IP int
}

func computer(input string, a int) string {
	sections := utils.GetSections(input)
	registers := strings.Split(sections[0][:len(sections[0])-1], "\n")
	program := sections[1][1:]
	instructions := strings.Split(strings.Split(program, " ")[1], ",")

	var A int
	if a == -1 {
		A = utils.Extract(strconv.Atoi(strings.Split(registers[0], " ")[2]))
	} else {
		A = a
	}

	B := utils.Extract(strconv.Atoi(strings.Split(registers[1], " ")[2]))
	C := utils.Extract(strconv.Atoi(strings.Split(registers[2], " ")[2]))
	IP := 0
	output := make([]int, 0)
	cache := make(map[state]bool)
	//fmt.Println(A, "|", B, "|", C, "|", IP)

	for IP < len(instructions) {
		_, in := cache[state{A, B, C, IP}]
		if in {
			return ""
		} else {
			cache[state{A, B, C, IP}] = true
		}

		instruction := utils.Extract(strconv.Atoi(instructions[IP]))
		operand := utils.Extract(strconv.Atoi(instructions[IP+1]))

		switch instruction {
		case 0: // adv
			numerator := A
			combo := comboOperand(operand, A, B, C)
			if combo == -1 {
				return ""
			}
			denominator := math.Pow(2, float64(combo))
			result := int(math.Floor(float64(numerator) / denominator))
			A = result
		case 1: // bxl
			B = B ^ operand
		case 2: // bst
			combo := comboOperand(operand, A, B, C)
			if combo == -1 {
				return ""
			}
			B = combo % 8
		case 3: // jnz
			if A != 0 {
				IP = operand
				continue
			}
		case 4: // bxc
			B = B ^ C
		case 5: // out
			combo := comboOperand(operand, A, B, C)
			if combo == -1 {
				return ""
			}
			output = append(output, combo%8)
		case 6: // bdv
			numerator := A
			combo := comboOperand(operand, A, B, C)
			if combo == -1 {
				return ""
			}
			denominator := math.Pow(2, float64(combo))
			result := int(math.Floor(float64(numerator) / denominator))
			B = result
		case 7: // cdv
			numerator := A
			combo := comboOperand(operand, A, B, C)
			if combo == -1 {
				return ""
			}
			denominator := math.Pow(2, float64(combo))
			result := int(math.Floor(float64(numerator) / denominator))
			C = result
		}

		IP += 2
	}

	//fmt.Println(A, "|", B, "|", C, "|", IP)
	out := ""
	for i := range output {
		out = out + strconv.Itoa(output[i]) + ","
	}
	return out[:len(out)-1]
}

func part1(input string) string {
	return computer(input, -1)
}

func part2S(input string) string {
	//sections := utils.GetSections(input)
	//registers := strings.Split(sections[0][:len(sections[0])-1], "\n")
	//program := sections[1][1:]
	//instructions := strings.Split(program, " ")[1]
	//A := utils.Extract(strconv.Atoi(strings.Split(registers[0], " ")[2]))
	//B := utils.Extract(strconv.Atoi(strings.Split(registers[1], " ")[2]))
	//C := utils.Extract(strconv.Atoi(strings.Split(registers[2], " ")[2]))

	for A := range 100 {
		B := (A % 8) ^ 1
		C := A >> B
		/*
			fmt.Println("A: ", A%8)
			fmt.Println("B: ", B)
			fmt.Println("C: ", C%8)
		*/

		out := ((B ^ 4) ^ C) % 8
		A = A / 8
		fmt.Println(A, ": ", out)
	}
	return ""
}

func part2(input string) string {
	sections := utils.GetSections(input)
	program := sections[1][1:]
	instructions := strings.Split(program, " ")[1]

	a := 562949953421312
	out := ""
	for out != instructions {
		a += 1
		fmt.Println("Checking: ", a)
		out = computer(input, a)
	}

	return input
}
