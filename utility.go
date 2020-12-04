package advent2020

import (
	"bufio"
	"os"
	"strconv"
)

// LoadStringInputFile is a utility function that attempts to load a file line-by-line as a slice of strings
func LoadStringInputFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(file)
	out := []string{}
	for sc.Scan() {
		out = append(out, sc.Text())
	}

	return out, sc.Err()
}

// LoadIntInputFile is a utility function that attempts to load a file of integers into a slice.
func LoadIntInputFile(fileName string) ([]int, error) {
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
