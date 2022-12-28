package adventOfCode2022

import (
	"fmt"
	"golang.org/x/exp/slices"
	"regexp"
	"strings"
)

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

type Current struct {
	id          string
	minutesLeft int
}

func FindMaxPressureAfterMinutes(fileName string, minutes int, workers int, verbose bool) int {
	flow, done := 0, 0
	valves := readValves(fileName)

	openValves := make([]string, 0)

	current := make([]Current, workers)
	for i := 0; i < workers; i++ {
		current[i] = Current{id: "AA", minutesLeft: minutes}
	}

	for {
		for workerId := 0; workerId < workers; workerId++ {
			if !stepCurrent(&current[workerId], valves, openValves, workerId) {
				done++
				continue
			}

			valve := valves[current[workerId].id]

			if verbose {
				fmt.Printf("Worker %d opens %s, minutes left: %d\n", workerId, valve.id, current[workerId].minutesLeft)
			}

			openValves = append(openValves, valve.id)
			flow += valve.flow * current[workerId].minutesLeft
		}
		if done >= workers {
			break
		}
	}

	return flow
}

func stepCurrent(current *Current, valves map[string]Valve, openValves []string, workerId int) bool {
	moves := findNextToOpen(valves, openValves, valves[current.id], current.minutesLeft)
	if len(moves) == 0 {
		return false
	}

	move := moves[0]
	if workerId == 1 && len(moves) > 2 {
		move = moves[1]
	}

	current.id = move.id

	// Move to valve
	current.minutesLeft -= move.moveTime

	// Open valve
	current.minutesLeft--
	if current.minutesLeft == 0 {
		return false
	}
	return true
}

func findNextToOpen(valves map[string]Valve, openValves []string, currentValve Valve, minutesLeft int) []Move {
	moves := make([]Move, 0)

	if minutesLeft <= 0 {
		return moves
	}

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
		nextMoves := findNextToOpen(valves, append(openValves, move.id), valve, minutesLeft-move.moveTime-1)
		if len(nextMoves) > 0 {
			move.score += nextMoves[0].score
		}

		moves = append(moves, move)
	}

	slices.SortFunc(
		moves,
		func(a Move, b Move) bool {
			return a.score > b.score
		})

	return moves
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
