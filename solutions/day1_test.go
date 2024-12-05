package solutions

import (
	"testing"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

const day1Input = `3   4
4   3
2   5
1   3
3   9
3   3
`

func TestDay1Part1(t *testing.T) {
	day := Day1{}
	result, err := day.Part1(day1Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "11")
}

func TestDay1Part2(t *testing.T) {
	day := Day1{}
	result, err := day.Part2(day1Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "31")
}
