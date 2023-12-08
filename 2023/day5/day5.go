package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start  int
	length int
}

func (r *Range) Print() {
	fmt.Printf("start: %d, length: %d\n", r.start, r.length)
}

func Execute(filePath string) {
	fmt.Println("Day 5")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var seedRanges []*Range

	for scanner.Scan() {
		line := scanner.Text()
		if line != "\n" && len(line) > 5 {
			if line[0:5] == "seeds" {
				// we're at the first line, want to get the seed numbers
				//seeds = getSeedNumbers(line)
				seedRanges = getSeedNumbersPart2(line)
				for _, r := range seedRanges {
					r.Print()
				}
			} else {
				//pass seeds as a slice so it can be updated
				seedRanges = readMapAsRange(scanner, seedRanges[:])
			}
		}

	}
	//fmt.Println(minVal(seeds))
}

func readMap(scanner *bufio.Scanner, seeds []int) []int {
	var found []int
	found = append(found, seeds...)
	scanner.Scan()

	for len(scanner.Text()) != 0 {
		line := scanner.Text()
		nums := strings.Fields(line)

		var sourceStart, destStart, length int

		destStart = getInt(nums[0])
		sourceStart = getInt(nums[1])
		length = getInt(nums[2])

		for i, val := range seeds {
			if val >= sourceStart && val < sourceStart+length {
				// add the new value
				found[i] = destStart + val - sourceStart
			}
		}
		scanner.Scan()
	}

	return found

}

func readMapAsRange(scanner *bufio.Scanner, seeds []*Range) []*Range {
	var found []*Range
	scanner.Scan()

	for len(scanner.Text()) != 0 {
		line := scanner.Text()
		nums := strings.Fields(line)
		fmt.Println(nums)
		/*
			var sourceStart, destStart, length int

			destStart = getInt(nums[0])
			sourceStart = getInt(nums[1])
			length = getInt(nums[2])

			for i, val := range seeds {
				if val >= sourceStart && val < sourceStart+length {
					// add the new value
					found[i] = destStart + val - sourceStart
				}
			}*/
		scanner.Scan()
	}

	return found
}

func minVal(arr []int) int {
	minVal := 0
	for _, num := range arr {
		if num < minVal || minVal == 0 {
			minVal = num
		}
	}
	return minVal
}

func getInt(number string) int {
	num, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal("unable to convert string to int")
	}
	return num
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

func getSeedNumbersPart2(input string) []*Range {
	var seeds []*Range

	data := strings.Split(input, ":")

	//seed info is in the second element
	seedStrings := strings.Fields(data[1])

	if len(seedStrings)%2 != 0 {
		log.Fatal("even number of elements required")
	}

	for i := 0; i < len(seedStrings); i = i + 2 {
		number := getInt(seedStrings[i])
		nextNumber := getInt(seedStrings[i+1])

		seedRange := &Range{start: number, length: nextNumber}

		seeds = append(seeds, seedRange)
	}

	return seeds
}
