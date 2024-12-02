package solutions

import (
	"math"
	"sort"
	"strconv"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

type Day1 struct{}

func parseInput(input string) ([]int, []int, error) {
	left := []int{}
	right := []int{}

	lines, err := utils.SplitLinesIntFields(input)

	if err != nil {
		return nil, nil, err
	}

	for i := 0; i < len(lines); i++ {
		left = append(left, lines[i][0])
		right = append(right, lines[i][1])
	}

	return left, right, nil
}

func (d Day1) Part1(input string) (string, error) {
	left, right, err := parseInput(input)

	if err != nil {
		return "", err
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	sum := 0.

	for i := 0; i < len(left); i++ {
		sum += math.Abs(float64(right[i] - left[i]))
	}

	return strconv.FormatFloat(sum, 'f', 0, 64), nil
}

func (d Day1) Part2(input string) (string, error) {
	left, right, err := parseInput(input)

	if err != nil {
		return "", err
	}

	freq := make(map[int]int)

	for i := 0; i < len(right); i++ {
		count, ok := freq[right[i]]

		if ok {
			freq[right[i]] = count + 1
		} else {
			freq[right[i]] = 1
		}
	}

	sum := 0

	for i := 0; i < len(left); i++ {
		count, ok := freq[left[i]]

		if !ok {
			continue
		}

		sum += left[i] * count
	}

	return strconv.Itoa(sum), nil
}
