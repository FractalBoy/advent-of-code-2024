package solutions

import (
	"strconv"
	"strings"
)

type Day11 struct{}

func blink(stones map[int]int) (map[int]int, error) {
	newStones := make(map[int]int)

	for stone := range stones {
		strStone := strconv.Itoa(stone)
		count := stones[stone]

		if stone == 0 {
			val, ok := newStones[1]

			if ok {
				newStones[1] = val + count
			} else {
				newStones[1] = count
			}
		} else if len(strStone)%2 == 0 {
			firstHalf := strStone[0 : len(strStone)/2]
			secondHalf := strStone[len(strStone)/2 : len(strStone)]

			first, err := strconv.Atoi(firstHalf)

			if err != nil {
				return nil, err
			}

			second, err := strconv.Atoi(secondHalf)

			if err != nil {
				return nil, err
			}

			val, ok := newStones[first]

			if ok {
				newStones[first] = val + count
			} else {
				newStones[first] = count
			}

			val, ok = newStones[second]

			if ok {
				newStones[second] = val + count
			} else {
				newStones[second] = count
			}
		} else {
			replacement := stone * 2024

			val, ok := newStones[replacement]

			if ok {
				newStones[replacement] = val + count
			} else {
				newStones[replacement] = count
			}
		}
	}

	return newStones, nil
}

func parseInput(input string) (map[int]int, error) {
	stones := make(map[int]int)

	input = strings.Replace(input, "\n", "", -1)
	for _, stone := range strings.Split(input, " ") {
		num, err := strconv.Atoi(stone)

		if err != nil {
			return nil, err
		}

		count, ok := stones[num]

		if ok {
			stones[num] = count + 1
		} else {
			stones[num] = 1
		}
	}

	return stones, nil
}

func (d Day11) Part1(input string) (string, error) {
	stones, err := parseInput(input)

	if err != nil {
		return "", err
	}

	for i := 0; i < 25; i++ {
		var err error
		stones, err = blink(stones)

		if err != nil {
			return "", err
		}
	}

	total := 0

	for _, count := range stones {
		total += count
	}

	return strconv.Itoa(total), nil
}

func (d Day11) Part2(input string) (string, error) {
	stones, err := parseInput(input)

	if err != nil {
		return "", err
	}

	for i := 0; i < 75; i++ {
		var err error
		stones, err = blink(stones)

		if err != nil {
			return "", err
		}
	}

	total := 0

	for _, count := range stones {
		total += count
	}

	return strconv.Itoa(total), nil
}
