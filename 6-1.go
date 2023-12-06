package main

import (
	"fmt"
)

func main() {
	result := 1
	numberOfRace := 4
	var dispose string

	times := make([]int, numberOfRace)
	fmt.Scanf("%s", &dispose)
	for i := 0; i < numberOfRace; i++ {
		fmt.Scanf("%d", &times[i])
	}
	records := make([]int, numberOfRace)
	fmt.Scanf("%s", &dispose)
	for i := 0; i < numberOfRace; i++ {
		fmt.Scanf("%d", &records[i])
	}

	for i := 0; i < numberOfRace; i++ {
		timeBeatRecord := 0
		for j := 1; j <= times[i]; j++ {
			distance := j * (times[i] - j)
			if distance > records[i] {
				timeBeatRecord++
			}
		}
		result *= timeBeatRecord
	}

	fmt.Println(result)
}
