package main

import (
	"bufio"
	"fmt"
	"math"
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
	joltage := 0
	for _, ln := range lines {
		first := 0
		second := 0
		for i := range ln {
			val := int(ln[i] - '0')
			if val > first {
				if i+1 < len(ln) {
					first = val
					second = int(ln[i+1] - '0')
				} else {
					second = val
				}
			} else if val > second {
				second = val
			}
		}
		joltage += 10 * first
		joltage += second
	}
	fmt.Printf("Part 1: %v\n", joltage)
}

func part_02(lines []string) {
	joltage := 0
	for _, ln := range lines {
		batteries := make([]int, 12)
		for i := range ln {
			val := int(ln[i] - '0')
			remaining_distance := len(ln) - (i + 1)

			for b_i, b_val := range batteries {
				batteries_space := len(batteries) - (b_i + 1)
				if val > b_val {
					if batteries_space <= remaining_distance {
						batteries[b_i] = val
						for k := b_i + 1; k < len(batteries); k++ {
							batteries[k] = 0
						}
						break
					}
				}
			}
		}
		for i := range batteries {
			joltage += int(math.Pow(10, float64(len(batteries)-(i+1)))) * batteries[i]
		}
	}
	fmt.Printf("Part 2: %v\n", joltage)
}
