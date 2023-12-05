package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ExecutePart2(filePath string) {
	fmt.Println("Day 4 Part 2")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var cards [204]int
	for i := range cards {
		cards[i] = 1
	}
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()

		card := strings.Split(line, ":")

		if len(card) != 2 {
			log.Fatal("Not a valid card")
		}

		gameData := card[1]

		numberGroups := strings.Split(gameData, "|")
		winners := strings.Fields(strings.Trim(numberGroups[0], " "))
		myNumbers := strings.Fields(strings.Trim(numberGroups[1], " "))

		score := 0

		for _, num := range myNumbers {
			for _, win := range winners {
				if num == win {
					// winning number!
					score += 1
				}
			}
		}

		for i := lineNumber + 1; i < lineNumber+score+1; i++ {

			cards[i] = cards[i] + cards[lineNumber]

		}

		lineNumber += 1
	}
	total := 0
	for _, val := range cards {
		total += val
	}

	fmt.Print(cards)
	fmt.Print(total)
}

func ExecutePart1(filePath string) {
	fmt.Println("Day 4 Part 2")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		cardData := strings.Split(line, ":")

		if len(cardData) != 2 {
			log.Fatal("invalide card data")
		}

		game := cardData[1]

		numberGroups := strings.Split(game, "|")
		winners := strings.Fields(strings.Trim(numberGroups[0], " "))
		myNumbers := strings.Fields(strings.Trim(numberGroups[1], " "))
		fmt.Println(winners)
		fmt.Println(myNumbers)
		score := 0
		for _, num := range myNumbers {
			for _, win := range winners {
				if num == win {
					// winning number!
					if score == 0 {
						score = 1
					} else {
						score = score * 2
					}
				}
			}
		}
		total += score
	}

	fmt.Printf("Score: %d", total)
}
