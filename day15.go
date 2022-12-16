package adventOfCode2022

import (
	"math"
	"regexp"
)

func parseSensorsFrom(fileName string) (int, int, int, int, map[Coordinate]Coordinate) {
	regex := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

	sensorMap := make(map[Coordinate]Coordinate)

	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt

	ForLinesIn(fileName, func(line string) {
		matches := regex.FindAllStringSubmatch(line, -1)
		if matches != nil {
			sensor := Coordinate{x: MustAtoi(matches[0][1]), y: MustAtoi(matches[0][2])}
			beacon := Coordinate{x: MustAtoi(matches[0][3]), y: MustAtoi(matches[0][4])}

			if sensor.x < minX {
				minX = sensor.x
			}
			if sensor.x > maxX {
				maxX = sensor.x
			}
			if sensor.y < minY {
				minY = sensor.y
			}
			if sensor.y > maxY {
				maxY = sensor.y
			}

			sensorMap[sensor] = beacon
		}
	})

	return minX, maxX, minY, maxY, sensorMap
}

func pointHasBeacon(sensorMap map[Coordinate]Coordinate, point Coordinate) bool {
	for _, beacon := range sensorMap {
		if beacon == point {
			return true
		}
	}
	return false
}

func canHaveUnknownBeacon(sensorMap map[Coordinate]Coordinate, point Coordinate) bool {
	for sensor, closestBeacon := range sensorMap {
		distanceToClosest := ManhattanDistance(closestBeacon, sensor)
		distanceToPoint := ManhattanDistance(point, sensor)

		if distanceToPoint <= distanceToClosest {
			return false
		}
	}

	return true
}

func CountCoordinatesWithoutMissingBeacon(fileName string, y int) int {
	minX, maxX, _, _, sensorMap := parseSensorsFrom(fileName)
	// TODO: Find a better way to define range?
	xStart := minX - maxX/5
	xEnd := maxX + maxX/5

	count := 0

	for x := xStart; x <= xEnd; x++ {
		point := Coordinate{x: x, y: y}
		if !canHaveUnknownBeacon(sensorMap, point) && !pointHasBeacon(sensorMap, point) {
			count++
		}
	}

	return count
}
