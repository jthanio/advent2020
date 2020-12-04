package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jthanio/advent2020"
)

const inputFile = "day2.txt"

func main() {
	input, err := advent2020.LoadStringInputFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	a1, err := SolveDay2Part1(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	a2, err := SolveDay2Part2(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The answer to the Day 2 part 1 puzzle: %d\n", a1)
	fmt.Printf("The answer to the Day 2 part 2 puzzle: %d\n", a2)
}

// SolveDay2Part1 finds the number of valid passwords in a slice of formatted strings
func SolveDay2Part1(input []string) (int, error) {
	var validCount int
	for _, s := range input {
		mainParts := strings.Split(s, ":") // The first substring specifies the password rules, the second specifies the password itself
		if len(mainParts) != 2 {
			return 0, fmt.Errorf("the input \"%s\" is invalid", s) // Invalid input
		}
		password := strings.TrimSpace(mainParts[1])

		// Parse the rule specification
		ruleParts := strings.Split(mainParts[0], " ")
		if len(ruleParts) != 2 {
			return 0, fmt.Errorf("the password rule specification \"%s\" is invalid", mainParts[0]) // Invalid password rule part
		}

		// Parse the character count rule
		countRuleParts := strings.Split(ruleParts[0], "-")
		if len(countRuleParts) != 2 {
			return 0, fmt.Errorf("the password character count rule specification \"%s\" is invalid", ruleParts[0]) // Invalid password character count rule part
		}
		lowerLimit, err := strconv.Atoi(countRuleParts[0])
		if err != nil {
			return 0, err
		}
		upperLimit, err := strconv.Atoi(countRuleParts[1])
		if err != nil {
			return 0, err
		}

		// Parse the character
		character := ruleParts[1]

		count := strings.Count(password, character)
		if count <= upperLimit && count >= lowerLimit {
			validCount++ // Password meets specified rules
		}
	}

	return validCount, nil
}

// SolveDay2Part2 finds the number of valid passwords in a slice of formatted strings
func SolveDay2Part2(input []string) (int, error) {
	var validCount int
	for _, s := range input {
		mainParts := strings.Split(s, ":") // The first substring specifies the password rules, the second specifies the password itself
		if len(mainParts) != 2 {
			return 0, fmt.Errorf("the input \"%s\" is invalid", s) // Invalid input
		}
		password := []rune(strings.TrimSpace(mainParts[1])) // Use as rune slice for index access

		// Parse the rule specification
		ruleParts := strings.Split(mainParts[0], " ")
		if len(ruleParts) != 2 {
			return 0, fmt.Errorf("the password rule specification \"%s\" is invalid", mainParts[0]) // Invalid password rule part
		}

		// Parse the character position rule
		posRuleParts := strings.Split(ruleParts[0], "-")
		if len(posRuleParts) != 2 {
			return 0, fmt.Errorf("the password character position rule specification \"%s\" is invalid", ruleParts[0]) // Invalid password character position rule part
		}
		lowerPos, err := strconv.Atoi(posRuleParts[0])
		if err != nil {
			return 0, err
		}
		upperPos, err := strconv.Atoi(posRuleParts[1])
		if err != nil {
			return 0, err
		}

		// Parse the character
		character := []rune(ruleParts[1])

		if (password[lowerPos-1] == character[0] && password[upperPos-1] != character[0]) || (password[lowerPos-1] != character[0] && password[upperPos-1] == character[0]) {
			validCount++ // Password meets specified rules
		}
	}

	return validCount, nil
}
