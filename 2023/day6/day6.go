package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func (r *Race) Print() {
	fmt.Printf("time: %d, distance: %d\n", r.time, r.distance)
}

func (r *Race) GetWinners() int {
	var waysToWin int

	for i := 1; i < r.time-1; i++ {
		if i*(r.time-i) > r.distance {
			waysToWin++
		}
	}

	return waysToWin

}

func Execute(filePath string) {
	fmt.Println("Day 6")
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	race := &Race{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")

		var num string

		for _, val := range strings.Fields(parts[1]) {
			num += val
		}
		if race.time == 0 {
			race.time = getInt(num)
		} else {
			race.distance = getInt(num)
		}
	}

	fmt.Println(race.GetWinners())

}

func getInt(number string) int {
	num, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal("unable to convert string to int")
	}
	return num
}
