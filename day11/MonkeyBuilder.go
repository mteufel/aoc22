package day11

func buildMonkeysSample() []Monkey {

	var m0 Monkey
	m0.Items = []int64{79, 98}
	m0.Operation = func(old int64) int64 {
		return old * 19
	}
	m0.DivisibleBy = 23
	m0.TrueToMonkey = 2
	m0.FalseToMonkey = 3

	var m1 Monkey
	m1.Items = []int64{54, 65, 75, 74}
	m1.Operation = func(old int64) int64 {
		return old + 6
	}
	m1.DivisibleBy = 19
	m1.TrueToMonkey = 2
	m1.FalseToMonkey = 0

	var m2 Monkey
	m2.Items = []int64{79, 60, 97}
	m2.Operation = func(old int64) int64 {
		return old * old
	}
	m2.DivisibleBy = 13
	m2.TrueToMonkey = 1
	m2.FalseToMonkey = 3

	var m3 Monkey
	m3.Items = []int64{74}
	m3.Operation = func(old int64) int64 {
		return old + 3
	}
	m3.DivisibleBy = 17
	m3.TrueToMonkey = 0
	m3.FalseToMonkey = 1

	monkeys := make([]Monkey, 0)
	monkeys = append(monkeys, m0)
	monkeys = append(monkeys, m1)
	monkeys = append(monkeys, m2)
	monkeys = append(monkeys, m3)

	return monkeys
}

func buildMonkeys() []Monkey {

	// Monkey 0 -----------------------------------------

	var m0 Monkey
	m0.Items = []int64{71, 56, 50, 73}
	m0.Operation = func(old int64) int64 {
		return old * 11
	}
	m0.DivisibleBy = 13
	m0.TrueToMonkey = 1
	m0.FalseToMonkey = 7

	// Monkey 1 -----------------------------------------

	var m1 Monkey
	m1.Items = []int64{70, 89, 82}
	m1.Operation = func(old int64) int64 {
		return old + 1
	}
	m1.DivisibleBy = 7
	m1.TrueToMonkey = 3
	m1.FalseToMonkey = 6

	// Monkey 2 -----------------------------------------

	var m2 Monkey
	m2.Items = []int64{52, 95}
	m2.Operation = func(old int64) int64 {
		return old * old
	}
	m2.DivisibleBy = 3
	m2.TrueToMonkey = 5
	m2.FalseToMonkey = 4

	// Monkey 3 -----------------------------------------

	var m3 Monkey
	m3.Items = []int64{94, 64, 69, 87, 70}
	m3.Operation = func(old int64) int64 {
		return old + 2
	}
	m3.DivisibleBy = 19
	m3.TrueToMonkey = 2
	m3.FalseToMonkey = 6

	// Monkey 4 -----------------------------------------

	var m4 Monkey
	m4.Items = []int64{98, 72, 98, 53, 97, 51}
	m4.Operation = func(old int64) int64 {
		return old + 6
	}
	m4.DivisibleBy = 5
	m4.TrueToMonkey = 0
	m4.FalseToMonkey = 5

	// Monkey 5 -----------------------------------------

	var m5 Monkey
	m5.Items = []int64{79}
	m5.Operation = func(old int64) int64 {
		return old + 7
	}
	m5.DivisibleBy = 2
	m5.TrueToMonkey = 7
	m5.FalseToMonkey = 0

	// Monkey 6 -----------------------------------------

	var m6 Monkey
	m6.Items = []int64{77, 55, 63, 93, 66, 90, 88, 71}
	m6.Operation = func(old int64) int64 {
		return old * 7
	}
	m6.DivisibleBy = 11
	m6.TrueToMonkey = 2
	m6.FalseToMonkey = 4

	// Monkey 7 -----------------------------------------

	var m7 Monkey
	m7.Items = []int64{54, 97, 87, 70, 59, 82, 59}
	m7.Operation = func(old int64) int64 {
		return old + 8
	}
	m7.DivisibleBy = 17
	m7.TrueToMonkey = 1
	m7.FalseToMonkey = 3

	monkeys := make([]Monkey, 0)
	monkeys = append(monkeys, m0)
	monkeys = append(monkeys, m1)
	monkeys = append(monkeys, m2)
	monkeys = append(monkeys, m3)
	monkeys = append(monkeys, m4)
	monkeys = append(monkeys, m5)
	monkeys = append(monkeys, m6)
	monkeys = append(monkeys, m7)

	return monkeys

}
