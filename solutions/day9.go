package solutions

import (
	"maps"
	"slices"
	"strconv"
)

type Day9 struct{}

type Sector struct {
	start int
	end   int
}

func parseDiskMap(input string) []int {
	disk := []int{}

	lastEnd := 0

	for i, char := range input {
		size := int(char - '0')
		end := lastEnd + size

		id := -1

		if i%2 == 0 {
			id = i / 2
		}

		for j := 0; j < size; j++ {
			disk = append(disk, id)
		}

		lastEnd = end
	}

	return disk
}

func (d Day9) Part1(input string) (string, error) {
	diskMap := parseDiskMap(input)

	for i := len(diskMap) - 1; i >= 0; i-- {
		if diskMap[i] == -1 {
			continue
		}

		for j := 0; j < i; j++ {
			if diskMap[j] == -1 {
				diskMap[j] = diskMap[i]
				diskMap[i] = -1
				break
			}
		}
	}

	checksum := 0

	for i, id := range diskMap {
		if id == -1 {
			continue
		}

		checksum += i * id
	}

	return strconv.Itoa(checksum), nil
}

func (d Day9) Part2(input string) (string, error) {
	diskMap := parseDiskMap(input)

	files := make(map[int]*Sector)
	free := []*Sector{}

	for i := 0; i < len(diskMap); {
		j := i

		for ; j < len(diskMap) && diskMap[j] == diskMap[i]; j++ {
		}

		if diskMap[i] == -1 {
			free = append(free, &Sector{start: i, end: j})
		} else {
			files[diskMap[i]] = &Sector{start: i, end: j}
		}

		i = j
	}

	for _, id := range slices.SortedFunc(maps.Keys(files), func(a int, b int) int { return b - a }) {
		file := files[id]
		size := file.end - file.start

		for i, sector := range free {
			if sector.end <= file.start && sector.end-sector.start >= size {
				file.start = sector.start
				file.end = file.start + size

				sector.start += size

				if sector.end == sector.start {
					// Remove free space
					free = slices.Delete(free, i, i+1)
				}

				break
			}
		}
	}

	checksum := 0

	for id := range files {
		file := files[id]

		for i := file.start; i < file.end; i++ {
			checksum += i * id
		}
	}

	return strconv.Itoa(checksum), nil
}
