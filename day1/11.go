package day1

import (
	"aicoc22/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Task11() {

	inputFile, err := os.Getwd()
	inputFile = inputFile + "/data/day1/input.txt"
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
	for _, line := range fileLines {
		if line == "" {
			if calories > maxCalories {
				maxCalories = calories
			}
			calories = 0
		} else {
			c, err := strconv.Atoi(line)
			util.CheckError(err)
			calories = calories + c
			fmt.Println("-> " + strconv.Itoa(calories))
		}
	}
	fmt.Println("Max Calories: " + strconv.Itoa(maxCalories))

}
