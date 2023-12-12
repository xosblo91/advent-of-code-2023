package main

import (
	"bufio"
	"github.com/xosblo91/advent-of-code-2023/internal/conv"
	"os"
	"strings"
)

// absolute fucking garbage (as per usual)

func part1(start [][]int) int {
	multi := make([][]int, 0)
	for _, s := range start {
		multi = append(multi, findStuff(s, false))
	}

	sum := 0
	for _, m := range multi {
		sum += doReverseStuff(m, false)
	}

	return sum
}

func part2(start [][]int) int {
	multi := make([][]int, 0)
	for _, s := range start {
		multi = append(multi, findStuff(s, true))
	}

	sum := 0
	for _, m := range multi {
		sum += doReverseStuff(m, true)
	}

	return sum
}

func doReverseStuff(values []int, isPart2 bool) int {
	cary := 0
	for i := len(values) - 1; i >= 0; i-- {
		if i == 0 {
			break
		}
		n1 := 0
		if i == len(values)-1 {
			n1 = values[i]
		} else {
			n1 = cary
		}

		n2 := values[i-1]
		if isPart2 {
			cary = n2 - n1
		} else {
			cary = n1 + n2
		}
	}

	return cary
}

func findStuff(input []int, isPart2 bool) []int {
	multi := [][]int{input}
	for {
		n := getDiff(input)
		multi = append(multi, n)
		if allZeros(n) {
			break
		}

		input = n
	}

	l := make([]int, 0)
	for i := 0; i < len(multi); i++ {
		if isPart2 {
			l = append(l, multi[i][0])
		} else {
			l = append(l, multi[i][len(multi[i])-1])
		}
	}

	return l
}

func getDiff(input []int) []int {
	n := make([]int, 0)
	for i := 0; i < len(input); i++ {
		if i == len(input)-1 {
			break
		}
		diff := input[i+1] - input[i]
		n = append(n, diff)
	}

	return n
}

func allZeros(input []int) bool {
	for _, i := range input {
		if i != 0 {
			return false
		}
	}

	return true
}

func readInput() [][]int {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("fuck")
	}

	stuff := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := make([]int, 0)
		split := strings.Split(scanner.Text(), " ")
		for _, s := range split {
			l = append(l, conv.Atoi[int](s))
		}
		stuff = append(stuff, l)
	}

	return stuff
}
