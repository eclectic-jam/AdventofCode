package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var numbers = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func Execute(filePath string) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int

	for scanner.Scan() {
		line := scanner.Text()

		var first string
		var last string
		for i, char := range line {
			if unicode.IsDigit(char) {
				if first == "" {
					first = string(char)
				}
				last = string(char)
			} else {
				for j, num := range numbers {
					if len(num) <= len(line)-i {
						substr := line[i : i+len(num)]
						if substr == num {
							if first == "" {
								first = strconv.Itoa(j + 1)
							}
							last = strconv.Itoa(j + 1)
						}
					}

				}
			}
		}
		numeral := first + last
		number, err := strconv.Atoi(numeral)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %d\n", line, number)
		total += number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(total)
}
