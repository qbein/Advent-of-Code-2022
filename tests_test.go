package adventOfCode2022

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 1, part 1

func TestFindMinIndex(t *testing.T) {
	var minIndex = findMinIndex([]int{12, 56, 1, 6})
	assert.Equal(t, 2, minIndex)
}

func Test0101Example(t *testing.T) {
	maxCalories := findMaxGroupsInIntList("input/day01_example", 1)[0]
	assert.Equal(t, 24000, maxCalories)
}

func Test0101(t *testing.T) {
	maxCalories := findMaxGroupsInIntList("input/day01", 1)[0]
	assert.Equal(t, 69883, maxCalories)
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 1, part 2

func Test0102Example(t *testing.T) {
	maxCalories := sum(findMaxGroupsInIntList("input/day01_example", 3))
	assert.Equal(t, 45000, maxCalories)
}

func Test0102(t *testing.T) {
	maxCalories := sum(findMaxGroupsInIntList("input/day01", 3))
	assert.Equal(t, 207576, maxCalories)
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 2, part 1

func TestEvaluateRockPaperScissors(t *testing.T) {
	assert.EqualValues(t, Win, evaluateRockPaperScissors(Rock, Scissors))
	assert.EqualValues(t, Win, evaluateRockPaperScissors(Scissors, Paper))
	assert.EqualValues(t, Win, evaluateRockPaperScissors(Paper, Rock))

	assert.EqualValues(t, Loose, evaluateRockPaperScissors(Scissors, Rock))
	assert.EqualValues(t, Loose, evaluateRockPaperScissors(Paper, Scissors))
	assert.EqualValues(t, Loose, evaluateRockPaperScissors(Rock, Paper))

	assert.EqualValues(t, Draw, evaluateRockPaperScissors(Rock, Rock))
	assert.EqualValues(t, Draw, evaluateRockPaperScissors(Scissors, Scissors))
	assert.EqualValues(t, Draw, evaluateRockPaperScissors(Paper, Paper))
}

func Test0201Example(t *testing.T) {
	score := scoreRockPaperScissors(charToRockPaperScissors('A'), charToRockPaperScissors('Y'))
	score += scoreRockPaperScissors(charToRockPaperScissors('B'), charToRockPaperScissors('X'))
	score += scoreRockPaperScissors(charToRockPaperScissors('C'), charToRockPaperScissors('Z'))

	assert.Equal(t, 15, score)
}

func Test0201(t *testing.T) {
	assert.Equal(t, 11841, readAndEvaluateRockPaperScissors("input/day02"))
}
