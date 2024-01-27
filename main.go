package main

import (
	"fmt"
	lem "lem/tools"
	"os"
)

// This program (funcion) checks command-line arguments, processes data, computes paths, and moves ants based on the input. It also handles errors if the number of arguments is not 2.
func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Error! Check the input.")
		os.Exit(1)
	}
	ants, start, end, links := lem.Data()
	paths := lem.Michi(start, end, links)
	lem.Move(ants, start, end, paths)
}
