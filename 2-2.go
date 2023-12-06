package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getListOfGames(line string) []string {
	parts := strings.Split(line, ": ")
	parts2 := strings.Split(parts[1], "; ")
	return parts2
}

func getListOfCubes(game string) []string {
	records := strings.Split(game, ", ")
	return records
}

func getRecord(cube string) (int, string) {
	var numberOfCube int
	var color string
	fmt.Sscanf(cube, "%d %s", &numberOfCube, &color)
	return numberOfCube, color
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	gameId := 0
	result := 0
	for scanner.Scan() {
		records := make(map[string]int)
		records["red"] = 0
		records["green"] = 0
		records["blue"] = 0

		gameId++

		line := scanner.Text()
		for _, game := range getListOfGames(line) {
			for _, cube := range getListOfCubes(game) {
				var numberOfCube int
				var color string
				fmt.Sscanf(cube, "%d %s", &numberOfCube, &color)
				if records[color] < numberOfCube {
					records[color] = numberOfCube
				}
			}
		}
		result += records["red"] * records["green"] * records["blue"]
	}

	fmt.Println(result)
}
