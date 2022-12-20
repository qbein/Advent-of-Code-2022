package adventOfCode2022

import (
	"fmt"
	"regexp"
	"strings"
)

func readValves(fileName string) map[string]Valve {
	valves := make(map[string]Valve)
	regex := regexp.MustCompile(`^Valve ([A-Z]{2}) has flow rate=(\d+); tunnels lead to valves ([A-Z, ]+)`)

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

func FindMaxPressureAfterMinutes(fileName string, minutes int) int {
	valves := readValves(fileName)

	startValve := valves["AA"]

	result := InspectionResult{flow: 0, openValves: ""}
	result = inspectValve(valves, startValve, result, minutes)

	return result.flow
}

func inspectValve(valves map[string]Valve, valve Valve, result InspectionResult, minutesLeft int) InspectionResult {
	if minutesLeft <= 0 {
		return result
	}

	// No need to open valve with no pressure
	if valve.flow > 2 && strings.Index(result.openValves, valve.id) == -1 {
		// Takes one minute to open valve
		minutesLeft--
		result.flow += valve.flow * minutesLeft
		result.openValves += ";" + valve.id

		fmt.Printf("Opening %s -> %s, flow %d, minutes %d\n", valve.id, result.openValves, result.flow, minutesLeft)
	}

	if minutesLeft <= 1 {
		return result
	}

	// Takes one minute to navigate to next valve
	descendantResults := make([]InspectionResult, 0)
	minutesLeft--
	for _, childId := range valve.links {
		// Valve is already open
		child := valves[childId]
		/*if result.isOpen(child) {
			continue
		}*/
		descendantResults = append(descendantResults, inspectValve(valves, child, result, minutesLeft))
	}

	for i, descendantResult := range descendantResults {
		if i == 0 || descendantResult.flow > result.flow {
			result = descendantResult
		}
	}

	return result
}

type Valve struct {
	id    string
	flow  int
	links []string
}

type InspectionResult struct {
	flow       int
	openValves string
}

func (c *InspectionResult) isOpen(valve Valve) bool {
	return strings.Index(c.openValves, valve.id) > -1
}
