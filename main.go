package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"

	"github.com/FractalBoy/advent-of-code-2024/aoc"
	"github.com/FractalBoy/advent-of-code-2024/solutions"
)

var dayNum = flag.Int("day", 0, "The day to execute")
var part1 = flag.Bool("part1", false, "Execute part 1")
var part2 = flag.Bool("part2", false, "Execute part 2")

func main() {
	flag.Parse()

	if *dayNum == 0 {
		fmt.Fprintln(os.Stderr, "Invalid day number.")
		os.Exit(1)
	}

	var days map[int]aoc.Day = make(map[int]aoc.Day)

	days[1] = solutions.Day1{}
	days[2] = solutions.Day2{}
	days[3] = solutions.Day3{}
	days[4] = solutions.Day4{}
	days[5] = solutions.Day5{}
	days[6] = solutions.Day6{}
	days[7] = solutions.Day7{}
	days[8] = solutions.Day8{}
	days[9] = solutions.Day9{}
	days[10] = solutions.Day10{}
	days[11] = solutions.Day11{}

	day, ok := days[*dayNum]

	if !ok {
		fmt.Fprintf(os.Stderr, "Day %d unimplemented\n", *dayNum)
		os.Exit(1)
	}

	input, err := aoc.DayInput(*dayNum)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if !*part1 && !*part2 {
		fmt.Fprintln(os.Stderr, "No part specified.")
		os.Exit(1)
	}

	if *part1 {
		executePart(day, 1, input)
	}

	if *part2 {
		executePart(day, 2, input)
	}
}

func executePart(day aoc.Day, part int, input string) {
	method := reflect.ValueOf(day).MethodByName(fmt.Sprintf("Part%d", part))

	out := method.Call([]reflect.Value{reflect.ValueOf(input)})

	solution := out[0].String()
	err := out[1].Interface()

	if err == nil {
		println(solution)
	} else {
		err := err.(error)

		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
