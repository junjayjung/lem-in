package tools

import "math"

func Dijkstra(matrix Matrix) map[string]int {
	distances := make(map[string]int)
	visited := make(map[string]bool)

	for _, v := range matrix.Vertices {
		distances[v] = math.MaxInt32
	}

	distances[matrix.StartRoom] = 0

	for i := 0; i < len(matrix.Vertices)-1; i++ {
		u := minDistance(distances, visited)
		visited[u] = true

		for _, v := range matrix.Edges[u] {
			if !visited[v] && distances[u]+1 < distances[v] {
				distances[v] = distances[u] + 1
			}
		}
	}

	return distances
}

func minDistance(distances map[string]int, visited map[string]bool) string {
	min := math.MaxInt32
	var minNode string

	for k, v := range distances {
		if !visited[k] && v <= min {
			min = v
			minNode = k
		}
	}

	return minNode
}

func ShortestPath(matrix Matrix) []string {
	distances := Dijkstra(matrix)
	path := []string{matrix.EndRoom}
	currentRoom := matrix.EndRoom

	for currentRoom != matrix.StartRoom {
		for _, room := range matrix.Edges[currentRoom] {
			if distances[room] == distances[currentRoom]-1 {
				path = append([]string{room}, path...)
				currentRoom = room
				break
			}
		}
	}

	return path
}
