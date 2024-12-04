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

func splitLinesFieldsWithFunc[T any](str string, delim string, conv func(string) (T, error)) ([][]T, error) {
	lines := SplitLines(str)
	linesFields := [][]T{}

	for i := 0; i < len(lines); i++ {
		fields := strings.Split(lines[i], delim)
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

func SplitLinesFieldsWithDelim(str string, delim string) [][]string {
	lines, _ := splitLinesFieldsWithFunc(str, delim, func(s string) (string, error) { return s, nil })
	return lines
}

func SplitLinesFields(str string) [][]string {
	return SplitLinesFieldsWithDelim(str, " ")
}

func SplitLinesIntFieldsWithDelim(str string, delim string) ([][]int, error) {
	return splitLinesFieldsWithFunc(str, delim, func(s string) (int, error) {
		return strconv.Atoi(s)
	})
}

func SplitLinesIntFields(str string) ([][]int, error) {
	return SplitLinesIntFieldsWithDelim(str, " ")
}

func SplitLinesFloatFieldsWithDelim(str string, delim string) ([][]float64, error) {
	return splitLinesFieldsWithFunc(str, delim, func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	})
}

func SplitLinesFloatFields(str string) ([][]float64, error) {
	return SplitLinesFloatFieldsWithDelim(str, " ")
}
