package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Execute(filePath string) {
	fmt.Println("Day 9")

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("Could not open file %s", filePath)
	}

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		var sets [][]int
		newSet := []int{}
		for _, num := range nums {
			newSet = append(newSet, getInt(num))
		}

		sets = append(sets, newSet)

		allZeros := false
		setNum := 1
		newSet = []int{}

		for !allZeros {
			allZeros = true
			for i := 0; i < len(sets[setNum-1])-1; i++ {
				newNum := sets[setNum-1][i+1] - sets[setNum-1][i]
				if newNum != 0 {
					allZeros = false
				}
				newSet = append(newSet, newNum)
			}
			sets = append(sets, newSet)
			newSet = []int{}
			setNum++
		}

		for i := len(sets) - 1; i >= 0; i-- {
			if i == len(sets)-1 {
				sets[i] = append([]int{0}, sets[i]...)
			} else {
				newNum := sets[i][0] - sets[i+1][0]
				if i == 0 {
					total += newNum
				}
				sets[i] = append([]int{newNum}, sets[i]...)
			}
		}

	}
	fmt.Println(total)
}

func getInt(number string) int {
	num, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal("unable to convert string to int")
	}
	return num
}
