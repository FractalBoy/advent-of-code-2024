package solutions

import (
	"math"
	"strconv"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

type Day8 struct{}

func isAntinode(lines [][]string, x int, y int) bool {
	for yIncr := len(lines) * -1; yIncr < len(lines); yIncr++ {
		for xIncr := len(lines) * -1; xIncr < len(lines); xIncr++ {
			if xIncr == 0 && yIncr == 0 {
				continue
			}

			frequencies := findFrequencies(lines, x, y, xIncr, yIncr)

			for _, dists := range frequencies {
				for i, dist1 := range dists {
					for j, dist2 := range dists {
						if i == j {
							continue
						}

						if dist2 == dist1*2 {
							return true
						}
					}
				}
			}
		}
	}

	return false
}

func isAntinodeWithResonantHarmonics(lines [][]string, x int, y int) bool {
	if lines[y][x] != "." {
		return true
	}

	for yIncr := len(lines) * -1; yIncr < len(lines); yIncr++ {
		for xIncr := len(lines) * -1; xIncr < len(lines); xIncr++ {
			if xIncr == 0 && yIncr == 0 {
				continue
			}

			frequencies := findFrequencies(lines, x, y, xIncr, yIncr)

			for _, dists := range frequencies {
				if len(dists) >= 2 {
					return true
				}
			}
		}
	}

	return false
}

func findFrequencies(lines [][]string, x int, y int, xIncr int, yIncr int) map[string][]int {
	xNudge := xIncr
	yNudge := yIncr

	frequencies := make(map[string][]int)

	for x+xNudge < len(lines[y]) && x+xNudge >= 0 && y+yNudge < len(lines) && y+yNudge >= 0 {
		newX := x + xNudge
		newY := y + yNudge

		node := lines[newY][newX]
		xDist := math.Abs(float64(xNudge))
		yDist := math.Abs(float64(yNudge))

		// Taxicab distance
		dist := int(xDist) + int(yDist)

		if node != "." {
			_, ok := frequencies[node]

			if ok {
				frequencies[node] = append(frequencies[node], dist)
			} else {
				frequencies[node] = []int{dist}
			}
		}

		xNudge += xIncr
		yNudge += yIncr
	}

	return frequencies
}

func (d Day8) Part1(input string) (string, error) {
	lines := utils.SplitLinesFieldsWithDelim(input, "")

	count := 0

	for y, line := range lines {
		for x := range line {
			if isAntinode(lines, x, y) {
				count++
			}
		}
	}

	return strconv.Itoa(count), nil
}

func (d Day8) Part2(input string) (string, error) {
	lines := utils.SplitLinesFieldsWithDelim(input, "")

	count := 0

	for y, line := range lines {
		for x := range line {
			if isAntinodeWithResonantHarmonics(lines, x, y) {
				count++
			}
		}
	}

	return strconv.Itoa(count), nil
}
