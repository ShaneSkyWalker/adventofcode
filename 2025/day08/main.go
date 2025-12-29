package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
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

// solvePart1 contains the logic for the first part of the puzzle.
type Connection struct {
	a, b int
	dist float64
}
func solvePart1(lines []string) int {
	var connections []Connection
	for i := 0; i < len(lines); i++ {
		dims := strings.Split(lines[i], ",")
		val0, _ := strconv.ParseFloat(strings.TrimSpace(dims[0]), 64)
		val1, _ := strconv.ParseFloat(strings.TrimSpace(dims[1]), 64)
		val2, _ := strconv.ParseFloat(strings.TrimSpace(dims[2]), 64)
		for j := i + 1; j < len(lines); j++ {
			dims2 := strings.Split(lines[j], ",")
			val3, _ := strconv.ParseFloat(strings.TrimSpace(dims2[0]), 64)
			val4, _ := strconv.ParseFloat(strings.TrimSpace(dims2[1]), 64)
			val5, _ := strconv.ParseFloat(strings.TrimSpace(dims2[2]), 64)

			square := math.Pow(val0 - val3, 2) + math.Pow(val1 - val4, 2) + math.Pow(val2 - val5, 2)
			connections = append(connections, Connection{i, j, square})
		}
	}

	// sort
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].dist < connections[j].dist
	})

	// Initialize DSU (Union-Find)
	parent := make([]int, len(lines))
	size := make([]int, len(lines))
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	// Helper functions for DSU
	var find func(int) int
	find = func(i int) int {
		if parent[i] == i {
			return i
		}
		parent[i] = find(parent[i])
		return parent[i]
	}
	
	union := func(i, j int) {
        rootI, rootJ := find(i), find(j)
        if rootI != rootJ {
            // Merge smaller into larger
            if size[rootI] < size[rootJ] {
                rootI, rootJ = rootJ, rootI
            }
            parent[rootJ] = rootI
            size[rootI] += size[rootJ]
            size[rootJ] = 0 // Optional: reset size of the merged-away root
        }
    }

	// Process first 1000 shortest connections
    limit := 1000
    if len(connections) < 1000 {
        limit = len(connections)
    }

	for k := 0; k < limit; k++ {
        c := connections[k]
		// fmt.Printf("%v\n", size)
        union(c.a, c.b)
    }
	var finalSizes []int
    for i := range size {
        if parent[i] == i {
            finalSizes = append(finalSizes, size[i])
        }
    }
	sort.Sort(sort.Reverse(sort.IntSlice(finalSizes)))
	// fmt.Printf("size0: %d, size1: %d, size2: %d", finalSizes[0], finalSizes[1], finalSizes[2])
	// fmt.Printf("find: %v", find)
	// fmt.Printf("size: %v", size)
	// fmt.Printf("finalsize: %v", finalSizes)

	total := finalSizes[0] * finalSizes[1] * finalSizes[2]
	return total
}

// solvePart2 contains the logic for the second part of the puzzle.
// It often builds upon or modifies the logic from Part 1.
//

func solvePart2(lines []string) int {
	var connections []Connection
	for i := 0; i < len(lines); i++ {
		dims := strings.Split(lines[i], ",")
		val0, _ := strconv.ParseFloat(strings.TrimSpace(dims[0]), 64)
		val1, _ := strconv.ParseFloat(strings.TrimSpace(dims[1]), 64)
		val2, _ := strconv.ParseFloat(strings.TrimSpace(dims[2]), 64)
		for j := i + 1; j < len(lines); j++ {
			dims2 := strings.Split(lines[j], ",")
			val3, _ := strconv.ParseFloat(strings.TrimSpace(dims2[0]), 64)
			val4, _ := strconv.ParseFloat(strings.TrimSpace(dims2[1]), 64)
			val5, _ := strconv.ParseFloat(strings.TrimSpace(dims2[2]), 64)

			square := math.Pow(val0 - val3, 2) + math.Pow(val1 - val4, 2) + math.Pow(val2 - val5, 2)
			connections = append(connections, Connection{i, j, square})
		}
	}

	// sort
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].dist < connections[j].dist
	})

	// Initialize DSU (Union-Find)
	parent := make([]int, len(lines))
	size := make([]int, len(lines))
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	numCircuits := len(lines) // Start with every box as its own circuit

	// Helper functions for DSU
	var find func(int) int
	find = func(i int) int {
		if parent[i] == i {
			return i
		}
		parent[i] = find(parent[i])
		return parent[i]
	}
	
	union := func(i, j int) {
        rootI, rootJ := find(i), find(j)
        if rootI != rootJ {
            // Merge smaller into larger
            if size[rootI] < size[rootJ] {
                rootI, rootJ = rootJ, rootI
            }
            parent[rootJ] = rootI
            size[rootI] += size[rootJ]
            size[rootJ] = 0 // Optional: reset size of the merged-away root
        }
    }


	fmt.Println(len(connections))
	for k := 0; k < len(connections); k++ {
        c := connections[k]
		rootA := find(c.a)
		rootB := find(c.b)
		
		if rootA != rootB {
			union(rootA, rootB)
			numCircuits--


			if numCircuits == 1{
				dims := strings.Split(lines[c.a], ",")
				vala, _ := strconv.ParseFloat(strings.TrimSpace(dims[0]), 64)
				dims = strings.Split(lines[c.b], ",")
				valb, _ := strconv.ParseFloat(strings.TrimSpace(dims[0]), 64)
				return int(vala * valb)
			}
		}
    }
	
	return 0
}

