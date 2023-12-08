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

	// Tree creation
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

	// Each starting node has a different number of steps to the end
	resultsList := make([]int, 0, len(startingNodes))
	for _, startingNode := range startingNodes {
		currentNode := startingNode
		stepsToEnd := 0
		for currentNode[2] != 'Z' {
			index := stepsToEnd % len(instruction)
			side := instruction[index]
			currentNode = tree[currentNode][side]
			stepsToEnd++
		}
		resultsList = append(resultsList, stepsToEnd)
	}

	// Find the LCM of all the steps
	result := 1
	for _, v := range resultsList {
		result = lcm(result, v)
	}

	fmt.Println(result)
}
