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
	items         []int
	operation     func(int) int
	divisor       int
	falseMonkeyId int
	trueMonkeyId  int
	monkeys       map[int]*Monkey
	throwCount    int
}

func (m *Monkey) AddItem(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) RemoveItem(item int) bool {
	for i, v := range m.items {
		if v == item {
			m.items = slices.Delete(m.items, i, i+1)
			return true
		}
	}
	return false
}

func (m *Monkey) CalculateNewWorryLevel(item int) int {
	return m.operation(item) / 3
}

func (m *Monkey) ThrowItem(item int) {
	//fmt.Println(m.CalculateNewWorryLevel(item))
	if m.CalculateNewWorryLevel(item)%m.divisor == 0 {
		m.ThrowItemTo(m.trueMonkeyId, item)
	} else {
		m.ThrowItemTo(m.falseMonkeyId, item)
	}
	m.throwCount++
}

func (m *Monkey) ThrowItemTo(monkeyId int, item int) {
	if m.RemoveItem(item) {
		(*m.monkeys[monkeyId]).AddItem(m.CalculateNewWorryLevel(item))
	} else {
		panic(fmt.Sprintf("Tried to remove unexisting item %d for monkey %d", item, m.id))
	}
}

func readMonkeys(fileName string) map[int]*Monkey {
	monkeys := make(map[int]*Monkey)

	monkeyRegex := regexp.MustCompile(`^Monkey (\d):`)
	operationRegex := regexp.MustCompile(`Operation: new = old ([+*]) (.*)`)

	monkeyScanIndex := 0
	currentMonkeyId := -1

	forLines(fileName, func(line string) {
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
			items := make([]int, len(p))
			monkeys[currentMonkeyId].items = items
			for i, v := range p {
				monkeys[currentMonkeyId].items[i] = int(MustAtoi(v))
			}
		case 2:
			m := operationRegex.FindAllStringSubmatch(line, -1)
			if m == nil {
				panic(fmt.Sprintf("Could not parse operation: %s", line))
			}
			valueAsString := strings.TrimSpace(m[0][2])
			switch m[0][1] {
			case "+":
				monkeys[currentMonkeyId].operation = func(item int) int {
					value := item
					if valueAsString != "old" {
						value = int(MustAtoi(valueAsString))
					}
					return item + value
				}
			case "*":
				monkeys[currentMonkeyId].operation = func(item int) int {
					value := item
					if valueAsString != "old" {
						value = int(MustAtoi(valueAsString))
					}
					return item * value
				}
			}
		case 3:
			monkeys[currentMonkeyId].divisor = int(MustAtoi(line[21:]))
		case 4:
			monkeys[currentMonkeyId].trueMonkeyId = MustAtoi(line[strings.LastIndex(line, " ")+1:])
		case 5:
			monkeys[currentMonkeyId].falseMonkeyId = MustAtoi(line[strings.LastIndex(line, " ")+1:])
		}
	})

	return monkeys
}

func playRoundOfMonkeyInMiddle(monkeys map[int]*Monkey, round int) {
	for i := 0; i < len(monkeys); i++ {
		itemsInPlay := make([]int, len(monkeys[i].items))
		copy(itemsInPlay, monkeys[i].items)
		for _, item := range itemsInPlay {
			monkeys[i].ThrowItem(item)
		}
	}
}

func printMonkeys(monkeys map[int]*Monkey) {
	fmt.Println("Monkeys:")
	for i := 0; i < len(monkeys); i++ {
		fmt.Printf("  Monkey %d inspected items %d times.\n", i, monkeys[i].throwCount)
	}
}

func executeDay11(fileName string, rounds int) int {
	monkeys := readMonkeys(fileName)

	for i := 0; i < rounds; i++ {
		playRoundOfMonkeyInMiddle(monkeys, i)
	}

	throwCounts := make([]int, len(monkeys))
	for i, m := range monkeys {
		throwCounts[i] = m.throwCount
	}

	sort.Ints(throwCounts)

	return throwCounts[len(throwCounts)-2] * throwCounts[len(throwCounts)-1]
}
