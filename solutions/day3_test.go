package solutions

import (
	"testing"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

func TestDay3Part1(t *testing.T) {
	day := Day3{}
	result, err := day.Part1("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "161")
}

func TestDay3Part2(t *testing.T) {
	day := Day3{}
	result, err := day.Part2("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "48")
}
