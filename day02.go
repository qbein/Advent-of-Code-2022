package adventOfCode2022

import (
	"fmt"
	"strings"
)

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

func EvaluateRockPaperScissors(you RockPaperScissors, opponent RockPaperScissors) RockPaperScissorsOutcome {
	if isRockPaperScissorsWin(you, opponent) {
		return Win
	} else if isRockPaperScissorsWin(opponent, you) {
		return Loose
	} else {
		return Draw
	}
}

func scoreRockPaperScissors(you RockPaperScissors, opponent RockPaperScissors) int {
	return int(EvaluateRockPaperScissors(you, opponent)) + int(you)
}

func CharToRockPaperScissors(input byte) RockPaperScissors {
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

func ReadAndEvaluateRockPaperScissors(fileName string) int {
	score := 0

	ForLinesIn(fileName, func(line string) {
		plays := strings.Split(line, " ")
		score += scoreRockPaperScissors(CharToRockPaperScissors(plays[1][0]), CharToRockPaperScissors(plays[0][0]))
	})

	return score
}

func ReadAndEvaluateRockPaperScissorsOutcome(fileName string) int {
	score := 0

	ForLinesIn(fileName, func(valueAsString string) {
		plays := strings.Split(valueAsString, " ")
		play := findRockPaperScissorsForOutcome(
			CharToRockPaperScissors(plays[0][0]),
			charToRockPaperScissorsOutcome(plays[1][0]))

		score += scoreRockPaperScissors(play, CharToRockPaperScissors(plays[0][0]))
	})

	return score
}
