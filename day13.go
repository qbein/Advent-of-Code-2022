package adventOfCode2022

import (
	"encoding/json"
	"strings"
)

func countOrderedIndices(fileName string) int {
	count := 0

	var left string
	var right string

	i := 0
	packetIndex := 0

	forLines(fileName, func(line string) {
		defer func() {
			i++
		}()

		line = strings.TrimSpace(line)

		if i%3 == 0 {
			left = line
		} else if i%3 == 1 {
			right = line

			packetIndex++
			if Compare(parseString(left), parseString(right)) < 0 {
				count += packetIndex
			}
		}
	})

	return count
}

// Compare the two sides. Returns negative numbers if left side
// is smaller (in order), 0 if sides are equal (in order), and
//positive numbers if right side is larger (out of order).
func Compare(left any, right any) int {
	_, okL := left.(float64)
	_, okR := right.(float64)

	// If both are numerics, compare and return
	if okL && okR {
		return int(left.(float64)) - int(right.(float64))
	}

	// If one is numeric, convert to array and compare
	if okL {
		left = []any{left}
	}
	if okR {
		right = []any{right}
	}

	leftSlice := left.([]any)
	rightSlice := right.([]any)

	// If one side ran out of items, check order
	if len(leftSlice) == 0 || len(rightSlice) == 0 {
		return len(leftSlice) - len(rightSlice)
	}

	// Both are arrays check order of first element
	result := Compare(leftSlice[0], rightSlice[0])

	// While sides are equal continue to compare
	if result == 0 {
		nextLeft := leftSlice[1:]
		nextRight := rightSlice[1:]

		if len(nextLeft) == 0 || len(nextRight) == 0 {
			return len(nextLeft) - len(nextRight)
		}

		return Compare(nextLeft, nextRight)
	}
	return result
}

func parseString(input string) []any {
	var data []any
	err := json.Unmarshal([]byte(input), &data)
	check(err)

	return data
}
