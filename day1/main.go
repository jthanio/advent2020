package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputFile = "input.txt"

func main() {
	input, err := loadInputFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The answer to the Day 1 puzzle: %d\n", Find2020EntriesAndMultiply(input))
}

// Find2020EntriesAndMultiply solves the Day 1 puzzle by finding two entries within the input slice and returning their product.
func Find2020EntriesAndMultiply(input []int) int {
	if len(input) <= 1 {
		fmt.Println("not enough entries in input data")
		return 0
	}

	// Initial naive approach
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				return input[i] * input[j]
			}
		}
	}

	return 0
}

// loadInputFile is a utility function that attempts to load a file of integers into a slice.
func loadInputFile(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(file)
	out := []int{}
	for sc.Scan() {
		i, err := strconv.Atoi(sc.Text()) // Parse the line
		if err != nil {
			return nil, err
		}
		out = append(out, i) // Add the parsed line to the output
	}

	return out, sc.Err()
}
