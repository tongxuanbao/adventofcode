package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	i, j int
}

func getChar(pipeMap []string, pos position) byte {
	if pos.i < 0 || pos.i >= len(pipeMap) || pos.j < 0 || pos.j >= len(pipeMap[0]) {
		return '.'
	}
	return pipeMap[pos.i][pos.j]
}

func getStartingNeighbors(pipeMap []string, pos position) []position {
	top := position{pos.i + 1, pos.j}
	bottom := position{pos.i - 1, pos.j}
	left := position{pos.i, pos.j - 1}
	right := position{pos.i, pos.j + 1}

	neighbors := make([]position, 0, 4)

	if getChar(pipeMap, top) == '|' || getChar(pipeMap, top) == '7' || getChar(pipeMap, top) == 'F' {
		neighbors = append(neighbors, top)
	}
	if getChar(pipeMap, bottom) == '|' || getChar(pipeMap, bottom) == 'J' || getChar(pipeMap, bottom) == 'L' {
		neighbors = append(neighbors, bottom)
	}
	if getChar(pipeMap, left) == '-' || getChar(pipeMap, left) == 'F' || getChar(pipeMap, left) == 'L' {
		neighbors = append(neighbors, left)
	}
	if getChar(pipeMap, right) == '-' || getChar(pipeMap, right) == '7' || getChar(pipeMap, right) == 'J' {
		neighbors = append(neighbors, right)
	}

	return neighbors
}

func getNeighbors(pipeMap []string, pos position) []position {
	top := position{pos.i - 1, pos.j}
	bottom := position{pos.i + 1, pos.j}
	left := position{pos.i, pos.j - 1}
	right := position{pos.i, pos.j + 1}

	neighbors := make([]position, 0, 4)

	switch pipeMap[pos.i][pos.j] {
	case '|':
		neighbors = append(neighbors, top, bottom)
	case '-':
		neighbors = append(neighbors, left, right)
	case 'L':
		neighbors = append(neighbors, top, right)
	case 'J':
		neighbors = append(neighbors, top, left)
	case '7':
		neighbors = append(neighbors, left, bottom)
	case 'F':
		neighbors = append(neighbors, right, bottom)
	case 'S':
		neighbors = append(neighbors, getStartingNeighbors(pipeMap, pos)...)
	}

	realNeighbors := make([]position, 0, 4)

	for _, neighbor := range neighbors {
		if !(neighbor.i < 0 || neighbor.i >= len(pipeMap) || neighbor.j < 0 || neighbor.j > len(pipeMap[0])) {
			realNeighbors = append(realNeighbors, neighbor)
		}
	}

	return realNeighbors
}

func bfs(pipeMap []string, startingPoint position) int {
	distance := make([][]int, len(pipeMap))
	for i := range distance {
		distance[i] = make([]int, len(pipeMap[0]))
		for j := range distance[i] {
			distance[i][j] = 10e9
		}
	}
	distance[startingPoint.i][startingPoint.j] = 0

	visited := make(map[position]bool)
	stack := make([]position, 0)
	stack = append(stack, startingPoint)

	visited[startingPoint] = true

	result := 0

	for len(stack) > 0 {
		currentNode := stack[0]
		//fmt.Println(currentNode)
		//fmt.Printf("%c\n", pipeMap[currentNode.i][currentNode.j])
		//fmt.Println(getNeighbors(pipeMap, currentNode))
		//fmt.Println(distance[currentNode.i][currentNode.j])
		//fmt.Println()
		stack = stack[1:]

		if result < distance[currentNode.i][currentNode.j] {
			result = distance[currentNode.i][currentNode.j]
		}

		for _, neighbor := range getNeighbors(pipeMap, currentNode) {
			if visited[neighbor] || pipeMap[neighbor.i][neighbor.j] == '.' {
				continue
			}
			visited[neighbor] = true
			stack = append(stack, neighbor)
			if distance[neighbor.i][neighbor.j] > distance[currentNode.i][currentNode.j]+1 {
				distance[neighbor.i][neighbor.j] = distance[currentNode.i][currentNode.j] + 1
			}
		}
	}

	//for _, line := range distance {
	//	fmt.Println(line)
	//}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pipeMap := make([]string, 0)

	for scanner.Scan() {
		pipeMap = append(pipeMap, scanner.Text())
	}

	startingPoint := position{0, 0}
	for i, line := range pipeMap {
		for j, char := range line {
			if char == 'S' {
				startingPoint = position{i, j}
			}
		}
	}

	result := bfs(pipeMap, startingPoint)

	fmt.Println(result)
	gm
}
