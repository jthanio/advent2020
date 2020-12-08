package day8

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
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
	acc, err := runProgram(input)
	if err == nil {
		return acc, errors.New("reached end of instructions without finding loop") // We actually want to find the loop
	}
	return acc, nil
}

// SolveDay8Part2 is not yet implemented.
func SolveDay8Part2(input []string) (int, error) {
	var acc int
	for i, line := range input {
		ins, err := parseInstruction(line)
		if err != nil {
			return 0, err
		}

		switch ins.ins {
		case insNOP:
			input[i] = strings.ReplaceAll(input[i], insNOP, insJMP) // Swap this to the other instruction
			acc, err := runProgram(input)
			if err == nil {
				return acc, nil // No loop encountered, that was the fix
			}
			input[i] = strings.ReplaceAll(input[i], insJMP, insNOP) // Swap it back
		case insJMP:
			input[i] = strings.ReplaceAll(input[i], insJMP, insNOP) // Swap this to the other instruction
			acc, err := runProgram(input)
			if err == nil {
				return acc, nil // No loop encountered, that was the fix
			}
			input[i] = strings.ReplaceAll(input[i], insNOP, insJMP) // Swap it back
		}
	}

	return acc, errors.New("reached end of instructions without fixing the loop")
}

func runProgram(instructions []string) (int, error) {
	var acc int
	visited := make(map[int]struct{}, len(instructions))
	for i := 0; i < len(instructions); {
		if _, ok := visited[i]; ok {
			return int(acc), errors.New("encountered infinite loop in program") // Found a loop
		}
		visited[i] = struct{}{}

		ins, err := parseInstruction(instructions[i])
		if err != nil {
			return 0, err
		}

		i, acc, err = evaluateInstruction(ins, i, acc)
		if err != nil {
			return 0, err
		}
	}
	return acc, nil
}

type instruction struct {
	ins string
	val int
}

func parseInstruction(line string) (instruction, error) {
	r := regexp.MustCompile(`(\w\w\w) ([+-]\d*)`)
	m := r.FindStringSubmatch(line)
	val, err := strconv.Atoi(m[2])
	if err != nil {
		return instruction{}, err
	}
	return instruction{
		ins: m[1],
		val: val,
	}, nil
}

func evaluateInstruction(ins instruction, index, acc int) (int, int, error) {
	switch ins.ins {
	case insNOP:
		return index + 1, acc, nil // Execute the next instruction
	case insACC:
		return index + 1, acc + ins.val, nil // Modify the accumulator by the value
	case insJMP:
		return index + ins.val, acc, nil // Modify the index based on the value
	}
	return 0, acc, errors.New("could not parse a valid boot code instruction")
}
