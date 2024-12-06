package solutions

import (
	"slices"
	"strconv"
	"strings"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

type Day5 struct{}

type OrderingRule struct {
	Before int
	After  int
}

func parseDay5Input(input string) ([]OrderingRule, [][]int, error) {
	lines := utils.SplitLines(input)

	rules := []OrderingRule{}
	updates := [][]int{}

	line, lines := lines[0], lines[1:]

	for line != "" {
		split := strings.Split(line, "|")
		before, err := strconv.Atoi(split[0])

		if err != nil {
			return nil, nil, err
		}

		after, err := strconv.Atoi(split[1])

		if err != nil {
			return nil, nil, err
		}

		rule := OrderingRule{Before: before, After: after}
		rules = append(rules, rule)
		line, lines = lines[0], lines[1:]
	}

	for _, line := range lines {
		split := strings.Split(line, ",")
		update := []int{}

		for _, elem := range split {
			num, err := strconv.Atoi(elem)

			if err != nil {
				return nil, nil, err
			}

			update = append(update, num)
		}

		updates = append(updates, update)
	}

	return rules, updates, nil
}

func sortUpdate(rules []OrderingRule, update []int) {
	slices.SortFunc(update, func(a int, b int) int {
		for _, rule := range rules {
			if rule.Before == a && rule.After == b {
				return -1
			}

			if rule.Before == b && rule.After == a {
				return 1
			}
		}

		return 0
	})
}

func (d Day5) Part1(input string) (string, error) {
	rules, updates, err := parseDay5Input(input)

	if err != nil {
		return "", err
	}

	sum := 0

	for _, update := range updates {
		before := slices.Clone(update)
		sortUpdate(rules, update)

		if slices.Equal(before, update) {
			sum += update[len(update)/2]
		}
	}

	return strconv.Itoa(sum), nil
}

func (d Day5) Part2(input string) (string, error) {
	rules, updates, err := parseDay5Input(input)

	if err != nil {
		return "", err
	}

	sum := 0

	for _, update := range updates {
		before := slices.Clone(update)
		sortUpdate(rules, update)

		if !slices.Equal(before, update) {
			sum += update[len(update)/2]
		}
	}

	return strconv.Itoa(sum), nil
}
