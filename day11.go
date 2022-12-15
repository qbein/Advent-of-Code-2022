package adventOfCode2022

import (
	"fmt"
	"golang.org/x/exp/slices"
	"regexp"
	"sort"
	"strings"
)

type Monkey struct {
	id            int
	items         []uint64
	operation     func(uint64) uint64
	divisor       uint64
	falseMonkeyId int
	trueMonkeyId  int
	monkeys       map[int]*Monkey
	throwCount    int
	mod           uint64
}

func (m *Monkey) AddItem(item uint64) {
	m.items = append(m.items, item)
}

func (m *Monkey) RemoveItem(item uint64) bool {
	for i, v := range m.items {
		if v == item {
			m.items = slices.Delete(m.items, i, i+1)
			return true
		}
	}
	return false
}

func (m *Monkey) CalculateNewWorryLevel(item uint64) uint64 {
	if m.mod > 0 {
		return m.operation(item) % m.mod
	} else {
		return m.operation(item) / 3
	}
}

func (m *Monkey) ThrowItem(item uint64) {
	if m.CalculateNewWorryLevel(item)%m.divisor == 0 {
		m.ThrowItemTo(m.trueMonkeyId, item)
	} else {
		m.ThrowItemTo(m.falseMonkeyId, item)
	}
	m.throwCount++
}

func (m *Monkey) ThrowItemTo(monkeyId int, item uint64) {
	if m.RemoveItem(item) {
		(*m.monkeys[monkeyId]).AddItem(m.CalculateNewWorryLevel(item))
	} else {
		panic(fmt.Sprintf("Tried to remove unexisting item %d for monkey %d", item, m.id))
	}
}

func readMonkeys(fileName string, part int) map[int]*Monkey {
	monkeys := make(map[int]*Monkey)

	monkeyRegex := regexp.MustCompile(`^Monkey (\d):`)
	operationRegex := regexp.MustCompile(`Operation: new = old ([+*]) (.*)`)

	monkeyScanIndex := 0
	currentMonkeyId := -1

	ForLinesIn(fileName, func(line string) {
		if monkeyRegex.MatchString(line) {
			m := monkeyRegex.FindAllStringSubmatch(line, 1)
			currentMonkeyId = MustAtoi(m[0][1])
			monkey := Monkey{id: currentMonkeyId, monkeys: monkeys}
			monkeys[currentMonkeyId] = &monkey
			monkeyScanIndex = 0
			return
		}

		monkeyScanIndex++
		switch monkeyScanIndex {
		case 1:
			itemIds := line[18:]
			p := strings.Split(itemIds, ", ")
			items := make([]uint64, len(p))
			monkeys[currentMonkeyId].items = items
			for i, v := range p {
				monkeys[currentMonkeyId].items[i] = uint64(MustAtoi(v))
			}
		case 2:
			m := operationRegex.FindAllStringSubmatch(line, -1)
			if m == nil {
				panic(fmt.Sprintf("Could not parse operation: %s", line))
			}
			valueAsString := strings.TrimSpace(m[0][2])
			switch m[0][1] {
			case "+":
				monkeys[currentMonkeyId].operation = func(item uint64) uint64 {
					value := item
					if valueAsString != "old" {
						value = uint64(MustAtoi(valueAsString))
					}
					return item + value
				}
			case "*":
				monkeys[currentMonkeyId].operation = func(item uint64) uint64 {
					value := item
					if valueAsString != "old" {
						value = uint64(MustAtoi(valueAsString))
					}
					return item * value
				}
			}
		case 3:
			monkeys[currentMonkeyId].divisor = uint64(MustAtoi(line[21:]))
		case 4:
			monkeys[currentMonkeyId].trueMonkeyId = MustAtoi(line[strings.LastIndex(line, " ")+1:])
		case 5:
			monkeys[currentMonkeyId].falseMonkeyId = MustAtoi(line[strings.LastIndex(line, " ")+1:])
		}
	})

	if part == 2 {
		mod := monkeys[0].divisor
		for i := 1; i < len(monkeys); i++ {
			if monkeys[i].divisor > 0 {
				mod *= monkeys[i].divisor
			}
		}

		for _, m := range monkeys {
			m.mod = mod
		}
	}

	return monkeys
}

func playRoundOfMonkeyInMiddle(monkeys map[int]*Monkey) {
	for i := 0; i < len(monkeys); i++ {
		itemsInPlay := make([]uint64, len(monkeys[i].items))
		copy(itemsInPlay, monkeys[i].items)
		for _, item := range itemsInPlay {
			monkeys[i].ThrowItem(item)
		}
	}
}

func ExecuteDay11(fileName string, rounds int, part int) int {
	monkeys := readMonkeys(fileName, part)

	for i := 0; i < rounds; i++ {
		playRoundOfMonkeyInMiddle(monkeys)
	}

	throwCounts := make([]int, len(monkeys))
	for i, m := range monkeys {
		throwCounts[i] = m.throwCount
	}

	sort.Ints(throwCounts)

	return throwCounts[len(throwCounts)-2] * throwCounts[len(throwCounts)-1]
}
