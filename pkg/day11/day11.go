package day11

// Now, you just need to model the people who will be arriving shortly. Fortunately, people are entirely predictable and always follow a simple set of rules.
// All decisions are based on the number of occupied seats adjacent to a given seat (one of the eight positions immediately up, down, left, right, or diagonal from the seat).
// The following rules are applied to every seat simultaneously:

// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
// Otherwise, the seat's state does not change.
// Floor (.) never changes; seats don't move, and nobody sits on the floor.

const (
	floorRune    = '.'
	emptyRune    = 'L'
	occupiedRune = '#'
)

// SolveDay11Part1 simulates seating rules and then counts the total number of occupied seats.
func SolveDay11Part1(input []string) (int, error) {
	prev := newGrid(input)
	next := prev.createNextGridState()

	for !gridEquals(prev, next) {
		prev = next
		next = prev.createNextGridState()
	}

	return next.numOccupied(), nil
}

// SolveDay11Part2 is not yet implemented.
func SolveDay11Part2(input []string) (int, error) {
	return 0, nil
}

type grid [][]rune

func newGrid(input []string) grid {
	g := make(grid, len(input))
	for y, row := range input {
		g[y] = []rune(row)
	}
	return g
}

func gridEquals(g1, g2 grid) bool {
	if g1 == nil {
		return g2 == nil // If they are both nil, they are equal
	}

	if g2 == nil {
		return g1 == nil
	}

	for y, row := range g1 {
		for x, r := range row {
			if r != g2[y][x] {
				return false
			}
		}
	}
	return true
}

func (g grid) numOccupied() (num int) {
	for y := range g {
		for x := range g[y] {
			if g[y][x] == occupiedRune {
				num++
			}
		}
	}
	return
}

func (g grid) createNextGridState() grid {
	next := make(grid, len(g))
	for y, row := range g {
		nextRow := make([]rune, len(row))
		for x, r := range row {
			nextRow[x] = getNextPosition(g.getAdjacentRunes(len(row)-1, x, y), r)
		}
		next[y] = nextRow
	}
	return next
}

func (g grid) getAdjacentRunes(xMax, x, y int) []rune {
	switch {
	// Top left corner
	case x-1 < 0 && y-1 < 0:
		return []rune{
			g[y][x+1],
			g[y+1][x],
			g[y+1][x+1],
		}
	// Top right corner
	case x+1 > xMax && y-1 < 0:
		return []rune{
			g[y][x-1],
			g[y+1][x-1],
			g[y+1][x],
		}
	// Bottom left corner
	case x-1 < 0 && y+1 >= len(g):
		return []rune{
			g[y-1][x],
			g[y-1][x+1],
			g[y][x+1],
		}
	// Bottom right corner
	case x+1 > xMax && y+1 >= len(g):
		return []rune{
			g[y-1][x-1],
			g[y-1][x],
			g[y][x-1],
		}
	// Left edge
	case x-1 < 0 && y-1 >= 0 && y+1 < len(g):
		return []rune{
			g[y-1][x],
			g[y-1][x+1],
			g[y][x+1],
			g[y+1][x],
			g[y+1][x+1],
		}
	// Right edge
	case x+1 > xMax && y-1 >= 0 && y+1 < len(g):
		return []rune{
			g[y-1][x-1],
			g[y-1][x],
			g[y][x-1],
			g[y+1][x-1],
			g[y+1][x],
		}
	// Top edge
	case y-1 < 0 && x-1 >= 0 && x+1 <= xMax:
		return []rune{
			g[y][x-1],
			g[y][x+1],
			g[y+1][x-1],
			g[y+1][x],
			g[y+1][x+1],
		}
	// Bottom edge
	case y+1 >= len(g) && x-1 >= 0 && x+1 <= xMax:
		return []rune{
			g[y-1][x-1],
			g[y-1][x],
			g[y-1][x+1],
			g[y][x-1],
			g[y][x+1],
		}
	}
	return []rune{
		g[y-1][x-1],
		g[y-1][x],
		g[y-1][x+1],
		g[y][x-1],
		g[y][x+1],
		g[y+1][x-1],
		g[y+1][x],
		g[y+1][x+1],
	}
}

func getNextPosition(adjacent []rune, current rune) rune {
	switch current {
	case occupiedRune:
		if countRune(adjacent, occupiedRune) >= 4 {
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
