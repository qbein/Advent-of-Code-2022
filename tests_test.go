package adventOfCode2022

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 1, part 1

func TestFindMinIndex(t *testing.T) {
	var minIndex = FindMinValueIndex([]int{12, 56, 1, 6})
	assert.Equal(t, 2, minIndex)
}

func Test0101Example(t *testing.T) {
	maxCalories := FindMaxGroupsInIntList("input/day01_example", 1)[0]
	assert.Equal(t, 24000, maxCalories)
}

func Test0101(t *testing.T) {
	maxCalories := FindMaxGroupsInIntList("input/day01", 1)[0]
	assert.Equal(t, 69883, maxCalories)
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 1, part 2

func Test0102Example(t *testing.T) {
	maxCalories := SumIntValues(FindMaxGroupsInIntList("input/day01_example", 3))
	assert.Equal(t, 45000, maxCalories)
}

func Test0102(t *testing.T) {
	maxCalories := SumIntValues(FindMaxGroupsInIntList("input/day01", 3))
	assert.Equal(t, 207576, maxCalories)
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 2, part 1

func TestEvaluateRockPaperScissors(t *testing.T) {
	assert.EqualValues(t, Win, EvaluateRockPaperScissors(Rock, Scissors))
	assert.EqualValues(t, Win, EvaluateRockPaperScissors(Scissors, Paper))
	assert.EqualValues(t, Win, EvaluateRockPaperScissors(Paper, Rock))

	assert.EqualValues(t, Loose, EvaluateRockPaperScissors(Scissors, Rock))
	assert.EqualValues(t, Loose, EvaluateRockPaperScissors(Paper, Scissors))
	assert.EqualValues(t, Loose, EvaluateRockPaperScissors(Rock, Paper))

	assert.EqualValues(t, Draw, EvaluateRockPaperScissors(Rock, Rock))
	assert.EqualValues(t, Draw, EvaluateRockPaperScissors(Scissors, Scissors))
	assert.EqualValues(t, Draw, EvaluateRockPaperScissors(Paper, Paper))
}

func Test0201Example(t *testing.T) {
	score := scoreRockPaperScissors(CharToRockPaperScissors('A'), CharToRockPaperScissors('Y'))
	score += scoreRockPaperScissors(CharToRockPaperScissors('B'), CharToRockPaperScissors('X'))
	score += scoreRockPaperScissors(CharToRockPaperScissors('C'), CharToRockPaperScissors('Z'))

	assert.Equal(t, 15, score)
}

func Test0201(t *testing.T) {
	assert.Equal(t, 11841, ReadAndEvaluateRockPaperScissors("input/day02"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 2, part 2

func Test0202(t *testing.T) {
	assert.Equal(t, 13022, ReadAndEvaluateRockPaperScissorsOutcome("input/day02"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 3, part 1

func Test0301Example(t *testing.T) {
	assert.Equal(t, 157, CalculateRucksackPriority("input/day03_example"))
}

func Test0301(t *testing.T) {
	assert.Equal(t, 8088, CalculateRucksackPriority("input/day03"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 3, part 2

func Test0302Example(t *testing.T) {
	assert.Equal(t, 70, CalculateRucksackChunkedPriority("input/day03_example", 3))
}

func Test0302(t *testing.T) {
	assert.Equal(t, 2522, CalculateRucksackChunkedPriority("input/day03", 3))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 4, part 1

func Test0401Example(t *testing.T) {
	assert.Equal(t, 2, FindCleaningAreaOverlap("input/day04_example", true))
}

func Test0401(t *testing.T) {
	assert.Equal(t, 462, FindCleaningAreaOverlap("input/day04", true))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 4, part 2

func Test0402Example(t *testing.T) {
	assert.Equal(t, 4, FindCleaningAreaOverlap("input/day04_example", false))
}

func Test0402(t *testing.T) {
	assert.Equal(t, 835, FindCleaningAreaOverlap("input/day04", false))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 5, part 1

func Test0501Example(t *testing.T) {
	assert.Equal(t, "CMZ", ReArrangeCrates("input/day05_example", new(CrateMover9000)))
}

func Test0501(t *testing.T) {
	assert.Equal(t, "CVCWCRTVQ", ReArrangeCrates("input/day05", new(CrateMover9000)))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 5, part 2

func Test0502Example(t *testing.T) {
	assert.Equal(t, "MCD", ReArrangeCrates("input/day05_example", new(CrateMover9001)))
}

func Test0502(t *testing.T) {
	assert.Equal(t, "CNSCZWLVT", ReArrangeCrates("input/day05", new(CrateMover9001)))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 6, part 1

func Test0601Example(t *testing.T) {
	assert.Equal(t, 7, FindStartOfPacket("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4))
	assert.Equal(t, 5, FindStartOfPacket("bvwbjplbgvbhsrlpgdmjqwftvncz", 4))
	assert.Equal(t, 6, FindStartOfPacket("nppdvjthqldpwncqszvftbrmjlhg", 4))
	assert.Equal(t, 10, FindStartOfPacket("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4))
	assert.Equal(t, 11, FindStartOfPacket("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4))
}

func Test0601(t *testing.T) {
	assert.Equal(t, 1275, FindStartOfPacketFromFile("input/day06", 4))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 6, part 2

func Test0602Example(t *testing.T) {
	assert.Equal(t, 19, FindStartOfPacket("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
	assert.Equal(t, 23, FindStartOfPacket("bvwbjplbgvbhsrlpgdmjqwftvncz", 14))
	assert.Equal(t, 23, FindStartOfPacket("nppdvjthqldpwncqszvftbrmjlhg", 14))
	assert.Equal(t, 29, FindStartOfPacket("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14))
	assert.Equal(t, 26, FindStartOfPacket("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14))
}

func Test0602(t *testing.T) {
	assert.Equal(t, 3605, FindStartOfPacketFromFile("input/day06", 14))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 7, part 1

func Test0701Example(t *testing.T) {
	assert.Equal(t, 95437, DirSizes("input/day07_example"))
}

func Test0701(t *testing.T) {
	assert.Equal(t, 1477771, DirSizes("input/day07"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 7, part 2

func Test0702Example(t *testing.T) {
	assert.Equal(t, 24933642, FindDirToDelete(70000000, 30000000, "input/day07_example"))
}

func Test0702(t *testing.T) {
	assert.Equal(t, 3579501, FindDirToDelete(70000000, 30000000, "input/day07"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 8, part 1

func Test0801Example(t *testing.T) {
	assert.Equal(t, 21, CountHiddenTrees("input/day08_example"))
}

func Test0801(t *testing.T) {
	assert.Equal(t, 1719, CountHiddenTrees("input/day08"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 8, part 2

func Test0802Example(t *testing.T) {
	x, y, score := FindMostScenicTree("input/day08_example")
	assert.Equal(t, 2, x)
	assert.Equal(t, 3, y)
	assert.Equal(t, 8, score)
}

func Test0802(t *testing.T) {
	x, y, score := FindMostScenicTree("input/day08")
	assert.Equal(t, 46, x)
	assert.Equal(t, 85, y)
	assert.Equal(t, 590824, score)
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 9, part 1

func Test0901Example(t *testing.T) {
	assert.Equal(t, 13, CountRopeTailPositions("input/day09_example", 2))
}

func Test0901(t *testing.T) {
	assert.Equal(t, 6284, CountRopeTailPositions("input/day09", 2))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 9, part 2

func Test0902Example1(t *testing.T) {
	assert.Equal(t, 1, CountRopeTailPositions("input/day09_example", 10))
}

func Test0902Example2(t *testing.T) {
	assert.Equal(t, 36, CountRopeTailPositions("input/day0902_example", 10))
}

func Test0902(t *testing.T) {
	assert.Equal(t, 2661, CountRopeTailPositions("input/day09", 10))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 10, part 1

func Test1001Example(t *testing.T) {
	assert.Equal(t, 13140, FindSignalStrength("input/day10_example"))
}

func Test1001(t *testing.T) {
	assert.Equal(t, 14920, FindSignalStrength("input/day10"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 10, part 2

func Test1002Example(t *testing.T) {
	RenderSignalToCrt("input/day10_example")
}

func Test1002(t *testing.T) {
	RenderSignalToCrt("input/day10")
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 11, part 1

func Test1101Example(t *testing.T) {
	assert.Equal(t, 10605, ExecuteDay11("input/day11_example", 20, 1))
}

func Test1101(t *testing.T) {
	assert.Equal(t, 111210, ExecuteDay11("input/day11", 20, 1))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 11, part 2

func Test1102Example(t *testing.T) {
	assert.Equal(t, 2713310158, ExecuteDay11("input/day11_example", 10000, 2))
}

func Test1102(t *testing.T) {
	assert.Equal(t, 15447387620, ExecuteDay11("input/day11", 10000, 2))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 12, part 1

func Test1201Example(t *testing.T) {
	assert.Equal(t, 31, FindShortestRoute("input/day12_example"))
}

func Test1201(t *testing.T) {
	assert.Equal(t, 456, FindShortestRoute("input/day12"))
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
	slice := ParsePacketString("[1,1,3,1,1]")

	assert.Equal(t, float64(1), slice[0])
	assert.Equal(t, float64(1), slice[1])
	assert.Equal(t, float64(3), slice[2])
	assert.Equal(t, float64(1), slice[3])
	assert.Equal(t, float64(1), slice[4])
}

func Test1301ParsePacket2(t *testing.T) {
	slice := ParsePacketString("[[8,7,6]]")

	switch v := slice[0].(type) {
	case []any:
		assert.Equal(t, float64(8), v[0])
		assert.Equal(t, float64(7), v[1])
		assert.Equal(t, float64(6), v[2])
	}
}

func Test1301Example(t *testing.T) {
	assert.Equal(t, 13, CountOrderedIndices("input/day13_example"))
}

func Test1301(t *testing.T) {
	assert.Equal(t, 5196, CountOrderedIndices("input/day13"))
}

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 13, part 2

func Test1302Example(t *testing.T) {
	assert.Equal(t, 140, FindDecoderKey("input/day13_example"))
}

func Test1302(t *testing.T) {
	assert.Equal(t, 22134, FindDecoderKey("input/day13"))
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

// ---------------------------------------------------------------------------
// Advent of Code 2022: Day 15, part 1

func TestManhattanDistance(t *testing.T) {
	assert.EqualValues(t, 0, ManhattanDistance(Coordinate{x: 0, y: 0}, Coordinate{x: 0, y: 0}))
	assert.EqualValues(t, 9, ManhattanDistance(Coordinate{x: 8, y: 7}, Coordinate{x: -1, y: 7}))
	assert.EqualValues(t, 9, ManhattanDistance(Coordinate{x: 8, y: 7}, Coordinate{x: 8, y: -2}))
	assert.EqualValues(t, 9, ManhattanDistance(Coordinate{x: 8, y: 7}, Coordinate{x: 8, y: 16}))
	assert.EqualValues(t, 9, ManhattanDistance(Coordinate{x: 8, y: 7}, Coordinate{x: 17, y: 7}))
	assert.EqualValues(t, 9, ManhattanDistance(Coordinate{x: 8, y: 7}, Coordinate{x: 9, y: 15}))
	assert.EqualValues(t, 9, ManhattanDistance(Coordinate{x: 8, y: 7}, Coordinate{x: 17, y: 7}))
	assert.EqualValues(t, 9, ManhattanDistance(Coordinate{x: 8, y: 7}, Coordinate{x: 0, y: 8}))
	assert.EqualValues(t, 9, ManhattanDistance(Coordinate{x: 8, y: 7}, Coordinate{x: 17, y: 7}))
}

func Test1501Example(t *testing.T) {
	assert.Equal(t, 26, CountCoordinatesWithoutMissingBeacon("input/day15_example", 10))
}

func Test1501(t *testing.T) {
	assert.Equal(t, 4737443, CountCoordinatesWithoutMissingBeacon("input/day15", 2000000))
}
