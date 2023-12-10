package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertLine(line string) (int, []int) {
	numbers := strings.Split(line, " ")
	n := len(numbers)
	list := make([]int, n)
	for i, number := range numbers {
		fmt.Sscanf(number, "%d", &list[i])
	}
	return n, list
}

func isAllZero(list []int) bool {
	for _, v := range list {
		if v != 0 {
			return false
		}
	}
	return true
}

func nextLine(line []int) []int {
	result := make([]int, len(line)-1)

	for i := 0; i < len(line)-1; i++ {
		result[i] = line[i+1] - line[i]
	}

	return result
}

func calculateLine(line []int) int {
	currentLine := line
	result := line[len(line)-1]
	if isAllZero(currentLine) {
		return result
	}
	return result + calculateLine(nextLine(currentLine))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		_, list := convertLine(line)
		result += calculateLine(list)
	}
	fmt.Println(result)
}
