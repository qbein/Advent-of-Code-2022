package adventOfCode2022

import (
	"fmt"
	"strings"
)

func parseTrees(fileName string) ([][]int, int, int) {
	trees := make([][]int, 0)
	ForLinesIn(fileName, func(line string) {
		trees = append(trees, treeRowToSlice(line))
	})

	// Ignore possibility of varying row sizes, and empty dataset
	height := len(trees)
	width := len(trees[0])

	return trees, height, width
}

func CountHiddenTrees(fileName string) int {
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
	if FindMaxValue(slice) < trees[row][col] {
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
		return calculateSliceScenicScore(outHeight, Reverse(slice))
	case 'w':
		return calculateSliceScenicScore(outHeight, Reverse(slice))
	case 's':
		return calculateSliceScenicScore(outHeight, slice)
	case 'e':
		return calculateSliceScenicScore(outHeight, slice)
	default:
		panic(fmt.Sprintf("Invalid direction %c (valid: n, s, e, w)", direction))
	}
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

func FindMostScenicTree(fileName string) (int, int, int) {
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
