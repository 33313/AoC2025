package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var signs []string
	var numbers [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		if _, err := strconv.Atoi(text[0]); err != nil {
			signs = text
			continue
		}
		for i, v := range text {
			if i >= len(numbers) {
				numbers = append(numbers, []int{})
			}
			num, _ := strconv.Atoi(v)
			numbers[i] = append(numbers[i], num)
		}
	}

	part_01(numbers, signs)
	part_02(numbers, signs)
}

func part_01(numbers [][]int, signs []string) {
	var sum int
	for i := range numbers {
		col_sum := 0
		sign := signs[i]
		for j := range numbers[i] {
			if j == 0 {
				col_sum += numbers[i][j]
			} else if sign == "*" {
				col_sum *= numbers[i][j]
			} else {
				col_sum += numbers[i][j]
			}
		}
		sum += col_sum
	}
	fmt.Printf("Part 1: %v\n", sum)
}

func part_02(numbers [][]int, signs []string) {
	var sum int
	fmt.Printf("Part 2: %v\n", sum)
}
