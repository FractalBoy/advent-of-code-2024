package solutions

import (
	"testing"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

const day17Input = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`

func TestDay17Part1(t *testing.T) {
	day := Day17{}
	result, err := day.Part1(day17Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "4,6,3,5,6,3,5,2,1,0")
}
