package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	FOC  = 7
	FrOC = 6
	FH   = 5
	TOC  = 4
	TP   = 3
	OP   = 2
	HC   = 1
)

var cardValue = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 0,
	"Q": 10,
	"K": 11,
	"A": 12,
}

type Hand struct {
	hand     string
	bet      int
	handType int
	score    int
}

func (h *Hand) Print() {
	fmt.Printf("Hand: %s, bet: %d, type: %d\n", h.hand, h.bet, h.handType)
}

func (h *Hand) Compare(c *Hand) int {
	if h.handType > c.handType {
		return 1
	}
	if h.handType < c.handType {
		return -1
	}
	//hand types are the same
	for i := range h.hand {
		if cardValue[string(h.hand[i])] > cardValue[string(c.hand[i])] {
			return 1
		}
		if cardValue[string(h.hand[i])] < cardValue[string(c.hand[i])] {
			return -1
		}
	}
	return 0
}

func (h *Hand) GetType() {
	hand := make(map[string]int)
	maxStringCount := 0
	maxString := ""
	for _, r := range h.hand {
		hand[string(r)] += 1
		if hand[string(r)] > maxStringCount && string(r) != "J" {
			maxString = string(r)
			maxStringCount = hand[string(r)]
		}
	}
	if _, ok := hand["J"]; ok {
		hand[maxString] += hand["J"]
		delete(hand, "J")
	}
	for _, n := range hand {
		if n == 5 {
			h.handType = FOC
		} else if n == 4 {
			h.handType = FrOC
		} else if n == 3 {
			if h.handType == OP {
				h.handType = FH
			} else {
				h.handType = TOC
			}
		} else if n == 2 {
			if h.handType == TOC {
				h.handType = FH
			} else if h.handType == OP {
				h.handType = TP
			} else {
				h.handType = OP
			}
		} else if n == 1 && h.handType == 0 {
			h.handType = HC
		}
	}
}

func Execute(filePath string) {
	fmt.Println("Day 7")

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("could not open file %s", filePath)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hands []*Hand

	for scanner.Scan() {
		line := scanner.Text()
		hand := &Hand{}

		data := strings.Fields(line)

		if len(data) != 2 {
			log.Fatal("should be 2 elements")
		}

		hand.hand = data[0]
		hand.bet = getInt(data[1])
		hand.GetType()

		hands = append(hands, hand)
	}

	sortedHands := QuickSortHands(hands)
	totalScore := 0
	for i, h := range sortedHands {
		h.score = h.bet * (i + 1)
		totalScore += h.score
	}

	fmt.Printf("Score: %d\n", totalScore)
}

func QuickSortHands(hands []*Hand) []*Hand {
	if len(hands) == 1 || len(hands) == 0 {
		return hands
	}
	pivot := hands[0]
	var g1 []*Hand
	var g2 []*Hand
	for i := 1; i < len(hands); i++ {
		if hands[i].Compare(pivot) >= 0 {
			g2 = append(g2, hands[i])
		} else {
			g1 = append(g1, hands[i])
		}
	}
	g2 = QuickSortHands(g2)
	g1 = QuickSortHands(g1)
	retArray := append(g1, pivot)
	return append(retArray, g2...)
}

func getInt(number string) int {
	num, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal("unable to convert string to int")
	}
	return num
}
