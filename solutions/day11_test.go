package solutions

import (
	"testing"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

const day11Input = "125 17\n"

func TestDay11Part1(t *testing.T) {
	day := Day11{}
	result, err := day.Part1(day11Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "55312")
}
