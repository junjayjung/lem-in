package tools

import (
	"sort"
)

// findPathsDFS explores all possible paths from 'start' to 'end' using DFS.
func findPathsDFS(start string, end string, links map[string][]string, visited map[string]bool, path []string, allPaths *[][]string) {
	if visited[start] || start == end {
		if start == end {
			*allPaths = append(*allPaths, append([]string(nil), path...))
		}
		return
	}
	visited[start] = true
	defer func() { visited[start] = false }()
	for _, next := range links[start] {
		findPathsDFS(next, end, links, visited, append(path, next), allPaths)
	}
}

// FindAllPaths initializes DFS search and sorts resulting paths by length.
func FindAllPaths(startRoom, endRoom string, links map[string][]string) [][]string {
	allPaths := make([][]string, 0)
	visited := make(map[string]bool)
	currentPath := []string{startRoom}
	findPathsDFS(startRoom, endRoom, links, visited, currentPath, &allPaths)

	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	return allPaths
}
