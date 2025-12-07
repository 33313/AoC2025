package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var lines []string
	var signs []string
	var numbers [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, reverse(scanner.Text()))
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
	part_02(lines)
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

func part_02(lines []string) {
	var sum int
	nums := lines[:len(lines)-1]
	signs := lines[len(lines)-1]

	gaps := make(map[int]bool)
	var last_rune rune = ' '
	for i, v := range signs {
		if last_rune == '*' || last_rune == '+' {
			gaps[i] = true
		}
		last_rune = v
	}

	parsed := make([]string, len(signs))
	for _, ln := range nums {
		for pos, char := range ln {
			if gaps[pos] {
				continue
			}
			parsed[pos] += string(char)
		}
	}

	temp := []string{}
	for i := range parsed {
		if gaps[i] {
			temp = []string{}
			continue
		}
		temp = append(temp, parsed[i])
		if signs[i] == '*' {
			temp_sum := 1
			for _, n_str := range temp {
				n_str = strings.TrimSpace(n_str)
				n, _ := strconv.Atoi(n_str)
				temp_sum *= n
			}
			sum += temp_sum
			continue
		} else if signs[i] == '+' {
			temp_sum := 0
			for _, n_str := range temp {
				n_str = strings.TrimSpace(n_str)
				n, _ := strconv.Atoi(n_str)
				temp_sum += n
			}
			sum += temp_sum
			continue
		}
	}
	fmt.Printf("Part 2: %v\n", sum)
}
