package day12

import (
	"aicoc22/util"
	"fmt"
	"os"
	"strings"
)

func Task121() {

	terrain := initializeTerrain("timo.txt")
	fmt.Printf("length=%d height=%d\n\n", terrain.Length, terrain.Height)
	terrain.PrintTerrain()

	// Build a Graph from the data, can be used for either Djkstra or Breitensuche
	terrain.BuildGraph()

	fmt.Println("End")

}

func djikstra(startX int, startY int, terrain Terrain) {

	// https://de.wikipedia.org/wiki/Dijkstra-Algorithmus

	visited := make([]int, .)
	//var weight []Node
	var active []Node

	active := append(active, terrain.GetNode(startX, startY))
	for len(active) > 0 {

	}

}

func initializeTerrain(fileName string) Terrain {

	inputFile, err := os.Getwd()
	inputFile = inputFile + "/data/day12/" + fileName
	util.CheckError(err)

	// Datei zeilenweise einlesen
	file, err := os.ReadFile(inputFile)
	util.CheckError(err)

	lines := strings.Split(string(file), "\n")
	var data [][]Node
	var temp []Node

	length := 0
	height := 0

	// den String zeilenweise durchgehen und ein 2d-Slice mit Koordinaten erstellen
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		length = len(l)
		height++
		temp = make([]Node, 0)
		for i, v := range l {
			temp = append(temp, NewNode(string(v), height-1, i))

		}
		data = append(data, [][]Node{temp}...)
	}

	var holder Terrain
	holder.Data = data
	holder.Length = length
	holder.Height = height
	return holder
}
