package main

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strings"
)

type card struct {
	CardNumber     string
	Numbers        []string
	WinningNumbers []string
}

func (c card) sumOfMatchingNumbers() int {
	matches := 0
	for _, n := range c.Numbers {
		if slices.Contains(c.WinningNumbers, n) {
			matches++
			continue
		}
	}

	return int(math.Pow(2, float64(matches)-1))
}

func (c card) matchingNumbers() int {
	matches := 0
	for _, n := range c.Numbers {
		if slices.Contains(c.WinningNumbers, n) {
			matches++
			continue
		}
	}

	return matches
}

func part1(cards []card) int {
	sum := 0

	for _, g := range cards {
		sum += g.sumOfMatchingNumbers()
	}

	return sum
}

func part2(cards []card) int {
	originalAndCopies := make([]int, 0, len(cards))

	// at least one of the original cards
	for range cards {
		originalAndCopies = append(originalAndCopies, 1)
	}

	for index, card := range cards {
		matchingCards := card.matchingNumbers()
		if matchingCards == 0 {
			continue
		}

		for i := index + 1; i <= index+matchingCards; i++ {
			originalAndCopies[i] += originalAndCopies[index]
		}
	}

	sum := 0
	for _, count := range originalAndCopies {
		sum += count
	}

	return sum
}

func readInput() ([]card, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	cards := make([]card, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split1 := strings.Split(scanner.Text(), ":")

		g := card{CardNumber: split1[0]}

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

		cards = append(cards, g)
	}

	return cards, nil
}
