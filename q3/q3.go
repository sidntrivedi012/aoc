package q3

func Solve(lines []string) int {
	sum := 0

	for _, line := range lines {
		maxVal := 0

		for i := 0; i < len(line); i++ {
			d1 := int(line[i] - '0')

			for j := i + 1; j < len(line); j++ {
				d2 := int(line[j] - '0')
				val := d1*10 + d2

				if val > maxVal {
					maxVal = val
				}
			}
		}

		sum += maxVal
	}

	return sum
}
