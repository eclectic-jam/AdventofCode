package day2

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numOfCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Execute(filePath string) {
	fmt.Println("Day 2")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int
	var powerTotal int

	//var idSum int
	for scanner.Scan() {
		possible := true
		line := scanner.Text()
		game := strings.Split(line, ":")
		if len(game) != 2 {
			log.Fatal("Game does not contain correct data")
		}
		gameNumber, err := getGameNumber(game[0])
		fmt.Printf("Game %d", gameNumber)

		if err != nil {
			log.Fatal("unable to retrieve game number")
		}

		powerTotal += minCubes(game[1])

		cubeSets := strings.Split(game[1], ";")

		for _, set := range cubeSets {
			if !analyzeGame(set) {
				possible = false
			}
		}

		if possible {
			total += gameNumber
		}

	}
	fmt.Println("total: ", total)
	fmt.Println("powerTotal: ", powerTotal)
}

// go throught game data and get the minimum number of cubes to make it a valid game.
// Return those numbers multiplied together
func minCubes(data string) int {
	fmt.Println(": ", data)

	var minimum = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	draws := strings.Split(data, ";")

	for _, cubes := range draws {
		colors := strings.Split(cubes, ",")

		for _, color := range colors {
			values := strings.Split(strings.Trim(color, " "), " ")

			if len(values) != 2 {
				log.Fatal("invalid draw")
			}

			value, err := strconv.Atoi(values[0])
			if err != nil {
				log.Fatal("number of cubes is not an int")
			}

			color := values[1]

			if minimum[color] < value {
				minimum[color] = value
			}
		}

	}
	fmt.Println(minimum)

	return minimum["red"] * minimum["green"] * minimum["blue"]
}

// go through game data and determine whether it is a valid game
func analyzeGame(data string) bool {
	colors := strings.Split(data, ",")

	for _, cube := range colors {
		drawn := strings.Split(strings.Trim(cube, " "), " ")
		if len(drawn) != 2 {
			log.Fatal("Invalid cube draw")
		}

		value, err := strconv.Atoi(drawn[0])
		if err != nil {
			log.Fatal("number of cubes is not an int")
		}
		color := drawn[1]
		if value > numOfCubes[color] {
			return false
		}
	}

	return true
}

// split the game name into two strings, "game" and the number itself
// return the number as an int
func getGameNumber(gameName string) (int, error) {
	words := strings.Split(gameName, " ")
	if len(words) != 2 {
		return 0, errors.New("invalid format of game label")
	}

	gameNumber, err := strconv.Atoi(words[1])

	if err != nil {
		return 0, errors.New("game label not in form of an Integer")
	}

	return gameNumber, nil

}
