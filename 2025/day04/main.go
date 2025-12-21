package main

import (
	"bufio"
	"fmt"
	"os"
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
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			// @@@		(i-1)(j-1) (i-1)( j ) (i-1)(j+1)
			// @@@		( i )(j-1) ( i )( j ) ( i )(j+1)
			// @@@		(i+1)(j-1) (i+1)( j ) (i+1)(j+1)
			if lines[i][j] == '@' {
				localSum := 0
				if i - 1 >= 0 && j - 1 >= 0 && lines[i-1][j-1] == '@' {
					localSum += 1
				}
				if i - 1 >= 0 && lines[i-1][j] == '@' {
					localSum += 1
				}
				if i - 1 >= 0 && j + 1 < len(lines[i]) && lines[i-1][j+1] == '@' {
					localSum += 1
				}
				if j - 1 >= 0 && lines[i][j-1] == '@' {
					localSum += 1
				}
				if j + 1 < len(lines[i]) && lines[i][j+1] == '@' {
					localSum += 1
				}
				if i + 1 < len(lines) && j - 1 >= 0 && lines[i+1][j-1] == '@' {
					localSum += 1
				}
				if i + 1 < len(lines) && lines[i+1][j] == '@' {
					localSum += 1
				}
				if i + 1 < len(lines) && j + 1 < len(lines[i+1]) && lines[i+1][j+1] == '@' {
					localSum += 1
				} 
				if localSum < 4 {
					total += 1	
					fmt.Printf("location: i: %d, j: %d\n", i, j)
				}
			}
		}
	}
	return total
}

// solvePart2 contains the logic for the second part of the puzzle.
// It often builds upon or modifies the logic from Part 1.
func solvePart2(lines []string) int {
	grid := make([][]rune, len(lines))
    for i, line := range lines {
        grid[i] = []rune(line)
    }
	total := 0
	for {

		innerSum := 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				// @@@		(i-1)(j-1) (i-1)( j ) (i-1)(j+1)
				// @@@		( i )(j-1) ( i )( j ) ( i )(j+1)
				// @@@		(i+1)(j-1) (i+1)( j ) (i+1)(j+1)
				if grid[i][j] == '@' {
					localSum := 0
					if i - 1 >= 0 && j - 1 >= 0 && grid[i-1][j-1] == '@' {
						localSum += 1
					}
					if i - 1 >= 0 && grid[i-1][j] == '@' {
						localSum += 1
					}
					if i - 1 >= 0 && j + 1 < len(grid[i]) && grid[i-1][j+1] == '@' {
						localSum += 1
					}
					if j - 1 >= 0 && grid[i][j-1] == '@' {
						localSum += 1
					}
					if j + 1 < len(grid[i]) && grid[i][j+1] == '@' {
						localSum += 1
					}
					if i + 1 < len(grid) && j - 1 >= 0 && grid[i+1][j-1] == '@' {
						localSum += 1
					}
					if i + 1 < len(grid) && grid[i+1][j] == '@' {
						localSum += 1
					}
					if i + 1 < len(grid) && j + 1 < len(grid[i+1]) && grid[i+1][j+1] == '@' {
						localSum += 1
					} 
					if localSum < 4 {
						innerSum += 1	
						// fmt.Printf("location: i: %d, j: %d\n", i, j)
						grid[i][j] = '.'
					}
				}
			}
		}
		if innerSum == 0 {
			break
		}
		total += innerSum
	}
	return total
}


