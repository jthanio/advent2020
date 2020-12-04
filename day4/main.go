package main

import (
	"fmt"
	"strings"

	"github.com/jthanio/advent2020"
)

const inputFile = "day4.txt"

func main() {
	input, err := advent2020.LoadStringInputFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The answer to the Day 4 part 1 puzzle: %d\n", SolveDay4Part1(input))
}

type passport struct {
	byr bool // Birth Year
	iyr bool // Issue Year
	eyr bool // Expiration Year
	hgt bool // Height
	hcl bool // Hair Color
	ecl bool // Eye Color
	pid bool // Passport ID
	cid bool // Country ID (Optional)
}

// validate checks for matching fields in the given input string and updates the passport for each match.
func (p *passport) validate(input string) {
	if strings.Contains(input, "byr:") {
		p.byr = true
	}
	if strings.Contains(input, "iyr:") {
		p.iyr = true
	}
	if strings.Contains(input, "eyr:") {
		p.eyr = true
	}
	if strings.Contains(input, "hgt:") {
		p.hgt = true
	}
	if strings.Contains(input, "hcl:") {
		p.hcl = true
	}
	if strings.Contains(input, "ecl:") {
		p.ecl = true
	}
	if strings.Contains(input, "pid:") {
		p.pid = true
	}
	if strings.Contains(input, "cid:") {
		p.cid = true
	}
}

// isValid checks if the passport has all of the required fields
func (p *passport) isValid() bool {
	return p.byr && p.iyr && p.eyr && p.hgt && p.hcl && p.ecl && p.pid
}

// SolveDay4Part1 finds the number of valid passports in the given batch input
func SolveDay4Part1(input []string) int {
	var count int
	p := &passport{}
	for _, s := range input {
		if s != "" {
			p.validate(s) // Check for matching fields
		} else {
			// Passport input ended, can attempt to validate
			if p.isValid() {
				count++
			}
			p = &passport{} // Start a new passport
		}
	}

	// Must check if the final passport is valid
	if p.isValid() {
		count++
	}

	return count
}
