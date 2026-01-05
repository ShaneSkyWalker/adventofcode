package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	// "sort"
	// "math"
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

	fmt.Println("--- Advent of Code 2025 - Day 10 ---")

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
	solution1 := func (lines []string) int {
		// solution1: bfs
		total := 0
		for _, line := range lines {
			eles := strings.Split(line, " ")
			lights := eles[0]
			buts := eles[1: len(eles) - 1]
			// fmt.Printf("%v\n", lights)
			// fmt.Printf("%v\n", buts)
			lightVal := 0
			lenLight := len(lights) - 2
			for _, li := range lights {
				lightVal *= 2
				if li == '#' {
					lightVal += 1
				}
			}
			lightVal /= 2
			// fmt.Println(lightVal)
			
			butVal := []int{}
			for _, but := range buts {
				val := 0
				valCh := strings.Split(but[1: len(but) - 1], ",")
				for _, v := range valCh {
					vInt, _ := strconv.Atoi(v)
					val += 1 << (lenLight - vInt - 1)
				}
				butVal = append(butVal, val)
			}
			// fmt.Printf("butVal: %v\n", butVal)

			// bfs
			d := map[int]int{}
			d[0] = 0
			q := []int{0}
			for len(q) > 0 {
				current := q[0]
				q = q[1:]
				// fmt.Printf("q: %v\n", q)

				if current == lightVal {
					total += d[lightVal]
					// fmt.Printf("d[lightVal]: %d\n", d[lightVal])
					break
				}
				
				set := make(map[int]struct{})
				for v, _ := range d {
					set[v] = struct{}{}
				}
				// fmt.Printf("set: %v\n", set)
				for _, b := range butVal {
					if _, ok := set[(current ^ b)]; !ok {
						d[current ^ b] = d[current] + 1
						q = append(q, current ^ b)
					}
				}
			}
		}
		return total
	}
	return solution1(lines)
}

// solvePart2 contains the logic for the second part of the puzzle.
// It often builds upon or modifies the logic from Part 1.
func solvePart2(lines []string) int {
	return 0
}