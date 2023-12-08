package main

import (
	"fmt"
)

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	var instruction string
	fmt.Scanf("%s\n\n", &instruction)

	tree := make(map[string](map[byte]string))
	startingNodes := make([]string, 0)

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
		if node[2] == 'A' {
			startingNodes = append(startingNodes, node)
		}
	}

	resultsList := make([]int, 0, len(startingNodes))

	for _, startingNode := range startingNodes {
		currentNode := startingNode
		stepsToEnd := 0
		for currentNode[2] != 'Z' {
			// fmt.Println(currentNode)
			index := stepsToEnd % len(instruction)
			if instruction[index] == 'L' {
				currentNode = tree[currentNode]['L']
			} else {
				currentNode = tree[currentNode]['R']
			}
			stepsToEnd++
		}
		resultsList = append(resultsList, stepsToEnd)
		// fmt.Println("end loop\n")

	}

	result := 1
	for _, v := range resultsList {
		result = lcm(result, v)
	}

	fmt.Println(result)
}
