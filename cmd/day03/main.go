package main

import (
	"bufio"
	"os"
	"strconv"
)

func part1(input [][]string) int {
	sum := 0
	adjacentSymbol := false
	combined := ""

	for y, row := range input {
		for x, value := range row {
			r, err := strconv.Atoi(value)
			if err != nil {
				r, _ = strconv.Atoi(combined)
				combined = ""
				if adjacentSymbol {
					sum += r
				}

				adjacentSymbol = false

				continue
			}

			combined += value

			// right
			if x+1 < len(row) {
				rightNeighbor := input[y][x+1]
				if isSymbol(rightNeighbor) {
					adjacentSymbol = true
					continue
				}
			}

			// bottom
			if y+1 < len(input) {
				bottomNeighbor := input[y+1][x]
				if isSymbol(bottomNeighbor) {
					adjacentSymbol = true
					continue
				}
			}

			// left
			if x-1 >= 0 {
				leftNeighbor := input[y][x-1]
				if isSymbol(leftNeighbor) {
					adjacentSymbol = true
					continue
				}
			}

			// top
			if y-1 >= 0 {
				topNeighbor := input[y-1][x]
				if isSymbol(topNeighbor) {
					adjacentSymbol = true
					continue
				}
			}

			// top-left
			if y-1 >= 0 && x-1 >= 0 {
				topLeftNeighbor := input[y-1][x-1]
				if isSymbol(topLeftNeighbor) {
					adjacentSymbol = true
					continue
				}
			}

			// top-right
			if y-1 >= 0 && x+1 < len(row) {
				topRightNeighbor := input[y-1][x+1]
				if isSymbol(topRightNeighbor) {
					adjacentSymbol = true
					continue
				}
			}

			// bottom-left
			if y+1 < len(input) && x-1 >= 0 {
				bottomLeftNeighbor := input[y+1][x-1]
				if isSymbol(bottomLeftNeighbor) {
					adjacentSymbol = true
					continue
				}
			}

			// bottom-right
			if y+1 < len(input) && x+1 < len(row) {
				bottomRightNeighbor := input[y+1][x+1]
				if isSymbol(bottomRightNeighbor) {
					adjacentSymbol = true
					continue
				}
			}

		}
	}

	return sum
}

type set struct {
	x int
	y int
}

type sets struct {
	s      []set
	number int
}

func part2(input [][]string) int {
	r := getGearPositions(input)
	ss := getNumberPositions(input)

	return findEngines(r, ss)
}

func findEngines(gears []set, numbers []sets) int {
	connections := make(map[string][]int)
	for _, g := range gears {
		for _, ss := range numbers {
			for _, s := range ss.s {
				if abs(s.x-g.x) <= 1 && abs(s.y-g.y) <= 1 {
					key := strconv.Itoa(g.x) + strconv.Itoa(g.y)
					value, exist := connections[key]
					if !exist {
						connections[key] = []int{ss.number}
					} else {
						temp := value
						temp = append(temp, ss.number)
						connections[key] = temp
					}
					break
				}
			}
		}
	}

	sum := 0
	for _, value := range connections {
		if len(value) == 2 {
			sum += value[0] * value[1]
		}
	}

	return sum
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getGearPositions(input [][]string) []set {
	positions := make([]set, 0)
	for y, row := range input {
		for x, value := range row {
			if value == "*" {
				positions = append(positions, set{
					x: x,
					y: y,
				})
			}
		}
	}

	return positions
}

// absolute trash
func getNumberPositions(input [][]string) []sets {
	all := make([]sets, 0)
	combined := ""
	for y, row := range input {
		index := make([]int, 0)
		positions := make([]set, 0)
		for x, value := range row {
			_, err := strconv.Atoi(value)

			if err != nil {
				r, _ := strconv.Atoi(combined)
				combined = ""
				for _, i := range index {
					positions = append(positions, set{
						x: i,
						y: y,
					})
				}

				if len(positions) > 0 {
					all = append(all, sets{
						s:      positions,
						number: r,
					})
				}

				index = nil
				positions = nil

				continue
			}

			combined += value
			index = append(index, x)

			if x == len(row)-1 {
				r, _ := strconv.Atoi(combined)
				combined = ""
				for _, i := range index {
					positions = append(positions, set{
						x: i,
						y: y,
					})
				}

				if len(positions) > 0 {
					all = append(all, sets{
						s:      positions,
						number: r,
					})
				}

			}
		}
	}

	return all
}

func isSymbol(value string) bool {
	if value == "." {
		return false
	}

	_, err := strconv.Atoi(value)
	if err == nil {
		return false
	}

	return true
}

func readInput() ([][]string, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	stuff := make([][]string, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input := scanner.Text()
		row := make([]string, 0)
		for _, char := range input {
			row = append(row, string(char))
		}

		stuff = append(stuff, row)
	}

	return stuff, nil
}
