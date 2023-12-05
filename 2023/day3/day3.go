package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var symbols string = "!@#$%^&*+-=/"

type directions struct {
	dx int
	dy int
}

var dirs = []directions{
	{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
}

var schematic [][]string
var length int
var lines int

func Execute(filePath string) {
	fmt.Println("Day 3")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		length = len(line)
		lines++

		schematic = append(schematic, strings.Split(line, ""))
	}

	var parts []string
	var gearRatioTotal int

	partNo := ""
	//symbol := ""

	for i := 0; i < lines; i++ {
		for j := 0; j < length; j++ {
			character := schematic[i][j]
			if character == "*" {
				adjParts := findAdjacentParts(i, j, schematic)
				fmt.Println(adjParts)
				if len(adjParts) == 2 {
					gear1, err := strconv.Atoi(adjParts[0])
					if err != nil {
						log.Fatal("Could not convert part number to int")
					}
					gear2, err := strconv.Atoi(adjParts[1])
					if err != nil {
						log.Fatal("Could not convert part number to int")
					}

					gearRatioTotal += gear1 * gear2
				}
				// it's a gear, we want to find adjacent part numbers
			} else if _, err := strconv.Atoi(character); err == nil {
				// it's an int
				partNo = partNo + character
			} else {
				//it's not an int, so end the part number if we've started one
				if partNo != "" {
					//fmt.Printf("partNo: %s at [%d][%d]\n", partNo, i, j)
					if isAdjacentToSymbol(i, j-len(partNo), len(partNo), schematic) {
						parts = append(parts, partNo)
					}
					partNo = ""
				}
			}
			if j == length-1 {
				// we're at the end of the line, so end the part number if we've started one
				if partNo != "" {
					if isAdjacentToSymbol(i, j-len(partNo), len(partNo), schematic) {
						parts = append(parts, partNo)
					}

					partNo = ""
				}

			}
		}
	}
	total := 0
	for _, num := range parts {
		value, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal("could not convert part number to int")
		}

		total += value
	}

	fmt.Printf("parts total: %d\n", total)
	fmt.Printf("gear ratio total: %d\n", gearRatioTotal)

}

func findAdjacentParts(row int, col int, schematic [][]string) []string {
	var parts []string
	n := len(schematic)
	m := len(schematic[0])

	for _, dir := range dirs {
		i := row + dir.dx
		j := col + dir.dy
		if i > -1 && i < n && j > -1 && j < m {
			//fmt.Printf("Checking [%d][%d]\n", row+dir.dx, col+dir.dy)
			if _, err := strconv.Atoi(schematic[i][j]); err == nil {
				//this element is part of a part number
				partNo := completePartNo(j, schematic[i])
				contains := false
				for _, part := range parts {
					if part == partNo {
						contains = true
					}
				}
				if !contains {
					parts = append(parts, partNo)
				}
			}
		}
	}

	return parts
}

func completePartNo(col int, schematicRow []string) string {
	partNo := schematicRow[col]
	for j := col - 1; j >= 0; j-- {
		if _, err := strconv.Atoi(schematicRow[j]); err == nil {
			partNo = schematicRow[j] + partNo
		} else {
			break
		}
	}
	for j := col + 1; j < len(schematicRow); j++ {
		if _, err := strconv.Atoi(schematicRow[j]); err == nil {
			partNo += schematicRow[j]
		} else {
			break
		}
	}
	return partNo
}

func isAdjacentToSymbol(row int, col int, length int, schematic [][]string) bool {
	for i := 0; i < length; i++ {
		if checkSurroundingSymbol(row, col+i, schematic) {
			return true
		}
	}
	return false
}

func checkSurroundingSymbol(row int, col int, schematic [][]string) bool {
	// We want to check
	// above 				(row - 1, col),
	// below 				(row + 1, col),
	// left 				(row		, col - 1),
	// left above 	(row - 1, col - 1) ,
	// left below 	(row + 1, col - 1).
	// right				(row		, col + 1)
	// right above 	(row - 1, col + 1),
	// right below  (row + 1, col + 1)
	n := len(schematic)    // number of rows
	m := len(schematic[0]) // number of columns
	//fmt.Printf("Checking surrounding elements of %s at [%d][%d]\n", schematic[row][col], row, col)
	for _, dir := range dirs {
		i := row + dir.dx
		j := col + dir.dy
		if i > -1 && i < n && j > -1 && j < m {
			//fmt.Printf("Checking [%d][%d]\n", row+dir.dx, col+dir.dy)
			if strings.Contains(symbols, schematic[i][j]) {
				return true
			}
		}
	}

	return false
}
