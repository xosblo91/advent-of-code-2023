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

type coordinates struct {
	x int
	y int
}

type part struct {
	indices []coordinates
	number  int
}

func part2(input [][]string) int {
	gears := getGearCoordinates(input)
	parts := getPartCoordinates(input)

	return findGearRatios(gears, parts)
}

func findGearRatios(gears []coordinates, parts []part) int {
	connections := make(map[string][]int)
	for _, gear := range gears {
		for _, part := range parts {
			for _, index := range part.indices {
				if abs(index.x-gear.x) <= 1 && abs(index.y-gear.y) <= 1 {
					key := strconv.Itoa(gear.x) + strconv.Itoa(gear.y)
					value, exist := connections[key]
					if !exist {
						connections[key] = []int{part.number}
					} else {
						temp := value
						temp = append(temp, part.number)
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

func getGearCoordinates(input [][]string) []coordinates {
	c := make([]coordinates, 0)
	for y, row := range input {
		for x, value := range row {
			if value == "*" {
				c = append(c, coordinates{
					x: x,
					y: y,
				})
			}
		}
	}

	return c
}

// absolute trash
func getPartCoordinates(input [][]string) []part {
	all := make([]part, 0)
	combined := ""
	for y, row := range input {
		index := make([]int, 0)
		c := make([]coordinates, 0)
		for x, value := range row {
			_, err := strconv.Atoi(value)

			if err != nil {
				r, _ := strconv.Atoi(combined)
				combined = ""
				for _, i := range index {
					c = append(c, coordinates{
						x: i,
						y: y,
					})
				}

				if len(c) > 0 {
					all = append(all, part{
						indices: c,
						number:  r,
					})
				}

				index = nil
				c = nil

				continue
			}

			combined += value
			index = append(index, x)

			if x == len(row)-1 {
				r, _ := strconv.Atoi(combined)
				combined = ""
				for _, i := range index {
					c = append(c, coordinates{
						x: i,
						y: y,
					})
				}

				if len(c) > 0 {
					all = append(all, part{
						indices: c,
						number:  r,
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
