package day10

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

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
				y = lineNumber
				x = i
			}
			newRow = append(newRow, string(r))
		}

		pipes = append(pipes, newRow)
		lineNumber++
	}
	prevX := x
	prevY := y
	step := 0
	for step == 0 || pipes[y][x] != "S" {
		if step == 0 {
			newDir, err := findDirections(x, y, pipes)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(newDir)
			x = newDir[0][1]
			y = newDir[0][0]

			step++
		} else {
			newX, newY, err := findSpecificDirection(pipes[y][x], x, y, prevX, prevY, pipes)
			if err != nil {
				log.Fatal(err)
			}
			prevX = x
			prevY = y
			x = newX
			y = newY
			step++
		}
	}
	fmt.Println(step / 2)
}

func findSpecificDirection(pipe string, x, y, prevX, prevY int, pipes [][]string) (int, int, error) {
	if pipe == "|" {
		//check north and south
		if prevY != y-1 && checkNorth(x, y, pipes) {
			return x, y - 1, nil
		}
		if prevY != y+1 && checkSouth(x, y, pipes) {
			return x, y + 1, nil
		}
	}
	if pipe == "-" {
		//check east and west
		if prevX != x+1 && checkEast(x, y, pipes) {
			return x + 1, y, nil
		}
		if prevX != x-1 && checkWest(x, y, pipes) {
			return x - 1, y, nil
		}
	}
	if pipe == "L" {
		// check north and east
		if prevY != y-1 && checkNorth(x, y, pipes) {
			return x, y - 1, nil
		}
		if prevX != x+1 && checkEast(x, y, pipes) {
			return x + 1, y, nil
		}
	}
	if pipe == "J" {
		// check North and west
		if prevY != y-1 && checkNorth(x, y, pipes) {
			return x, y - 1, nil
		}
		if prevX != x-1 && checkWest(x, y, pipes) {
			return x - 1, y, nil
		}
	}
	if pipe == "7" {
		// check south and west
		if prevY != y+1 && checkSouth(x, y, pipes) {
			return x, y + 1, nil
		}
		if prevX != x-1 && checkWest(x, y, pipes) {
			return x - 1, y, nil
		}
	}
	if pipe == "F" {
		// check south and east
		if prevY != y+1 && checkSouth(x, y, pipes) {
			return x, y + 1, nil
		}
		if prevX != x+1 && checkEast(x, y, pipes) {
			return x + 1, y, nil
		}
	}

	return 0, 0, errors.New("Could not find direction")
}

func checkNorth(x, y int, pipes [][]string) bool {
	if isValidCell(x, y-1, pipes) {
		pipe := pipes[y-1][x]
		if pipe == "S" || pipe == "|" || pipe == "F" || pipe == "7" {
			return true
		}
	}
	return false
}
func checkSouth(x, y int, pipes [][]string) bool {
	if isValidCell(x, y+1, pipes) {
		pipe := pipes[y+1][x]
		if pipe == "S" || pipe == "|" || pipe == "J" || pipe == "L" {
			return true
		}
	}
	return false
}
func checkEast(x, y int, pipes [][]string) bool {
	if isValidCell(x+1, y, pipes) {
		pipe := pipes[y][x+1]
		if pipe == "S" || pipe == "-" || pipe == "J" || pipe == "7" {
			return true
		}
	}
	return false
}

func checkWest(x, y int, pipes [][]string) bool {
	if isValidCell(x-1, y, pipes) {
		pipe := pipes[y][x-1]
		if pipe == "S" || pipe == "-" || pipe == "F" || pipe == "L" {
			return true
		}
	}
	return false
}
func findDirections(x, y int, pipes [][]string) ([][]int, error) {
	// x and y are the coordinates of our starting position
	var newDirs [][]int
	// check west
	if isValidCell(x-1, y, pipes) && (pipes[y][x-1] == "-" || pipes[y][x-1] == "L" || pipes[y][x-1] == "F") {
		newDirs = append(newDirs, []int{y, x - 1})
	}
	// check east
	if isValidCell(x+1, y, pipes) && (pipes[y][x+1] == "-" || pipes[y][x+1] == "J" || pipes[y][x+1] == "7") {
		newDirs = append(newDirs, []int{y, x + 1})
	}
	// check north
	if isValidCell(x, y-1, pipes) && (pipes[y-1][x] == "|" || pipes[y-1][x] == "7" || pipes[y-1][x] == "F") {

		newDirs = append(newDirs, []int{y - 1, x})
	}
	// check south
	if isValidCell(x, y+1, pipes) && (pipes[y+1][x] == "|" || pipes[y+1][x] == "J" || pipes[y+1][x] == "L") {
		newDirs = append(newDirs, []int{y + 1, x})
	}
	if len(newDirs) == 0 {

		return nil, errors.New("Could not find starting direction")
	}
	return newDirs, nil

}

func isValidCell(x, y int, pipes [][]string) bool {
	height := len(pipes)
	length := len(pipes[0])

	if y < 0 || y >= height {
		return false
	}

	if x < 0 || x >= length {
		return false
	}
	return true
}
