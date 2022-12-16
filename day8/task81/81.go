package task81

import (
	"aicoc22/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Task81() {

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

	// Bäume, außer die an den Rändern
	var resultAussen []int
	// Durchlauf 1 - die Waldgrenze
	for x := 0; x <= height; x++ {
		for y := 0; y <= length; y++ {
			var erg = "nicht sichtbar"
			baum := forest[x][y]
			if x-1 == -1 || x == height {
				erg = "sichtbar"
			}
			if y-1 == -1 || y == length {
				erg = "sichtbar"
			}
			if erg == "sichtbar" {
				resultAussen = append(resultAussen, baum)
			}
		}
	}
	// Durchlauf 2 - im Wald
	var resultInnen []int
	for x := 1; x <= (height - 1); x++ {
		for y := 1; y <= (length - 1); y++ {
			baum := forest[x][y]
			oben := nachOben(forest, baum, 0, x, y)
			unten := nachUnten(forest, baum, height, x, y)
			links := nachLinks(forest, baum, 0, x, y)
			rechts := nachRechts(forest, baum, length, x, y)

			if oben || unten || links || rechts {
				resultInnen = append(resultInnen, baum)
			}

		}
	}
	fmt.Println("Ergebnis Aussen:")
	fmt.Println("------------------------")
	fmt.Println(resultAussen)
	fmt.Println(len(resultAussen))
	fmt.Println("Ergebnis Innen:")
	fmt.Println("------------------------")
	fmt.Println(resultInnen)
	fmt.Println(len(resultInnen))
	resultInnen = append(resultInnen, resultAussen...)
	fmt.Println("Ergebnis Gesamt:")
	fmt.Println("------------------------")
	fmt.Println(len(resultInnen))

}

func nachOben(forest [][]int, myBaum int, max int, x int, y int) bool {
	result := true
	for x > max {
		x--
		if forest[x][y] >= myBaum {
			result = false
			break
		}
	}
	return result
}

func nachUnten(forest [][]int, myBaum int, max int, x int, y int) bool {
	result := true
	for x < max {
		x++
		if forest[x][y] >= myBaum {
			result = false
			break
		}
	}
	return result
}

func nachLinks(forest [][]int, myBaum int, max int, x int, y int) bool {
	result := true
	for y > max {
		y--
		if forest[x][y] >= myBaum {
			result = false
			break
		}
	}
	return result
}

func nachRechts(forest [][]int, myBaum int, max int, x int, y int) bool {
	result := true
	for y < max {
		y++
		if forest[x][y] >= myBaum {
			result = false
			break
		}
	}
	return result
}
