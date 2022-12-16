package day11

import (
	"aicoc22/util"
	"fmt"
	"sort"
)

func Task111() {
	fmt.Println("\nDay 11, Part 1: Monkey in the middle")
	fmt.Println("===============================================")

	//monkeys := buildMonkeysSample()
	monkeys := buildMonkeys()
	round := 1
	maxRounds := 20

	for round <= maxRounds {
		fmt.Println("-------------------- -- -   -")
		fmt.Printf("Round: %d\n", round)
		fmt.Println("-------------------- -- -   -")
		processRoundPart1(monkeys)
		round++
	}

	fmt.Printf("\nMonkey Business-Level=%d", calculateMonkeyBusinessLevel(monkeys))

}

func Task112() {
	fmt.Println("\nDay 11, Part 2: Monkey in the middle")
	fmt.Println("===============================================")

	//monkeys := buildMonkeysSample()
	monkeys := buildMonkeys()
	round := 1
	maxRounds := 10000
	var optimization int64 = 1
	for _, m := range monkeys {
		optimization = optimization * m.DivisibleBy
	}

	for round <= maxRounds {
		fmt.Println("-------------------- -- -   -")
		fmt.Printf("Round: %d\n", round)
		fmt.Println("-------------------- -- -   -")
		processRoundPart2(monkeys, optimization)
		round++
	}

	fmt.Printf("\nMonkey Business-Level=%d", calculateMonkeyBusinessLevel(monkeys))

}

func processRoundPart1(monkeys []Monkey) {

	for monkeyNumber, monkey := range monkeys {
		fmt.Printf("Processing Monkey %d\n", monkeyNumber)
		valuesToRemove := make([]int64, 0)
		for _, worryLevel := range monkey.Items {
			monkey.Item = worryLevel
			monkey.Operate()
			monkey.Item = monkey.Item / 3
			if monkey.Item%monkey.DivisibleBy == 0 {
				monkeys[monkey.TrueToMonkey].AddItem(monkey.Item)
				monkey.Item = 0
			} else {
				monkeys[monkey.FalseToMonkey].AddItem(monkey.Item)
				monkey.Item = 0
			}
			valuesToRemove = append(valuesToRemove, worryLevel)
		}

		for _, value := range valuesToRemove {
			monkey.Items = util.RemoveValueFromIntSlice(monkey.Items, value)
		}
		//monkey.Items = util.RemoveIndex(monkey.Items, 0)
		monkeys[monkeyNumber] = monkey
	}

	for monkeyNumber, monkey := range monkeys {
		monkey.ToString(monkeyNumber)
	}
}

func processRoundPart2(monkeys []Monkey, optimization int64) {

	for monkeyNumber, monkey := range monkeys {
		fmt.Printf("Processing Monkey %d\n", monkeyNumber)
		valuesToRemove := make([]int64, 0)
		for _, worryLevel := range monkey.Items {
			monkey.Item = worryLevel
			monkey.Operate()
			monkey.Item = monkey.Item / 1 % optimization
			if monkey.Item%monkey.DivisibleBy == 0 {
				monkeys[monkey.TrueToMonkey].AddItem(monkey.Item)
				monkey.Item = 0
			} else {
				monkeys[monkey.FalseToMonkey].AddItem(monkey.Item)
				monkey.Item = 0
			}
			valuesToRemove = append(valuesToRemove, worryLevel)
		}

		for _, value := range valuesToRemove {
			monkey.Items = util.RemoveValueFromIntSlice(monkey.Items, value)
		}
		//monkey.Items = util.RemoveIndex(monkey.Items, 0)
		monkeys[monkeyNumber] = monkey
	}

	for monkeyNumber, monkey := range monkeys {
		monkey.ToString(monkeyNumber)
	}
}

func calculateMonkeyBusinessLevel(monkeys []Monkey) int64 {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspections > monkeys[j].Inspections
	})
	return monkeys[0].Inspections * monkeys[1].Inspections
}
