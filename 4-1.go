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
	result := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		matches := countMatches(line)

		if matches > 0 {
			result += 1 << (matches - 1) // 2^(matches - 1)
		}
	}

	fmt.Println(result)
}
