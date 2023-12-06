package main

type race struct {
	time     int
	distance int
}

func part1() int {
	races := []race{
		{time: 59, distance: 430},
		{time: 70, distance: 1218},
		{time: 78, distance: 1213},
		{time: 78, distance: 1276},
	}

	return findRecordTimes(races)
}

func part2() int {
	races := []race{
		{time: 59707878, distance: 430121812131276},
	}

	return findRecordTimes(races)
}

func findRecordTimes(races []race) int {
	total := make([]int, 0)
	for _, r := range races {
		wins := 0

		for i := 1; i <= r.time; i++ {
			if (r.time-i)*i > r.distance {
				wins++
			}
		}
		total = append(total, wins)

	}

	result := 1
	for _, w := range total {
		result *= w
	}

	return result
}
