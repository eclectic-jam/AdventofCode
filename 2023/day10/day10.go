package day10

import (
	"bufio"
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

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}

}
