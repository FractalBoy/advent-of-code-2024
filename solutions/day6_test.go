package solutions

import (
	"testing"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

const day6Input = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

func TestDay6Part1(t *testing.T) {
	day := Day6{}
	result, err := day.Part1(day6Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "41")
}

func TestDay6Part2(t *testing.T) {
	day := Day6{}
	result, err := day.Part2(day6Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "")
}
