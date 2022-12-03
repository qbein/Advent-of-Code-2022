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

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		action(strings.TrimSpace(scanner.Text()))
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
	var current int = 0
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

func resolveRucksackPriority(fileName string) int {
	priority := 0

	forLines(fileName, func(line string) {
		var compA, compB = line[0 : len(line)/2], line[len(line)/2:]

		r := regexp.MustCompile("[" + compA + "]")
		match := r.FindStringSubmatch(compB)

		if len(match) != 1 || len(match[0]) != 1 {
			panic("Expected single byte match")
		}

		commonItem := match[0][0]

		subtract := 0
		if int(commonItem) < int('a') {
			subtract = int('A') - 27
		} else {
			subtract = int('a') - 1
		}

		itemPriority := int(match[0][0]) - subtract
		priority += itemPriority
	})

	return priority
}
