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

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 2, part 2

func Test0202(t *testing.T) {
	assert.Equal(t, 13022, readAndEvaluateRockPaperScissorsOutcome("input/day02"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 3, part 1

func Test0301Example(t *testing.T) {
	assert.Equal(t, 157, resolveRucksackPriority("input/day03_example"))
}

func Test0301(t *testing.T) {
	assert.Equal(t, 8088, resolveRucksackPriority("input/day03"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 3, part 2

func Test0302Example(t *testing.T) {
	assert.Equal(t, 70, resolveRucksackChunkedPriority("input/day03_example", 3))
}

func Test0302(t *testing.T) {
	assert.Equal(t, 2522, resolveRucksackChunkedPriority("input/day03", 3))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 4, part 1

func Test0401Example(t *testing.T) {
	assert.Equal(t, 2, findCleaningAreaOverlap("input/day04_example", true))
}

func Test0401(t *testing.T) {
	assert.Equal(t, 462, findCleaningAreaOverlap("input/day04", true))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 4, part 2

func Test0402Example(t *testing.T) {
	assert.Equal(t, 4, findCleaningAreaOverlap("input/day04_example", false))
}

func Test0402(t *testing.T) {
	assert.Equal(t, 835, findCleaningAreaOverlap("input/day04", false))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 5, part 1

func Test0501Example(t *testing.T) {
	assert.Equal(t, "CMZ", rearrangeCrates("input/day05_example", new(CrateMover9000)))
}

func Test0501(t *testing.T) {
	assert.Equal(t, "CVCWCRTVQ", rearrangeCrates("input/day05", new(CrateMover9000)))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 5, part 2

func Test0502Example(t *testing.T) {
	assert.Equal(t, "MCD", rearrangeCrates("input/day05_example", new(CrateMover9001)))
}

func Test0502(t *testing.T) {
	assert.Equal(t, "CNSCZWLVT", rearrangeCrates("input/day05", new(CrateMover9001)))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 6, part 1

func Test0601Example(t *testing.T) {
	assert.Equal(t, 7, findStartOfPacket("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4))
	assert.Equal(t, 5, findStartOfPacket("bvwbjplbgvbhsrlpgdmjqwftvncz", 4))
	assert.Equal(t, 6, findStartOfPacket("nppdvjthqldpwncqszvftbrmjlhg", 4))
	assert.Equal(t, 10, findStartOfPacket("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4))
	assert.Equal(t, 11, findStartOfPacket("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4))
}

func Test0601(t *testing.T) {
	assert.Equal(t, 1275, findStartOfPacketFromFile("input/day06", 4))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 6, part 2

func Test0602Example(t *testing.T) {
	assert.Equal(t, 19, findStartOfPacket("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
	assert.Equal(t, 23, findStartOfPacket("bvwbjplbgvbhsrlpgdmjqwftvncz", 14))
	assert.Equal(t, 23, findStartOfPacket("nppdvjthqldpwncqszvftbrmjlhg", 14))
	assert.Equal(t, 29, findStartOfPacket("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14))
	assert.Equal(t, 26, findStartOfPacket("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14))
}

func Test0602(t *testing.T) {
	assert.Equal(t, 3605, findStartOfPacketFromFile("input/day06", 14))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 7, part 1

func Test0701Example(t *testing.T) {
	assert.Equal(t, 95437, dirSizes("input/day07_example"))
}

func Test0701(t *testing.T) {
	assert.Equal(t, 1477771, dirSizes("input/day07"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 7, part 2

func Test0702Example(t *testing.T) {
	assert.Equal(t, 24933642, findDirToDelete(70000000, 30000000, "input/day07_example"))
}

func Test0702(t *testing.T) {
	assert.Equal(t, 3579501, findDirToDelete(70000000, 30000000, "input/day07"))
}
