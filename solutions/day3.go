package solutions

import (
	"strconv"
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

	sum := 0

	for i := 0; i < len(input); i++ {
		if allowDisable && i < len(input)-1 && input[i:i+2] == "do" {
			if i < len(input)-4 && input[i+2:i+5] == "n't" {
				disableCalc = true
				i += 4

				continue
			} else {
				disableCalc = false
				i++
			}
		}

		if disableCalc {
			continue
		}

		if !sawMulOpenParen {
			if i < len(input)-3 && input[i:i+4] == "mul(" {
				sawMulOpenParen = true
				i += 3
			}
		} else if !sawNum1 || sawNum1 && !sawComma {
			if input[i] >= '0' && input[i] <= '9' {
				sawNum1 = true
				currNum1 += string(input[i])
			} else if input[i] == ',' {
				sawComma = true
			} else {
				resetState()
			}
		} else {
			if input[i] >= '0' && input[i] <= '9' {
				currNum2 += string(input[i])
			} else if input[i] == ')' {
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

	return strconv.Itoa(sum), nil
}

func (d Day3) Part1(input string) (string, error) {
	return calculateResults(input, false)
}

func (d Day3) Part2(input string) (string, error) {
	return calculateResults(input, true)
}
