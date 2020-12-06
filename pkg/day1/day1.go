package day1

import (
	"fmt"
)

// SolveDay1Part1 solves the Day 1 puzzle by finding two entries within the input slice that sum to 2020 and then returning their product.
func SolveDay1Part1(input []int) int {
	if len(input) <= 1 {
		fmt.Println("not enough entries in input data")
		return 0
	}

	// Store input in a map for faster lookup
	mappedInput := make(map[int]struct{}, len(input))
	for _, i := range input {
		mappedInput[i] = struct{}{}
	}

	for _, i := range input {
		diff := 2020 - i
		if _, ok := mappedInput[diff]; ok {
			return i * diff // If diff matches an entry, then the pair has been found
		}
	}

	return 0
}

// SolveDay1Part2 solves the Day 1 puzzle by finding three entries within the input slice that sum to 2020 and then returning their product.
func SolveDay1Part2(input []int) int {
	// Naive approach
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					return input[i] * input[j] * input[k]
				}
			}
		}
	}

	return 0
}
