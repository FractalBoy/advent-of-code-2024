package solutions

import (
	"slices"
	"strconv"

	"github.com/FractalBoy/advent-of-code-2024/utils"
)

type Day6 struct{}

const (
	NORTH = iota
	SOUTH = iota
	EAST  = iota
	WEST  = iota
)

type Guard struct {
	Orientation int
	Location    Point
}

type Point struct {
	X int
	Y int
}

type Map struct {
	Size      int
	Obstacles []Point
}

func parseDay6Input(input string) (Guard, Map) {
	lines := utils.SplitLinesFieldsWithDelim(input, "")

	// Assume a square map
	myMap := Map{Size: len(lines), Obstacles: []Point{}}
	guard := Guard{Orientation: NORTH, Location: Point{X: 0, Y: 0}}

	for y, line := range lines {
		for x, cell := range line {
			if cell == "#" {
				myMap.Obstacles = append(myMap.Obstacles, Point{X: x, Y: y})
			} else if cell == "^" {
				guard.Location.X = x
				guard.Location.Y = y
			}
		}
	}

	return guard, myMap
}

func (d Day6) Part1(input string) (string, error) {
	guard, myMap := parseDay6Input(input)

	visited := []Point{}

	for guard.Location.X >= 0 && guard.Location.X < myMap.Size && guard.Location.Y >= 0 && guard.Location.Y < myMap.Size {
		if !slices.Contains(visited, guard.Location) {
			visited = append(visited, guard.Location)
		}

		switch guard.Orientation {
		case NORTH:
			guard.Location.Y--
		case SOUTH:
			guard.Location.Y++
		case EAST:
			guard.Location.X++
		case WEST:
			guard.Location.X--
		}

		if slices.Index(myMap.Obstacles, guard.Location) != -1 {
			switch guard.Orientation {
			case NORTH:
				guard.Orientation = EAST
				guard.Location.Y++
			case SOUTH:
				guard.Orientation = WEST
				guard.Location.Y--
			case EAST:
				guard.Orientation = SOUTH
				guard.Location.X--
			case WEST:
				guard.Orientation = NORTH
				guard.Location.X++
			}
		}
	}

	return strconv.Itoa(len(visited)), nil
}

func (d Day6) Part2(input string) (string, error) {
	return "", nil
}
