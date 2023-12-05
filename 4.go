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

func main4() {
	n := 198
	current := 0
	result := 0

	accumulate := make([]int, 200)
	accumulate[1] = 1

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		matches := 0
		parts := strings.Split(scanner.Text(), ": ")
		parts2 := strings.Split(parts[1], " | ")
		winningNumbers := convertToMap(parts2[0])
		owningNumbers := strings.Split(parts2[1], " ")

		for _, v := range owningNumbers {
			if winningNumbers[v] {
				matches++
			}
		}

		current++

		accumulate[current] += accumulate[current-1]
		result += accumulate[current]

		to := min(current+matches, n)

		accumulate[current+1] += accumulate[current]
		accumulate[to+1] -= accumulate[current]

	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
