package common

import (
	"regexp"
	"strconv"
)

func ExtractNumbersFromLine(line string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(line, -1)

	if len(matches) == 0 {
		return []int{}
	}

	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil
		}
		numbers = append(numbers, num)
	}

	return numbers
}
