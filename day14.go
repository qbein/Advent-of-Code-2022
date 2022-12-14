package adventOfCode2022

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func SandBoxFromFile(fileName string) *sandBox {
	sandbox := new(sandBox)

	toCoordinate := func(segment string) Coordinate {
		p := strings.Split(segment, ",")
		return Coordinate{x: MustAtoi(p[0]), y: MustAtoi(p[1])}
	}

	addLineToMatrix := func(offsetX int, offsetY int, a Coordinate, b Coordinate) {
		if a.x == b.x {
			from := min(a.y, b.y)
			to := max(a.y, b.y)
			for y := from; y <= to; y++ {
				sandbox.matrix[y-offsetY][a.x-offsetX] = 1
			}
		} else if a.y == b.y {
			from := min(a.x, b.x)
			to := max(a.x, b.x)
			for x := from; x <= to; x++ {
				sandbox.matrix[a.y-offsetY][x-offsetX] = 1
			}
		}
	}

	minX, minY, maxX, maxY := math.MaxInt, 0, math.MinInt, math.MinInt

	segments := make([][]Coordinate, 0)
	forLines(fileName, func(line string) {
		p := strings.Split(line, " -> ")
		segments = append(segments, make([]Coordinate, 0))
		last := len(segments) - 1
		for i := 0; i < len(p); i++ {
			c := toCoordinate(p[i])
			segments[last] = append(segments[last], c)
			if c.y < minY {
				minY = c.y
			}
			if c.x < minX {
				minX = c.x
			}
			if c.y > maxY {
				maxY = c.y
			}
			if c.x > maxX {
				maxX = c.x
			}
		}
	})

	// Initialize matrix
	sandbox.matrix = make([][]int, maxY-minY+1)
	for y := 0; y < len(sandbox.matrix); y++ {
		sandbox.matrix[y] = make([]int, maxX-minX+1)
	}

	for _, s := range segments {
		for i := 0; i < len(s)-1; i++ {
			addLineToMatrix(minX, minY, s[i], s[i+1])
		}
	}

	sandbox.dropCoordinate = Coordinate{x: 500 - minX, y: 0}

	return sandbox
}

type sandBox struct {
	matrix         [][]int
	dropCoordinate Coordinate
	count          int
}

func (s *sandBox) print() {
	fmt.Printf("Result after %d drops:\n", s.count)
	for y, row := range s.matrix {
		for x, v := range row {
			c := "."
			if s.dropCoordinate.x == x && s.dropCoordinate.y == y {
				c = "+"
			} else if v == 1 {
				c = "#"
			} else if v == 2 {
				c = "o"
			}
			fmt.Printf("%s", c)
		}
		fmt.Println()
	}
	fmt.Println()
}

// dropSandUnit from the given coordinate. Function returns false if
// the unit overflows the matrix and falls of the map.
func (s *sandBox) drop(verbose bool) bool {
	position := s.dropCoordinate

	if verbose {
		defer s.print()
	}

	for {
		p, err := s.step(position)
		if err != nil {
			return false
		}
		// Nowhere to go
		if p == position {
			s.matrix[p.y][p.x] = 2
			s.count++
			return true
		}
		position = p
	}
}

func (s *sandBox) step(position Coordinate) (Coordinate, error) {
	// If we're at the bottom of the matrix, abort
	if position.y+1 >= len(s.matrix) {
		return position, errors.New("overflow bottom")
	}
	// If next field is not fixed, step to next
	if s.matrix[position.y+1][position.x] == 0 {
		return Coordinate{x: position.x, y: position.y + 1}, nil
	}
	// We're dropping out on left
	if position.x == 0 {
		return Coordinate{x: -1, y: position.y}, errors.New("dropping out on left")
	}
	// Drop down to left
	if s.matrix[position.y+1][position.x-1] == 0 {
		return Coordinate{x: position.x - 1, y: position.y + 1}, nil
	}
	// Drop out to right
	if len(s.matrix[position.y+1]) == position.x+1 {
		return Coordinate{x: position.x + 1, y: position.y + 1}, errors.New("dropping out on right")
	}
	// Drop down to right
	if s.matrix[position.y+1][position.x+1] == 0 {
		return Coordinate{x: position.x + 1, y: position.y + 1}, nil
	}
	// Nowhere to go
	return position, nil
}

func (s *sandBox) solve(verbose bool) int {
	for {
		if !s.drop(verbose) {
			break
		}
	}

	if verbose {
		s.print()
	}

	return s.count
}
