package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	from int
	to   int
}

func convertToSliceOfPair(line string) []pair {
	list := strings.Split(line, " ")
	result := make([]pair, len(list)/2)
	cur := 0
	for i := 0; i < len(list); i += 2 {
		result[cur].from, _ = strconv.Atoi(list[i])
		result[cur].to, _ = strconv.Atoi(list[i+1])
		result[cur].to = result[cur].from + result[cur].to - 1
		cur++
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	parts := strings.Split(scanner.Text(), ": ")
	seeds := convertToSliceOfPair(parts[1])

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

		transformed := make([]bool, len(seeds))
		transformedSeeds := make([]pair, 0)
		for scanner.Scan() {
			line := scanner.Text()
			var destination, source, rang int
			n, err := fmt.Sscanf(line, "%d %d %d\n", &destination, &source, &rang)
			if n == 0 || err != nil {
				break
			}

			// Transform seeds
			for i := 0; i < len(seeds); i++ {
				seed := seeds[i]
				if transformed[i] {
					continue
				}
				if !(seed.from <= source+rang-1 && seed.to >= source) {
					continue
				}

				if seed.from < source && seed.to >= source {
					// Remove seed
					transformed[i] = true

					// Add first half of seed
					seeds = append(seeds, pair{seed.from, source - 1})
					transformed = append(transformed, false)

					// Transform second half
					transformedSeeds = append(transformedSeeds, pair{destination, seed.to + (destination - source)})
				} else if seed.from <= source+rang-1 && seed.to > source+rang-1 {
					// Remove seed
					transformed[i] = true

					// Add second half of seed
					seeds = append(seeds, pair{source + rang, seed.to})
					transformed = append(transformed, false)

					// Transform first half
					transformedSeeds = append(transformedSeeds, pair{destination + seed.from - source, destination + rang - 1})
				} else if seed.from < source && seed.to > source+rang-1 {
					// Remove seed
					transformed[i] = true

					// Add first half of seed
					seeds = append(seeds, pair{seed.from, source - 1})
					transformed = append(transformed, false)

					// Add second half of seed
					seeds = append(seeds, pair{source + rang, seed.to})
					transformed = append(transformed, false)

					// Transform middle
					transformedSeeds = append(transformedSeeds, pair{destination, destination + rang - 1})
				} else if seed.from >= source && seed.to <= source+rang-1 {
					// Remove seed
					transformed[i] = true

					// Transform seed
					transformedSeeds = append(transformedSeeds, pair{destination + seed.from - source, destination + seed.to - source})
				}

				test := make([]pair, 0)
				for i, seed := range seeds {
					if !transformed[i] {
						test = append(test, seed)
					}
				}
			}
		}

		for i, seed := range seeds {
			if !transformed[i] {
				transformedSeeds = append(transformedSeeds, seed)
			}
		}

		seeds = transformedSeeds
	}

	result := seeds[0].from

	for _, v := range seeds {
		if v.from < result {
			result = v.from
		}
	}

	fmt.Println(result)
}
