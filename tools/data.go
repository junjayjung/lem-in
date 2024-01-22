package tools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	StartRoom string
	EndRoom   string
	Edges     map[string][]string
	Vertices  []string
}

type Ant struct {
	ID      int
	Current string
}

type Dot struct {
	NumAnts int
}

func Read(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: Unable to open the file - %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	var fileLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	if len(fileLines) == 0 {
		fmt.Println("Error: The file is empty.")
		os.Exit(1)
	}

	return fileLines
}

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
