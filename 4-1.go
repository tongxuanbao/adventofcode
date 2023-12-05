package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func init() {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	f, _ := os.Create("file")
	f.Write([]byte(input))
	f.Close()
	os.Stdin, _ = os.Open("file")
}

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
