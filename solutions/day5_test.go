package solutions

import (
	"testing"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

const day5Input = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

func TestDay5Part1(t *testing.T) {
	day := Day5{}
	result, err := day.Part1(day5Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "143")
}

func TestDay5Part2(t *testing.T) {
	day := Day5{}
	result, err := day.Part2(day5Input)

	utils.AssertNil(t, err)
	utils.AssertEqual(t, result, "123")
}
