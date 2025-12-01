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
	dial := Dial(50)
	count := 0

	fp := "input.txt"
	file, _ := os.Open(fp)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
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
	fmt.Println(count)
}
