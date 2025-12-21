package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
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

	fmt.Println("--- Advent of Code 2025 - Day 02 ---")

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

	rangeStrs := strings.Split(lines[0], ",")
	return rangeStrs, nil
}

// solvePart1 contains the logic for the first part of the puzzle.
func solvePart1(pairs []string) int64 {
	var total int64 = 0
	for _, pairStr := range pairs {
		parts := strings.Split(pairStr, "-")
		if len(parts) != 2 {
			continue
		}
		start, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return -1
		}
		end, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return -1
		}
		i := start
		for i <= end {
			s := strconv.FormatInt(i, 10)
			idRunes := []rune(s)
			lenRunes := len(idRunes)
			if lenRunes % 2 != 0 {
				// for all i in start to lenRunes of 9s, there is no what 
				// we are looking for.
				i = int64(math.Pow(10.0, float64(lenRunes)))
				continue
			}

			// only for length of i is even, may have what is asked for
			firstHalfRunes := idRunes[:lenRunes / 2]
			firstHalfStr := string(firstHalfRunes)
			firstHalfInt, err := strconv.ParseInt(firstHalfStr, 10, 64)
			if err != nil {
				fmt.Printf("Error converting string to int: %v\n", err)
				return -1
			}
			pssblInvalidId := firstHalfInt * int64(math.Pow(10.0, float64(lenRunes / 2))) + firstHalfInt
			if pssblInvalidId < i {
				firstHalfInt += 1
				pssblInvalidId := firstHalfInt * int64(math.Pow(10.0, float64(lenRunes / 2))) + firstHalfInt
				i = pssblInvalidId
			} else if pssblInvalidId <= end {
				// in the scope, adds up and go on
				total += pssblInvalidId
				firstHalfInt += 1
				pssblInvalidId := firstHalfInt * int64(math.Pow(10.0, float64(lenRunes / 2))) + firstHalfInt
				i = pssblInvalidId
			} else {
				break
			}
		}
	}
	return total 
}

func fillin(itr int64, lens int64, divisor int) int64 {
	sum := itr

	quotient := lens / int64(divisor)
	for j := int64(0); j < quotient - 1; j++ {
		sum = sum * int64(math.Pow(10.0, float64(divisor))) + itr
	}

	return sum
}

func solvePairs(start int64, end int64) int64 {
	var total int64 = 0
	primeNumArr := []int{2, 3, 5, 7, 11, 13, 17, 19}
	primeNum := make(map[int]bool)

	for _, number := range primeNumArr {
		primeNum[number] = true
	}
	cmpstNum := map[int][]int {
		4: []int{2},
		6: []int{2, 3},
		8: []int{2, 4},
		9: []int{3},
		10: []int{2, 5},
		12: []int{2, 3, 4, 6},
		14: []int{2, 7},
		15: []int{3, 5},
		16: []int{2, 4, 8},
		18: []int{2, 3, 6, 9},
	}
	
	// tidy up input
	s := strconv.FormatInt(start, 10)
	idRunes := []rune(s)
	lenRunes := len(idRunes)
	// _, existsInPrimeNum := primeNum[lenRunes]
	divisors := []int{}
	// fmt.Printf("lenRunes: %s, cmpstNum[lenRunes]: %v, len(cmpstNum[lenRunes]): %d\n", lenRunes, cmpstNum[lenRunes], len(cmpstNum[lenRunes]))

	if lenRunes == 1 {
		// single digit always cannot be invalid ID.
		return 0
	} else if len(cmpstNum[lenRunes]) > 0 {
		// divisors := cmpstNum[lenRunes]
		divisors = append(divisors, cmpstNum[lenRunes]...)
	}
	divisors = append(divisors, 1)
	// fmt.Printf("divisors: %v\n", divisors)
	pssblInvalidIdMap := make(map[int64]bool)
	for _, divisor := range divisors {
		firstDivisorStr := string(idRunes[:divisor])
		// fmt.Printf("firstDivisorStr: %s\n", firstDivisorStr)
		firstDivisorInt, err := strconv.ParseInt(firstDivisorStr, 10, 64)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return -1
		}

		i := start
		for i <= end {
			pssblInvalidId := fillin(firstDivisorInt, int64(lenRunes), divisor)
			if pssblInvalidId < start {
				firstDivisorInt += 1
				i = fillin(firstDivisorInt, int64(lenRunes), divisor)
			} else if pssblInvalidId <= end {
				// in the scope
				pssblInvalidIdMap[pssblInvalidId] = true
				firstDivisorInt += 1
				i = fillin(firstDivisorInt, int64(lenRunes), divisor)
			} else {
				break
			}
			if len(strconv.FormatInt(firstDivisorInt, 10)) > len(firstDivisorStr) {
				// firstDivisorInt overflow
				break
			}
		}
	
	}
	for ptntlInvalidId, _ := range pssblInvalidIdMap {
		total += ptntlInvalidId
	}
	return total
}

// solvePart2 contains the logic for the second part of the puzzle.
// It often builds upon or modifies the logic from Part 1.
func solvePart2(pairs []string) int64 {
	var total int64 = 0
	for _, pairStr := range pairs {
		parts := strings.Split(pairStr, "-")
		if len(parts) != 2 {
			continue
		}
		start, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return -1
		}
		end, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return -1
		}

		// compare length of start and end, if not same, divide 
		if  len(parts[1]) - len(parts[0]) == 0 {
			total += solvePairs(start, end)
		} else if len(parts[1]) - len(parts[0]) == 1 {
			// get end length 
			middle := int64(math.Pow(10.0, float64(len(parts[0]))))
			total += solvePairs(start, middle - 1)
			total += solvePairs(middle, end)
		} else if len(parts[1]) - len(parts[0]) > 1 {
			// not happen according to current data
			fmt.Println("len(end) - len(start) > 1")
		}
	}
	return  total
}
