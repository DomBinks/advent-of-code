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
	data, err := os.ReadFile("day05.txt")
	check(err)
	input := string(data)
	part := getPart()

	if part == "1" {
		fmt.Println(part1(input))
	} else if part == "2" {
		fmt.Println(part2(input))
	}
}

func getRules(section string) map[string][]string {
	lines := strings.Split(section, "\n")
	rules := make(map[string][]string)

	for i := range lines {
		rule := strings.Split(lines[i], "|")

		_, in := rules[rule[1]]
		if in {
			rules[rule[1]] = append(rules[rule[1]], rule[0])
		} else {
			rules[rule[1]] = []string{rule[0]}
		}
	}

	return rules
}

func getUpdates(section string) [][]string {
	lines := strings.Split(section, "\n")
	updates := make([][]string, len(lines))

	for i := range lines {
		updates[i] = strings.Split(lines[i], ",")
	}

	return updates
}

func part1(input string) string {
	sections := strings.Split(input, "X")
	rules := getRules(sections[0])
	updates := getUpdates(sections[1])
	total := 0

	for i := range updates {
		update := updates[i]
		seen := make(map[string]bool)
		doUpdate := true

		for j := range update {
			reqs, in := rules[update[j]]
			if in {
				for k := range reqs {
					_, met := seen[reqs[k]]
					if slices.Contains(update, reqs[k]) && !met {
						doUpdate = false
						break
					}
				}
			}

			if !doUpdate {
				break
			} else {
				seen[update[j]] = true
			}
		}

		if doUpdate {
			mid, err := strconv.Atoi(update[len(update)/2])
			check(err)
			total += mid
		}
	}

	out := strconv.Itoa(total)
	return out
}

func part2(input string) string {
	sections := strings.Split(input, "X")
	rules := getRules(sections[0])
	updates := getUpdates(sections[1])
	total := 0

	for i := range updates {
		update := updates[i]
		seen := make(map[string]bool)
		fixed := false

		j := 0
		for j < len(update) {
			toPass := j

			reqs, in := rules[update[j]]
			if in {
				for k := range reqs {
					_, met := seen[reqs[k]]
					if slices.Contains(update, reqs[k]) && !met {
						toPass = max(toPass, slices.Index(update, reqs[k]))
					}
				}
			}

			if toPass == j {
				seen[update[j]] = true
				j += 1
			} else {
				fixed = true

				var end []string
				if toPass == len(update)-1 {
					end = []string{}
				} else {
					end = update[toPass+1:]
				}

				update = slices.Concat(update[:j], update[j+1:toPass+1], []string{update[j]}, end)
			}
		}

		if fixed {
			mid, err := strconv.Atoi(update[len(update)/2])
			check(err)
			total += mid
		}
	}

	out := strconv.Itoa(total)
	return out
}
