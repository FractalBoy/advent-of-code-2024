package solutions

import (
	"testing"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

const day4Input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

func TestDay4Part1(t *testing.T) {
	day := Day4{}
	result, err := day.Part1(day4Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "18")
}

func TestDay4Part2(t *testing.T) {
	day := Day4{}
	result, err := day.Part2(day4Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "9")
}
