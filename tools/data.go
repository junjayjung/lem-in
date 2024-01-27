package lem

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Data() (int, string, string, map[string][]string) {
	var numAnts int
	var start, end string
	links := make(map[string][]string)

	rawText, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	inputLines := strings.Split(string(rawText), "\n")
	numAnts, err = strconv.Atoi(inputLines[0])
	if err != nil || numAnts <= 0 {
		log.Fatal("ERROR: Invalid number of ants.")
	}

	for index, line := range inputLines {
		if line == "##start" || line == "##end" {
			roomInfo := strings.Split(inputLines[index+1], " ")
			if line == "##start" {
				start = roomInfo[0]
			} else {
				end = roomInfo[0]
			}
		} else if parts := strings.Fields(line); len(parts) > 1 {
			links[parts[0]] = []string{}
		} else if strings.Contains(line, "-") {
			linkInfo := strings.Split(line, "-")
			links[linkInfo[0]] = append(links[linkInfo[0]], linkInfo[1])
			links[linkInfo[1]] = append(links[linkInfo[1]], linkInfo[0])
		}
	}

	if start == "" || end == "" || len(links[start]) == 0 || len(links[end]) == 0 {
		log.Fatal("ERROR: invalid data format.")
	}
	return numAnts, start, end, links
}
