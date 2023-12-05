package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/xosblo91/advent-of-code-2023/internal/sliceutils"
)

type mapping struct {
	destination uint64
	source      uint64
	length      uint64
}

func (m mapping) find(n uint64) uint64 {
	if n >= m.source && n <= m.source+m.length-1 {
		return (n - m.source) + m.destination
	}

	return n
}

func part1(seeds []uint64, mappings [][]mapping) uint64 {
	return findLowestLocationNumber(seeds, mappings)
}

func findLowestLocationNumber(seeds []uint64, mappings [][]mapping) uint64 {
	var lowest uint64 = math.MaxUint64
	for _, seed := range seeds {
		translation := seed
		for _, category := range mappings {
			for _, m := range category {
				if result := m.find(translation); result != translation {
					translation = result
					break
				}
			}
		}
		if translation < lowest {
			lowest = translation
		}
	}

	return lowest
}

func part2(initialSeeds []uint64, mappings [][]mapping) uint64 {
	allSeeds := makeSeedRanges(initialSeeds)

	var lowest uint64 = math.MaxUint64
	for _, p := range allSeeds {
		result := findLowestLocationNumber(p, mappings)
		if result < lowest {
			lowest = result
		}
	}

	return lowest
}

func makeSeedRanges(initialSeeds []uint64) [][]uint64 {
	seedBatches := sliceutils.MakeBatches[uint64](initialSeeds, 2)

	allSeeds := make([][]uint64, 0)
	for _, batch := range seedBatches {
		seeds := make([]uint64, 0)
		for i := batch[0]; i < batch[0]+batch[1]; i++ {
			seeds = append(seeds, i)
		}
		allSeeds = append(allSeeds, seeds)
	}

	return allSeeds
}

func atoi(s string) uint64 {
	n, _ := strconv.Atoi(s)
	return uint64(n)
}

// oh lawd
func readInput() ([]uint64, [][]mapping) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	seeds := make([]uint64, 0)

	split := strings.Split(strings.Join(lines, "\n"), "\n\n")

	seedsStr := split[0]
	seedsNumbersStr := regexp.MustCompile("[0-9]+").FindAllString(seedsStr, -1)

	for _, seedNumberStr := range seedsNumbersStr {
		seedNumber, _ := strconv.Atoi(seedNumberStr)
		seeds = append(seeds, uint64(seedNumber))
	}

	everything := make([][]mapping, 0)
	for _, split := range split[1:] {
		elementsWithoutTitle := strings.Split(split, "\n")
		mappings := make([]mapping, 0)

		for _, e := range elementsWithoutTitle[1:] {
			split := strings.Split(e, " ")
			mappings = append(mappings, mapping{
				destination: atoi(split[0]),
				source:      atoi(split[1]),
				length:      atoi(split[2]),
			})
		}
		everything = append(everything, mappings)
	}

	return seeds, everything
}
