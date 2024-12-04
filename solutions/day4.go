package solutions

import (
	"strconv"
	"strings"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

type Day4 struct{}

const XMAS = "XMAS"
const SAMX = "SAMX"
const MAS = "MAS"
const SAM = "SAM"

func (d Day4) Part1(input string) (string, error) {
	grid := utils.SplitLinesFieldsWithDelim(input, "")
	count := 0

	for i := range grid {
		for j := range grid[i] {
			if j+3 < len(grid[i]) {
				right := strings.Join(grid[i][j:j+4], "")

				if right == XMAS || right == SAMX {
					count++
				}

				if i-3 >= 0 {
					upRight := grid[i][j] + grid[i-1][j+1] + grid[i-2][j+2] + grid[i-3][j+3]

					if upRight == XMAS || upRight == SAMX {
						count++
					}
				}

				if i+3 < len(grid) {
					downRight := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2] + grid[i+3][j+3]

					if downRight == XMAS || downRight == SAMX {
						count++
					}
				}
			}

			if i+3 < len(grid) {
				down := grid[i][j] + grid[i+1][j] + grid[i+2][j] + grid[i+3][j]

				if down == XMAS || down == SAMX {
					count++
				}
			}
		}
	}

	return strconv.Itoa(count), nil
}

func (d Day4) Part2(input string) (string, error) {
	grid := utils.SplitLinesFieldsWithDelim(input, "")

	count := 0

	for i := range grid {
		for j := range grid[i] {
			if i > 0 && i < len(grid)-1 && j > 0 && j < len(grid[i])-1 {
				diag1 := grid[i-1][j-1] + grid[i][j] + grid[i+1][j+1]
				diag2 := grid[i+1][j-1] + grid[i][j] + grid[i-1][j+1]

				if (diag1 == MAS || diag1 == SAM) && (diag2 == MAS || diag2 == SAM) {
					count++
				}
			}
		}
	}

	return strconv.Itoa(count), nil
}
