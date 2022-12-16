package task102

import (
	"aicoc22/util"
	"fmt"
	"strconv"
	"strings"
)

func Task102() {
	fmt.Println("Day 10, Part 2: Cathode-Ray Tube")
	fmt.Println("-------------------------------------")

	programm := util.ReadFile("input3.txt")

	cycleCounter := 0
	programCounter := 0
	totalStrength := 0
	strength := 0
	x := 1
	crt := "........................................" +
		"........................................" +
		"........................................" +
		"........................................" +
		"........................................" +
		"........................................"

	for {
		cmd, val := parse(programm[programCounter])

		switch {
		case cmd == "noop":
			crt = setPixel(cycleCounter, x, crt)
			cycleCounter, strength = increaseCycle(cycleCounter, x)
			totalStrength = totalStrength + strength
			fmt.Printf("Executed: %s [%d] Cyle=%d X=%d strength=%d totalStrength=%d\n", cmd, val, cycleCounter, x, strength, totalStrength)
			crt = setPixel(cycleCounter, x, crt)
		case cmd == "addx":
			crt = setPixel(cycleCounter, x, crt)
			cycleCounter, strength = increaseCycle(cycleCounter, x)
			totalStrength = totalStrength + strength
			fmt.Printf("Executed: %s [%d] Cyle=%*d X=%d strength=%d totalStrength=%d\n", cmd, val, cycleCounter, x, strength, totalStrength)
			crt = setPixel(cycleCounter, x, crt)
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
	fmt.Println(crt[0:39])
	fmt.Println(crt[40:79])
	fmt.Println(crt[80:119])
	fmt.Println(crt[120:159])
	fmt.Println(crt[160:199])
	fmt.Println(crt[200:239])

}

func setPixel(cycle int, x int, crt string) string {
	raster := "........................................"
	raster = replaceAtIndexRaster(raster, '#', x-1)
	raster = replaceAtIndexRaster(raster, '#', x)
	raster = replaceAtIndexRaster(raster, '#', x+1)
	if string(raster[cycle%40]) == "#" {
		crt = replaceAtIndex(crt, '#', cycle)
	}
	return crt
}
func replaceAtIndexRaster(str string, replacement rune, index int) string {
	if index >= 0 && index <= 39 {
		return str[:index] + string(replacement) + str[index+1:]
	}
	return str
}

func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
	return str
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
