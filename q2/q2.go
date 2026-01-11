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
			// Check if there's any subsequence, if yes find and add all the subsequences.
			if isInvalidID(i) {
				sum += i
			}
		}
	}
	return sum
}

// isInvalidID creates a lsp array and returns
// true if the id has repeated pattern.
func isInvalidID(n int) bool {
	num := strconv.Itoa(n)
	l := len(num)
	lps := make([]int, l)

	i := 1
	length := 0

	for i < l {
		if num[i] == num[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps[l-1] > 0 && l%(l-lps[l-1]) == 0
}
