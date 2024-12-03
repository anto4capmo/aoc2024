package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func extractNumbers(line string) ([]int, error) {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(line, -1)

	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

func calculateTotalDistance(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		totalDistance += int(math.Abs(float64(left[i] - right[i])))
	}

	return totalDistance
}

func calculateSimilarityScore(left, right []int) int {
	frequencyMap := make(map[int]int)
	for _, num := range right {
		frequencyMap[num]++
	}

	similarityScore := 0
	for _, num := range left {
		similarityScore += num * frequencyMap[num]
	}

	return similarityScore
}

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	var left, right []int

	for _, line := range lines {
		numbers, _ := extractNumbers(line)
		if len(numbers) != 2 {
			continue
		}

		left = append(left, numbers[0])
		right = append(right, numbers[1])
	}

	totalDistance := calculateTotalDistance(left, right)
	similarityScore := calculateSimilarityScore(left, right)

	fmt.Printf("Total distance: %d\n", totalDistance)
	fmt.Printf("Similarity score: %d\n", similarityScore)
}
