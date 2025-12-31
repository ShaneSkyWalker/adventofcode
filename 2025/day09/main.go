package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
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
			xDist := math.Abs(float64(dots[i].x - dots[j].x)) + 1
			yDist := math.Abs(float64(dots[i].y - dots[j].y)) + 1
			pair := Pairs{i, j, int(xDist * yDist)}
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

type Point struct {
	X, Y int
}

type Edge struct {
	P1, P2 Point
}

func solvePart2(lines []string) int64 {
	points := []Point{}
	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(coords[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(coords[1]))
		points = append(points, Point{x, y})
	}

	n := len(points)
	edges := make([]Edge, n)
	for i := 0; i < n; i++ {
		edges[i] = Edge{points[i], points[(i+1)%n]}
	}

	var maxArea int64 = 0

	// Iterate through every pair of red tiles as opposite corners
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1, p2 := points[i], points[j]

			xmin, xmax := min(p1.X, p2.X), max(p1.X, p2.X)
			ymin, ymax := min(p1.Y, p2.Y), max(p1.Y, p2.Y)

			width := int64(xmax - xmin) + 1
			height := int64(ymax - ymin) + 1
			area := width * height

			if area <= maxArea || width == 0 || height == 0 {
				continue
			}

			// 1. Midpoint Check (Ray Casting)
			midX := float64(xmin+xmax) / 2.0
			midY := float64(ymin+ymax) / 2.0

			if isInside(midX, midY, edges) {
				// 2. Slicing Check (Ensure no walls pass through the rectangle)
				if !isSliced(xmin, xmax, ymin, ymax, edges) {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

// isInside uses the Ray Casting algorithm to check if a point is within the polygon
func isInside(x, y float64, edges []Edge) bool {
	inside := false
	for _, e := range edges {
		y1, y2 := float64(e.P1.Y), float64(e.P2.Y)
		x1, x2 := float64(e.P1.X), float64(e.P2.X)

		if (y1 > y) != (y2 > y) {
			intersectX := (x2-x1)*(y-y1)/(y2-y1) + x1
			if x < intersectX {
				inside = !inside
			}
		}
	}
	return inside
}

// isSliced checks if any vertical or horizontal polygon edge cuts the rectangle interior
func isSliced(xmin, xmax, ymin, ymax int, edges []Edge) bool {
	for _, e := range edges {
		exMin, exMax := min(e.P1.X, e.P2.X), max(e.P1.X, e.P2.X)
		eyMin, eyMax := min(e.P1.Y, e.P2.Y), max(e.P1.Y, e.P2.Y)

		if e.P1.X == e.P2.X { // Vertical Edge
			if e.P1.X > xmin && e.P1.X < xmax {
				// Check if Y-ranges overlap
				if max(ymin, eyMin) < min(ymax, eyMax) {
					return true
				}
			}
		} else { // Horizontal Edge
			if e.P1.Y > ymin && e.P1.Y < ymax {
				// Check if X-ranges overlap
				if max(xmin, exMin) < min(xmax, exMax) {
					return true
				}
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func max(a, b int) int {
	if a > b { return a }
	return b
}