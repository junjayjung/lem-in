package lem

import (
	"sort"
)

func Michi(start string, end string, links map[string][]string) [][]string {
	var paths [][]string
	var currentPath []string
	visited := make(map[string]bool)
	calculateRoutes(&currentPath, start, end, links, visited, &paths)
	Nagasa(paths)
	if len(paths) > 10 {
		return paths[:10]
	}
	return paths
}

func calculateRoutes(path *[]string, currentRoom string, endRoom string, links map[string][]string, visited map[string]bool, paths *[][]string) {
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
				calculateRoutes(path, nextRoom, endRoom, links, visited, paths)
			}
		}
	}
}

func Nagasa(paths [][]string) {
	sort.Slice(paths, func(i, j int) bool { return len(paths[i]) < len(paths[j]) })
}
