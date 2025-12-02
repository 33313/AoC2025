package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAX_VAL = 100

type Dial int

func (i *Dial) Decr() {
	curr := *i
	temp := (curr - 1 + MAX_VAL) % MAX_VAL
	*i = Dial(temp)
}

func (i *Dial) Incr() {
	curr := *i
	temp := (curr + 1 + MAX_VAL) % MAX_VAL
	*i = Dial(temp)
}

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
	dial := 50
	count := 0
	for _, ln := range lines {
		direction := ln[0]
		magnitude, _ := strconv.Atoi(ln[1:])
		if direction == 'L' {
			dial = (dial - magnitude + MAX_VAL) % MAX_VAL
		} else {
			dial = (dial + magnitude + MAX_VAL) % MAX_VAL
		}
		if dial == 0 {
			count++
		}
	}
	fmt.Printf("Part 1: %v\n", count)
}

func part_02(lines []string) {
	dial := Dial(50)
	count := 0
	for _, ln := range lines {
		direction := ln[0]
		magnitude, _ := strconv.Atoi(ln[1:])
		for range magnitude {
			if direction == 'L' {
				dial.Decr()
			} else {
				dial.Incr()
			}
			if dial == 0 {
				count++
			}
		}
	}
	fmt.Printf("Part 2: %v\n", count)
}
