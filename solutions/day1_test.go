package solutions

import "testing"

var day1Input = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestDay1Part1(t *testing.T) {
	day := Day1{}
	result, err := day.Part1(day1Input)

	if err != nil {
		t.Fatal(err)
	}

	if result != "11" {
		t.Fatalf("result is incorrect, was %s, expected 11", result)
	}
}

func TestDay1Part2(t *testing.T) {
	day := Day1{}
	result, err := day.Part2(day1Input)

	if err != nil {
		t.Fatal(err)
	}

	if result != "31" {
		t.Fatalf("result is incorrect, was %s, expected 31", result)
	}
}
