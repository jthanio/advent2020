package day10

import (
	"fmt"
)

// SolveDay10Part1 tests that there is a path from the starting joltage to the last adaptor among all adaptors and returns a product of the differences per test.
func SolveDay10Part1(input []int) (int, error) {
	if len(input) == 0 {
		return 0, nil
	}

	adaptors := make(map[int]struct{}, len(input))
	for _, v := range input {
		adaptors[v] = struct{}{}
	}

	next, err := selectNextAdaptor(0, adaptors) // The initial joltage to adapt from is 0
	if err != nil {
		return 0, err
	}

	// Maintain a count of the 1-jolt and 3-jolt differences
	threes := 1 // The device's built-in adaptor is always higher than the highest adaptor in the list
	ones := 0
	for prev := 0; err == nil; next, err = selectNextAdaptor(next, adaptors) {
		diff := next - prev
		switch diff {
		case 1:
			ones++
		case 3:
			threes++
		}
		prev = next
	}

	return ones * threes, nil
}

func SolveDay10Part2(input []int) (int, error) {
	return 0, nil
}

// selectNextAdaptor chooses the next appropriate adaptor in the collection that can match with the current jolts.
func selectNextAdaptor(current int, adaptors map[int]struct{}) (int, error) {
	for i := 1; i < 4; i++ {
		if _, ok := adaptors[current+i]; ok {
			return current + i, nil
		}
	}
	return 0, fmt.Errorf("found no appropriate adaptor for %d jolts", current)
}
