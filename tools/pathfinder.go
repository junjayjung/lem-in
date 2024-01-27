package lem

import (
	"sort"
)

// The Michi algorithm calculates routes from a start to an end using a map of links, and returns the top 10 paths if there are more than 10, otherwise it returns all paths.
// Michi （道；みち） means 'path' or 'way' in Japanese.
func Michi(start string, end string, links map[string][]string) [][]string {
	var paths [][]string
	var currentPath []string
	visited := make(map[string]bool)
	Keisan(&currentPath, start, end, links, visited, &paths)
	Nagasa(paths)
	if len(paths) > 10 {
		return paths[:10]
	}
	return paths
}

// The Keisan algorithm defines a recursive function Keisan to find all paths between currentRoom and endRoom in a graph, using backtracking and storing the paths in a slice.
// Keisan (計算；けいさん) means 'calculate' or 'compute' in Japanese.
func Keisan(path *[]string, currentRoom string, endRoom string, links map[string][]string, visited map[string]bool, paths *[][]string) {
	*path = append(*path, currentRoom)
	visited[currentRoom] = true
	defer func() {
		*path = (*path)[:len(*path)-1]
		visited[currentRoom] = false
	}()

	if currentRoom == endRoom {
		newPath := make([]string, len(*path))
		copy(newPath, *path)
		*paths = append(*paths, newPath)
	} else {
		for _, nextRoom := range links[currentRoom] {
			if !visited[nextRoom] {
				Keisan(path, nextRoom, endRoom, links, visited, paths)
			}
		}
	}
}

// The Nagasa algorithm sorts the paths in a slice based on the length of each path.
// Nagasa (長さ；ながさ) means 'length' or 'length of' in Japanese.
func Nagasa(paths [][]string) {
	sort.Slice(paths, func(i, j int) bool { return len(paths[i]) < len(paths[j]) })
}
