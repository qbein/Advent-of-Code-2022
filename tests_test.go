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
