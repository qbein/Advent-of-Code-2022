package adventOfCode2022

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/maps"
	"os"
	"regexp"
	"sort"
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
	id int
	x  int
	y  int
}

func (c *Coordinate) ToString() string {
	return fmt.Sprintf("id %d = %d:%d", c.id, c.x, c.y)
}

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
	var current = 0
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

func flattenUnique(input [][]string) string {
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

func findCommonChars(input []string) string {
	if len(input) < 1 {
		return ""
	}

	common := input[0]

	for i := 0; i < len(input)-1; i++ {
		r := regexp.MustCompile("[" + common + "]")
		common = flattenUnique(r.FindAllStringSubmatch(input[i+1], -1))
	}

	return common
}

func resolveRucksackPriority(fileName string) int {
	priority := 0

	forLines(fileName, func(line string) {
		var compA, compB = line[0 : len(line)/2], line[len(line)/2:]

		match := findCommonChars([]string{compA, compB})
		if len(match) != 1 {
			panic("Expected single byte match")
		}
		priority += calculatePriority(match[0])
	})

	return priority
}

func calculatePriority(input byte) int {
	subtract := 0
	if int(input) < int('a') {
		subtract = int('A') - 27
	} else {
		subtract = int('a') - 1
	}
	return int(input) - subtract
}

func resolveRucksackChunkedPriority(fileName string, chunkSize int) int {
	priority := 0

	i := 0
	chunk := make([]string, chunkSize)

	forLines(fileName, func(line string) {
		defer func() {
			i++
		}()

		chunk[i%chunkSize] = line

		if i%chunkSize == chunkSize-1 {
			common := findCommonChars(chunk[:])
			if len(common) != 1 {
				panic(fmt.Sprintf("Incorrect common items found %s, %q\n", common, chunk))
			}
			priority += calculatePriority(common[0])
		}
	})

	return priority
}

func MustAtoi(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprintf("Cannot convert %s to int", input))
	}
	return output
}

func findCleaningAreaOverlap(fileName string, requireFullOverlap bool) int {
	overlapCount := 0

	r := regexp.MustCompile(`[-,]`)
	forLines(fileName, func(line string) {
		p := r.Split(line, -1)
		s1, e1, s2, e2 := MustAtoi(p[0]), MustAtoi(p[1]), MustAtoi(p[2]), MustAtoi(p[3])
		if requireFullOverlap {
			if s1 >= s2 && e1 <= e2 {
				overlapCount++
			} else if s2 >= s1 && e2 <= e1 {
				overlapCount++
			}
		} else {
			if s1 >= s2 && s1 <= e2 {
				overlapCount++
			} else if s2 >= s1 && s2 <= e1 {
				overlapCount++
			}
		}
	})

	return overlapCount
}

func getStackCount(numericRow string) int {
	numericRow = strings.TrimSpace(numericRow)
	i := strings.LastIndex(numericRow, " ") + 1
	return MustAtoi(numericRow[i:])
}

func parseStacks(stacksStr []string) [][]rune {
	stackCnt := getStackCount(stacksStr[len(stacksStr)-1])
	stacks := make([][]rune, stackCnt)

	for row := len(stacksStr) - 2; row >= 0; row-- {
		for col := 0; col < stackCnt; col++ {
			rowStr := stacksStr[row]

			// Only works for single digit columns
			chrIdx := 1
			if col > 0 {
				chrIdx = col + 4 + (col-1)*3
			}

			if chrIdx < len(rowStr) {
				value := rowStr[chrIdx]
				if value != ' ' && value != 0 {
					stacks[col] = append(stacks[col], rune(value))
				}
			}
		}
	}

	return stacks
}

func rearrangeCrates(fileName string, crane CrateMover) string {
	var stacks [][]rune
	var stacksStr []string
	operationRegexp := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

	forLines(fileName, func(line string) {
		// Parse stack configurations until first empty line
		if stacks == nil && len(strings.TrimSpace(line)) > 0 {
			stacksStr = append(stacksStr, line)
			return
		}

		// All stacks are read, parse into slices
		if stacks == nil {
			stacks = parseStacks(stacksStr)
		}

		m := operationRegexp.FindAllStringSubmatch(line, -1)

		// Skip lines without operations
		if m == nil {
			return
		}

		num, fromIdx, toIdx := MustAtoi(m[0][1]), MustAtoi(m[0][2])-1, MustAtoi(m[0][3])-1

		crane.MoveCrates(stacks, num, fromIdx, toIdx)
	})

	output := strings.Builder{}
	for _, s := range stacks {
		output.WriteRune(s[len(s)-1])
	}
	return output.String()
}

