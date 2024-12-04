package main

import (
	_ "embed"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}