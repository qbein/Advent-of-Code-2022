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

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 8, part 1

func Test0801Example(t *testing.T) {
	assert.Equal(t, 21, countHiddenTrees("input/day08_example"))
}

func Test0801(t *testing.T) {
	assert.Equal(t, 1719, countHiddenTrees("input/day08"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 8, part 2

func Test0802Example(t *testing.T) {
	x, y, score := findMostScenicTree("input/day08_example")
	assert.Equal(t, 2, x)
	assert.Equal(t, 3, y)
	assert.Equal(t, 8, score)
}

func Test0802(t *testing.T) {
	x, y, score := findMostScenicTree("input/day08")
	assert.Equal(t, 46, x)
	assert.Equal(t, 85, y)
	assert.Equal(t, 590824, score)
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 9, part 1

func Test0901Example(t *testing.T) {
	assert.Equal(t, 13, countRopeTailPositions("input/day09_example", 2))
}

func Test0901(t *testing.T) {
	assert.Equal(t, 6284, countRopeTailPositions("input/day09", 2))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 9, part 2

func Test0902Example1(t *testing.T) {
	assert.Equal(t, 1, countRopeTailPositions("input/day09_example", 10))
}

func Test0902Example2(t *testing.T) {
	assert.Equal(t, 36, countRopeTailPositions("input/day0902_example", 10))
}

func Test0902(t *testing.T) {
	assert.Equal(t, 2661, countRopeTailPositions("input/day09", 10))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 10, part 1

func Test1001Example(t *testing.T) {
	assert.Equal(t, 13140, findSignalStrength("input/day10_example"))
}

func Test1001(t *testing.T) {
	assert.Equal(t, 14920, findSignalStrength("input/day10"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 10, part 2

func Test1002Example(t *testing.T) {
	renderSignalToCrt("input/day10_example")
}

func Test1002(t *testing.T) {
	renderSignalToCrt("input/day10")
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 11, part 1

func Test1101Example(t *testing.T) {
	assert.Equal(t, 10605, executeDay11("input/day11_example", 20, 1))
}

func Test1101(t *testing.T) {
	assert.Equal(t, 111210, executeDay11("input/day11", 20, 1))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 11, part 2

func Test1102Example(t *testing.T) {
	assert.Equal(t, 2713310158, executeDay11("input/day11_example", 10000, 2))
}

func Test1102(t *testing.T) {
	assert.Equal(t, 15447387620, executeDay11("input/day11", 10000, 2))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 12, part 1

func Test1201Example(t *testing.T) {
	assert.Equal(t, 31, findShortestRoute("input/day12_example"))
}

func Test1201(t *testing.T) {
	assert.Equal(t, 456, findShortestRoute("input/day12"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 12, part 2

func Test1202Example(t *testing.T) {
	assert.Equal(t, 29, findShortestRouteFromAny("input/day12_example", 'a'))
}

/*
func Test1202(t *testing.T) {
	assert.Equal(t, 454, findShortestRouteFromAny("input/day12", 'a'))
}
*/

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 13, part 1

func Test1301ParsePacket(t *testing.T) {
	slice := parsePacketString("[1,1,3,1,1]")

	assert.Equal(t, float64(1), slice[0])
	assert.Equal(t, float64(1), slice[1])
	assert.Equal(t, float64(3), slice[2])
	assert.Equal(t, float64(1), slice[3])
	assert.Equal(t, float64(1), slice[4])
}

func Test1301ParsePacket2(t *testing.T) {
	slice := parsePacketString("[[8,7,6]]")

	switch v := slice[0].(type) {
	case []any:
		assert.Equal(t, float64(8), v[0])
		assert.Equal(t, float64(7), v[1])
		assert.Equal(t, float64(6), v[2])
	}
}

func Test1301Example(t *testing.T) {
	assert.Equal(t, 13, countOrderedIndices("input/day13_example"))
}

func Test1301(t *testing.T) {
	assert.Equal(t, 5196, countOrderedIndices("input/day13"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 13, part 2

func Test1302Example(t *testing.T) {
	assert.Equal(t, 140, findDecoderKey("input/day13_example"))
}

func Test1302(t *testing.T) {
	assert.Equal(t, 22134, findDecoderKey("input/day13"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 14, part 1

func Test1401Example(t *testing.T) {
	assert.Equal(t, 24, SandBoxFromFile("input/day14_example", -1).solve(false))
}

func Test1401(t *testing.T) {
	assert.Equal(t, 892, SandBoxFromFile("input/day14", -1).solve(false))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 14, part 2

func Test1402Example(t *testing.T) {
	assert.Equal(t, 93, SandBoxFromFile("input/day14_example", 2).solve(false))
}

func Test1402(t *testing.T) {
	assert.Equal(t, 27155, SandBoxFromFile("input/day14", 2).solve(false))
}
