package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

type Day17 struct{}

type Computer struct {
	ip           int
	a            int
	b            int
	c            int
	instructions []int
}

const (
	ADV = 0
	BXL = 1
	BST = 2
	JNZ = 3
	BXC = 4
	OUT = 5
	BDV = 6
	CDV = 7
)

func assemble(input string) (Computer, error) {
	lines := utils.SplitLinesFieldsWithDelim(input, ": ")
	computer := Computer{}

	for _, fields := range lines {
		if len(fields) != 2 {
			continue
		}

		if fields[0] == "Register A" {
			a, err := strconv.Atoi(fields[1])

			if err != nil {
				return computer, err
			}

			computer.a = a
		} else if fields[0] == "Register B" {
			b, err := strconv.Atoi(fields[1])

			if err != nil {
				return computer, err
			}

			computer.b = b
		} else if fields[0] == "Register C" {
			c, err := strconv.Atoi(fields[1])

			if err != nil {
				return computer, err
			}

			computer.c = c
		} else if fields[0] == "Program" {
			program := strings.Split(fields[1], ",")
			computer.instructions = make([]int, len(program))

			for i, inst := range program {
				inst, err := strconv.Atoi(inst)

				if err != nil {
					return computer, err
				}

				computer.instructions[i] = inst
			}
		}
	}

	return computer, nil
}

func (c Computer) run() string {
	out := []string{}

	for c.ip < len(c.instructions)-1 {
		instruction := c.instructions[c.ip]
		operand := c.instructions[c.ip+1]

		inc := 2

		switch instruction {
		case ADV:
			c.a = c.xdv(operand)
		case BXL:
			c.b = c.b ^ operand
		case BST:
			c.b = c.comboOperand(operand) % 8
		case JNZ:
			if c.a != 0 {
				c.ip = operand
				inc = 0
			}
		case BXC:
			c.b ^= c.c
		case OUT:
			out = append(out, strconv.Itoa(c.comboOperand(operand)%8))
		case BDV:
			c.b = c.xdv(operand)
		case CDV:
			c.c = c.xdv(operand)
		}

		c.ip += inc
	}

	return strings.Join(out, ",")
}

func (c Computer) comboOperand(operand int) int {
	if operand >= 0 && operand <= 3 {
		return operand
	}

	if operand == 4 {
		return c.a
	}

	if operand == 5 {
		return c.b
	}

	if operand == 6 {
		return c.c
	}

	panic(fmt.Sprintf("The combo operand %d is invalid", operand))
}

func (c Computer) xdv(operand int) int {
	numerator := c.a
	denominator := 1 << c.comboOperand(operand)

	return numerator / denominator
}

func (d Day17) Part1(input string) (string, error) {
	computer, err := assemble(input)

	if err != nil {
		return "", err
	}

	output := computer.run()

	return output, nil
}

func (d Day17) Part2(input string) (string, error) {
	return "", nil
}
