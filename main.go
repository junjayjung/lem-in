package main

import (
	"fmt"
	"lem-in/tools"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No input file provided.")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fileLines := tools.Read(filePath)
	matrix, dot := tools.ParseInputData(fileLines)

	fmt.Println("Ants:", dot.NumAnts)
	fmt.Println("Rooms:", matrix.Vertices)
	fmt.Println("Edges:", matrix.Edges)

	tools.MoveAnts(matrix, dot)
}
