package tools

import (
	"fmt"
	"sort"
	"strings"
)

func MoveAnts(matrix Matrix, dot Dot) {
	allPaths := FindAllPaths(matrix.StartRoom, matrix.EndRoom, matrix.Edges)
	selectedPaths := DistributeAnts(SelectPaths(allPaths), dot.NumAnts)

	// Initialize ants and assign them to paths.
	ants := initializeAnts(dot.NumAnts, selectedPaths)
	roomOccupancy := make(map[string]int)

	// Enhanced debugging: Print paths and initial room occupancy
	fmt.Println("Selected Paths and Initial Assignments:")
	for _, path := range selectedPaths {
		fmt.Printf("Path: %v, Ants: %v\n", path.Path, path.Ants)
	}

	for !allAntsAtEnd(ants, matrix.EndRoom) {
		var movements []string
		// Reset room occupancy except for the end room
		for room := range roomOccupancy {
			if room != matrix.EndRoom {
				roomOccupancy[room] = 0
			}
		}
		for i, ant := range ants {
			if ant.CurrentRoom == matrix.EndRoom {
				continue // Skip ants already at the end
			}
			nextRoom := getNextRoom(ant, roomOccupancy, matrix.EndRoom)
			if nextRoom != "" {
				moveAnt(&ants[i], nextRoom, roomOccupancy)
				movements = append(movements, fmt.Sprintf("L%d-%s", ant.ID, nextRoom))
			}
		}
		if len(movements) == 0 && !allAntsAtEnd(ants, matrix.EndRoom) {
			fmt.Println("Detected a deadlock situation. Review path assignments and occupancy logic.")
			break
		}
		fmt.Println(strings.Join(movements, " "))
	}
}

// Initialize ants and assign them to paths.
func initializeAnts(numAnts int, paths []PathWithAnts) []Ant {
	ants := make([]Ant, numAnts)
	for i := range ants {
		ants[i].ID = i + 1
		for _, path := range paths {
			if containsAntID(path.Ants, ants[i].ID) {
				ants[i].Path = path.Path
				ants[i].CurrentRoom = path.Path[0]
				break
			}
		}
	}
	return ants
}

func containsAntID(antIDs []int, id int) bool {
	for _, antID := range antIDs {
		if antID == id {
			return true
		}
	}
	return false
}

// Move an ant to the next room and update room occupancy.
// Move an ant to the next room and update room occupancy.
func moveAnt(ant *Ant, nextRoom string, occupancy map[string]int) {
	if ant.CurrentRoom != "" {
		occupancy[ant.CurrentRoom] = 0 // Vacate current room
	}
	ant.CurrentRoom, ant.CurrentIndex = nextRoom, ant.CurrentIndex+1
	occupancy[nextRoom] = ant.ID // Occupy new room
}

// Determine if all ants have reached the end.
func allAntsAtEnd(ants []Ant, endRoom string) bool {
	for _, ant := range ants {
		if ant.CurrentRoom != endRoom {
			return false
		}
	}
	return true
}

// Select the next room for an ant to move to.
// Select the next room for an ant to move to.
func getNextRoom(ant Ant, occupancy map[string]int, endRoom string) string {
	if ant.CurrentIndex >= len(ant.Path)-1 {
		return ""
	}
	nextRoom := ant.Path[ant.CurrentIndex+1]
	if nextRoom == endRoom || occupancy[nextRoom] == 0 {
		return nextRoom
	}
	return ""
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
