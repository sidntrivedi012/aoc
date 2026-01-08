package q1

import (
	"strconv"
)

// Solve returns the solutions for part 1 and part 2
func Solve(lines []string) int {
	current := 50
	count := 0
	passZero := false

	for _, line := range lines {
		passZero = false
		dir := line[:1]
		value, _ := strconv.Atoi(line[1:])

		switch dir {
		case "L":
			distance := value % 100
			count = count + (value / 100)

			current = current - distance
			if current < 0 {
				if current+distance != 0 {
					passZero = true
				}
				current = 100 + current
			}
			if current != 0 && passZero {
				count++
			}
			if current == 0 {
				count++
			}
			break
		case "R":
			distance := value % 100
			current = current + distance
			count = count + (value / 100)

			if current > 99 {
				passZero = true
				current = current % 100
			}
			if current != 0 && passZero {
				count++
			}
			if current == 0 {
				count++
			}
			break
		}
	}

	return count
}
