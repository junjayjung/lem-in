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
		//for more readability
		if i == len(filelines)-1 {
			fmt.Println()
		}
	}
	matrix, dot := tools.ParseInputData(filelines)
	allPaths := tools.FindAllPaths(matrix.StartRoom, matrix.EndRoom, matrix.Edges)
	if len(allPaths) == 0 {
		tools.ExitWithError("Error: invalid data format\n", nil)
		return
	}
	tools.MoveAnts(matrix, dot)
}
