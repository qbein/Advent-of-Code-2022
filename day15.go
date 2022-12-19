package adventOfCode2022

import (
	"fmt"
	"regexp"
)

type Beacon struct {
	coordinate Coordinate
	distance   int
}

func parseSensorsFrom(fileName string) map[Coordinate]Beacon {
	regex := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

	sensorMap := make(map[Coordinate]Beacon)

	ForLinesIn(fileName, func(line string) {
		matches := regex.FindAllStringSubmatch(line, -1)
		if matches != nil {
			sensor := Coordinate{x: MustAtoi(matches[0][1]), y: MustAtoi(matches[0][2])}
			beacon := Coordinate{x: MustAtoi(matches[0][3]), y: MustAtoi(matches[0][4])}
			distance := ManhattanDistance(sensor, beacon)
			sensorMap[sensor] = Beacon{beacon, distance}
		}
	})

	return sensorMap
}

func hasBeaconAtCoordinate(beacons map[Coordinate]bool, point Coordinate) bool {
	_, ok := beacons[point]
	return ok
}

func canHaveUnknownBeacon(sensorMap map[Coordinate]Beacon, point Coordinate) bool {
	for sensor, closestBeacon := range sensorMap {
		distanceToClosest := closestBeacon.distance
		distanceToPoint := ManhattanDistance(point, sensor)

		if distanceToPoint <= distanceToClosest {
			return false
		}
	}

	return true
}

func uniqueBeacons(sensorMap map[Coordinate]Beacon) map[Coordinate]bool {
	beaconMap := make(map[Coordinate]bool)
	for _, b := range sensorMap {
		beaconMap[b.coordinate] = true
	}
	return beaconMap
}

func CountCoordinatesWithoutMissingBeacon(fileName string, y int, minX int, maxX int) int {
	sensorMap := parseSensorsFrom(fileName)
	count := 0

	beacons := uniqueBeacons(sensorMap)

	for x := minX; x <= maxX; x++ {
		point := Coordinate{x: x, y: y}
		if !canHaveUnknownBeacon(sensorMap, point) && !hasBeaconAtCoordinate(beacons, point) {
			count++
		}
	}

	return count
}

func printProgress(message string, count int, max int) {
	percent := 0.0
	if count > 0 {
		percent = float64(count) / float64(max) * 100.0
	}
	fmt.Printf("%s: %f%%\n", message, percent)
}

func FindTuningFrequency(fileName string, max int, verbose bool) int {
	sensorMap := parseSensorsFrom(fileName)

	searchPoints := make([]Coordinate, 0)

	count := 0

	// The missing beacon has to be just outside the range of other sensors
	// Attempt to locate the beacon

	for sensor, beacon := range sensorMap {
		count++

		// Add search coordinates northeast of sensor range
		for x := 0; x <= beacon.distance+1; x++ {
			// Find y coordinate from x and manhattan distance
			y := beacon.distance + 1 - x

			if y < 0 || sensor.y+y > max || sensor.x+x > max {
				continue
			}

			coordinate := Coordinate{x: sensor.x + x, y: sensor.y + y}
			searchPoints = append(searchPoints, coordinate)
		}

		// Found correct coordinate in this location, no need to continue...
	}

	count = 0

	// Check the search locations for possible unknown beacons
	for _, point := range searchPoints {
		count++
		if canHaveUnknownBeacon(sensorMap, point) {
			return point.x*4000000 + point.y
		}
	}

	return -1
}
