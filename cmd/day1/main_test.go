package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input, err := readInput()
	require.NoError(t, err)

	sum := calibration1(input)
	t.Log(sum)
}

func TestPart2(t *testing.T) {
	input, err := readInput()
	require.NoError(t, err)

	sum := calibration2(input)
	t.Log(sum)
}
