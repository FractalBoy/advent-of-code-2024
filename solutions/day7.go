package solutions

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

type Day7 struct{}

type Equation struct {
	LeftHandSide  int
	RightHandSide []int
}

func parseDay7Input(input string) ([]Equation, error) {
	lines := utils.SplitLines(input)
	equations := make([]Equation, len(lines))

	for i, line := range lines {
		split := strings.Split(line, ": ")
		leftHandSide, err := strconv.Atoi(split[0])

		if err != nil {
			return nil, err
		}

		split2 := strings.Split(split[1], " ")
		equation := Equation{LeftHandSide: leftHandSide, RightHandSide: make([]int, len(split2))}

		for j, strnum := range split2 {
			num, err := strconv.Atoi(strnum)

			if err != nil {
				return nil, err
			}

			equation.RightHandSide[j] = num
		}

		equations[i] = equation
	}

	return equations, nil
}

func (equation Equation) possiblyTrue() bool {
	first := equation.RightHandSide[0]
	second := equation.RightHandSide[1]
	tail := equation.RightHandSide[2:]

	results := possibleResults(first, second, tail)

	return slices.Contains(results, equation.LeftHandSide)
}

func possibleResults(first int, second int, tail []int) []int {
	if len(tail) == 0 {
		return []int{first * second, first + second}
	}

	results := []int{}
	results = append(results, possibleResults(first+second, tail[0], tail[1:])...)
	results = append(results, possibleResults(first*second, tail[0], tail[1:])...)

	return results
}

func (d Day7) Part1(input string) (string, error) {
	equations, err := parseDay7Input(input)

	if err != nil {
		return "", err
	}

	sum := 0

	for _, equation := range equations {
		if equation.possiblyTrue() {
			sum += equation.LeftHandSide
		}
	}

	return strconv.Itoa(sum), nil
}

func (equation Equation) possiblyTrueWithConcatenation() bool {
	first := equation.RightHandSide[0]
	second := equation.RightHandSide[1]
	tail := equation.RightHandSide[2:]

	results := possibleResultsWithConcatenation(first, second, tail)

	return slices.Contains(results, equation.LeftHandSide)
}

func possibleResultsWithConcatenation(first int, second int, tail []int) []int {
	digits := int(math.Floor(math.Log10(float64(second)))) + 1
	concat := first*int(math.Pow10(digits)) + second

	if len(tail) == 0 {
		return []int{first * second, first + second, concat}
	}

	results := []int{}
	results = append(results, possibleResultsWithConcatenation(first+second, tail[0], tail[1:])...)
	results = append(results, possibleResultsWithConcatenation(first*second, tail[0], tail[1:])...)
	results = append(results, possibleResultsWithConcatenation(concat, tail[0], tail[1:])...)

	return results
}

func (d Day7) Part2(input string) (string, error) {
	equations, err := parseDay7Input(input)

	if err != nil {
		return "", err
	}

	sum := 0

	for _, equation := range equations {
		if equation.possiblyTrueWithConcatenation() {
			sum += equation.LeftHandSide
		}
	}

	return strconv.Itoa(sum), nil
}
