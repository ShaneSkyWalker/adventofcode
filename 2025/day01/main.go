package main

import (
	"bufio"
	"fmt"
	"os"
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

	fmt.Println("--- Advent of Code 2025 - Day 01 ---")

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
	init := 50
	total := 0
	for _, line := range lines {
		if len(line) == 0 {
			return total
		}
		charSlice := []rune(line)
		if prefix := charSlice[0]; prefix == 'L' {
			intVal, _ := strconv.Atoi(string(charSlice[1:]))
			init -= intVal
		} else if prefix == 'R' {
			intVal, _ := strconv.Atoi(string(charSlice[1:]))
			init += intVal
		} else {
			return total
		}
		for {
			if init < 0 {
				init += 100
			} else if init > 99 {
				init -= 100
			} else {
				if init == 0 {
					total += 1
				}
				break
			}
		}
	}
	return total 
}

// solvePart2 contains the logic for the second part of the puzzle.
// It often builds upon or modifies the logic from Part 1.
func solvePart2(lines []string) int {
	init := 50
	total := 0
	// one condition should not be considered to plus one,
	// currently dial is at 0 and turned left.
	temp := 0
	for _, line := range lines {
		if len(line) == 0 {
			return total
		}
		charSlice := []rune(line)
		if prefix := charSlice[0]; prefix == 'L' {
			intVal, _ := strconv.Atoi(string(charSlice[1:]))
			if init == 0 {
				temp = -1
			}
			init -= intVal
		} else if prefix == 'R' {
			intVal, _ := strconv.Atoi(string(charSlice[1:]))
			init += intVal
		} else {
			return total
		}
		quotient := init / 100
		remainder := init % 100

		if init < 0 {
			init = (100 + remainder) % 100
			total += 1 - quotient
			if temp != 0 {
				total += temp
				temp = 0
			}
		} else if init > 99 {
			init = remainder
			total += quotient
		} else if init == 0 {
			total += 1
		}
	}
	return  total
}
