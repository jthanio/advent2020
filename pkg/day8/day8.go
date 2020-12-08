package day8

import (
	"errors"
	"regexp"
	"strconv"
)

// Define the instructions used by the boot code.
const (
	insNOP = "nop"
	insACC = "acc"
	insJMP = "jmp"
)

// SolveDay8Part1 parses computer instructions and returns the value of the accumulator before it enters an infinite loop.
func SolveDay8Part1(input []string) (int, error) {
	var acc int
	visited := make(map[int]struct{}, len(input))
	for i := 0; i < len(input); {
		if _, ok := visited[i]; ok {
			return int(acc), nil // Found the loop
		}
		visited[i] = struct{}{}

		var err error
		i, acc, err = evaluateLine(input[i], i, acc)
		if err != nil {
			return 0, err
		}
	}
	return int(acc), errors.New("reached end of instructions without finding loop")
}

type accumulator int

func evaluateLine(line string, index, acc int) (int, int, error) {
	r := regexp.MustCompile(`(\w\w\w) ([+-]\d*)`)
	m := r.FindStringSubmatch(line)
	ins := m[1] // The actual instruction
	val, err := strconv.Atoi(m[2])
	if err != nil {
		return 0, acc, err
	}

	switch ins {
	case insNOP:
		return index + 1, acc, nil // Execute the next instruction
	case insACC:
		return index + 1, acc + val, nil // Modify the accumulator by the value
	case insJMP:
		return index + val, acc, nil // Modify the index based on the value
	}
	return 0, acc, errors.New("could not parse a valid boot code instruction")
}

// SolveDay8Part2 is not yet implemented.
func SolveDay8Part2(input []string) (int, error) {
	return 0, nil
}
