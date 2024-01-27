package lem

import (
	"fmt"
	"sort"
)

func Move(numberOfAnts int, startRoom string, endRoom string, paths [][]string) {
	antLocations := make(map[int]string, numberOfAnts)
	for i := 1; i <= numberOfAnts; i++ {
		antLocations[i] = startRoom
	}

	for len(antLocations) > 0 {
		var ants []int
		for ant := range antLocations {
			ants = append(ants, ant)
		}
		sort.Ints(ants)

		var directLinkUsed bool
		for _, ant := range ants {
			if antLocations[ant] == endRoom {
				delete(antLocations, ant)
			} else {
				roomsAvailable := NextHeya(antLocations[ant], &directLinkUsed, startRoom, endRoom, antLocations, paths)
				if len(roomsAvailable) > 0 {
					antLocations[ant] = roomsAvailable[0]
					fmt.Printf("L%d-%s ", ant, roomsAvailable[0])
				}
			}
		}
		if len(antLocations) != 0 {
			fmt.Println()
		}
	}
}

func NextHeya(antLoc string, directLinkUsed *bool, startRoom string, endRoom string, antLocations map[int]string, paths [][]string) []string {
	var nextRooms []string
	for _, path := range paths {
		for index, content := range path {
			if content == antLoc {
				nextRooms = append(nextRooms, path[index+1])
			}
		}
	}

	var roomsAvailable []string
	if len(antLocations) < 5 && antLoc == startRoom && *directLinkUsed {
		return roomsAvailable
	}

	for _, room := range nextRooms {
		var occupied bool
		if room == endRoom && !*directLinkUsed {
			roomsAvailable = append(roomsAvailable, room)
			if antLoc == startRoom {
				*directLinkUsed = true
			}
		} else {
			for _, loc := range antLocations {
				if room == loc {
					occupied = true
					break
				}
			}
		}
		if !occupied {
			roomsAvailable = append(roomsAvailable, room)
		}
	}

	return roomsAvailable
}
