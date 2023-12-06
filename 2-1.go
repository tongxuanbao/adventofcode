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

func checkGame(game string) bool {
	records := strings.Split(game, ", ")
	for _, record := range records {
		var numberOfCube int
		var color string
		fmt.Sscanf(record, "%d %s", &numberOfCube, &color)
		if color == "red" && numberOfCube > 12 {
			return false
		}
		if color == "green" && numberOfCube > 13 {
			return false
		}
		if color == "blue" && numberOfCube > 14 {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	gameId := 0
	result := 0
	for scanner.Scan() {
		gameId++
		line := scanner.Text()
		isPossible := true
		for _, game := range getListOfGames(line) {
			isPossible = isPossible && checkGame(game)
		}
		if isPossible {
			result += gameId
		}
	}
	fmt.Println(result)
}
