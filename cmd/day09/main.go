package main

func part1() {
	start := []int{0, 3, 6, 9, 12, 15}

	multi := [][]int{start}

	for {
		n := hej(start)
		multi = append(multi, n)
		if isAllZeros(n) {
			break
		}

		start = n
	}

	println(multi)
}

func hej(input []int) []int {
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

func isAllZeros(input []int) bool {
	for _, i := range input {
		if i != 0 {
			return false
		}
	}

	return true
}
