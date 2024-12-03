package solutions

import (
	"math"
	"strconv"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

type Day2 struct{}

func isReportSafe(level []int) bool {
	sign := 0

	for j := 1; j < len(level); j++ {
		currLevel := level[j]
		prevLevel := level[j-1]

		diff := math.Abs(float64(prevLevel - currLevel))

		if diff == 0 || diff > 3 {
			return false
		}

		if sign == 0 {
			if prevLevel > currLevel {
				sign = 1
			} else if prevLevel < currLevel {
				sign = -1
			}
		} else if sign == 1 {
			if prevLevel < currLevel {
				return false
			}
		} else {
			if prevLevel > currLevel {
				return false
			}
		}
	}

	return true
}

func (d Day2) Part1(input string) (string, error) {
	reports, err := utils.SplitLinesIntFields(input)

	if err != nil {
		return "", err
	}

	safeCount := 0

	for i := 0; i < len(reports); i++ {
		if isReportSafe(reports[i]) {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount), nil
}

func (d Day2) Part2(input string) (string, error) {
	reports, err := utils.SplitLinesIntFields(input)

	if err != nil {
		return "", err
	}

	safeCount := 0

	for i := 0; i < len(reports); i++ {
		report := reports[i]
		for j := 0; j < len(report); j++ {
			permutation := []int{}

			// Build a report with one level missing at a time
			for k := 0; k < len(report); k++ {
				if k == j {
					continue
				}

				permutation = append(permutation, report[k])
			}

			// Found one permutation that is safe. Stop checking.
			if isReportSafe(permutation) {
				safeCount++
				break
			}
		}
	}

	return strconv.Itoa(safeCount), nil
}
