package q4

import "strings"

func Solve(lines []string) int {
	sum := 0

	rows := len(lines)
	cols := len(lines[0])

	// Convert input into grid
	grid := make([][]string, rows)
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	// Directions for 8 neighbors
	dirs := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != "@" {
				continue
			}

			count := 0

			// Check all 8 neighbors
			for _, d := range dirs {
				ni := i + d[0]
				nj := j + d[1]

				if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
					if grid[ni][nj] == "@" {
						count++
					}
				}
			}

			if count < 4 {
				sum++
			}
		}
	}

	return sum
}
