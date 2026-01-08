package q1

import (
	"strconv"
	"strings"
)

// Solve returns the solutions for part 1 and part 2
func Solve(lines []string) (int, int) {
	// Example: Let's say Part 1 asks to sum all numbers in the input
	// and Part 2 asks to sum only even numbers

	part1 := 0
	part2 := 0

	// Process each line
	for _, line := range lines {
		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Example: Parse the line as a number
		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			// Handle parsing errors if needed
			continue
		}

		// Part 1: Sum all numbers
		part1 += num

		// Part 2: Sum only even numbers
		if num%2 == 0 {
			part2 += num
		}
	}

	return part1, part2
}
