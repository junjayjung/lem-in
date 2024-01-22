package tools

import (
	"fmt"
	"strings"
)

func MoveAnts(matrix Matrix, dot Dot) {
	path := FindAllPaths(matrix)
	ants := initializeAnts(dot.NumAnts)
	turns := 0

	for !allAntsAtEnd(ants, matrix.EndRoom) {
		moveAntsAlongPath(ants, path)
		printAntsPositions(ants, turns)
		turns++
	}
}

func initializeAnts(numAnts int) []Ant {
	ants := make([]Ant, numAnts)
	for i := range ants {
		ants[i] = Ant{ID: i + 1, Current: ""}
	}
	return ants
}

func moveAntsAlongPath(ants []Ant, path [][]string) {
	for i := range ants {
		for j, room := range path {
			if ants[i].Current == room[j] && j < len(path)-1 {
				ants[i].Current = path[j+1]
				break
			}
		}
	}
}

func printAntsPositions(ants []Ant, turn int) {
	var movements []string
	for _, ant := range ants {
		if ant.Current != "" {
			movements = append(movements, fmt.Sprintf("L%d-%s", ant.ID, ant.Current))
		}
	}

	if len(movements) > 0 {
		fmt.Printf("Turn %d: %s\n", turn, strings.Join(movements, " "))
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