func isAllUnique(input string) bool {
	for i := 0; i < len(input); i++ {
		if strings.LastIndex(input, string(input[i])) != i {
			return false
		}
	}
	return true
}

func findStartOfPacket(input string, frameSize int) int {
	for i := frameSize; i < len(input); i++ {
		frame := input[i-frameSize : i]
		if isAllUnique(frame) {
			return i
		}
	}
	return -1
}

func findStartOfPacketFromFile(fileName string, frameSize int) int {
	start := -1
	forLines(fileName, func(line string) {
		start = findStartOfPacket(line, frameSize)
	})
	return start
}

func buildDirSizes(fileName string) map[string]int {
	dir := make(map[string]int)

	cmdRegex := regexp.MustCompile(`^\$ ([a-z]+)(.*)$`)
	fileRegex := regexp.MustCompile(`^(\d+) (\S+)$`)

	cd := "/"

	forLines(fileName, func(line string) {
		if cmdRegex.MatchString(line) {
			cmdMatch := cmdRegex.FindAllStringSubmatch(line, -1)
			switch cmdMatch[0][1] {
			case "cd":
				dirName := strings.TrimSpace(cmdMatch[0][2])

				if dirName == "/" {
					cd = "/"
				} else if dirName == ".." {
					cd = cd[0 : strings.LastIndex(cd, "/")-1]
				} else {
					cd = fmt.Sprintf("%s%s/", cd, dirName)
				}
			case "ls":
				// Noop
			}
			return
		}

		if fileRegex.MatchString(line) {
			fileMatch := fileRegex.FindAllStringSubmatch(line, -1)
			fileSize := MustAtoi(fileMatch[0][1])
			dir[cd] += fileSize

			addSizeToParentDirs(cd, dir, fileSize)
		}
	})

	return dir
}

func dirSizes(fileName string) int {
	dir := buildDirSizes(fileName)

	sum := 0
	for _, value := range dir {
		if value < 100000 {
			sum += value
		}
	}
	return sum
}

func addSizeToParentDirs(cd string, dir map[string]int, fileSize int) {
	for {
		i := strings.LastIndexByte(cd[0:len(cd)-1], '/') + 1
		if i <= 0 {
			break
		}
		cd = cd[0:i]
		dir[cd] += fileSize
	}
}

func findDirToDelete(total int, required int, fileName string) int {
	dir := buildDirSizes(fileName)

	available := total - dir["/"]
	values := maps.Values(dir)
	sort.Ints(values)

	for _, value := range values {
		if available+value > required {
			return value
		}
	}

	return -1
}

func parseTrees(fileName string) ([][]int, int, int) {
	trees := make([][]int, 0)
	forLines(fileName, func(line string) {
		trees = append(trees, treeRowToSlice(line))
	})

	// Ignore possibility of varying row sizes, and empty dataset
	height := len(trees)
	width := len(trees[0])

	return trees, height, width
}

func countHiddenTrees(fileName string) int {
	trees, height, width := parseTrees(fileName)

	visibleCount := width*2 + height*2 - 4

	// Offsetting start and stop since all edges are visible
	for row := 1; row < height-1; row++ {
		for col := 1; col < width-1; col++ {
			if isTreeVisible(trees, row, col) {
				visibleCount++
			}
		}
	}

	return visibleCount
}

func treeRowToSlice(row string) []int {
	slice := make([]int, len(strings.TrimSpace(row)))
	for i, r := range row {
		slice[i] = int(r - '0')
	}
	return slice
}

func isTreeVisible(trees [][]int, row int, col int) bool {
	return isTreeVisibleInDirection(trees, row, col, 'n') ||
		isTreeVisibleInDirection(trees, row, col, 's') ||
		isTreeVisibleInDirection(trees, row, col, 'e') ||
		isTreeVisibleInDirection(trees, row, col, 'w')
}

func isTreeVisibleInDirection(trees [][]int, row int, col int, direction rune) bool {
	slice := findTreesInDirection(trees, row, col, direction)
	if maxInSlice(slice) < trees[row][col] {
		return true
	}
	return false
}

func findTreesInDirection(trees [][]int, row int, col int, direction rune) []int {
	slice := make([]int, 0)
	switch direction {
	case 'n':
		slice = make([]int, row)
		for i := 0; i < row; i++ {
			slice[i] = trees[i][col]
		}
	case 's':
		slice = make([]int, len(trees)-row-1)
		j := 0
		for i := row + 1; i < len(trees); i++ {
			slice[j] = trees[i][col]
			j++
		}
	case 'e':
		slice = trees[row][col+1:]
	case 'w':
		slice = trees[row][0:col]
	default:
		panic(fmt.Sprintf("Invalid direction %c (valid: n, s, e, w)", direction))
	}

	return slice
}

