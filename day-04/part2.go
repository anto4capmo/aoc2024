package main

import (
	_ "embed"
	"fmt"
	"strings"
)

func part2(input string) {
	lines := strings.Split(input, "\n")
	matrixSize := len(lines)
	matrix := make([][]rune, matrixSize)
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	directions := [][]int{
		{-1, -1}, // Up-Left
		{-1, 1},  // Up-Right
		{1, -1},  // Down-Left
		{1, 1},   // Down-Right
	}

	wordCount := 0

	checkX := func(x, y int) bool {
		matrixSize := len(matrix)

		letters := make([]rune, 4)

		for i, dir := range directions {
			nextX := x + dir[0]
			nextY := y + dir[1]
			if nextX < 0 || nextX >= matrixSize || nextY < 0 || nextY >= len(matrix[nextX]) {
				return false
			}
			letters[i] = matrix[nextX][nextY]
		}

		edges := [][]rune{
			{letters[0], letters[3]}, // Up-Left and Down-Right
			{letters[1], letters[2]}, // Up-Right and Down-Left
		}

		for _, pair := range edges {
			if (pair[0] == 'M' && pair[1] == 'S') || (pair[0] == 'S' && pair[1] == 'M') {
				continue
			} else {
				return false
			}
		}

		return true
	}

	for x := 0; x < matrixSize; x++ {
		for y := 0; y < len(matrix[x]); y++ {
			if matrix[x][y] == 'A' {
				found := checkX(x, y)
				if found {
					fmt.Printf("Found MAS in X pattern at center (%d, %d)\n", x, y)
					wordCount++
				}
			}
		}
	}

	fmt.Printf("Total number of MAS found in X pattern: %d\n", wordCount)
}
