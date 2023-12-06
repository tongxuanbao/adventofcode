package main

import (
	"fmt"
)

func numberOfDigit(number int) int {
	digit := 0
	for number > 0 {
		number /= 10
		digit++
	}
	return digit
}

func power(number int, digit int) int {
	result := 1
	for i := 0; i < digit; i++ {
		result *= number
	}
	return result
}

func arrayToInt(array []int) int {
	result := 0
	base := 1
	for i := len(array) - 1; i >= 0; i-- {
		result += array[i] * base
		base *= power(10, numberOfDigit(array[i]))
	}
	return result
}

func main() {
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
	time := arrayToInt(times)
	record := arrayToInt(records)

	result := 0
	for i := 1; i <= time; i++ {
		distance := i * (time - i)
		if distance > record {
			result++
		}
	}

	fmt.Println(result)
}
