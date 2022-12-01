package adventOfCode2022

import "testing"

func TestDecember1stExample(t *testing.T) {
	maxCalories := findHigestNumberInGroupedList("input/day01_example")
	if maxCalories != 24000 {
		t.Fatalf(`Expected the elf with most calories to carry %v, got %v`, 12, maxCalories)
	}
}

func TestDecember1st(t *testing.T) {
	maxCalories := findHigestNumberInGroupedList("input/day01")
	if maxCalories != 69883 {
		t.Fatalf(`Expected the elf with most calories to carry %v, got %v`, 12, maxCalories)
	}
}
