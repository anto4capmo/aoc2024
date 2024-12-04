package main

import (
	_ "embed"
	"fmt"
	"strings"
)

func part1(input string) {
	word := "XMAS"
	wordVector := []rune(word)

	lines := strings.Split(input, "\n")
	matrixSize := len(lines)
	matrix := make([][]rune, matrixSize)
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	directions := [][]int{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // Up-Left
		{-1, 1},  // Up-Right
		{1, -1},  // Down-Left
		{1, 1},   // Down-Right
	}

	wordCount := 0

	checkWord := func(x, y int, direction []int) bool {
		for k := 0; k < len(wordVector); k++ {
			newX := x + direction[0]*k
			newY := y + direction[1]*k
			if newX < 0 || newX >= matrixSize || newY < 0 || newY >= len(matrix[newX]) {
				return false
			}
			if matrix[newX][newY] != wordVector[k] {
				return false
			}
		}
		return true
	}

	for x := 0; x < matrixSize; x++ {
		for y := 0; y < len(matrix[x]); y++ {
			if matrix[x][y] == wordVector[0] {
				for _, direction := range directions {
					if checkWord(x, y, direction) {
						fmt.Printf("Found word '%s' starting at (%d, %d) in direction (%d, %d)\n",
							word, x, y, direction[0], direction[1])
						wordCount++
					}
				}
			}
		}
	}

	fmt.Printf("Word '%s' found %d times\n", word, wordCount)
}
