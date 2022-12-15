package adventOfCode2022

import (
	"fmt"
	"strings"
)

func renderSignal(fileName string, renderer func(int, int)) {
	cycle := 0
	x := 1

	ForLinesIn(fileName, func(line string) {
		p := strings.Split(line, " ")
		instruction := strings.TrimSpace(p[0])
		var argument int
		if len(p) > 1 {
			argument = MustAtoi(p[1])
		}

		cyclesToComplete := 1
		if instruction == "addx" {
			cyclesToComplete = 2
		}

		cycleUntil := cycle + cyclesToComplete
		for ; cycle < cycleUntil; cycle++ {
			renderer(cycle, x)
		}

		if instruction == "addx" {
			x += argument
		}
	})
}

func FindSignalStrength(fileName string) int {
	signalStrengthSum := 0

	renderSignal(fileName, func(cycle int, x int) {
		c := cycle + 1
		if c > 0 && c == 20 || (c > 40 && (c-20)%40 == 0) {
			signalStrength := x * c
			signalStrengthSum += signalStrength
		}
	})

	return signalStrengthSum
}

func RenderSignalToCrt(fileName string) {
	fmt.Printf("\nCRT output from %s:\n\n", fileName)

	renderSignal(fileName, func(cycle int, x int) {
		col := cycle % 40
		if cycle > 0 && col == 0 {
			fmt.Println()
		}
		if col >= x-1 && col <= x+1 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
	})

	fmt.Println()
	fmt.Println()
}
