package tools

import (
	"fmt"
	"strings"
)

func MoveAnts(matrix Matrix, dot Dot) {
	path := ShortestPath(matrix)
	if len(path) == 0 {
		fmt.Println("No path found from start to end")
		return
	}

	ants := initializeAnts(dot.NumAnts)
	for _, ant := range ants {
		ant.Current = matrix.StartRoom
	}

	for !allAntsAtEnd(ants, matrix.EndRoom) {
		moveAntsAlongPath(ants, path, matrix)
		printAntsPositions(ants, matrix)
	}
}

func initializeAnts(numAnts int) []Ant {
	ants := make([]Ant, numAnts)
	for i := range ants {
		ants[i] = Ant{ID: i + 1}
	}
	return ants
}

func moveAntsAlongPath(ants []Ant, path []string, matrix Matrix) {
	for i := range ants {
		if ants[i].Current == matrix.EndRoom {
			continue
		}

		currentIndex := findRoomIndex(path, ants[i].Current)
		if currentIndex < len(path)-1 {
			nextRoom := path[currentIndex+1]
			if !isRoomOccupied(ants, nextRoom) || nextRoom == matrix.EndRoom {
				ants[i].Current = nextRoom
			}
		}
	}
}

func findRoomIndex(path []string, room string) int {
	for i, r := range path {
		if r == room {
			return i
		}
	}
	return -1
}

func isRoomOccupied(ants []Ant, room string) bool {
	for _, ant := range ants {
		if ant.Current == room {
			return true
		}
	}
	return false
}

func printAntsPositions(ants []Ant, matrix Matrix) {
	var movements []string
	for _, ant := range ants {
		if ant.Current != matrix.StartRoom && ant.Current != matrix.EndRoom {
			movements = append(movements, fmt.Sprintf("L%d-%s", ant.ID, ant.Current))
		}
	}

	if len(movements) > 0 {
		fmt.Println(strings.Join(movements, " "))
	}
}

func allAntsAtEnd(ants []Ant, endRoom string) bool {
	for _, ant := range ants {
		if ant.Current != endRoom {
			return false
		}
	}
	return true
}
