package adventOfCode2022

import "fmt"

func CalculateRucksackPriority(fileName string) int {
	priority := 0

	ForLinesIn(fileName, func(line string) {
		var compA, compB = line[0 : len(line)/2], line[len(line)/2:]

		match := FindCommonChars([]string{compA, compB})
		if len(match) != 1 {
			panic("Expected single byte match")
		}
		priority += calculatePriority(match[0])
	})

	return priority
}

func CalculateRucksackChunkedPriority(fileName string, chunkSize int) int {
	priority := 0

	i := 0
	chunk := make([]string, chunkSize)

	ForLinesIn(fileName, func(line string) {
		defer func() {
			i++
		}()

		chunk[i%chunkSize] = line

		if i%chunkSize == chunkSize-1 {
			common := FindCommonChars(chunk[:])
			if len(common) != 1 {
				panic(fmt.Sprintf("Incorrect common items found %s, %q\n", common, chunk))
			}
			priority += calculatePriority(common[0])
		}
	})

	return priority
}

func calculatePriority(input byte) int {
	subtract := 0
	if int(input) < int('a') {
		subtract = int('A') - 27
	} else {
		subtract = int('a') - 1
	}
	return int(input) - subtract
}
