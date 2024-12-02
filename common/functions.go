package common

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return lines
}

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
