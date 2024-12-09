package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(inpath string) (string, error) {
	bytes, err := os.ReadFile(inpath)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(bytes[:]), "\n"), nil
}

func checksum(disk []int) int {
	output := 0
	for i, f := range disk {
		if f >= 0 {
			output += i * f
		}
	}
	return output
}

func part1(line string) int {
	leng := 0
	for _, c := range line {
		leng += int(c - '0')
	}
	disk := make([]int, 0, leng)
	for i, c := range line {
		n := int(c - '0')
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				id := i / 2
				disk = append(disk, id)
			} else {
				disk = append(disk, -1)
			}
		}
	}

	e := len(disk) - 1
	for i := 0; i < e; i++ {
		if disk[i] < 0 {
			disk[i] = disk[e]
			disk[e] = -1
			for disk[e] < 0 {
				e--
			}
		}
	}

	return checksum(disk)
}

type space struct {
	index, size, file int
}

func part2(line string) int {
	leng := 0
	for _, c := range line {
		leng += int(c - '0')
	}
	disk := make([]int, 0, leng)
	spaces := []space{}
	for i, c := range line {
		size := int(c - '0')
		if i%2 == 0 {
			file := i / 2
			spaces = append(spaces, space{len(disk), size, file})
			for j := 0; j < size; j++ {
				disk = append(disk, file)
			}
		} else {
			spaces = append(spaces, space{len(disk), size, -1})
			for j := 0; j < size; j++ {
				disk = append(disk, -1)
			}
		}
	}

	for f := len(spaces) - 1; f > 0; f -= 2 {
		file := &spaces[f]
		for g := 1; g < f; g += 2 {
			gap := &spaces[g]
			if gap.size >= file.size {
				for j := 0; j < file.size; j++ {
					disk[gap.index+j] = file.file
					disk[file.index+j] = -1
				}
				gap.size -= file.size
				gap.index += file.size
				break
			}
		}
	}

	return checksum(disk)
}

func expect(expected, actual int) bool {
	if expected != actual {
		fmt.Printf("got %d but expected %d\n", actual, expected)
	}
	return expected != actual
}

func main() {
	inEx, err := readFile("day09/example.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	if expect(1928, part1(inEx)) {
		return
	}
	if expect(2858, part2(inEx)) {
		return
	}

	inCh, err := readFile("day09/challenge.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("p1:", part1(inCh))
	fmt.Println("p2:", part2(inCh))
}
