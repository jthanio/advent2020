package day6

// SolveDay6Part1 reads group inputs and finds the total sum of all customs declarations.
func SolveDay6Part1(input []string) int {
	var sum int
	group := map[rune]struct{}{}
	for _, s := range input {
		if s == "" {
			// Group input ended
			sum += len(group)
			group = map[rune]struct{}{} // Start a new group
		} else {
			for _, r := range []rune(s) {
				group[r] = struct{}{}
			}
		}
	}

	sum += len(group) // Must do last group

	return sum
}

// SolveDay6Part2 is not yet implemented.
func SolveDay6Part2(input []string) int {
	return 0
}
