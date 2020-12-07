package day7

import (
	"fmt"
	"regexp"
	"strconv"
)

const bagColorShinyGold = "shiny gold" // The target color

// SolveDay7Part1 parses the bag rules and counts the number of bag colors that can hold a shiny gold bag.
func SolveDay7Part1(input []string) (int, error) {
	graph := bagGraph{
		bags:  make(map[string]struct{}),
		edges: make(map[string][]edge),
	}
	for _, i := range input {
		color, err := parseBagColor(i)
		if err != nil {
			return 0, err
		}
		graph.addBag(color) // Add the bag identified by its color

		contents, err := parseBagContents(i)
		if err != nil {
			return 0, err
		}

		for bag, count := range contents {
			graph.addEdge(color, bag, count)
		}
	}

	return graph.countContainingBags(bagColorShinyGold), nil
}

type bagGraph struct {
	bags  map[string]struct{}
	edges map[string][]edge
}

func (b *bagGraph) addEdge(parent, bag string, count int) {
	b.edges[parent] = append(b.edges[parent], edge{color: bag, count: count})
}

func (b *bagGraph) addBag(bag string) {
	b.bags[bag] = struct{}{}
}

// countContainingBags will traverse the graph to find all bags that can contain the requested color
func (b *bagGraph) countContainingBags(target string) int {
	var count int
	for color := range b.bags {
		if color == target {
			continue // Ignore the target, since there are no cyclic bags
		}
		if b.dfs(target, color) {
			count++
		}
	}
	return count
}

func (b *bagGraph) dfs(target, color string) bool {
	return b.dfsHelper(make(map[string]struct{}), target, color, 0) > 0
}

func (b *bagGraph) dfsHelper(visited map[string]struct{}, target, color string, count int) int {
	if _, ok := visited[color]; ok {
		return count
	}
	visited[color] = struct{}{} // Mark this color as visited
	if color == target {
		count++
	}
	for _, edge := range b.edges[color] {
		count = b.dfsHelper(visited, target, edge.color, count)
	}
	return count
}

type edge struct {
	color string
	count int
}

// SolveDay7Part2 is not yet implemented
func SolveDay7Part2(input []string) (int, error) {
	return 0, nil
}

func parseBagColor(line string) (string, error) {
	parentRegex := regexp.MustCompile(`(\w* \w*) bags contain`) // The parent bag
	parentMatch := parentRegex.FindAllStringSubmatch(line, -1)
	if len(parentMatch) == 0 {
		return "", fmt.Errorf("could not parse \"%s\", invalid bag syntax", line)
	}
	return parentMatch[0][1], nil
}

func parseBagContents(line string) (map[string]int, error) {
	var contents = make(map[string]int)
	contentsRegex := regexp.MustCompile(`(\d) (\w* \w*) bag`)       // Any contained bags
	noContentsRegex := regexp.MustCompile(`contain no other bags.`) // Contains no bags
	contentsMatch := contentsRegex.FindAllStringSubmatch(line, -1)

	if noContentsRegex.MatchString(line) {
		return nil, nil // Exit early if no contents
	}

	if len(contentsMatch) == 0 {
		return nil, fmt.Errorf("could not parse \"%s\", invalid contained bag syntax", line)
	}

	for _, m := range contentsMatch {
		n, err := strconv.Atoi(m[1]) // Parse the number of bags
		if err != nil {
			return nil, err
		}
		contents[m[2]] = n // Store the count for the bag of the given key
	}
	return contents, nil
}
