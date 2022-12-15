package adventOfCode2022

import "regexp"

func FindCleaningAreaOverlap(fileName string, requireFullOverlap bool) int {
	overlapCount := 0

	r := regexp.MustCompile(`[-,]`)
	ForLinesIn(fileName, func(line string) {
		p := r.Split(line, -1)
		s1, e1, s2, e2 := MustAtoi(p[0]), MustAtoi(p[1]), MustAtoi(p[2]), MustAtoi(p[3])
		if requireFullOverlap {
			if s1 >= s2 && e1 <= e2 {
				overlapCount++
			} else if s2 >= s1 && e2 <= e1 {
				overlapCount++
			}
		} else {
			if s1 >= s2 && s1 <= e2 {
				overlapCount++
			} else if s2 >= s1 && s2 <= e1 {
				overlapCount++
			}
		}
	})

	return overlapCount
}
