package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	GameNumber int
	Sets       []Set
}

type Set struct {
	Cubes map[string]int
}

func part1(games []Game) int {
	m := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	for _, game := range games {
		possible := true
		for _, set := range game.Sets {
			if set.Cubes["red"] > m["red"] {
				possible = false
			}
			if set.Cubes["green"] > m["green"] {
				possible = false
			}
			if set.Cubes["blue"] > m["blue"] {
				possible = false
			}
		}
		if possible {
			sum += game.GameNumber
		}
	}

	return sum
}

func part2(games []Game) int {
	sum := 0
	for _, game := range games {
		m := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, set := range game.Sets {
			if set.Cubes["red"] > m["red"] {
				m["red"] = set.Cubes["red"]
			}
			if set.Cubes["green"] > m["green"] {
				m["green"] = set.Cubes["green"]
			}
			if set.Cubes["blue"] > m["blue"] {
				m["blue"] = set.Cubes["blue"]
			}
		}

		sum += m["red"] * m["green"] * m["blue"]
	}

	return sum
}

func readInput() ([]Game, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	games := make([]Game, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		g := Game{}
		s1 := strings.Split(scanner.Text(), ":")
		g.GameNumber, _ = strconv.Atoi(strings.Split(s1[0], " ")[1])
		sets := strings.Split(s1[1], ";")

		if g.Sets == nil {
			g.Sets = make([]Set, 0)
		}
		for _, set := range sets {
			cubes := strings.Split(set, ",")

			m := make(map[string]int)
			for _, cube := range cubes {
				final := strings.Split(cube, " ")

				c, _ := strconv.Atoi(final[1])
				m[final[2]] = c
			}
			g.Sets = append(g.Sets, Set{
				Cubes: m,
			})
		}

		games = append(games, g)
	}

	return games, nil
}
