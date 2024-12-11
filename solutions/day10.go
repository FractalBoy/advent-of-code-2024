package solutions

import (
	"slices"
	"strconv"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

type Day10 struct{}

func findHikingTrails(trailMap [][]int, y int, x int, reachable *[]int) {
	height := trailMap[y][x]

	if height == 9 {
		*reachable = append(*reachable, y*len(trailMap)+x)
		return
	}

	if y-1 >= 0 && trailMap[y-1][x] == height+1 {
		findHikingTrails(trailMap, y-1, x, reachable)
	}

	if y+1 < len(trailMap) && trailMap[y+1][x] == height+1 {
		findHikingTrails(trailMap, y+1, x, reachable)
	}

	if x-1 >= 0 && trailMap[y][x-1] == height+1 {
		findHikingTrails(trailMap, y, x-1, reachable)
	}

	if x+1 < len(trailMap[y]) && trailMap[y][x+1] == height+1 {
		findHikingTrails(trailMap, y, x+1, reachable)
	}
}

func (d Day10) Part1(input string) (string, error) {
	trailMap, err := utils.SplitLinesIntFieldsWithDelim(input, "")

	if err != nil {
		return "", err
	}

	count := 0

	for y, row := range trailMap {
		for x, col := range row {
			if col == 0 {
				reachable := []int{}
				findHikingTrails(trailMap, y, x, &reachable)
				slices.Sort(reachable)
				reachable = slices.Compact(reachable)
				count += len(reachable)
			}
		}
	}

	return strconv.Itoa(count), nil
}

func (d Day10) Part2(input string) (string, error) {
	trailMap, err := utils.SplitLinesIntFieldsWithDelim(input, "")

	if err != nil {
		return "", err
	}

	rating := 0

	for y, row := range trailMap {
		for x, col := range row {
			if col == 0 {
				reachable := []int{}
				findHikingTrails(trailMap, y, x, &reachable)
				rating += len(reachable)
			}
		}
	}

	return strconv.Itoa(rating), nil
}
