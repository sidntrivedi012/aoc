package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads a file and returns its lines as a slice of strings
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// ReadFile reads entire file content as a string
func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ParseInts converts a slice of strings to a slice of integers
func ParseInts(lines []string) ([]int, error) {
	var nums []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

// SplitByEmptyLines splits input into groups separated by empty lines
func SplitByEmptyLines(lines []string) [][]string {
	var groups [][]string
	var current []string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			if len(current) > 0 {
				groups = append(groups, current)
				current = []string{}
			}
		} else {
			current = append(current, line)
		}
	}

	if len(current) > 0 {
		groups = append(groups, current)
	}

	return groups
}

// ExtractDigits extracts all single digits from a string
func ExtractDigits(s string) []int {
	var digits []int
	for _, char := range s {
		if char >= '0' && char <= '9' {
			digits = append(digits, int(char-'0'))
		}
	}
	return digits
}

// ExtractNumbers extracts all numbers (including multi-digit and negative) from a string
func ExtractNumbers(s string) []int {
	var numbers []int
	var current strings.Builder
	var hasDigit bool

	for i, char := range s {
		if char >= '0' && char <= '9' {
			current.WriteRune(char)
			hasDigit = true
		} else if char == '-' && !hasDigit && i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
			current.WriteRune(char)
		} else {
			if hasDigit {
				if num, err := strconv.Atoi(current.String()); err == nil {
					numbers = append(numbers, num)
				}
				current.Reset()
				hasDigit = false
			}
		}
	}

	// Don't forget the last number
	if hasDigit {
		if num, err := strconv.Atoi(current.String()); err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

// ParseGrid converts lines into a 2D grid of characters
func ParseGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

// ParseIntGrid converts lines into a 2D grid of integers (each char is a digit)
func ParseIntGrid(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			if char >= '0' && char <= '9' {
				grid[i][j] = int(char - '0')
			}
		}
	}
	return grid
}

// ParseKeyValue parses lines in "key: value" format
func ParseKeyValue(line, separator string) (string, string) {
	parts := strings.SplitN(line, separator, 2)
	if len(parts) == 2 {
		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	}
	return "", ""
}
