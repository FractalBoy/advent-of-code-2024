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

	for _, line := range lines {
		left = append(left, line[0])
		right = append(right, line[1])
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

	for _, distance := range right {
		count, ok := freq[distance]

		if ok {
			freq[distance] = count + 1
		} else {
			freq[distance] = 1
		}
	}

	sum := 0

	for _, distance := range left {
		count, ok := freq[distance]

		if !ok {
			continue
		}

		sum += distance * count
	}

	return strconv.Itoa(sum), nil
}
