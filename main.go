package main

import (
	"fmt"
	lem "lem/tools"
	"os"
)

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
