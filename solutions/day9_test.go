package solutions

import (
	"testing"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

const day9Input = "2333133121414131402"

func TestDay9Part1(t *testing.T) {
	day := Day9{}
	result, err := day.Part1(day9Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "1928")
}

func TestDay9Part2(t *testing.T) {
	day := Day9{}
	result, err := day.Part2(day9Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "2858")
}
