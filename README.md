# Advent of Code Solutions

Solutions for Advent of Code challenges written in Go.

## Structure

```
aoc/
├── utils/          # Common utility functions
│   └── utils.go    # File reading, parsing helpers
├── q1/             # Day 1 solution
│   ├── q1.go       # Solution code
│   ├── input.txt   # Actual puzzle input
│   └── test_input.txt  # Sample/test input
├── q2/             # Day 2 solution
│   └── ...
└── go.mod
```

## Usage

### Building the Binary

From the project root:

```bash
go build -o aoc
```

### Running a Solution

```bash
# Run question 1 with actual input
./aoc -q 1

# Run question 1 with test input
./aoc -q 1 -test

# Run question 2 with actual input
./aoc -q 2

# Run question 2 with test input
./aoc -q 2 -test
```

You can also use `go run`:

```bash
# Run with actual input
go run main.go -q 1

# Run with test input
go run main.go -q 1 -test
```

### Adding a New Solution

1. Create a new directory (e.g., `q2/`)
2. Create `q2/q2.go` with the following template:

```go
package q2

// Solve returns the solutions for part 1 and part 2
func Solve(lines []string) (int, int) {
    // Part 1 solution
    part1 := 0

    // Part 2 solution
    part2 := 0

    return part1, part2
}
```

3. Add your puzzle input to `q2/input.txt`
4. Add sample input to `q2/test_input.txt`
5. Update `main.go` to import and call your new solution:

```go
import (
    // ... existing imports
    "github.com/sidntrivedi012/aoc/q2"
)

// In the switch statement, add:
case 2:
    part1, part2 = q2.Solve(lines)
```

6. Rebuild: `go build -o aoc`
7. Run: `./aoc -q 2 -test`

## Utility Functions

The `utils` package provides:

- `ReadLines(filename)` - Read file as slice of strings (one per line)
- `ReadFile(filename)` - Read entire file as a single string
- `ParseInts(lines)` - Convert string slice to int slice
- `SplitByEmptyLines(lines)` - Split input into groups separated by blank lines

## Testing Workflow

1. Add sample input from the problem to `test_input.txt`
2. Run with `-test` flag to verify logic with sample data
3. Add actual puzzle input to `input.txt`
4. Run without flags to get the final answer

## Quick Start

1. **Build the binary:**
   ```bash
   go build -o aoc
   ```

2. **Add your inputs:**
   - Put sample input in `q1/test_input.txt`
   - Put actual puzzle input in `q1/input.txt`

3. **Test with sample input:**
   ```bash
   ./aoc -q 1 -test
   ```

4. **Run with actual input:**
   ```bash
   ./aoc -q 1
   ```

