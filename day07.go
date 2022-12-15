package adventOfCode2022

import (
	"fmt"
	"golang.org/x/exp/maps"
	"regexp"
	"sort"
	"strings"
)

func buildDirSizes(fileName string) map[string]int {
	dir := make(map[string]int)

	cmdRegex := regexp.MustCompile(`^\$ ([a-z]+)(.*)$`)
	fileRegex := regexp.MustCompile(`^(\d+) (\S+)$`)

	cd := "/"

	ForLinesIn(fileName, func(line string) {
		if cmdRegex.MatchString(line) {
			cmdMatch := cmdRegex.FindAllStringSubmatch(line, -1)
			switch cmdMatch[0][1] {
			case "cd":
				dirName := strings.TrimSpace(cmdMatch[0][2])

				if dirName == "/" {
					cd = "/"
				} else if dirName == ".." {
					cd = cd[0 : strings.LastIndex(cd, "/")-1]
				} else {
					cd = fmt.Sprintf("%s%s/", cd, dirName)
				}
			case "ls":
				// Noop
			}
			return
		}

		if fileRegex.MatchString(line) {
			fileMatch := fileRegex.FindAllStringSubmatch(line, -1)
			fileSize := MustAtoi(fileMatch[0][1])
			dir[cd] += fileSize

			addSizeToParentDirs(cd, dir, fileSize)
		}
	})

	return dir
}

func DirSizes(fileName string) int {
	dir := buildDirSizes(fileName)

	sum := 0
	for _, value := range dir {
		if value < 100000 {
			sum += value
		}
	}
	return sum
}

func addSizeToParentDirs(cd string, dir map[string]int, fileSize int) {
	for {
		i := strings.LastIndexByte(cd[0:len(cd)-1], '/') + 1
		if i <= 0 {
			break
		}
		cd = cd[0:i]
		dir[cd] += fileSize
	}
}

func FindDirToDelete(total int, required int, fileName string) int {
	dir := buildDirSizes(fileName)

	available := total - dir["/"]
	values := maps.Values(dir)
	sort.Ints(values)

	for _, value := range values {
		if available+value > required {
			return value
		}
	}

	return -1
}
