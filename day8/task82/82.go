package task82

import (
	"aicoc22/util"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Task82() {

	inputFile, err := os.Getwd()
	inputFile = inputFile + "/data/day8/input2.txt"
	util.CheckError(err)

	// Datei zeilenweise einlesen
	file, err := os.ReadFile(inputFile)
	util.CheckError(err)

	lines := strings.Split(string(file), "\n")
	var forest [][]int
	var temp []int

	length := 0
	height := 0

	// den String zeilenweise durchgehen und ein 2d-Slice mit Koordinaten erstellen
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		length = len(l)
		height++
		temp = make([]int, 0)
		for _, value := range l {
			n, err := strconv.Atoi(string(value))
			util.CheckError(err)
			temp = append(temp, n)

		}
		forest = append(forest, [][]int{temp}...)
	}
	length--
	height--
	fmt.Printf("Length: %d Height: %d \n", length, height)
	// komplettes slice ausgeben
	fmt.Println(forest)

	// Durchlauf 2 - im Wald
	var result []int
	for x := 1; x <= (height - 1); x++ {
		for y := 1; y <= (length - 1); y++ {

			if x == 1 && y == 2 {
				fmt.Println("Stop")
			}

			baum := forest[x][y]
			oben := nachOben(forest, baum, 0, x, y)
			unten := nachUnten(forest, baum, height, x, y)
			links := nachLinks(forest, baum, 0, x, y)
			rechts := nachRechts(forest, baum, length, x, y)

			score := oben * links * rechts * unten
			result = append(result, score)

		}
	}
	fmt.Println("Vor Sortierung: ")
	fmt.Println(result)
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	fmt.Println("Nach Sortierung: ")
	fmt.Println(result)
	fmt.Println("Ergebnis: ")
	fmt.Println(result[0])

}

func nachOben(forest [][]int, myBaum int, max int, x int, y int) int {
	result := 0
	for x > max {
		x--
		result++
		if forest[x][y] >= myBaum {
			break
		}
	}
	return result
}

func nachUnten(forest [][]int, myBaum int, max int, x int, y int) int {
	result := 0
	for x < max {
		x++
		result++
		if forest[x][y] >= myBaum {
			break
		}
	}
	return result
}

func nachLinks(forest [][]int, myBaum int, max int, x int, y int) int {
	result := 0
	for y > max {
		y--
		result++
		if forest[x][y] >= myBaum {
			break
		}
	}
	return result
}

func nachRechts(forest [][]int, myBaum int, max int, x int, y int) int {
	result := 0
	for y < max {
		y++
		result++
		if forest[x][y] >= myBaum {
			break
		}
	}
	return result
}
