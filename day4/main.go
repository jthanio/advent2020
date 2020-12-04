package main

import (
	"fmt"

	"github.com/jthanio/advent2020"
	"github.com/jthanio/advent2020/day4/passport"
)

const inputFile = "day4.txt"

func main() {
	input, err := advent2020.LoadStringInputFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The answer to the Day 4 part 1 puzzle: %d\n", SolveDay4Part1(input))
	fmt.Printf("The answer to the Day 4 part 2 puzzle: %d\n", SolveDay4Part2(input))
}

// SolveDay4Part1 finds the number of valid passports in the given batch input
func SolveDay4Part1(input []string) int {
	var count int
	p := &passport.Passport{}
	for _, s := range input {
		if s == "" {
			// Passport input ended, can attempt to validate
			if p.HasAllRequiredFields() {
				count++
			}
			p = &passport.Passport{} // Start a new passport
		} else {
			p.ValidateSimple(s) // Check for matching fields
		}
	}

	// Must check if the final passport is valid
	if p.HasAllRequiredFields() {
		count++
	}

	return count
}

// SolveDay4Part2 finds the number of valid passports in the given batch input, with stricter parsing rules
func SolveDay4Part2(input []string) int {
	var count int
	p := &passport.Passport{}
	for _, s := range input {
		if s == "" {
			// Passport input ended, can attempt to validate
			if p.HasAllValidFields() {
				count++
			}
			p = &passport.Passport{} // Start a new passport
		} else {
			p.ValidateStrict(s) // Check for matching fields
		}
	}

	// Must check if the final passport is valid
	if p.HasAllValidFields() {
		count++
	}

	return count
}
