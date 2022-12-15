package adventOfCode2022

import "strconv"

func updateGroups(groups *[]int, current int) {
	var minIndex = FindMinValueIndex(*groups)
	if current > (*groups)[minIndex] {
		(*groups)[minIndex] = current
	}
}

func FindMaxGroupsInIntList(fileName string, num int) []int {
	var current = 0
	var groups = make([]int, num)

	ForLinesIn(fileName, func(line string) {
		if line == "" {
			updateGroups(&groups, current)
			current = 0
			return
		}

		value, err := strconv.Atoi(line)
		check(err)

		current += value
	})

	updateGroups(&groups, current)

	return groups
}
