package solutions

import (
	"testing"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

const day7Input = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func TestDay7Part1(t *testing.T) {
	day := Day7{}
	result, err := day.Part1(day7Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "3749")
}

func TestDay7Part2(t *testing.T) {
	day := Day7{}
	result, err := day.Part2(day7Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "11387")
}
