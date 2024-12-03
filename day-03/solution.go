package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatchIndex(input, -1)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	sum1 := 0
	sum2 := 0

	isEnabled := true

	for _, match := range matches {
		doMatch := reDo.FindAllStringIndex(input[:match[0]], -1)
		dontMatch := reDont.FindAllStringIndex(input[:match[0]], -1)

		if doMatch != nil && (dontMatch == nil || doMatch[len(doMatch)-1][0] > dontMatch[len(dontMatch)-1][0]) {
			isEnabled = true
		} else if dontMatch != nil {
			isEnabled = false
		}

		num1, _ := strconv.Atoi(input[match[2]:match[3]])
		num2, _ := strconv.Atoi(input[match[4]:match[5]])
		sum1 += num1 * num2

		if isEnabled {
			sum2 += num1 * num2
		}
	}
	fmt.Printf("Sum1: %d\n", sum1)
	fmt.Printf("Sum2: %d\n", sum2)
}
