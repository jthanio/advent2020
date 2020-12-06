package day5

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
)

// SolveDay5Part1 reads boarding passes to find the highest seat ID from the input. It takes advantage of the problem specification by interpreting the input as binary numbers.
func SolveDay5Part1(input []string) (int, error) {
	ids, err := convertAndSortIDs(input)
	if err != nil {
		return 0, err
	}

	if len(ids) == 0 {
		return 0, nil
	}

	return ids[len(ids)-1], nil
}

// SolveDay5Part2 reads boarding passes and determines the missing boarding pass ID.
func SolveDay5Part2(input []string) (int, error) {
	ids, err := convertAndSortIDs(input)
	if err != nil {
		return 0, err
	}

	lastID := ids[0]
	for _, i := range ids {
		if i-lastID > 1 {
			return lastID + 1, nil
		}
		lastID = i
	}

	return 0, errors.New("no gap in IDs found, perhaps you have no seat after all")
}

func convertAndSortIDs(input []string) ([]int, error) {
	if len(input) == 0 {
		return nil, nil
	}
	ids := make([]int, 0, len(input))
	for _, s := range input {
		s = regexp.MustCompile(`F|L`).ReplaceAllString(s, "0")
		s = regexp.MustCompile(`B|R`).ReplaceAllString(s, "1")
		id, err := strconv.ParseInt(s, 2, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, int(id))
	}
	sort.Ints(ids)
	return ids, nil
}
