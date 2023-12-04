package main

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strings"
)

type game struct {
	Card           string
	Numbers        []string
	WinningNumbers []string
}

func (g game) sumOfMatchingNumbers() int {
	matches := 0
	for _, n := range g.Numbers {
		if slices.Contains(g.WinningNumbers, n) {
			matches++
			continue
		}
	}

	return int(math.Pow(2, float64(matches)-1))
}

func (g game) matchingNumbers() int {
	matches := 0
	for _, n := range g.Numbers {
		if slices.Contains(g.WinningNumbers, n) {
			matches++
			continue
		}
	}

	return matches
}

func part1(games []game) int {
	sum := 0

	for _, g := range games {
		sum += g.sumOfMatchingNumbers()
	}

	return sum
}

func part2(games []game) int {
	copies := make([]int, len(games))
	// deck := make([]int, 1024)
	// ncard := 0
	// for i := range copies {
	// 	copies[i] = 1
	// }

	for index, g := range games {
		matchingCards := g.sumOfMatchingNumbers()
		if matchingCards == 0 {
			continue
		}

		startIdx := index + 1
		endIdx := min(len(games), index+matchingCards)

		c := games[startIdx : endIdx-1]

		println(c)

		currentCardCopies := copies[index]
		for i := startIdx; i <= endIdx; i++ {
			// increase count by adding copy of current card
			copies = append(copies, currentCardCopies)
		}
	}

	sum := 0
	for _, v := range copies {
		sum += v
	}

	return sum
}

func count() {

}

func readInput() ([]game, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	games := make([]game, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split1 := strings.Split(scanner.Text(), ":")

		g := game{Card: split1[0]}

		split2 := strings.Split(split1[1], "|")

		numbers := strings.Split(strings.TrimSpace(split2[1]), " ")
		winningNumbers := strings.Split(strings.TrimSpace(split2[0]), " ")

		for _, n := range numbers {
			if n != "" {
				g.Numbers = append(g.Numbers, n)
			}
		}

		for _, n := range winningNumbers {
			if n != "" {
				g.WinningNumbers = append(g.WinningNumbers, n)
			}
		}

		games = append(games, g)
	}

	return games, nil
}
