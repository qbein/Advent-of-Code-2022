package adventOfCode2022

import (
	"golang.org/x/exp/maps"
	"strconv"
	"strings"
)

func CountRopeTailPositions(fileName string, knotCount int) int {
	moves := readRopeHeadMoves(fileName)

	// Initialize and place all knot in center of grid
	knots := make([]Coordinate, knotCount)

	// Map for keeping track of tail positions
	tailPositions := make(map[string]bool)
	tailPositions[knots[0].ToString()] = true

	lastPos := ""

	for _, move := range moves {
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
	ForLinesIn(fileName, func(line string) {
		moves = append(moves, strings.TrimSpace(line))
	})
	return moves
}
