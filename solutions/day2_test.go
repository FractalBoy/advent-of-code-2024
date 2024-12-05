package solutions

import (
	"testing"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

const day2Input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestDay2Part1(t *testing.T) {
	day := Day2{}
	result, err := day.Part1(day2Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "2")
}

func TestDay2Part2(t *testing.T) {
	day := Day2{}
	result, err := day.Part2(day2Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "4")
}
