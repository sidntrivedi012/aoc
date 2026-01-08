package q1

import (
	"strconv"
)

// Solve returns the solutions for part 1 and part 2
func Solve(lines []string) (int, int) {
	part2 := 0
	current := 50
	count := 0

	for _, line := range lines {
		dir := line[:1]
		value, _ := strconv.Atoi(line[1:])

		switch dir {
		case "L":
			distance := value % 100
			current = current - distance
			if current < 0 {
				current = 100 + current
			}
			if current == 0 {
				count++
			}
		case "R":
			current = current + value
			if current > 99 {
				current = current % 100
			}
			if current == 0 {
				count++
			}
		}
	}

	return count, part2
}
