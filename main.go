package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sidntrivedi012/aoc/q1"
	"github.com/sidntrivedi012/aoc/q2"
	"github.com/sidntrivedi012/aoc/q3"
	"github.com/sidntrivedi012/aoc/q4"
	"github.com/sidntrivedi012/aoc/utils"
)

func main() {
	// Flags
	question := flag.Int("q", 1, "Question number to run (e.g., 1 for q1)")
	useTest := flag.Bool("test", false, "Use test_input.txt instead of input.txt")
	flag.Parse()

	// Determine input file based on question and test flag
	inputFile := fmt.Sprintf("q%d/input.txt", *question)
	if *useTest {
		inputFile = fmt.Sprintf("q%d/test_input.txt", *question)
	}

	// Check if file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		log.Fatalf("Input file %s does not exist", inputFile)
	}

	// Read input
	lines, err := utils.ReadLines(inputFile)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Run the appropriate solution
	var ans int
	switch *question {
	case 1:
		ans = q1.Solve(lines)
	case 2:
		ans = q2.Solve(lines)
	case 3:
		ans = q3.Solve(lines)
	case 4:
		ans = q4.Solve(lines)
	default:
		log.Fatalf("Solution for question %d not implemented yet", *question)
	}

	// Print results
	fmt.Printf("Question %d Results:\n", *question)
	fmt.Printf("Answer: %d\n", ans)
}