func maxInSlice(slice []int) int {
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

func calculateScenicScore(trees [][]int, row int, col int) int {
	return calculateScenicScoreInDirection(trees, row, col, 'n') *
		calculateScenicScoreInDirection(trees, row, col, 's') *
		calculateScenicScoreInDirection(trees, row, col, 'e') *
		calculateScenicScoreInDirection(trees, row, col, 'w')
}

func calculateScenicScoreInDirection(trees [][]int, row int, col int, direction rune) int {
	slice := findTreesInDirection(trees, row, col, direction)
	outHeight := trees[row][col]

	switch direction {
	case 'n':
		return calculateSliceScenicScore(outHeight, reverseSlice(slice))
	case 'w':
		return calculateSliceScenicScore(outHeight, reverseSlice(slice))
	case 's':
		return calculateSliceScenicScore(outHeight, slice)
	case 'e':
		return calculateSliceScenicScore(outHeight, slice)
	default:
		panic(fmt.Sprintf("Invalid direction %c (valid: n, s, e, w)", direction))
	}
}

func reverseSlice(slice []int) []int {
	l := len(slice)
	reversed := make([]int, l)
	for i, v := range slice {
		reversed[l-i-1] = v
	}
	return reversed
}

func calculateSliceScenicScore(ourHeight int, slice []int) int {
	score := 0
	highest := 0

	for _, h := range slice {
		score++
		if h >= highest {
			highest = h
		}
		if h >= ourHeight {
			break
		}
	}
	return score
}

func findMostScenicTree(fileName string) (int, int, int) {
	x, y, score := 0, 0, -1

	trees, height, width := parseTrees(fileName)

	// Offsetting start and stop since all edges are visible
	for row := 1; row < height-1; row++ {
		for col := 1; col < width-1; col++ {
			s := calculateScenicScore(trees, row, col)
			if s > score {
				score = s
				x = col
				y = row
			}
		}
	}

	return x, y, score
}

func countRopeTailPositions(fileName string, knotCount int) int {
	moves := readRopeHeadMoves(fileName)

	// Initialize and place all knot in center of grid
	knots := make([]Coordinate, knotCount)
	for i, knot := range knots {
		knot.id = i
	}

	// Map for keeping track of tail positions
	tailPositions := make(map[string]bool)
	tailPositions[knots[0].ToString()] = true

	lastPos := ""

	for mi, move := range moves {
		p := strings.Split(move, " ")
		d := p[0]
		s, _ := strconv.Atoi(p[1])

		// Not bothering with size sanitizing
		for i := 0; i < s; i++ {
			stepKnot(d, &knots[0])
			for j := 0; j < len(knots)-1; j++ {
				if !followKnot(&knots[j], &knots[j+1]) {
					break
				}
			}

			last := knots[len(knots)-1]
			pos := last.ToString()
			fmt.Printf("tail: %s, mi: %d\n", pos, mi)
			if lastPos != pos {
				tailPositions[pos] = true
				lastPos = pos
			}
		}
	}

	return len(maps.Keys(tailPositions))
}

func stepKnot(d string, knot *Coordinate) {
	switch d {
	case "R":
		knot.x += 1
	case "L":
		knot.x -= 1
	case "U":
		knot.y -= 1
	case "D":
		knot.y += 1
	}
}

func followKnot(head *Coordinate, tail *Coordinate) bool {
	if head.x > tail.x+1 {
		tail.x = head.x - 1
		if tail.y > head.y {
			tail.y--
		} else if tail.y < head.y {
			tail.y++
		}
		return true
	} else if head.x < tail.x-1 {
		tail.x = head.x + 1
		if tail.y > head.y {
			tail.y--
		} else if tail.y < head.y {
			tail.y++
		}
		return true
	} else if head.y > tail.y+1 {
		tail.y = head.y - 1
		if tail.x > head.x {
			tail.x--
		} else if tail.x < head.x {
			tail.x++
		}
		return true
	} else if head.y < tail.y-1 {
		tail.y = head.y + 1
		if tail.x > head.x {
			tail.x--
		} else if tail.x < head.x {
			tail.x++
		}
		return true
	} else {
		return false
	}
}

func readRopeHeadMoves(fileName string) []string {
	moves := make([]string, 0)
	forLines(fileName, func(line string) {
		moves = append(moves, strings.TrimSpace(line))
	})
	return moves
}
