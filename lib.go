package adventOfCode2022

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------------
// Types

type RockPaperScissors int8

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

type RockPaperScissorsOutcome int8

const (
	Loose = 0
	Draw  = 3
	Win   = 6
)

type CrateMover interface {
	// MoveCrates moves the given number of crates from one stack to another.
	// Should be ok to modify stacks without directly: https://stackoverflow.com/a/39993797/738002
	MoveCrates(stacks [][]rune, num int, fromIdx int, toIdx int)
}

type CrateMover9000 struct {
}

func (p *CrateMover9000) MoveCrates(stacks [][]rune, num int, fromIdx int, toIdx int) {
	for i := 0; i < num; i++ {
		removed := stacks[fromIdx][len(stacks[fromIdx])-1]
		stacks[fromIdx] = stacks[fromIdx][0 : len(stacks[fromIdx])-1]
		stacks[toIdx] = append(stacks[toIdx], removed)
	}
}

type CrateMover9001 struct {
}

func (p *CrateMover9001) MoveCrates(stacks [][]rune, num int, fromIdx int, toIdx int) {
	removed := stacks[fromIdx][len(stacks[fromIdx])-num:]
	stacks[fromIdx] = stacks[fromIdx][0 : len(stacks[fromIdx])-num]
	stacks[toIdx] = append(stacks[toIdx], removed...)
}

// ---------------------------------------------------------------------------
// Functions

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func forLines(fileName string, action func(string)) {
	file, err := os.Open(fileName)
	check(err)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		action(scanner.Text())
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
	var current = 0
	var groups = make([]int, num)

	forLines(fileName, func(line string) {
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

func isRockPaperScissorsWin(you RockPaperScissors, opponent RockPaperScissors) bool {
	return you == Rock && opponent == Scissors ||
		you == Paper && opponent == Rock ||
		you == Scissors && opponent == Paper
}

func findRockPaperScissorsForOutcome(opponent RockPaperScissors, outcome RockPaperScissorsOutcome) RockPaperScissors {
	if outcome == Win {
		if opponent == Scissors {
			return Rock
		} else if opponent == Rock {
			return Paper
		} else {
			return Scissors
		}
	}
	if outcome == Loose {
		if opponent == Scissors {
			return Paper
		} else if opponent == Rock {
			return Scissors
		} else {
			return Rock
		}
	} else {
		return opponent
	}
}

func evaluateRockPaperScissors(you RockPaperScissors, opponent RockPaperScissors) RockPaperScissorsOutcome {
	if isRockPaperScissorsWin(you, opponent) {
		return Win
	} else if isRockPaperScissorsWin(opponent, you) {
		return Loose
	} else {
		return Draw
	}
}

func scoreRockPaperScissors(you RockPaperScissors, opponent RockPaperScissors) int {
	return int(evaluateRockPaperScissors(you, opponent)) + int(you)
}

func charToRockPaperScissors(input byte) RockPaperScissors {
	if input == 'A' || input == 'X' {
		return Rock
	} else if input == 'B' || input == 'Y' {
		return Paper
	} else {
		return Scissors
	}
}

func charToRockPaperScissorsOutcome(input byte) RockPaperScissorsOutcome {
	if input == 'X' {
		return Loose
	} else if input == 'Y' {
		return Draw
	} else if input == 'Z' {
		return Win
	} else {
		panic(fmt.Sprintf("Unexpected input %c", input))
	}
}

func readAndEvaluateRockPaperScissors(fileName string) int {
	score := 0

	forLines(fileName, func(line string) {
		plays := strings.Split(line, " ")
		score += scoreRockPaperScissors(charToRockPaperScissors(plays[1][0]), charToRockPaperScissors(plays[0][0]))
	})

	return score
}

func readAndEvaluateRockPaperScissorsOutcome(fileName string) int {
	score := 0

	forLines(fileName, func(valueAsString string) {
		plays := strings.Split(valueAsString, " ")
		play := findRockPaperScissorsForOutcome(
			charToRockPaperScissors(plays[0][0]),
			charToRockPaperScissorsOutcome(plays[1][0]))

		score += scoreRockPaperScissors(play, charToRockPaperScissors(plays[0][0]))
	})

	return score
}

func flattenUnique(input [][]string) string {
	m := make(map[string]string)
	for _, el := range input {
		m[el[0]] = el[0]
	}
	s := make([]string, len(m))
	for _, el := range m {
		s = append(s, el)
	}
	return strings.Join(s, "")
}

func findCommonChars(input []string) string {
	if len(input) < 1 {
		return ""
	}

	common := input[0]

	for i := 0; i < len(input)-1; i++ {
		r := regexp.MustCompile("[" + common + "]")
		common = flattenUnique(r.FindAllStringSubmatch(input[i+1], -1))
	}

	return common
}

func resolveRucksackPriority(fileName string) int {
	priority := 0

	forLines(fileName, func(line string) {
		var compA, compB = line[0 : len(line)/2], line[len(line)/2:]

		match := findCommonChars([]string{compA, compB})
		if len(match) != 1 {
			panic("Expected single byte match")
		}
		priority += calculatePriority(match[0])
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

func resolveRucksackChunkedPriority(fileName string, chunkSize int) int {
	priority := 0

	i := 0
	chunk := make([]string, chunkSize)

	forLines(fileName, func(line string) {
		defer func() {
			i++
		}()

		chunk[i%chunkSize] = line

		if i%chunkSize == chunkSize-1 {
			common := findCommonChars(chunk[:])
			if len(common) != 1 {
				panic(fmt.Sprintf("Incorrect common items found %s, %q\n", common, chunk))
			}
			priority += calculatePriority(common[0])
		}
	})

	return priority
}

func MustAtoi(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprintf("Cannot convert %s to int", input))
	}
	return output
}

func findCleaningAreaOverlap(fileName string, requireFullOverlap bool) int {
	overlapCount := 0

	r := regexp.MustCompile(`[-,]`)
	forLines(fileName, func(line string) {
		p := r.Split(line, -1)
		s1, e1, s2, e2 := MustAtoi(p[0]), MustAtoi(p[1]), MustAtoi(p[2]), MustAtoi(p[3])
		if requireFullOverlap {
			if s1 >= s2 && e1 <= e2 {
				overlapCount++
			} else if s2 >= s1 && e2 <= e1 {
				overlapCount++
			}
		} else {
			if s1 >= s2 && s1 <= e2 {
				overlapCount++
			} else if s2 >= s1 && s2 <= e1 {
				overlapCount++
			}
		}
	})

	return overlapCount
}

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

func rearrangeCrates(fileName string, crane CrateMover) string {
	var stacks [][]rune
	var stacksStr []string
	operationRegexp := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

	forLines(fileName, func(line string) {
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
