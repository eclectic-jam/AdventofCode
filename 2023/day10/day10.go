package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type directions struct {
	dx int
	dy int
}

var dirs = []directions{
	{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
}

func Execute(filePath string) {
	fmt.Println("Day 10")

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("unable to open file")
	}

	scanner := bufio.NewScanner(file)

	lineNumber := 0
	var pipes [][]string
	var x, y int

	for scanner.Scan() {
		line := scanner.Text()
		var newRow []string
		for i, r := range line {
			if string(r) == "S" {
				x = lineNumber
				x = i
			}
			newRow = append(newRow, string(r))
		}

		pipes = append(pipes, newRow)
		lineNumber++
	}

	step := 0
	for step == 0 || pipes[x][y] != "S" {
		if step == 0 {

		}
	}

}
