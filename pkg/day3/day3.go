package day3

// SolveDay3Part1 finds the number of tree collisions in for the given input course.
func SolveDay3Part1(input []string) int {
	const xSlope = 3
	const ySlope = 1

	return simulateToboggan(input, xSlope, ySlope)
}

// SolveDay3Part2 finds the number of tree collisions in for the given input course for each of the slopes and then multiplies the result together.
func SolveDay3Part2(input []string) int {
	slopes := []struct {
		xSlope int
		ySlope int
	}{
		{xSlope: 1, ySlope: 1},
		{xSlope: 3, ySlope: 1},
		{xSlope: 5, ySlope: 1},
		{xSlope: 7, ySlope: 1},
		{xSlope: 1, ySlope: 2},
	}

	var collisionProduct = 1
	for _, s := range slopes {
		collisionProduct *= simulateToboggan(input, s.xSlope, s.ySlope) // The solution is only concerned with the product of the collisions
	}

	return collisionProduct
}

// simulateToboggan counts the number of tree collisions in the input course based on the given slope.
func simulateToboggan(input []string, xSlope, ySlope int) int {
	var curXPos, count int
	for i := 0; i < len(input); i += ySlope {
		line := []rune(input[i]) // Treat the current line as a slice
		// Check if the char at the current pos is a tree (#)
		if line[curXPos] == '#' {
			count++
		}

		// Move to the next position laterally
		curXPos += xSlope
		if curXPos >= len(line) {
			curXPos -= len(line) // Wrap around back to the beginning of the line
		}
	}

	return count
}
