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

