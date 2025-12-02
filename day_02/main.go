package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	input_str := string(bytes)
	arr := strings.Split(input_str, ",")
	fmt.Printf("Part 1: %v\n", inRangeDo(arr, part_01))
	fmt.Printf("Part 2: %v\n", inRangeDo(arr, part_02))
}

func inRangeDo(arr []string, fn func(int, string, *uint64)) uint64 {
	sum := uint64(0)
	for i := range arr {
		from_to := strings.Split(arr[i], "-")
		from, _ := strconv.Atoi(from_to[0])
		to, _ := strconv.Atoi(from_to[1])
		for n := from; n <= to; n++ {
			n_str := strconv.Itoa(n)
			fn(n, n_str, &sum)
		}
	}
	return sum
}

func part_01(n int, n_str string, sum *uint64) {
	if len(n_str)%2 != 0 {
		return
	}
	half := len(n_str) / 2
	left := n_str[:half]
	right := n_str[half:]
	if left == right {
		*sum += uint64(n)
	}
}

func part_02(n int, n_str string, sum *uint64) {
	for j := 1; j <= len(n_str)/2; j++ {
		if len(n_str)%j != 0 {
			continue
		}
		pattern := n_str[:j]
		str := strings.Repeat(pattern, len(n_str)/j)
		if str == n_str {
			*sum += uint64(n)
			break
		}
	}
}
