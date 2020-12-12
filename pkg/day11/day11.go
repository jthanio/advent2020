package day11

const (
	floorRune      = '.'
	emptyRune      = 'L'
	occupiedRune   = '#'
	tolerancePart1 = 4
	tolerancePart2 = 5
)

// SolveDay11Part1 simulates seating rules and then counts the total number of occupied seats.
func SolveDay11Part1(input []string) (int, error) {
	prev := newGrid(input)
	next := createNextGridState(prev, tolerancePart1)

	for !gridStateEquals(prev, next) {
		prev = next
		next = createNextGridState(prev, tolerancePart1)
	}

	return numOccupied(next), nil
}

// SolveDay11Part2 is not yet implemented.
func SolveDay11Part2(input []string) (int, error) {
	return 0, nil
}

func newGrid(input []string) [][]rune {
	g := make([][]rune, len(input))
	for y, row := range input {
		g[y] = []rune(row)
	}
	return g
}

func gridStateEquals(s1, s2 [][]rune) bool {
	if s1 == nil {
		return s2 == nil // If they are both nil, they are equal
	}

	if s2 == nil {
		return s1 == nil
	}

	for y, row := range s1 {
		for x, r := range row {
			if r != s2[y][x] {
				return false
			}
		}
	}
	return true
}

func numOccupied(grid [][]rune) (num int) {
	for _, row := range grid {
		for _, r := range row {
			if r == occupiedRune {
				num++
			}
		}
	}
	return
}

func createNextGridState(grid [][]rune, tolerance int) [][]rune {
	next := make([][]rune, len(grid))
	for y, row := range grid {
		nextRow := make([]rune, len(row))
		for x, r := range row {
			nextRow[x] = getNextPosition(getAdjacentRunes(grid, 0, len(row)-1, 0, len(grid)-1, x, y), r, tolerance)
		}
		next[y] = nextRow
	}
	return next
}

func getAdjacentRunes(r [][]rune, xMin, xMax, yMin, yMax, x, y int) []rune {
	switch {
	// Top left corner
	case x-1 < xMin && y-1 < yMin:
		return []rune{
			r[y][x+1],
			r[y+1][x],
			r[y+1][x+1],
		}
	// Top right corner
	case x+1 > xMax && y-1 < yMin:
		return []rune{
			r[y][x-1],
			r[y+1][x-1],
			r[y+1][x],
		}
	// Bottom left corner
	case x-1 < xMin && y+1 > yMax:
		return []rune{
			r[y-1][x],
			r[y-1][x+1],
			r[y][x+1],
		}
	// Bottom right corner
	case x+1 > xMax && y+1 > yMax:
		return []rune{
			r[y-1][x-1],
			r[y-1][x],
			r[y][x-1],
		}
	// Left edge
	case x-1 < xMin && y-1 >= yMin && y+1 <= yMax:
		return []rune{
			r[y-1][x],
			r[y-1][x+1],
			r[y][x+1],
			r[y+1][x],
			r[y+1][x+1],
		}
	// Right edge
	case x+1 > xMax && y-1 >= yMin && y+1 <= yMax:
		return []rune{
			r[y-1][x-1],
			r[y-1][x],
			r[y][x-1],
			r[y+1][x-1],
			r[y+1][x],
		}
	// Top edge
	case x-1 >= xMin && x+1 <= xMax && y-1 < yMin:
		return []rune{
			r[y][x-1],
			r[y][x+1],
			r[y+1][x-1],
			r[y+1][x],
			r[y+1][x+1],
		}
	// Bottom edge
	case x-1 >= xMin && x+1 <= xMax && y+1 > yMax:
		return []rune{
			r[y-1][x-1],
			r[y-1][x],
			r[y-1][x+1],
			r[y][x-1],
			r[y][x+1],
		}
	}
	return []rune{
		r[y-1][x-1],
		r[y-1][x],
		r[y-1][x+1],
		r[y][x-1],
		r[y][x+1],
		r[y+1][x-1],
		r[y+1][x],
		r[y+1][x+1],
	}
}

func getNextPosition(adjacent []rune, current rune, crowdTolerance int) rune {
	switch current {
	case occupiedRune:
		if countRune(adjacent, occupiedRune) >= crowdTolerance {
			return emptyRune // Too crowded, I'm leaving
		}
	case emptyRune:
		if countRune(adjacent, occupiedRune) == 0 {
			return occupiedRune // Perfect spot
		}
	}
	return current // This is a floor tile, or something else...
}

func countRune(in []rune, r rune) (count int) {
	for _, i := range in {
		if i == r {
			count++
		}
	}
	return
}
