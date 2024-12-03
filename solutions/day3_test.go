package solutions

import "testing"

func TestDay3Part1(t *testing.T) {
	day := Day3{}
	result, err := day.Part1("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")

	if err != nil {
		t.Fatal(err)
	}

	if result != "161" {
		t.Fatalf("result is incorrect, was %s, expected 161", result)
	}
}

func TestDay3Part2(t *testing.T) {
	day := Day3{}
	result, err := day.Part2("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")

	if err != nil {
		t.Fatal(err)
	}

	if result != "48" {
		t.Fatalf("result is incorrect, was %s, expected 48", result)
	}
}
