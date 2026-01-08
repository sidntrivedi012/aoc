package q2

import (
	"strconv"
	"strings"
)

func Solve(lines []string) int {
	sum := 0
	lines = strings.Split(lines[0], ",")
	for _, line := range lines {
		entry := strings.Split(line, "-")
		start := entry[0]
		end := entry[1]

		beginIndex, _ := strconv.Atoi(start)
		endIndex, _ := strconv.Atoi(end)

		for i := beginIndex; i <= endIndex; i++ {
			if len(strconv.Itoa(i))%2 != 0 {
				continue
			}
			middle := len(strconv.Itoa(i)) / 2

			firstHalf := strconv.Itoa(i)[:middle]
			secondHalf := strconv.Itoa(i)[middle:]
			// fmt.Printf("First half: %s, Second half: %s", firstHalf, secondHalf)

			if strings.EqualFold(firstHalf, secondHalf) {
				sum = sum + i
			}
		}
	}
	return sum
}
