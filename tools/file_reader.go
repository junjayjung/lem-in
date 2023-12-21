package main

import (
	"bufio"
	"fmt"
	"os"
)

func read(filePath string) []string {
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
