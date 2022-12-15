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

type Coordinate struct {
	x int
	y int
}

func (c *Coordinate) ToString() string {
	return fmt.Sprintf("%d:%d", c.x, c.y)
}

// ---------------------------------------------------------------------------
// Functions

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func AbsInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// ForLinesIn calls the provided function with each line in the file.
func ForLinesIn(fileName string, action func(string)) {
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

func FindMaxValue(slice []int) int {
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

func FindMinValueIndex(slice []int) int {
	minIndex := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func SumIntValues(values []int) int {
	var sum = 0
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return sum
}

func Reverse(slice []int) []int {
	l := len(slice)
	reversed := make([]int, l)
	for i, v := range slice {
		reversed[l-i-1] = v
	}
	return reversed
}

func FlattenUnique(input [][]string) string {
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

// FindCommonChars in the provided input string.
func FindCommonChars(input []string) string {
	if len(input) < 1 {
		return ""
	}

	common := input[0]

	for i := 0; i < len(input)-1; i++ {
		r := regexp.MustCompile("[" + common + "]")
		common = FlattenUnique(r.FindAllStringSubmatch(input[i+1], -1))
	}

	return common
}

// MustAtoi converts the provided string to an int and panics if the conversion could not be done.
func MustAtoi(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprintf("Cannot convert %s to int", input))
	}
	return output
}

func AllLettersUnique(input string) bool {
	for i := 0; i < len(input); i++ {
		if strings.LastIndex(input, string(input[i])) != i {
			return false
		}
	}
	return true
}
