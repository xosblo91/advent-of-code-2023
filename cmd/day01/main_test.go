package day01

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input, err := readInput()
	require.NoError(t, err)

	sum, err := calibration1(input)
	require.NoError(t, err)

	t.Log(sum)
}

func TestPart2(t *testing.T) {
	input, err := readInput()
	require.NoError(t, err)

	sum, err := calibration2(input)
	require.NoError(t, err)

	t.Log(sum)
}
