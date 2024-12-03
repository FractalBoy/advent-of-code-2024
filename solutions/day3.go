package solutions

import (
	"strconv"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

type Day3 struct{}

func calculateResults(input string, allowDisable bool) (string, error) {
	sawMulOpenParen := false
	sawNum1 := false
	sawComma := false

	disableCalc := false

	currNum1 := ""
	currNum2 := ""

	resetState := func() {
		sawMulOpenParen = false
		sawNum1 = false
		sawComma = false

		disableCalc = false

		currNum1 = ""
		currNum2 = ""
	}

	lines := utils.SplitLines(input)
	sum := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if allowDisable && j < len(lines[i])-1 && lines[i][j:j+2] == "do" {
				if j < len(lines[i])-4 && lines[i][j+2:j+5] == "n't" {
					disableCalc = true
					j += 4

					continue
				} else {
					disableCalc = false
					j++
				}
			}

			if disableCalc {
				continue
			}

			if !sawMulOpenParen {
				if j < len(lines[i])-3 && lines[i][j:j+4] == "mul(" {
					sawMulOpenParen = true
					j += 3
				}
			} else if !sawNum1 || sawNum1 && !sawComma {
				if lines[i][j] >= '0' && lines[i][j] <= '9' {
					sawNum1 = true
					currNum1 += string(lines[i][j])
				} else if lines[i][j] == ',' {
					sawComma = true
				} else {
					resetState()
				}
			} else {
				if lines[i][j] >= '0' && lines[i][j] <= '9' {
					currNum2 += string(lines[i][j])
				} else if lines[i][j] == ')' {
					num1, err := strconv.Atoi(currNum1)

					if err != nil {
						return "", nil
					}
					num2, err := strconv.Atoi(currNum2)

					if err != nil {
						return "", nil
					}

					sum += num1 * num2

					resetState()
				} else {
					resetState()
				}
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func (d Day3) Part1(input string) (string, error) {
	return calculateResults(input, false)
}

func (d Day3) Part2(input string) (string, error) {
	return calculateResults(input, true)
}
