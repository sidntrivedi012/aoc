package q3

import "strconv"

func Solve(lines []string) int {
	sum := 0

	for _, line := range lines {
		res, _ := strconv.Atoi(getMaxJoltage(line, 12))
		sum = sum + res
	}

	return sum
}

// getMaxJoltage returns the max joltage for a
// input string using the greedy algorithm.

// Start iterating the numbers
// Push the first number into the stack
// If the next number is greater than the top element of stack and still elements can be removed,
// pop the top of stack and push the greater number until the top of the stack has greater number
// than the next number. Keep decreasing value of remove variable as well.
func getMaxJoltage(s string, keep int) string {
	n := len(s)
	remove := n - keep

	stack := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		d := s[i]

		for remove > 0 && len(stack) > 0 && stack[len(stack)-1] < d {
			stack = stack[:len(stack)-1]
			remove--
		}

		stack = append(stack, d)
	}

	if remove > 0 {
		stack = stack[:len(stack)-remove]
	}
	return string(stack[:keep])
}
