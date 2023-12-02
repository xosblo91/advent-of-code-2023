package day01

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var match = map[string]string{
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func calibration1(input []string) (int, error) {
	sum := 0
	for _, calibration := range input {
		current := ""
		for _, r := range []rune(calibration) {
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				current += string(r)
			}
		}

		v, err := strconv.Atoi(current[0:1] + current[len(current)-1:])
		if err != nil {
			return 0, err
		}
		sum += v
	}

	return sum, nil
}

type set struct {
	index int
	value string
}

func calibration2(input []string) (int, error) {
	sum := 0
	for _, calibration := range input {
		minimum := 999999
		maximum := 0

		rr := make([]set, 0)

		first, last := "", ""
		for k, m := range match {
			re := regexp.MustCompile(k)
			result := re.FindAllStringIndex(calibration, -1)

			for _, r := range result {
				rr = append(rr, set{
					index: r[0],
					value: m,
				})

			}
		}

		for _, s := range rr {
			if s.index <= minimum {
				first = s.value
				minimum = s.index
			}
			if s.index >= maximum {
				last = s.value
				maximum = s.index
			}
		}

		v, err := strconv.Atoi(first + last)
		if err != nil {
			return 0, err
		}

		sum += v
	}

	return sum, nil
}

func readInput() ([]string, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	stuff := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		stuff = append(stuff, scanner.Text())
	}

	return stuff, nil
}
