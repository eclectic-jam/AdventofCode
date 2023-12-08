package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

func (n *Node) ToString() string {
	retVal := ""
	retVal += n.left
	retVal += " "
	retVal += n.right
	return retVal
}

func Execute(filePath string) {
	fmt.Println("Day 8")

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("could not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var dirs string

	nodeMap := make(map[string]*Node)

	var nodes []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if dirs == "" && (string(line[0]) == "R" || string(line[0]) == "L") {
			dirs = line
			//Skip the empty line
		} else {
			treeData := strings.Split(line, "=")
			nodeValue := strings.Trim(treeData[0], " ")
			if string(nodeValue[2]) == "A" {
				nodes = append(nodes, nodeValue)
			}

			children := strings.Split(strings.Trim(treeData[1], " "), ",")
			leftNode := strings.Trim(children[0], "( ")
			rightNode := strings.Trim(children[1], ") ")
			nodeMap[nodeValue] = &Node{left: leftNode, right: rightNode}
		}
	}
	dirLength := len(dirs)
	var steps []int
	for _, n := range nodes {
		step := 0
		nextNode := n
		for string(nextNode[2]) != "Z" {
			if string(dirs[step%dirLength]) == "L" {

				nextNode = nodeMap[nextNode].left

			} else {
				nextNode = nodeMap[nextNode].right
			}
			step++
		}
		steps = append(steps, step)
	}

	if len(steps) >= 2 {
		fmt.Println(lcm(steps...))
	}

}

func lcm(numbers ...int) int {
	result := numbers[0] * numbers[1] / gcd(numbers[0], numbers[1])

	for i := 2; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
