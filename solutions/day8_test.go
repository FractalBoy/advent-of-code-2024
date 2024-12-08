package solutions

import (
	"testing"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

const day8Input = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`

func TestDay8Part1(t *testing.T) {
	day := Day8{}
	result, err := day.Part1(day8Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "14")
}

func TestDay8Part2(t *testing.T) {
	day := Day8{}
	result, err := day.Part2(day8Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "34")
}
