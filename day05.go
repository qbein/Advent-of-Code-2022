package adventOfCode2022

import (
	"regexp"
	"strings"
)

func getStackCount(numericRow string) int {
	numericRow = strings.TrimSpace(numericRow)
	i := strings.LastIndex(numericRow, " ") + 1
	return MustAtoi(numericRow[i:])
}

func parseStacks(stacksStr []string) [][]rune {
	stackCnt := getStackCount(stacksStr[len(stacksStr)-1])
	stacks := make([][]rune, stackCnt)

	for row := len(stacksStr) - 2; row >= 0; row-- {
		for col := 0; col < stackCnt; col++ {
			rowStr := stacksStr[row]

			// Only works for single digit columns
			chrIdx := 1
			if col > 0 {
				chrIdx = col + 4 + (col-1)*3
			}

			if chrIdx < len(rowStr) {
				value := rowStr[chrIdx]
				if value != ' ' && value != 0 {
					stacks[col] = append(stacks[col], rune(value))
				}
			}
		}
	}

	return stacks
}

func ReArrangeCrates(fileName string, crane CrateMover) string {
	var stacks [][]rune
	var stacksStr []string
	operationRegexp := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

	ForLinesIn(fileName, func(line string) {
		// Parse stack configurations until first empty line
		if stacks == nil && len(strings.TrimSpace(line)) > 0 {
			stacksStr = append(stacksStr, line)
			return
		}

		// All stacks are read, parse into slices
		if stacks == nil {
			stacks = parseStacks(stacksStr)
		}

		m := operationRegexp.FindAllStringSubmatch(line, -1)

		// Skip lines without operations
		if m == nil {
			return
		}

		num, fromIdx, toIdx := MustAtoi(m[0][1]), MustAtoi(m[0][2])-1, MustAtoi(m[0][3])-1

		crane.MoveCrates(stacks, num, fromIdx, toIdx)
	})

	output := strings.Builder{}
	for _, s := range stacks {
		output.WriteRune(s[len(s)-1])
	}
	return output.String()
}
