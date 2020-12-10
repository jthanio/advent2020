package day10

import (
	"fmt"
	"sort"
)

// SolveDay10Part1 tests that there is a path from the starting joltage to the last adaptor among all adaptors and returns a product of the differences per test.
func SolveDay10Part1(input []int) (int, error) {
	if len(input) == 0 {
		return 0, nil
	}

	adaptors := createIntSet(input)
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

// SolveDay10Part2 calculates the total number of combinations that the adaptors can be arranged in.
func SolveDay10Part2(input []int) (int, error) {
	if len(input) == 0 {
		return 0, nil
	}
	sort.Ints(input)
	return countPermutations(input, 0, input[len(input)-1]), nil
}

func countPermutations(input []int, from, to int) int {
	memo := map[int]int{
		to: 1, // The last hop is always 1
	}
	adaptors := createIntSet(input)
	return countPermutationsHelper(memo, adaptors, from)
}

func countPermutationsHelper(memo map[int]int, adaptors map[int]struct{}, from int) int {
	if v, ok := memo[from]; ok {
		return v
	}
	var total int
	for i := 1; i <= 3; i++ {
		if _, ok := adaptors[from+i]; ok {
			total += countPermutationsHelper(memo, adaptors, from+i)
		}
	}
	memo[from] = total
	return total
}

// selectNextAdaptor chooses the next appropriate adaptor in the collection that can match with the current jolts.
func selectNextAdaptor(current int, adaptors map[int]struct{}) (int, error) {
	for i := 1; i <= 3; i++ {
		if _, ok := adaptors[current+i]; ok {
			return current + i, nil
		}
	}
	return 0, fmt.Errorf("found no appropriate adaptor for %d jolts", current)
}

func createIntSet(s []int) map[int]struct{} {
	set := make(map[int]struct{})
	for _, v := range s {
		set[v] = struct{}{}
	}
	return set
}
