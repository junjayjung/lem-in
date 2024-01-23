package tools

import (
	"sort"
)

// findPathsDFS explores all possible paths from 'current' to 'endRoom' using DFS.
func findPathsDFS(current string, endRoom string, links map[string][]string, visited map[string]bool, path []string, allPaths *[][]string) {
	if visited[current] {
		return
	}
	if current == endRoom {
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		*allPaths = append(*allPaths, pathCopy)
		return
	}
	visited[current] = true
	for _, next := range links[current] {
		findPathsDFS(next, endRoom, links, visited, append(path, next), allPaths)
	}
	visited[current] = false
}

// FindAllPaths initializes the DFS search and sorts the resulting paths by length.
func FindAllPaths(startRoom string, endRoom string, links map[string][]string) [][]string {
	allPaths := make([][]string, 0)
	findPathsDFS(startRoom, endRoom, links, make(map[string]bool), []string{startRoom}, &allPaths)

	// Sort paths by length.
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	return allPaths
}
