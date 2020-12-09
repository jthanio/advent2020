package day9

import (
	"errors"
	"sort"
)

const defaultPreambleLen = 25

// SolveDay9Part1 finds the first number in the list that does not follow the XMAS rules.
func SolveDay9Part1(input []int) (int, error) {
	return findInvalidXMASNumber(input, defaultPreambleLen)
}

// SolveDay9Part2 finds the xmas vulnerability number and then a contiguous set of numbers in the last that sum to that number.
func SolveDay9Part2(input []int) (int, error) {
	target, err := findInvalidXMASNumber(input, defaultPreambleLen)
	if err != nil {
		return 0, err
	}
	set := findContiguousSummation(target, input)
	if len(set) < 2 {
		return 0, errors.New("could not find large enough summation")
	}
	sort.Ints(set)
	return set[0] + set[len(set)-1], nil
}

func findContiguousSummation(target int, numSlice []int) []int {
	for i := 0; i < len(numSlice); i++ {
		var sum int
		for j := i; j < len(numSlice); j++ {
			s := numSlice[i:j]
			sum = sumSlice(numSlice[i:j])
			if sum == target {
				return s
			}
		}
	}
	return nil
}

func findInvalidXMASNumber(input []int, preambleLength int) (int, error) {
	if len(input) < preambleLength {
		return 0, errors.New("preamble length exceeds length of input")
	}
	var pos int
	for _, target := range input[preambleLength:] {
		if !isSumFromSlicePair(target, input[:preambleLength+pos]) {
			return target, nil
		}
		pos++
	}
	return 0, errors.New("could not find the invalid number")
}

func isSumFromSlicePair(target int, numSlice []int) bool {
	// Make a map containing each unique number
	numSet := make(map[int]struct{}, len(numSlice))
	for _, x := range numSlice {
		numSet[x] = struct{}{}
	}

	// For each number, see if any pairs sum to the target
	for x := range numSet {
		if _, ok := numSet[target-x]; ok {
			return true // A pair has been found
		}
	}
	return false
}

func sumSlice(s []int) (y int) {
	for _, x := range s {
		y += x
	}
	return
}
