package day11

import (
	"fmt"
	"strconv"
	"strings"
)

type Monkey struct {
	Items         []int64
	Item          int64
	DivisibleBy   int64
	TrueToMonkey  int64
	FalseToMonkey int64
	Operation     OperationDef
	Inspections   int64
}

type OperationDef func(int64) int64

func (m *Monkey) Operate() {
	m.Item = m.Operation(m.Item)
	m.Inspections++
}

func (m *Monkey) ToString(monkeyNumber int) {
	output := make([]string, 0)
	for _, value := range m.Items {
		output = append(output, strconv.FormatInt(value, 10))
	}
	fmt.Println("Monkey " + strconv.Itoa(monkeyNumber) + ": " + strings.Join(output, ", ") + " (" + strconv.FormatInt(m.Inspections, 10) + " Inspections)")
}

func (m *Monkey) AddItem(item int64) {
	m.Items = append(m.Items, item)
}
