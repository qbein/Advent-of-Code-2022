package adventOfCode2022

import (
	"fmt"
	"strings"
)

// Solver is an A* path finding solver.
// Ref. https://en.wikipedia.org/wiki/Pathfinding
type Solver struct {
	heightMap [][]rune
	start Coordinate
	end Coordinate
}

type Node struct {
	x int
	y int
	count int
	parent *Node
}

func NodeFromCoordinate(coordinate Coordinate, count int, parent *Node) Node {
	return Node{x: coordinate.x, y: coordinate.y, count: count, parent: parent}
}

func CoordinateFromNode(node Node) Coordinate {
	return Coordinate{x: node.x, y: node.y}
}

func (s *Solver) FindShortest() int {
	queue := make([]Node, 0)
	queue = append(queue, NodeFromCoordinate(s.start, 0, nil))

	fmt.Printf("  0123456789")
	for y, r := range s.heightMap {
		fmt.Printf("%d ", y)
		for _, c := range r {
			fmt.Printf("%s", string(c))
		}
		fmt.Println()
	}

	for i:=0; ; i++ {
		if len(queue) == 0 {
			break
		}

		current := queue[i]

		if current.x == s.end.x && current.y == s.end.y {
			fmt.Printf("Found exist from x:%d y:%d count:%count\n", current.x, current.y, current.count)
			return current.count
		}

		adjacentItems := make([]Coordinate, 0)
		adjacentItems = s.appendIfValid(queue, adjacentItems, current, Coordinate{x: current.x, y: current.y+1})
		adjacentItems = s.appendIfValid(queue, adjacentItems, current, Coordinate{x: current.x+1, y: current.y})
		adjacentItems = s.appendIfValid(queue, adjacentItems, current, Coordinate{x: current.x - 1, y: current.y})
		adjacentItems = s.appendIfValid(queue, adjacentItems, current, Coordinate{x: current.x, y: current.y-1})

		for _, v := range adjacentItems {
			count := current.count+1
			queue = append(queue, NodeFromCoordinate(v, count, &current))
			fmt.Printf("Appending x:%d, y:%d, count: %d (parent x:%d y:%d)\n", v.x, v.y, count, current.x, current.y)
		}
	}

	return -1
}

func (s *Solver) appendIfValid(queue []Node, adjacent []Coordinate, current Node, next Coordinate) []Coordinate {
	// Coordinate must be in map
	if next.y > len(s.heightMap)-1 ||
		next.x > len(s.heightMap[0])-1 ||
		next.x < 0 ||
		next.y < 0 {
		return adjacent
	}

	currentHeight := s.heightMap[current.y][current.x]
	nextHeight := s.heightMap[next.y][next.x]

	valid := false

	// 1 step incline is valid
	if currentHeight+1 == nextHeight {
		valid = true
	}

	// Can go down steep declines
	if currentHeight >= nextHeight {
		valid = true
	}

	// Do not add duplicates
	if valid {
		for _, v := range queue {
			if v.x == next.x && v.y == next.y {
				return adjacent
			}
		}

		return append(adjacent, next)
	}

	return adjacent
}

func SolverForFile(fileName string) Solver {
	heightMap := make([][]rune, 0)
	var start Coordinate
	var end Coordinate

	y := 0
	forLines(fileName, func(line string) {
		line = strings.TrimSpace(line)
		heightMap = append(heightMap, make([]rune, len(line)))
		for x, c := range line {
			switch c {
			case 'S':
				start = Coordinate{x, y}
				heightMap[y][x] = 'a'
			case 'E':
				end = Coordinate{x, y}
				heightMap[y][x] = 'z'
			default:
				heightMap[y][x] = c
			}
		}
		y++
	})

	return Solver{
		heightMap,
		start,
		end,
	}
}


func findShortestRoute(fileName string) int {
	heightMap := SolverForFile(fileName)
	return heightMap.FindShortest()
}