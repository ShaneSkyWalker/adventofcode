package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

// Global variable for the input file path relative to the day directory
const inputFile = "input.txt"

func main() {
	// Read input and handle potential errors
	scopes, ingres, err := readInput(inputFile)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	fmt.Println("--- Advent of Code 2025 - Day 04 ---")

	// Execute Part 1
	result1 := solvePart1(scopes, ingres)
	fmt.Printf("Part 1 Result: %d\n", result1)

	// Execute Part 2
	result2 := solvePart2(scopes)
	fmt.Printf("Part 2 Result: %d\n", result2)
}

// readInput reads a file line-by-line and returns a slice of strings.
// It is designed to be reusable for all days.
func readInput(filename string) ([]string, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	var scopes []string
	var ingres []string
	var isScope bool = true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			isScope = false
			continue
		}
		if isScope {
			scopes = append(scopes, text)
		} else {
			ingres = append(ingres, text)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error during file scan: %w", err)
	}

	return scopes, ingres, nil
}

// solvePart1 contains the logic for the first part of the puzzle.
func solvePart1(scopes []string, ingres []string) int {

	total := 0
	for i := 0; i < len(ingres); i++ {
		ingreUInt, _ := strconv.ParseUint(ingres[i], 10, 0)
		for j := 0; j < len(scopes); j++ {
			parts := strings.Split(scopes[j], "-")
			leftUInt, _  := strconv.ParseUint(parts[0], 10, 0)
			rightUInt, _ := strconv.ParseUint(parts[1], 10, 0)
			if ingreUInt >= leftUInt && ingreUInt <= rightUInt {
				total++
				break
			}
		}
	}
	return total
}

type Scope struct {
	start int64
	end   int64
}
// solvePart2 contains the logic for the second part of the puzzle.
// It often builds upon or modifies the logic from Part 1.
func solvePart2(scopes []string) int64 {
	var (
		scopeRange []Scope
		total int64
	)

	for i := 0; i < len(scopes); i++ {
		parts := strings.Split(scopes[i], "-")
		leftInt, _  := strconv.ParseInt(parts[0], 10, 64)
		rightInt, _ := strconv.ParseInt(parts[1], 10, 64)
		scopeRange = append(scopeRange, Scope{leftInt, rightInt})
	}

	sort.Slice(scopeRange, func(i, j int) bool { return scopeRange[i].start < scopeRange[j].start })
	var mergedScope []Scope
	for {
		if len(scopeRange) == 1 {
			mergedScope = append(mergedScope, scopeRange...)
			break
		}
		first  := scopeRange[0]
		second := scopeRange[1]
		scopeRange = scopeRange[2:]
		
		if second.start > first.end {
			// not intersect
			mergedScope = append(mergedScope, first)
			scopeRange = append([]Scope{second}, scopeRange...)
		} else {
			scopeRange = append([]Scope{{first.start, max(first.end, second.end)}}, scopeRange...)
		}
	} 

	for _, r := range mergedScope {
		total += r.end - r.start + 1
	}
	return int64(total)
}
// func solvePart2(scopes []string) int64 {
// 	// Initialize wideleft and wideright
// 	parts_0 := strings.Split(scopes[0], "-")
// 	wideLeftInt, _  := strconv.ParseInt(parts_0[0], 10, 64)
// 	wideRightInt, _ := strconv.ParseInt(parts_0[1], 10, 64)
//
// 	total := int64(0)
//
// 	// fmt.Printf("total: %d\n", total)
//
// 	for i := 1; i < len(scopes); i++ {
//
// 		parts := strings.Split(scopes[i], "-")
// 		leftInt, _  := strconv.ParseInt(parts[0], 10, 64)
// 		rightInt, _ := strconv.ParseInt(parts[1], 10, 64)
// 		gap := int64(0)
//
// 		// condition 1: there is gap, no overlap
// 		if rightInt <= wideLeftInt{
// 			if wideLeftInt == rightInt {
// 				gap = 0
// 			} else {
// 				gap = wideLeftInt - rightInt - 1
// 			}
// 			wideLeftInt = leftInt
// 		}
//
// 		// condition 2: only overlap
// 		if leftInt <= wideLeftInt && wideLeftInt <= rightInt && rightInt <= wideRightInt {
// 			wideLeftInt = leftInt
// 		}
//
// 		// condition 3: only overlap
// 		if leftInt <= wideLeftInt && wideRightInt <= rightInt {
// 			wideLeftInt = leftInt
// 			wideRightInt = rightInt
// 		}
//
// 		// condition 4: only overlap
// 		if wideLeftInt <= leftInt && rightInt <= wideRightInt {
// 		}
//
// 		// condition 5: only overlap
// 		if wideLeftInt <= leftInt && leftInt <= wideRightInt && wideRightInt <= rightInt {
// 			wideRightInt = rightInt
// 		}
//
// 		// condition 6: only gap
// 		if wideRightInt <= leftInt {
// 			if wideRightInt == leftInt {
// 				gap = 0
// 			} else {
// 				gap = leftInt - wideRightInt - 1
// 			}
// 			wideRightInt = rightInt
// 		}
//
// 		// ignore overlay, but consider gap
// 		// fmt.Printf("left: %d, right: %d, total: %d, gap: %d\n", wideLeftInt, wideRightInt, total, gap)
// 		total = wideRightInt - wideLeftInt + 1 - gap
// 	}
// 	return total
// }


