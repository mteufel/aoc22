package day12

import (
	"fmt"
	"github.com/aybabtme/rgbterm"
)

type Terrain struct {
	Data   [][]Node
	Graph  [][]Node
	Length int
	Height int
}

func (t *Terrain) BuildGraph() {
	var graph [][]Node
	var nodes []Node
	var x, y int

	graph = make([][]Node, 0)

	for x < t.Height {
		for y < t.Length {

			node := t.GetNode(x, y)
			nodes = make([]Node, 0)

			var nodeUp, nodeDown, nodeLeft, nodeRight Node
			var checkUp, checkDown, checkLeft, checkRight bool

			checkUp, nodeUp = t.CheckAndGetNode(x-1, y, node.ToNumber())
			checkDown, nodeDown = t.CheckAndGetNode(x+1, y, node.ToNumber())
			checkLeft, nodeLeft = t.CheckAndGetNode(x, y-1, node.ToNumber())
			checkRight, nodeRight = t.CheckAndGetNode(x, y+1, node.ToNumber())

			if checkUp {
				nodes = append(nodes, nodeUp)
			}
			if checkDown {
				nodes = append(nodes, nodeDown)
			}
			if checkLeft {
				nodes = append(nodes, nodeLeft)
			}
			if checkRight {
				nodes = append(nodes, nodeRight)
			}
			graph = append(graph, nodes)

			y++
		}

		x++
		y = 0
	}
	t.Graph = graph
}

func (t *Terrain) GetNode(x int, y int) Node {
	if x < 0 || y < 0 || x >= t.Height || y >= t.Length {
		return Node{X: -1, Y: -1}
	}
	return t.Data[x][y]
}

func (t *Terrain) CheckAndGetNode(x int, y int, elevation int) (bool, Node) {
	node := t.GetNode(x, y)
	if node.X > -1 && node.Y > -1 {

		// this is hip
		//elevation+1 >= node.ToNumber() :)

		// this is easier to understand
		if elevation == node.ToNumber() ||
			elevation+1 == node.ToNumber() ||
			node.ToNumber() < elevation {
			return true, node
		}

	}
	return false, Node{X: -1, Y: -1}
}

/*
func (t *Terrain) FindPosition(position string) Point {
	var x, y int = 0, 0
	for x < t.Height {
		for y < t.Length {
			if t.Data[x][y].Height == position {
				return Point{X: x, Y: y}
			}
			y++
		}
		x++
		y = 0
	}
	return Point{X: -1, Y: -1}
}
*/

func (t *Terrain) PrintTerrain() {
	var x, y int = 0, 0
	for x < t.Height {
		for y < t.Length {
			t.Data[x][y].PrintWithColor()
			y++
		}
		fmt.Println()
		x++
		y = 0
	}
}

type Node struct {
	X      int
	Y      int
	Height string
}

func NewNode(height string, x int, y int) Node {
	return Node{Height: height, X: x, Y: y}
}

func (n *Node) ToNumber() int {
	var height int = 0
	switch {
	case n.Height == "X":
		height = 0
	case n.Height == "S":
		height = 1
	case n.Height == "E":
		height = 26
	case n.Height == "a":
		height = 1
	case n.Height == "b":
		height = 2
	case n.Height == "c":
		height = 3
	case n.Height == "d":
		height = 4
	case n.Height == "e":
		height = 5
	case n.Height == "f":
		height = 6
	case n.Height == "g":
		height = 7
	case n.Height == "h":
		height = 8
	case n.Height == "i":
		height = 9
	case n.Height == "j":
		height = 10
	case n.Height == "k":
		height = 11
	case n.Height == "l":
		height = 12
	case n.Height == "m":
		height = 13
	case n.Height == "n":
		height = 14
	case n.Height == "o":
		height = 15
	case n.Height == "p":
		height = 16
	case n.Height == "q":
		height = 17
	case n.Height == "r":
		height = 18
	case n.Height == "s":
		height = 19
	case n.Height == "t":
		height = 20
	case n.Height == "u":
		height = 21
	case n.Height == "v":
		height = 22
	case n.Height == "w":
		height = 23
	case n.Height == "x":
		height = 24
	case n.Height == "y":
		height = 25
	case n.Height == "z":
		height = 26
	}
	return height
}

func (n *Node) PrintWithColor() {
	var out string
	switch {
	case n.Height == "S":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 255, 0, 0))
	case n.Height == "E":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 255, 0, 0))
	case n.Height == "a":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 80, 80, 80))
	case n.Height == "b":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 100, 100, 100))
	case n.Height == "c":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 110, 110, 110))
	case n.Height == "d":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 120, 120, 120))
	case n.Height == "e":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 130, 130, 130))
	case n.Height == "f":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 140, 140, 140))
	case n.Height == "g":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 145, 145, 145))
	case n.Height == "h":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 150, 150, 150))
	case n.Height == "i":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 155, 155, 155))
	case n.Height == "j":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 160, 160, 160))
	case n.Height == "k":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 165, 165, 165))
	case n.Height == "l":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 170, 170, 170))
	case n.Height == "m":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 175, 175, 175))
	case n.Height == "n":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 180, 180, 180))
	case n.Height == "o":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 185, 185, 185))
	case n.Height == "p":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 190, 190, 190))
	case n.Height == "q":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 195, 195, 195))
	case n.Height == "r":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 200, 200, 200))
	case n.Height == "s":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 205, 205, 205))
	case n.Height == "t":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 210, 210, 210))
	case n.Height == "u":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 215, 215, 215))
	case n.Height == "v":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 220, 220, 220))
	case n.Height == "w":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 225, 225, 225))
	case n.Height == "x":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 240, 240, 240))
	case n.Height == "y":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 245, 245, 245))
	case n.Height == "z":
		fmt.Print(rgbterm.String(n.Height, 0, 0, 0, 255, 255, 255))

	default:
		fmt.Print(n.Height)
	}
	fmt.Print(out)
}
