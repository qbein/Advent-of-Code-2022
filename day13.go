package adventOfCode2022

import (
	"encoding/json"
	"sort"
	"strings"
)

func FindDecoderKey(fileName string) int {
	packets := readPackets(fileName)
	packets = append(packets, "[[2]]")
	packets = append(packets, "[[6]]")

	sort.SliceStable(packets, func(i int, j int) bool {
		return Compare(ParsePacketString(packets[i]), ParsePacketString(packets[j])) < 0
	})

	key := 1

	for i, v := range packets {
		if v == "[[2]]" || v == "[[6]]" {
			key *= i + 1
		}
	}

	return key
}

func CountOrderedIndices(fileName string) int {
	count := 0
	packetIndex := 0

	packets := readPackets(fileName)

	for i := 0; i < len(packets)-1; i += 2 {
		left := packets[i]
		right := packets[i+1]

		packetIndex++
		if Compare(ParsePacketString(left), ParsePacketString(right)) < 0 {
			count += packetIndex
		}
	}

	return count
}

func readPackets(fileName string) []string {
	packets := make([]string, 0)

	ForLinesIn(fileName, func(line string) {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			packets = append(packets, strings.TrimSpace(line))
		}
	})

	return packets
}

// Compare the two sides. Returns negative numbers if left side
// is smaller (in order), 0 if sides are equal (in order), and
// positive numbers if right side is larger (out of order).
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

func ParsePacketString(input string) []any {
	var data []any
	err := json.Unmarshal([]byte(input), &data)
	check(err)

	return data
}
