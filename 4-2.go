package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertToMap(line string) map[string]bool {
	list := strings.Split(line, " ")
	result := map[string]bool{}
	for _, v := range list {
		if v != "" {
			result[v] = true
		}
	}
	return result
}

func countMatches(line string) int {
	matches := 0
	parts := strings.Split(line, ": ")
	parts2 := strings.Split(parts[1], " | ")
	winningNumbers := convertToMap(parts2[0])
	owningNumbers := strings.Split(parts2[1], " ")

	for _, v := range owningNumbers {
		if winningNumbers[v] {
			matches++
		}
	}
	return matches
}

func main() {
	numberOfCards := 198
	result := 0
	currentCardIndex := 0

	// Setup accumulate array
	accumulate := make([]int, 200)
	accumulate[1] = 1

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		matches := countMatches(line)

		// Calculate the number of current card
		currentCardIndex++
		accumulate[currentCardIndex] += accumulate[currentCardIndex-1]

		// Add to result
		result += accumulate[currentCardIndex]

		// Mark the number of cards will be added
		from := currentCardIndex + 1
		to := min(currentCardIndex+matches+1, numberOfCards+1)
		accumulate[from] += accumulate[currentCardIndex]
		accumulate[to] -= accumulate[currentCardIndex]
	}

	fmt.Println(result)
}
