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

	fmt.Println("--- Advent of Code 2025 - Day 03 ---")

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

	for _, line := range lines {
		bigDgt := map[int]int{
			0: 0,
			1: 0,
		}
		
		lineLen := len(line)

		for i := 0; i < lineLen - 1; i++ {
			currentDgt := int(line[i] - '0')
			// fmt.Printf("%d", currentDgt)
			
			if currentDgt > bigDgt[0] {
				bigDgt[0] = currentDgt
				bigDgt[1] = 0
			} else if currentDgt == bigDgt[0] {
				bigDgt[1] = currentDgt
			} else {
				if currentDgt >= bigDgt[1] {
					bigDgt[1] = currentDgt
				}
			}

		}
		lastDgt := int(line[lineLen - 1] - '0')
		if bigDgt[1] < lastDgt {
			bigDgt[1] = lastDgt
		}
		current := bigDgt[0] * 10 + bigDgt[1]
		// fmt.Printf("number of this line is %d\n", current)
		total += current 
	}
	return total 
}

// solvePart2 contains the logic for the second part of the puzzle.
// It often builds upon or modifies the logic from Part 1.
func solvePart2(lines []string) int {
	total := 0

	for _, line := range lines {
		current := 0
		bigDgt := map[int]int{
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
			6: 0,
			7: 0,
			8: 0,
			9: 0,
			10: 0,
			11: 0,
		}
		
		lineLen := len(line)

		for i:= 0; i < lineLen; i++ {
// 5373475263753258336423442254746263332334232217334431337464342726873125223932312363675175435324343745

			currentDgt := int(line[i] - '0')

			for j := 0; j < len(bigDgt); j++ {
				// fmt.Printf("key: %d, value: %d\n", j, bigDgt[j])
				// fmt.Printf("currentDgt is %d, lenbigdgt: %d\n", currentDgt, len(bigDgt))
				// passkey, for key bigger than passkey, reset correspondent value to 0
				// 81111111111111[9]
				// lineLen=15, i=15, resLen=2, key=0 lineLen-i=0 >= reslen-key-2=0 no
				// lineLen=15, i=15, resLen=2, key=1 lineLen-i=0 >= reslen-key-2=-1 yes
				// 8111111111111[9]1
				// lineLen=15, i=14, reslen=2, key=0 lineLen-i=1 >= reslen-key-2=0 yes
				// 8[1]1111111111144
				// lineLen=15, i=1, reslen=2, key=1 lineLen-i=14 >= reslen-key-2=-1 yes
				if currentDgt > bigDgt[j] && lineLen-i >= 12-j {
					bigDgt[j] = currentDgt
					for k := j + 1; k < len(bigDgt); k++ {
						bigDgt[k] = 0
					}
					break
				}
			}
			// fmt.Printf("bigDgt: %v\n", bigDgt)

		}

		for i := 0; i < len(bigDgt); i++ {
			// fmt.Printf("value: %d, current: %d\n", bigDgt[i], current)
			current = current * 10 + bigDgt[i]
			// fmt.Printf("current: %d\n", current)
		}
		// fmt.Printf("number of this line is %d\n", current)
		total += current
	}
	return total
}


