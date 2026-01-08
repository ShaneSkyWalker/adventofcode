package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "strconv"
	// "regexp"
	// "sort"
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

	fmt.Println("--- Advent of Code 2025 - Day 11 ---")

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
	devices := make(map[string]int)
	total := 0
	var search func(device string)
	search = func(device string) {
		var match []string
		if value, ok := devices[device]; ok {
			line := lines[value]
			match = strings.Split(line, ":")
		} else {
			for i, line := range lines {
				match = strings.Split(line, ":")
				if match[0] == device {
					devices[match[0]] = i
					break
				}
			}
		}
		subDevices := strings.Split(strings.Trim(match[1], " "), " ")
		for _, subDevice := range subDevices {
			if subDevice == "out" {
				total += 1
				return
			}
			search(subDevice)
		}
	}
	search("you")
	return total
}

// solvePart2 contains the logic for the second part of the puzzle.
// It often builds upon or modifies the logic from Part 1.

func solvePart2(lines []string) int {
	devices := make(map[string]int)
	total := 0


	// 2. Setup Memoization
	// Key: "nodeName-hasFFT-hasDAC"
	memo := make(map[string]int)

	var search func(device string, is_fft, is_dac bool)
	search = func(device string, is_fft, is_dac bool) {
		var match []string
		if value, ok := devices[device]; ok {
			line := lines[value]
			match = strings.Split(line, ":")
			return
		} else {
			for i, line := range lines {
				match = strings.Split(line, ":")
				if match[0] == device {
					devices[match[0]] = i
					break
				}
			}
		}
		subDevices := strings.Split(strings.Trim(match[1], " "), " ")
		// fmt.Println(subDevices)
		for _, subDevice := range subDevices {
			if subDevice == "out" {
				if is_fft && is_dac {
					total += 1
				}
				return
			}
			if match[0] == "fft" {
				is_fft = true
			}
			if match[0] == "dac" {
				is_dac = true
			}
			search(subDevice, is_fft, is_dac)
		}

		stateKey := fmt.Sprintf("%s-%t-%t", device, is_fft, is_dac)
		if _, exists := memo[stateKey]; exists {
			return
		}

		memo[stateKey] = 1
		
	}
	search("svr", false, false)
	return total
}