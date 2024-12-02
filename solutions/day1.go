package solutions

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day1 struct{}

func parseInput(input string) ([]int, []int, error) {
	lines := strings.Split(input, "\n")

	left := []int{}
	right := []int{}

	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")

		for j := 0; j < len(parts); j++ {
			if parts[j] == "" {
				continue
			}

			num, err := strconv.Atoi(parts[j])

			if err != nil {
				return nil, nil, fmt.Errorf("failed to convert \"%s\" to int: %s", parts[j], err)
			}

			if len(left) == i {
				left = append(left, num)
			} else if len(right) == i {
				right = append(right, num)
			}
		}
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
