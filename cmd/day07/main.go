package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/xosblo91/advent-of-code-2023/internal/conv"
)

var cardWeights = map[string]int{
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

// const (
// 	HC = "HighCard"
// 	OP = "OnePair"
// 	TP = "TwoPair"
// 	TH = "ThreeOfAKind"
// 	FH = "FullHouse"
// 	FO = "FourOfAKind"
// 	FI = "FiveOfAKind"
// )
//
// var handStrength = map[string]int{
// 	HC: 1, // 1 2 3 4 5 - 5
// 	OP: 2, // 1 1 3 4 5 - 4
// 	TP: 3, // 1 1 2 2 5 - 3
// 	TH: 4, // 1 1 1 4 5 - 3
// 	FH: 5, // 1 1 1 2 2 - 2
// 	FO: 6, // 1 1 1 1 2 - 2
// 	FI: 7, // 1 1 1 1 1 - 1
// }

type game struct {
	cards    string
	betSize  int
	strength int
}

func (g *game) setHandStrength() {
	count := make(map[string]int)
	for _, card := range g.cards {
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

func part1(games []*game) int {
	for _, g := range games {
		g.setHandStrength()
	}

	return 0
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
