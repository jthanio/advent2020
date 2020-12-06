package day4

// SolveDay4Part1 finds the number of valid passports in the given batch input
func SolveDay4Part1(input []string) int {
	var count int
	p := &Passport{}
	for _, s := range input {
		if s == "" {
			// Passport input ended, can attempt to validate
			if p.HasAllRequiredFields() {
				count++
			}
			p = &Passport{} // Start a new passport
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
	p := &Passport{}
	for _, s := range input {
		if s == "" {
			// Passport input ended, can attempt to validate
			if p.HasAllValidFields() {
				count++
			}
			p = &Passport{} // Start a new passport
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
