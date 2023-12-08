package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    readFile, err := os.Open("input")

    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    result := 0

    listofstring := []string {"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", 
                            "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}
    valueofstring := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
                            1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
    
    for fileScanner.Scan() {
        filelines := fileScanner.Text()
        var first, last int
        firstIndex := len(filelines)
        lastIndex := -1

        for i := 0; i < len(filelines); i++ {
            for j := 0; j < len(listofstring); j++ {
                lastSubstrIndex := strings.LastIndex(filelines, listofstring[j])
                if  lastSubstrIndex >= lastIndex && lastSubstrIndex != -1{
                        lastIndex = lastSubstrIndex
                        last = valueofstring[j]
                }

                firstSubstrIndex := strings.Index(filelines, listofstring[j])
                if firstSubstrIndex <= firstIndex && firstSubstrIndex != -1 {
                   firstIndex = firstSubstrIndex 
                    first = valueofstring[j]
                }
            }
        }


        result += first*10+last
    }
    fmt.Println(result)
}
