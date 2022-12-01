package adventOfCode2022

import (
	"testing"
)

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 1, part 1

func TestFindMinIndex(t *testing.T) {
	var minIndex = findMinIndex([]int{12, 56, 1, 6})
	if minIndex != 2 {
		t.Fatalf(`Expected index 2 to be returned as minimum, got %v`, minIndex)
	}
}

func Test0101Example(t *testing.T) {
	maxCalories := findMaxGroupsInIntList("input/day01_example", 1)[0]
	if maxCalories != 24000 {
		t.Fatalf(`Expected the elf with most calories to carry %v, got %v`, 12, maxCalories)
	}
}

func Test0101(t *testing.T) {
	maxCalories := findMaxGroupsInIntList("input/day01", 1)[0]
	if maxCalories != 69883 {
		t.Fatalf(`Expected the elf with most calories to carry %v, got %v`, 12, maxCalories)
	}
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 1, part 2

func Test0102Example(t *testing.T) {
	maxCalories := sum(findMaxGroupsInIntList("input/day01_example", 3))
	if maxCalories != 45000 {
		t.Fatalf(`Expected the top three elfs with most calories to carry %v, got %v`, 12, maxCalories)
	}
}

func Test0102(t *testing.T) {
	maxCalories := sum(findMaxGroupsInIntList("input/day01", 3))
	if maxCalories != 207576 {
		t.Fatalf(`Expected the top three elfs with most calories to carry %v, got %v`, 12, maxCalories)
	}
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 2, part 1
