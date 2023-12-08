package main

import (
	"fmt"
)

func main() {
	var instruction string
	fmt.Scanf("%s\n\n", &instruction)

	tree := make(map[string](map[byte]string))

	for {
		var node, left, right string
		_, err := fmt.Scanf("%3s = (%3s, %3s)\n", &node, &left, &right)
		if err != nil {
			break
		}
		tree[node] = map[byte]string{
			'L': left,
			'R': right,
		}
	}
	result := 0
	currentNode := "AAA"

	for currentNode != "ZZZ" {
		index := result % len(instruction)
		if instruction[index] == 'L' {
			currentNode = tree[currentNode]['L']
		} else {
			currentNode = tree[currentNode]['R']
		}
		result++
	}

	fmt.Println(result)
}
