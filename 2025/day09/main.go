package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
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

type Dots struct {
	x, y int
}
type Pairs struct {
	a, b int
	dist int
}
// solvePart1 contains the logic for the first part of the puzzle.
func solvePart1(lines []string) int {
	dots := []Dots{}
	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(coords[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(coords[1]))
		dots = append(dots, Dots{x, y})
	}
	pairs := []Pairs{}
	for i := 0; i < len(dots); i++ {
		for j := i + 1; j < len(dots); j++ {
			xDist, yDist := 0, 0
			if dots[i].x > dots[j].x {
				xDist = dots[i].x - dots[j].x + 1
			} else {
				xDist = dots[j].x - dots[i].x + 1
			}
			if dots[i].y > dots[j].y {
				yDist = dots[i].y - dots[j].y + 1
			} else {
				yDist = dots[j].y - dots[i].y + 1
			}
			pair := Pairs{i, j, xDist * yDist}
			pairs = append(pairs, pair)
		}
	}
	// fmt.Println(pairs)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].dist < pairs[j].dist
	})
	return pairs[len(pairs) - 1].dist
}

// solvePart2 contains the logic for the second part of the puzzle.
// It often builds upon or modifies the logic from Part 1.
//
// compression + flood fill?
// ray casting
// green theorem
func solvePart2(lines []string) int {
	dots := []Dots{}
	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(coords[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(coords[1]))
		dots = append(dots, Dots{x, y})
	}
	pairs := []Pairs{}
	for i := 0; i < len(dots); i++ {
		for j := i + 1; j < len(dots); j++ {
			xDist, yDist := 0, 0
			if dots[i].x > dots[j].x {
				xDist = dots[i].x - dots[j].x + 1
			} else {
				xDist = dots[j].x - dots[i].x + 1
			}
			if dots[i].y > dots[j].y {
				yDist = dots[i].y - dots[j].y + 1
			} else {
				yDist = dots[j].y - dots[i].y + 1
			}
			pair := Pairs{i, j, xDist * yDist}
			pairs = append(pairs, pair)
		}
	}
	// fmt.Println(pairs)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].dist < pairs[j].dist
	})
	for i, _ range pairs {
		pairs[i]
	}
	return pairs[len(pairs) - 1].dist
	return 0
}

