package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func Aufgabe8() {

	inputFile, err := os.Getwd()
	inputFile = inputFile + "/day8/input.txt"
	CheckError(err)

	file, err := os.Open(inputFile)
	CheckError(err)

	br := bufio.NewReader(file)

	for {
		b, err := br.ReadByte()

		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}

		if b == 13 {
			fmt.Println("xx")
		}
		if b == 10 {
			fmt.Println("Umbruch")
		}
		//fmt.Println(strconv.Atoi(string(b)))
		fmt.Println(b)

		if err != nil {
			// end of file
			break
		}
	}

}
