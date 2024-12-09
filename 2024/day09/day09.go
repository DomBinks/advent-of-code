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
	data, err := os.ReadFile("day09.txt")
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
	diskMapStrs := strings.Split(input, "\n")[0]
	diskMap := make([]int, len(diskMapStrs))
	size := 0
	for i := 0; i < len(diskMapStrs); i++ {
		elem, err := strconv.Atoi(string(diskMapStrs[i]))
		check(err)
		diskMap[i] = elem
		size += elem
	}

	system := make([]int, size)
	index := 0
	id := 0
	for i := 0; i < len(diskMap); i++ {
		if i%2 == 0 {
			for _ = range diskMap[i] {
				system[index] = id
				index += 1
			}
			id += 1
		} else {
			for _ = range diskMap[i] {
				system[index] = -1
				index += 1
			}
		}
	}

	for i := range len(system) {
		for i < len(system) && system[i] == -1 {
			system[i] = system[len(system)-1]
			system = system[:len(system)-1]
		}
	}

	total := 0
	for i := range system {
		total += i * system[i]
	}

	out := strconv.Itoa(total)
	return out
}

func part2(input string) string {
	diskMapStrs := strings.Split(input, "\n")[0]
	diskMap := make([]int, len(diskMapStrs))
	size := 0
	for i := 0; i < len(diskMapStrs); i++ {
		elem, err := strconv.Atoi(string(diskMapStrs[i]))
		check(err)
		diskMap[i] = elem
		size += elem
	}

	system := make([]int, size)
	gaps := make([][]int, 0)  // length, start
	files := make([][]int, 0) // length, start, id
	index := 0
	id := 0
	for i := 0; i < len(diskMap); i++ {
		if i%2 == 0 {
			files = append(files, []int{diskMap[i], index, id})
			for _ = range diskMap[i] {
				system[index] = id
				index += 1
			}
			id += 1
		} else {
			if diskMap[i] > 0 {
				gaps = append(gaps, []int{diskMap[i], index})
			}
			for _ = range diskMap[i] {
				system[index] = -1
				index += 1
			}
		}
	}

	slices.Reverse(files)

	for i := range files {
		for j := range gaps {
			if files[i][0] <= gaps[j][0] && gaps[j][1] < files[i][1] {
				for k := gaps[j][1]; k < gaps[j][1]+files[i][0]; k++ {
					system[k] = files[i][2]
				}
				for k := files[i][1]; k < files[i][1]+files[i][0]; k++ {
					system[k] = -1
				}
				if gaps[j][0] == files[i][0] {
					gaps = slices.Concat(gaps[:j], gaps[j+1:])
				} else {
					gaps[j][0] -= files[i][0]
					gaps[j][1] += files[i][0]
				}
				break
			}
		}
	}

	total := 0
	for i := range system {
		if system[i] != -1 {
			total += i * system[i]
		}
	}

	out := strconv.Itoa(total)
	return out
}
