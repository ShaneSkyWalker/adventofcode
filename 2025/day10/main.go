package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"regexp"
	// "sort"
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

var minPresses int
func solvePart2(lines []string) int {
	total := 0
	for _, line := range lines {
		reButtons := regexp.MustCompile(`\((.*?)\)`)
		reTarget := regexp.MustCompile(`\{(.*?)\}`)

		buttonMatches := reButtons.FindAllStringSubmatch(line, -1)
		targetMatch := reTarget.FindStringSubmatch(line)
		// fmt.Println(targetMatch[1])

		targetStr := strings.Split(targetMatch[1], ",")
		rows := len(targetStr)
		targets := make([]int, rows)
		for i, s := range targetStr {
			targets[i], _ = strconv.Atoi(strings.TrimSpace(s))
		}

		cols := len(buttonMatches) + 1

		// 1. Convert to float64 for RREF
		matrix := make([][]float64, rows)
		for i := range matrix {
			matrix[i] = make([]float64, cols)
			matrix[i][cols-1] = float64(targets[i])
		}

		for bIdx, match := range buttonMatches {
			affectedStr := strings.Split(match[1], ",")
			// fmt.Printf("%v\n", affectedStr)
			for _, s := range affectedStr {
				s = strings.TrimSpace(s)
				if s == "" { continue }
				cIdx, _ := strconv.Atoi(s)
				if cIdx < rows {
					matrix[cIdx][bIdx] = float64(1)
				}
			}
		}
		// fmt.Printf("%v\n", matrix)

		// 2. Perform Gauss-Jordan Elimination (RREF)
		pivotRow := 0
		pivotCols := make([]int, 0)
		for j := 0; j < cols - 1 && pivotRow < rows; j++ {
			sel := pivotRow
			for i := pivotRow; i < rows; i++ {
				if math.Abs(matrix[i][j]) > math.Abs(matrix[sel][j]) {
					sel = i
				}
			}
			if math.Abs(matrix[sel][j]) < 1e-9 {
				// ==0
				continue
			}
			matrix[pivotRow], matrix[sel] = matrix[sel], matrix[pivotRow]
			pivotCols = append(pivotCols, j)

			divisor := matrix[pivotRow][j]
			for k := j; k < cols; k++ {
				matrix[pivotRow][k] /= divisor
			}

			for i := 0; i < rows; i++ {
				if i != pivotRow {
					factor := matrix[i][j]
					for k := j; k < cols; k++ {
						matrix[i][k] -= factor * matrix[pivotRow][k]
					}
				}
			}
			pivotRow++

		}

		isPivot := make(map[int]bool)
		for _, c := range pivotCols {
			isPivot[c] = true
		}
		// 3. Identify Free Variables
		freeVars := make([]int, 0)
		for j := 0; j < cols - 1; j++ {
			if !isPivot[j] {
				freeVars = append(freeVars, j)
			}
		}

		// fmt.Printf("Main: %v, freeVars: %v\n", pivotCols, freeVars)
		// fmt.Println(matrix)
		
		minPresses = math.MaxInt32
		freeVals := make([]int, len(freeVars))

		// Recurssion or Dijkstra Search
		backtrack(0, freeVals, freeVars, pivotCols, matrix, cols)
		if minPresses == math.MaxInt32 {
			minPresses = 0
			fmt.Printf("%v\n", matrix)
			fmt.Printf("%s\n", line)
		}
		total += minPresses
	}
	return total
}

func backtrack(idx int, freeVals []int, freeVars []int, 
	pivotCols []int, mat [][]float64, totalCols int) {

	if idx == len(freeVars) {
		checkSolution(freeVals, freeVars, pivotCols, mat, totalCols)
		return
	}

	for v := 0; v <= 200; v++ {
		// issue with this number. Bigger the better
		freeVals[idx] = v
		backtrack(idx+1, freeVals, freeVars, pivotCols, mat, totalCols)
	}
}

func checkSolution(freeVals []int, freeVars []int, pivotCols []int, 
	mat [][]float64, totalCols int) {

	currentSum := 0
	for _, v := range freeVals {
		currentSum += v
	}

	if currentSum > minPresses {
		return
	}

	for i, _ := range pivotCols {
		val := mat[i][totalCols-1]
		for fIdx, fCol := range freeVars {
			val -= mat[i][fCol] * float64(freeVals[fIdx])
		}

		rounded := math.Round(val)
		if val < -1e-9 || math.Abs(val-rounded) > 1e-9 {
			return
		}
		currentSum += int(rounded)
	}

	if currentSum < minPresses {
		minPresses = currentSum
	}
}