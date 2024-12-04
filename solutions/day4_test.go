package solutions

import "testing"

var day4Input = `MMMSXXMASM
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

	if err != nil {
		t.Fatal(err)
	}

	if result != "18" {
		t.Fatalf("result is incorrect, was %s, expected 18", result)
	}
}

func TestDay4Part2(t *testing.T) {
	day := Day4{}
	result, err := day.Part2(day4Input)

	if err != nil {
		t.Fatal(err)
	}

	if result != "9" {
		t.Fatalf("result is incorrect, was %s, expected 9", result)
	}
}
