package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertToSlice(line string) []int {
	list := strings.Split(line, " ")
	result := make([]int, len(list))
	for i, v := range list {
		n, _ := strconv.Atoi(v)
		result[i] = n
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	parts := strings.Split(scanner.Text(), ": ")
	seeds := convertToSlice(parts[1])

	for transform := 0; transform < 7; transform++ {
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				continue
			}
			if line[len(line)-1] == ':' {
				break
			}
		}

		isTransform := make([]bool, len(seeds))
		for scanner.Scan() {
			line := scanner.Text()
			var destination, source, rang int
			n, err := fmt.Sscanf(line, "%d %d %d\n", &destination, &source, &rang)
			if n == 0 || err != nil {
				break
			}

			// fmt.Printf("%d %d %d\n", destination, source, rang)
			for i, v := range seeds {
				if v >= source && v < source+rang && isTransform[i] == false {
					seeds[i] = destination + v - source
					isTransform[i] = true
				}
			}
		}
		// fmt.Println(seeds)
		// fmt.Println()
	}

	result := seeds[0]
	for _, v := range seeds[1:] {
		result = min(result, v)
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
