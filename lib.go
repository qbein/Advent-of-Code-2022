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

func findMinIndex(slice []int) int {
	minIndex := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func updateGroups(groups *[]int, current int) {
	var minIndex = findMinIndex(*groups)
	if current > (*groups)[minIndex] {
		(*groups)[minIndex] = current
	}
}

func sum(values []int) int {
	var sum = 0
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return sum
}

func findMaxGroupsInIntList(fileName string, num int) []int {
	file, err := os.Open(fileName)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var current int = 0
	var groups = make([]int, num)

	for scanner.Scan() {
		valueAsString := strings.TrimSpace(scanner.Text())

		if valueAsString == "" {
			updateGroups(&groups, current)
			current = 0
			continue
		}

		value, err := strconv.Atoi(valueAsString)
		check(err)

		current += value
	}

	updateGroups(&groups, current)

	return groups
}
