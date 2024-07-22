package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	newMap := make(map[string]int, 0)

	for points, letters := range in {
		for _, letter := range letters {
			newMap[strings.ToLower(letter)] = points
		}
	}

	return newMap
}
