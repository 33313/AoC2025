package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAX_VAL = 100

func main() {
	dial_index := 50
	count := 0

	fp := "input.txt"
	file, _ := os.Open(fp)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		direction := ln[0]
		magnitude, _ := strconv.Atoi(ln[1:])
		if direction == 'L' {
			dial_index = (dial_index - magnitude + MAX_VAL) % MAX_VAL
		} else {
			dial_index = (dial_index + magnitude + MAX_VAL) % MAX_VAL
		}
		if dial_index == 0 {
			count++
		}
	}
	fmt.Print(count)
}
