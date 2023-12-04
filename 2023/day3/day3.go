package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var symbols string = "!@#$%^&*"

var schematic []string
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

		schematic = append(schematic, strings.Split(line, "")...)
	}

	fmt.Println(schematic)

}
