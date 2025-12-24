package main

import (
	"bufio"
	"fmt"
	"os"
)

// Global variable for the input file path relative to the day directory
const inputFile = "input2.txt"

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
	beams := make(map[int]struct{})
	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == 'S' {
			beams[i] = struct{}{}
			break
		}
	}
	for i := 1; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == '^' {
				if _, exists := beams[j]; exists {
					delete(beams, j)
					if j != 0 {
						beams[j-1] = struct{}{}
					}
					if j != len(lines[0]) - 1 {
						beams[j+1] = struct{}{}
					}
					total += 1
				}

			}
		}
	}
	return total
}

// solvePart2 contains the logic for the second part of the puzzle.
// It often builds upon or modifies the logic from Part 1.
//
type Beam struct {
	x int
	y int
	rep int
}
type BeamKey struct {
	x, y int
}
func solvePart2(lines []string) int {
	total := 0
	beamMap := make(map[BeamKey]*Beam)
	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == 'S' {
			beam := &Beam{1, i, 1}
			beamMap[BeamKey{beam.x, beam.y}] = beam
		}
	}
	for i := 2; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == '^' {
				// fmt.Println(beamMap)
				if beam, exists := beamMap[BeamKey{i-1, j}]; exists {
					//fmt.Printf("beam:%v", beamMap)
					leftBeam := &Beam{}
					rightBeam := &Beam{}
					if curBeam, subExists := beamMap[BeamKey{i, j-1}]; subExists {
						leftBeam = &Beam{i, j-1, beam.rep + curBeam.rep}
						delete(beamMap, BeamKey{i, j-1})
					} else {
						leftBeam = &Beam{i, j-1, beam.rep}
					}
					if curBeam, subExists := beamMap[BeamKey{i, j+1}]; subExists {
						rightBeam = &Beam{i, j+1, beam.rep + curBeam.rep}
						delete(beamMap, BeamKey{i, j+1})
					} else {
						rightBeam = &Beam{i, j+1, beam.rep}
					}
					beamMap[BeamKey{i, j-1}] = leftBeam
					beamMap[BeamKey{i, j+1}] = rightBeam
					fmt.Printf("%v\n", leftBeam)
					fmt.Printf("%v\n", rightBeam)
				}
			} else {
				if beam, exists := beamMap[BeamKey{i-1, j}]; exists {
					beamMap[BeamKey{i, j}] = &Beam{i, j, beam.rep}
				}
			}
		}
	}
	for j := 0; j < len(lines[0]); j++ {
		if beam, exists := beamMap[BeamKey{len(lines) - 1, j}]; exists {
			total += beam.rep
		}
	}
	return total
}

