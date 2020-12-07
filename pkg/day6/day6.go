package day6

// SolveDay6Part1 reads group inputs and finds the total sum of all unique customs declarations in each group.
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

// SolveDay6Part2 counts all customs declarations that every member in a group has made instead of just the total sum.
func SolveDay6Part2(input []string) int {
	var total int
	var sum int
	group := map[rune]int{}
	for _, s := range input {
		if s == "" {
			sum += countMax(group, total)
			group = map[rune]int{} // Start a new group
			total = 0
		} else {
			for _, r := range []rune(s) {
				group[r]++ // Increment the entry by one
			}
			total++ // Increment the count of participants (lines)
		}
	}
	sum += countMax(group, total)
	return sum
}

func countMax(input map[rune]int, total int) int {
	var sum int
	for _, i := range input {
		if i == total {
			sum++
		}
	}
	return sum
}
