package conv

import (
	"strconv"
)

func Atoi[T int | int32](s string) T {
	n, _ := strconv.Atoi(s)
	return T(n)
}
