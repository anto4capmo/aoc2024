package main

import (
	"aoc2024/common"
	"fmt"
	"math"
	"path/filepath"
	"runtime"
)

func isSafe(levels ...int) bool {
	if len(levels) < 2 {
		return true
	}

	badLevels := getBadLevels(levels)

	return badLevels == 0
}

func isDampenerSafe(levels ...int) bool {
	if len(levels) < 2 {
		return true
	}

	if isSafe(levels...) {
		return true
	}

	for i := range levels {
		newLevels := append([]int{}, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)

		if isSafe(newLevels...) {
			return true
		}
	}

	return false
}

func getBadLevels(levels []int) int {
	badLevels := 0
	prevDir := 0 // Holds the direction of change: 1 (up), -1 (down), 0 (none)

	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i] - levels[i+1]

		dir := 0
		if diff > 0 {
			dir = 1
		} else if diff < 0 {
			dir = -1
		}

		if prevDir != 0 && dir != 0 && prevDir != dir {
			badLevels++
			continue
		}

		prevDir = dir

		absDiff := math.Abs(float64(diff))
		if absDiff > 3 || absDiff == 0 {
			badLevels++
		}
	}
	return badLevels
}

func main() {
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Error: Unable to determine current file location")
		return
	}
	dirPath := filepath.Dir(filePath)
	inputPath := filepath.Join(dirPath, "input.txt")

	lines := common.ReadFile(inputPath)
	if lines == nil {
		fmt.Println("Error: Unable to read input file")
		return
	}

	safeReports := 0
	dampenerSafe := 0

	for _, line := range lines {
		numbers := common.ExtractNumbersFromLine(line)
		if len(numbers) == 0 {
			fmt.Printf("Skipping line with no valid numbers: %s\n", line)
			continue
		}

		if isSafe(numbers...) {
			safeReports++
		}
		if isDampenerSafe(numbers...) {
			dampenerSafe++
		}
	}

	fmt.Printf("Safe reports: %d\n", safeReports)
	fmt.Printf("Dampener Safe reports: %d\n", dampenerSafe)
}
