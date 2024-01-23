package tools

import (
	"fmt"
	"sort"
	"strings"
)

func MoveAnts(matrix Matrix, dot Dot) {
	allPaths := FindAllPaths(matrix.StartRoom, matrix.EndRoom, matrix.Edges)
	selectedPaths := DistributeAnts(SelectPaths(allPaths), dot.NumAnts)
	ants := initializeAnts(dot.NumAnts, selectedPaths)
	roomOccupancy := make(map[string]int)

	turn := 0
	for !allAntsAtEnd(ants, matrix.EndRoom) {
		var movements []string
		for i := range ants {
			if ants[i].CurrentRoom == matrix.EndRoom {
				continue // Skip ants that have reached the end
			}

			nextRoom := getNextRoom(ants[i], roomOccupancy, matrix.EndRoom)
			if nextRoom != "" && (roomOccupancy[nextRoom] == 0 || nextRoom == matrix.EndRoom) {
				if ants[i].CurrentRoom != matrix.StartRoom { // Leave start room occupancy unchanged
					roomOccupancy[ants[i].CurrentRoom] = 0
				}
				ants[i].CurrentRoom = nextRoom
				roomOccupancy[nextRoom] = ants[i].ID
				ants[i].CurrentIndex++
				movements = append(movements, fmt.Sprintf("L%d-%s", ants[i].ID, nextRoom))
			}
		}

		if len(movements) > 0 {
			fmt.Println(strings.Join(movements, " "))
		}
		turn++
	}
}

func getNextRoom(ant Ant, occupancy map[string]int, endRoom string) string {
	if ant.CurrentIndex+1 < len(ant.Path) {
		nextRoom := ant.Path[ant.CurrentIndex+1]
		// Check occupancy only if nextRoom is not the end room
		if nextRoom != endRoom && occupancy[nextRoom] != 0 {
			return ""
		}
		return nextRoom
	}
	return ant.CurrentRoom
}

func initializeAnts(numAnts int, paths []PathWithAnts) []Ant {
	ants := make([]Ant, numAnts)
	for i := range ants {
		ants[i].ID = i + 1
		// Assign paths to ants
		for _, path := range paths {
			if contains(path.Ants, ants[i].ID) {
				ants[i].Path = path.Path
				break
			}
		}
	}
	return ants
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func allAntsAtEnd(ants []Ant, endRoom string) bool {
	for _, ant := range ants {
		if ant.CurrentRoom != endRoom {
			return false
		}
	}
	return true
}
func SelectPaths(allPaths [][]string) [][]string {
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})
	return allPaths
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
