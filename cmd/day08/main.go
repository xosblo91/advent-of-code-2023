package main

import (
	"os"
	"strings"
)

type set struct {
	left  string
	right string
}

func part1(nodes map[string]set, instructions []string) int {
	return find1(nodes, instructions)
}

func part2(nodes map[string]set, instructions []string) int {
	startingNodes := make([]string, 0)
	for key := range nodes {
		if strings.HasSuffix(key, "A") {
			startingNodes = append(startingNodes, key)
		}
	}

	stepsTaken := make([]int, 0)
	for _, currentNode := range startingNodes {
		steps := find2(currentNode, nodes, instructions)
		stepsTaken = append(stepsTaken, steps)
	}

	return lcm(stepsTaken)
}

func find1(nodes map[string]set, instructions []string) int {
	count := 0
	first := "AAA"

	for {
		current := first
		for _, i := range instructions {
			count++
			node := nodes[current]
			if i == "L" {
				current = node.left
			}
			if i == "R" {
				current = node.right
			}
			if current == "ZZZ" {
				return count
			}
		}

		first = current
	}
}

func find2(start string, nodes map[string]set, instructions []string) int {
	count := 0

	for {
		current := start
		for _, i := range instructions {
			count++
			node := nodes[current]
			if i == "L" {
				current = node.left
			}
			if i == "R" {
				current = node.right
			}
			if strings.HasSuffix(current, "Z") {
				return count
			}
		}

		start = current
	}
}

/*    Don't really understand this    */
func lcm(steps []int) int {
	if len(steps) == 2 {
		return result(steps[0], steps[1])
	}

	var step = steps[0]
	return result(step, lcm(steps[1:]))
}

func result(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}

	return a
}

/***************************************/

func readInput() (map[string]set, []string) {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic("fuck")
	}

	everything := strings.Split(string(f), "\r\n")
	instructions := make([]string, 0)
	for _, s := range everything[0] {
		instructions = append(instructions, string(s))
	}

	nodeStrings := everything[2:]

	nodes := make(map[string]set)
	for _, s := range nodeStrings {
		nodes[s[0:3]] = set{
			left:  s[7:10],
			right: s[12:15],
		}
	}

	return nodes, instructions
}
