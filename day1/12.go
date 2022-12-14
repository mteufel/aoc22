package day1

import (
	"aicoc22/util"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Task12() {
	inputFile, err := os.Getwd()
	inputFile = inputFile + "/data/day1/input2.txt"
	util.CheckError(err)

	readFile, err := os.Open(inputFile)
	util.CheckError(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	var calories int = 0
	var maxCalories int = 0

	var sorted []int

	for _, line := range fileLines {
		if line == "" {
			if calories > maxCalories {
				maxCalories = calories
			}
			sorted = append(sorted, calories)
			calories = 0
		} else {
			c, err := strconv.Atoi(line)
			util.CheckError(err)
			calories = calories + c
		}
	}
	sorted = append(sorted, calories)
	fmt.Println("Max Calories Aufgabe 1=" + strconv.Itoa(maxCalories))
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	fmt.Println(sorted)
	fmt.Println("Result Aufgabe 2=" + strconv.Itoa(sorted[0]+sorted[1]+sorted[2]))
}
