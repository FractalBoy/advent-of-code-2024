package solutions

import "testing"

var input = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1(t *testing.T) {
	day := Day1{}
	result := day.Part1(input)

	if result != "11" {
		t.Fatalf("result is incorrect, was %s, expected 11", result)
	}
}

func TestPart2(t *testing.T) {
	day := Day1{}
	result := day.Part2(input)

	if result != "31" {
		t.Fatalf("result is incorrect, was %s, expected 31", result)
	}
}
