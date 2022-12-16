package task101

import (
	"aicoc22/util"
	"fmt"
	"strconv"
	"strings"
)

func Task101() {
	fmt.Println("Day 10, Part 1: Cathode-Ray Tube")
	fmt.Println("-------------------------------------")

	programm := util.ReadFile("input3.txt")

	cycleCounter := 0
	programCounter := 0
	totalStrength := 0
	strength := 0
	x := 1
	for {
		cmd, val := parse(programm[programCounter])

		switch {
		case cmd == "noop":
			cycleCounter, strength = increaseCycle(cycleCounter, x)
			totalStrength = totalStrength + strength
			fmt.Printf("Executed: %s [%d] Cyle=%d X=%d strength=%d totalStrength=%d\n", cmd, val, cycleCounter, x, strength, totalStrength)
		case cmd == "addx":
			cycleCounter, strength = increaseCycle(cycleCounter, x)
			totalStrength = totalStrength + strength
			fmt.Printf("Executed: %s [%d] Cyle=%d X=%d strength=%d totalStrength=%d\n", cmd, val, cycleCounter, x, strength, totalStrength)
			cycleCounter, strength = increaseCycle(cycleCounter, x)
			totalStrength = totalStrength + strength
			x = x + val
			fmt.Printf("Executed: %s [%d] Cyle=%d X=%d strength=%d totalStrength=%d\n", cmd, val, cycleCounter, x, strength, totalStrength)
		}

		programCounter++
		if programCounter >= len(programm) {
			break
		}
	}

}

func increaseCycle(cycle int, x int) (int, int) {
	// returns new Cycle, calculated strength
	cycle++
	strength := 0
	if cycle == 20 {
		strength = cycle * x
	} else {
		if (cycle-20)%40 == 0 {
			strength = cycle * x
		}
	}

	return cycle, strength
}

func parse(operation string) (string, int) {
	parsed := strings.Fields(operation)
	cmd := parsed[0]
	if len(parsed) > 1 {
		val, err := strconv.Atoi(parsed[1])
		util.CheckError(err)
		return cmd, val
	} else {
		return cmd, 0
	}
}
