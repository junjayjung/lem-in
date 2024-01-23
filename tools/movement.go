package tools

import (
	"fmt"
	"sort"
)

type PathWithAnts struct {
	Path []string
	Ants []int
}

func MoveAnts(matrix Matrix, dot Dot) {
	allPaths := FindAllPaths(matrix.StartRoom, matrix.EndRoom, matrix.Edges)
	selectedPaths := DistributeAnts(SelectPaths(allPaths), dot.NumAnts)
	for _, path := range selectedPaths {
		fmt.Printf("Path: %v Ants: %v\n", path.Path, path.Ants)
	}

}
func SelectPaths(allPaths [][]string) [][]string {
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})
	return allPaths // This simplistic approach selects all paths; adjust as needed.
}

func DistributeAnts(paths [][]string, numAnts int) []PathWithAnts {
	distributions := make([]PathWithAnts, len(paths))
	for i, path := range paths {
		distributions[i] = PathWithAnts{Path: path}
	}
	antID := 1
	for antID <= numAnts {
		for i := range distributions {
			if antID > numAnts {
				break
			}
			distributions[i].Ants = append(distributions[i].Ants, antID)
			antID++
		}
	}
	return distributions
}
