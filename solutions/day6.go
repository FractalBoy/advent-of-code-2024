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

func getVisitedLocations(guard Guard, map_ Map) []Point {
	visited := []Point{}

	for guard.Location.X >= 0 && guard.Location.X < map_.Size && guard.Location.Y >= 0 && guard.Location.Y < map_.Size {
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

		if slices.Index(map_.Obstacles, guard.Location) != -1 {
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

	return visited
}

func (d Day6) Part1(input string) (string, error) {
	guard, map_ := parseDay6Input(input)

	visited := getVisitedLocations(guard, map_)

	return strconv.Itoa(len(visited)), nil
}

func (d Day6) Part2(input string) (string, error) {
	guard, map_ := parseDay6Input(input)

	visited := getVisitedLocations(guard, map_)
	cycles := 0

	for _, point := range visited {
		clonedGuard := Guard{Location: Point{X: guard.Location.X, Y: guard.Location.Y}, Orientation: guard.Orientation}

		obstacles := slices.Clone(map_.Obstacles)
		obstacles = append(obstacles, point)

		guards := make(map[int]map[int]map[int]bool)

		for clonedGuard.Location.X >= 0 && clonedGuard.Location.X < map_.Size && clonedGuard.Location.Y >= 0 && clonedGuard.Location.Y < map_.Size {
			switch clonedGuard.Orientation {
			case NORTH:
				clonedGuard.Location.Y--
			case SOUTH:
				clonedGuard.Location.Y++
			case EAST:
				clonedGuard.Location.X++
			case WEST:
				clonedGuard.Location.X--
			}

			if slices.Index(obstacles, clonedGuard.Location) != -1 {
				switch clonedGuard.Orientation {
				case NORTH:
					clonedGuard.Orientation = EAST
					clonedGuard.Location.Y++
				case SOUTH:
					clonedGuard.Orientation = WEST
					clonedGuard.Location.Y--
				case EAST:
					clonedGuard.Orientation = SOUTH
					clonedGuard.Location.X--
				case WEST:
					clonedGuard.Orientation = NORTH
					clonedGuard.Location.X++
				}
			} else {
				map1, ok := guards[clonedGuard.Location.X]

				if ok {
					map2, ok := map1[clonedGuard.Location.Y]

					if ok {
						exists, ok := map2[clonedGuard.Orientation]

						if ok && exists {
							cycles++
							break
						}

						guards[clonedGuard.Location.X][clonedGuard.Location.Y][clonedGuard.Orientation] = true
					} else {
						guards[clonedGuard.Location.X][clonedGuard.Location.Y] = make(map[int]bool)
						guards[clonedGuard.Location.X][clonedGuard.Location.Y][clonedGuard.Orientation] = true
					}
				} else {
					guards[clonedGuard.Location.X] = make(map[int]map[int]bool)
					guards[clonedGuard.Location.X][clonedGuard.Location.Y] = make(map[int]bool)
					guards[clonedGuard.Location.X][clonedGuard.Location.Y][clonedGuard.Orientation] = true
				}
			}
		}
	}

	return strconv.Itoa(cycles), nil
}
