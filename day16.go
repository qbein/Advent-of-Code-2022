package adventOfCode2022

import (
	"golang.org/x/exp/slices"
	"regexp"
	"strings"
)

const MINUTES = 30

var moveTimeCache map[string]int

func readValves(fileName string) map[string]Valve {
	moveTimeCache = make(map[string]int)

	valves := make(map[string]Valve)
	regex := regexp.MustCompile(`^Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

	ForLinesIn(fileName, func(line string) {
		match := regex.FindAllStringSubmatch(line, -1)
		if match == nil {
			return
		}

		valve := Valve{
			id:    match[0][1],
			flow:  MustAtoi(match[0][2]),
			links: strings.Split(match[0][3], ", "),
		}

		valves[valve.id] = valve
	})

	return valves
}

func FindMaxPressureAfterMinutes(fileName string) int {
	valves := readValves(fileName)

	openValves := make([]string, 0)

	currentValve := valves["AA"]
	minutesLeft := MINUTES
	flow := 0

	for {
		move := findNextToOpen(valves, openValves, currentValve, minutesLeft)
		if move.id == "" {
			break
		}

		currentValve = valves[move.id]

		// Move to valve
		minutesLeft -= move.moveTime

		// Open valve
		minutesLeft--
		if minutesLeft == 0 {
			break
		}
		flow += currentValve.flow * minutesLeft
		openValves = append(openValves, move.id)
	}

	return flow
}

func findNextToOpen(valves map[string]Valve, openValves []string, currentValve Valve, minutesLeft int) Move {
	if minutesLeft <= 0 {
		return Move{}
	}

	var bestMove Move

	for _, valve := range valves {
		// Skip open or stuck valves
		i := slices.Index(openValves, valve.id)
		if i > -1 || valve.flow == 0 || valve.id == currentValve.id {
			continue
		}

		// Find time to valve
		move := Move{
			id:       valve.id,
			moveTime: TimeTo(valves, currentValve.id, valve.id),
		}

		// If there is no time to open valve, abort
		if minutesLeft-move.moveTime <= 0 {
			continue
		}

		move.score = (minutesLeft - move.moveTime) * valve.flow

		// Make sure we're on the optimal path by also inspecting subsequent moves
		nextMove := findNextToOpen(valves, append(openValves, move.id), valve, minutesLeft-move.moveTime-1)
		if nextMove.id != "" {
			move.score += nextMove.score
		}

		if move.score > bestMove.score {
			bestMove = move
		}
	}

	return bestMove
}

func TimeTo(valves map[string]Valve, from string, to string) int {
	if from == to {
		return 0
	}
	cacheKey := from + to
	time, ok := moveTimeCache[cacheKey]
	if ok {
		return time
	}

	path := []string{from}
	path = recurseTo(valves, from, to, path)
	time = len(path)

	moveTimeCache[cacheKey] = time
	return time
}

func recurseTo(valves map[string]Valve, from string, to string, path []string) []string {
	if from == to {
		panic("cannot recurse to self")
	}

	valve := valves[from]

	var bestChildPath []string

	for _, childId := range valve.links {
		// Don't revisit valves
		if slices.Index(path, childId) > -1 {
			continue
		}

		childPath := append(path, childId)
		if childId == to {
			return path
		}

		childPath = recurseTo(valves, childId, to, childPath)

		if childPath != nil && (bestChildPath == nil || len(childPath) < len(bestChildPath)) {
			bestChildPath = childPath
		}
	}

	return bestChildPath
}

type Move struct {
	moveTime int
	score    int
	id       string
}

type Valve struct {
	id    string
	flow  int
	links []string
}
