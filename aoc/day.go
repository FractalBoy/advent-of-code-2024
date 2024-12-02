package aoc

type Day interface {
	Part1(input string) (string, error)
	Part2(input string) (string, error)
}
