package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInputData(data []string) (Matrix, Dot) {
	var matrix Matrix
	var dot Dot
	matrix.Edges = make(map[string][]string)

	numAnts, err := strconv.Atoi(data[0])
	if err != nil {
		exitWithError("Error: Failed to convert the number of ants to an integer", err)
	}
	validateNumAnts(numAnts)
	dot.NumAnts = numAnts

	for i := 1; i < len(data); i++ {
		line := data[i]
		if strings.HasPrefix(line, "##start") {
			matrix.StartRoom = extractVertex(data[i+1])
		} else if strings.HasPrefix(line, "##end") {
			matrix.EndRoom = extractVertex(data[i+1])
		} else if strings.Contains(line, "-") {
			addConnection(&matrix, line)
		} else if strings.Contains(line, " ") {
			addRoom(&matrix, line)
		}
	}

	return matrix, dot
}

func exitWithError(message string, err error) {
	fmt.Printf("%s\n", message)
	os.Exit(1)
}

func validateNumAnts(numAnts int) {
	if numAnts < 1 {
		exitWithError("Error: No ants specified\n", nil)
	}
}

func extractVertex(data string) string {
	fields := strings.Fields(data)
	return fields[0]
}

func addConnection(matrix *Matrix, connection string) {
	endpoints := strings.Split(connection, "-")
	from, to := endpoints[0], endpoints[1]

	if from == "" || to == "" {
		exitWithError("Error: Invalid connection format", nil)
	}

	matrix.Edges[from] = append(matrix.Edges[from], to)
	matrix.Edges[to] = append(matrix.Edges[to], from)
}

func addRoom(matrix *Matrix, roomData string) {
	fields := strings.Fields(roomData)
	roomName := fields[0]

	if strings.HasPrefix(roomName, "L") || strings.HasPrefix(roomName, "#") {
		exitWithError("Error: Invalid room name", nil)
	}

	matrix.Vertices = append(matrix.Vertices, roomName)
}
