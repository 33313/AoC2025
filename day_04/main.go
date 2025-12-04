package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	part_01(lines)
	part_02(lines)
}


func part_01(lines []string) {
	accessible := 0
	for i, ln := range lines {
		for j, pos := range ln {
			if !(pos == '@') {
				continue
			}
			rolls := 0
			countRolls(&rolls, lines, i, ln, j)
			if rolls < 4 {
				accessible += 1
			}
		}
	}
	fmt.Printf("Part 1: %v\n", accessible)
}

func part_02(lines []string) {
	accessible := 0
	last := accessible
	next := lines
	for {
		next = part_02_recurse(next, &accessible)
		if accessible == last {
			break
		}
		last = accessible
	}
	fmt.Printf("Part 2: %v\n", accessible)
}

func part_02_recurse(lines []string, accessible *int) []string {
	var arr []string
	for i, ln := range lines {
		new_line := ""
		for j, pos := range ln {
			if !(pos == '@') {
				new_line += string(pos)
				continue
			}
			rolls := 0
			countRolls(&rolls, lines, i, ln, j)
			if rolls < 4 {
				new_line += string('.')
				*accessible++
			} else {
				new_line += string(pos)
			}
		}
		arr = append(arr, new_line)
	}
	return arr
}

func countIfRoll(x byte, counter *int) {
	if x == '@' {
		*counter++
	}
}

func countRolls(rolls *int, lines []string, i int, ln string, j int) {
	has_north := i > 0
	has_west := j > 0
	has_south := (i+1 < len(lines))
	has_east := (j+1 < len(ln))
	if has_north {
		countIfRoll(lines[i-1][j], rolls)
		if has_west {
			countIfRoll(lines[i-1][j-1], rolls)
		}
		if has_east {
			countIfRoll(lines[i-1][j+1], rolls)
		}
	}
	if has_west {
		countIfRoll(ln[j-1], rolls)
	}
	if has_south {
		countIfRoll(lines[i+1][j], rolls)
		if has_west {
			countIfRoll(lines[i+1][j-1], rolls)
		}
		if has_east {
			countIfRoll(lines[i+1][j+1], rolls)
		}
	}
	if has_east {
		countIfRoll(ln[j+1], rolls)
	}
}
