package utils

import (
	"strconv"
	"strings"
)

func SplitLines(str string) []string {
	lines := strings.Split(str, "\n")
	nonEmptyLines := []string{}

	for i := 0; i < len(lines); i++ {
		if lines[i] != "" {
			nonEmptyLines = append(nonEmptyLines, lines[i])
		}
	}

	return nonEmptyLines
}

func splitLinesFieldsWithFunc[T any](str string, conv func(string) (T, error)) ([][]T, error) {
	lines := SplitLines(str)
	linesFields := [][]T{}

	for i := 0; i < len(lines); i++ {
		fields := strings.Split(lines[i], " ")
		nonEmptyFields := []T{}

		for j := 0; j < len(fields); j++ {
			if fields[j] != "" {
				val, err := conv(fields[j])

				if err != nil {
					return nil, err
				}

				nonEmptyFields = append(nonEmptyFields, val)
			}
		}

		linesFields = append(linesFields, nonEmptyFields)
	}

	return linesFields, nil
}

func SplitLinesFields(str string) ([][]string, error) {
	return splitLinesFieldsWithFunc(str, func(s string) (string, error) { return s, nil })
}

func SplitLinesIntFields(str string) ([][]int, error) {
	return splitLinesFieldsWithFunc(str, func(s string) (int, error) {
		return strconv.Atoi(s)
	})
}

func SplitLinesFloatFields(str string) ([][]float64, error) {
	return splitLinesFieldsWithFunc(str, func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	})
}
