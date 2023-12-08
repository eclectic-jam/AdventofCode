package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Execute(filePath string) {
	fmt.Println("Day 8")

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("could not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}

}
