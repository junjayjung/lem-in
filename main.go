package main

import (
	"fmt"
	"lem-in/tools"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}
	filelines := tools.Read(os.Args[1])
	for i := range filelines {
		fmt.Println(filelines[i])
	}
	matrix, dot := tools.ParseInputData(filelines)
	allPaths := tools.FindAllPaths(matrix.StartRoom, matrix.EndRoom, matrix.Edges)
	if len(allPaths) == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	tools.MoveAnts(matrix, dot)
}
