package tools

func FindAllPaths(matrix Matrix) [][]string {
	var paths [][]string
	visited := make(map[string]bool)
	dfsFindAllPaths(matrix.StartRoom, matrix.EndRoom, matrix, []string{}, &paths, visited)
	return paths
}

func dfsFindAllPaths(current, end string, matrix Matrix, path []string, paths *[][]string, visited map[string]bool) {
	path = append(path, current)
	visited[current] = true

	if current == end {
		newPath := make([]string, len(path))
		copy(newPath, path)
		*paths = append(*paths, newPath)
	} else {
		for _, next := range matrix.Edges[current] {
			if !visited[next] {
				dfsFindAllPaths(next, end, matrix, path, paths, visited)
			}
		}
	}

	visited[current] = false  // Backtrack
	path = path[:len(path)-1] // Backtrack
}
