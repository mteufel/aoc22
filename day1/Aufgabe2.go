package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Aufgabe2() {
	inputFile, err := os.Getwd()
	inputFile = inputFile + "/day1/input.txt"
	CheckError(err)

	readFile, err := os.Open(inputFile)
	CheckError(err)
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
			CheckError(err)
			calories = calories + c
			fmt.Println("-> " + strconv.Itoa(calories))
		}
	}
	fmt.Println("Max Calories: " + strconv.Itoa(maxCalories))
	fmt.Println(sorted)
}
