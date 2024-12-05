package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	rulesRe := regexp.MustCompile(`\d+\|\d+`)
	rulesMatches := rulesRe.FindAllString(input, -1)

	rulesMap := make(map[string]struct{})
	for _, rule := range rulesMatches {
		rulesMap[rule] = struct{}{}
	}

	updatesRe := regexp.MustCompile(`\d+,`)
	updatesMatches := updatesRe.FindAllStringIndex(input, -1)
	firstUpdate := updatesMatches[0]

	updatesSection := input[firstUpdate[0]:]
	updates := strings.Split(strings.TrimSpace(updatesSection), "\n")

	updateNumbers := make([][]int, len(updates))
	for i, update := range updates {
		updateNumberStrings := strings.Split(strings.TrimSpace(update), ",")
		for _, updateNumberString := range updateNumberStrings {
			val, err := strconv.Atoi(strings.TrimSpace(updateNumberString))
			if err == nil {
				updateNumbers[i] = append(updateNumbers[i], val)
			}
		}
	}

	correctSum := 0
	incorrectSum := 0

	sortUpdateWithRules := func(update []int) {
		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update); j++ {
				pi := update[i]
				pj := update[j]

				_, piBeforePj := rulesMap[fmt.Sprintf("%d|%d", pi, pj)]
				_, pjBeforePi := rulesMap[fmt.Sprintf("%d|%d", pj, pi)]

				if piBeforePj {
					continue
				} else if pjBeforePi {
					update[i], update[j] = update[j], update[i]
				}
			}
		}
	}

	for _, update := range updateNumbers {
		original := make([]int, len(update))
		copy(original, update)

		sortUpdateWithRules(update)

		if slices.Equal(original, update) {
			midVal := update[len(update)/2]
			correctSum += midVal
		} else {
			midVal := update[len(update)/2]
			incorrectSum += midVal
		}
	}

	fmt.Printf("Part one: %d\n", correctSum)
	fmt.Printf("Part two: %d\n", incorrectSum)
}
