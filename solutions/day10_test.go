package solutions

import (
	"testing"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

const day10Input = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

func TestDay10Part1(t *testing.T) {
	day := Day10{}
	result, err := day.Part1(day10Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "36")
}

func TestDay10Part2(t *testing.T) {
	day := Day10{}
	result, err := day.Part2(day10Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "81")
}
