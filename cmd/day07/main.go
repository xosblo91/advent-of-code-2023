package main

import (
	"bufio"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/xosblo91/advent-of-code-2023/internal/conv"
)

var cardWeights1 = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

var cardWeights2 = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 11,
	"K": 12,
	"A": 13,
}

type game struct {
	cards          string
	optimizedCards string
	betSize        int
	strength       int
}

func (g *game) setHandStrength(optimize bool) {
	count := make(map[string]int)
	to := ""
	if optimize {
		to = g.optimizedCards
	} else {
		to = g.cards
	}

	for _, card := range to {
		if _, exist := count[string(card)]; !exist {
			count[string(card)] = 1
		} else {
			t := count[string(card)]
			t++
			count[string(card)] = t
		}
	}

	switch len(count) {
	case 1:
		g.strength = 7
	case 2:
		for _, c := range count {
			if c == 4 {
				g.strength = 6
				return
			}
		}
		g.strength = 5
	case 3:
		for _, c := range count {
			if c == 3 {
				g.strength = 4
				return
			}
		}
		g.strength = 3
	case 4:
		g.strength = 2
	case 5:
		g.strength = 1
	}
}

func (g *game) optimizeWithJoker() {
	count := make(map[string]int)
	v := ""
	h := 0
	for _, card := range g.cards {
		if _, exist := count[string(card)]; !exist {
			count[string(card)] = 1
		} else {
			t := count[string(card)]
			t++
			count[string(card)] = t
		}

		if count[string(card)] >= h && string(card) != "J" {
			v = string(card)
			h = count[string(card)]
		}
	}

	temp := g.cards
	for _, card := range g.cards {
		if string(card) == "J" {
			g.optimizedCards = strings.ReplaceAll(temp, "J", v)
		}
	}
	if g.optimizedCards == "" {
		g.optimizedCards = g.cards
	}
}

func part1(games []*game) int {
	for _, g := range games {
		g.setHandStrength(false)
	}

	return summarizeScore(sortHands(games, cardWeights1))
}

func part2(games []*game) int {
	for _, g := range games {
		g.optimizeWithJoker()
		g.setHandStrength(true)
	}

	return summarizeScore(sortHands(games, cardWeights2))
}

func sortHands(games []*game, weights map[string]int) []*game {
	sort.Slice(games, func(i, j int) bool {
		if games[i].strength == games[j].strength {
			for ii := 0; ii < len(games[ii].cards); ii++ {
				if weights[string(games[i].cards[ii])] > weights[string(games[j].cards[ii])] {
					return true
				}
				if weights[string(games[i].cards[ii])] == weights[string(games[j].cards[ii])] {
					continue
				}
				return false
			}
		}

		return games[i].strength > games[j].strength
	})

	slices.Reverse(games)

	return games
}

func summarizeScore(sorted []*game) int {
	sum := 0
	for i, g := range sorted {
		sum += (i + 1) * g.betSize
	}

	return sum
}

func readInput() []*game {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("fuck")
	}

	games := make([]*game, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		games = append(games, &game{
			cards:   split[0],
			betSize: conv.Atoi[int](split[1]),
		})
	}

	return games
}
