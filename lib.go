package adventOfCode2022

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findHigestNumberInGroupedList(fileName string) int {
	file, err := os.Open(fileName)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var max, current int = 0, 0

	for scanner.Scan() {
		valueAsString := strings.TrimSpace(scanner.Text())

		if valueAsString == "" {
			if current > max {
				max = current
			}
			current = 0
			continue
		}

		value, err := strconv.Atoi(valueAsString)
		check(err)

		current += value
	}

	if current > max {
		max = current
	}

	return max
}
