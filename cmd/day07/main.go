package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/xosblo91/advent-of-code-2023/internal/conv"
)

var cardWeights = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

const (
	FI = "FiveOfAKind"
	FO = "FourOfAKind"
	FH = "FullHouse"
	TH = "ThreeOfAKind"
	TP = "TwoPair"
	OP = "OnePair"
	HC = "HighCard"
)

var handStrength = map[string]int{
	HC: 1,
	OP: 2,
	TP: 3,
	TH: 4,
	FH: 5,
	FO: 6,
	FI: 7,
}

type game struct {
	cards   string
	betSize int
}

func (g *game) identifyHand() {

}

func part1(games []game) int {
	return 0
}

func readInput() []game {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("fuck")
	}

	games := make([]game, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		games = append(games, game{
			cards:   split[0],
			betSize: conv.Atoi[int](split[1]),
		})
	}

	return games
}
