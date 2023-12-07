package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Execute(filePath string) {
	fmt.Println("Day 5")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var seeds []int

	for scanner.Scan() {
		line := scanner.Text()
		if line != "\n" && len(line) > 5 {
			if line[0:5] == "seeds" {
				// we're at the first line, want to get the seed numbers
				seeds = getSeedNumbers(line)
			} else {
				//pass seeds as a slice so it can be updated
				readMap(scanner, seeds[:])
			}
		}

	}
	fmt.Println(seeds)
}

func readMap(input *bufio.Scanner, seeds []int) {
	for i, seed := range seeds {
		seeds[i] = seed + 1
	}
}

func getSeedNumbers(input string) []int {
	var seeds []int

	data := strings.Split(input, ":")

	//seed info is in the second element
	seedStrings := strings.Fields(data[1])

	for _, seed := range seedStrings {
		seedNum, err := strconv.Atoi(seed)
		if err != nil {
			log.Fatal("Could not conver seed number to int")
		}
		seeds = append(seeds, seedNum)
	}

	return seeds
}
