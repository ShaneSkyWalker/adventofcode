package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// Global variable for the input file path relative to the day directory
const inputFile = "input.txt"

func main() {
	// Read input and handle potential errors
	lines, err := readInput(inputFile)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	fmt.Println("--- Advent of Code 2025 - Day 04 ---")

	// Execute Part 1
	result1 := solvePart1(lines)
	fmt.Printf("Part 1 Result: %d\n", result1)

	// Execute Part 2
	result2 := solvePart2(lines)
	fmt.Printf("Part 2 Result: %d\n", result2)
}

// readInput reads a file line-by-line and returns a slice of strings.
// It is designed to be reusable for all days.
func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error during file scan: %w", err)
	}

	return lines, nil
}

// solvePart1 contains the logic for the first part of the puzzle.
func solvePart1(lines []string) int {
	total := 0
	operatorLine := lines[len(lines) - 1]
	operators := strings.Fields(operatorLine)
	size := len(operators)
	fig := []int{}

	for i := 0; i < len(lines) - 1; i++ {
		var digits []string = strings.Fields(lines[i])
		for j := 0; j < len(digits); j++ {
			res, _ := strconv.Atoi(digits[j])
			fig = append(fig, res)
		}
	}
	for i := 0; i < size; i++ {
		if operators[i] == "+" {
			total += fig[i] + fig[i + size] + fig[i + 2 * size] + fig[i + 3 * size]

		} else if operators[i] == "*" {
			total += fig[i] * fig[i + size] * fig[i + 2 * size] * fig[i + 3 * size]
		}
	}
	return total
}

// solvePart2 contains the logic for the second part of the puzzle.
// It often builds upon or modifies the logic from Part 1.
func solvePart2(lines []string) int {
	total := 0
	operatorLine := lines[len(lines) - 1]
	colFigLen := []int{}
	return total
}


