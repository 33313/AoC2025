package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	var ranges [][]int64
	var products []string
	scanner := bufio.NewScanner(f)
	is_listing_products := false
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			is_listing_products = true
			continue
		}
		if is_listing_products {
			products = append(products, text)
		} else {
			r := strings.Split(text, "-")
			start, _ := strconv.Atoi(r[0])
			stop, _ := strconv.Atoi(r[1])
			ranges = append(ranges, []int64{int64(start), int64(stop)})
		}
	}
	part_01(ranges, products)
	part_02(ranges)
}

func part_01(ranges [][]int64, products []string) {
	fresh := 0
	for _, p := range products {
		product, _ := strconv.Atoi(p)
		for _, v := range ranges {
			if v[0] <= int64(product) && int64(product) <= v[1] {
				fresh++
				break
			}
		}
	}
	fmt.Printf("Part 1: %v\n", fresh)
}

func part_02(ranges [][]int64) {
	var fresh int64
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	start := ranges[0][0]
	end := ranges[0][1]
	for i := 1; i < len(ranges); i++ {
		next_start := ranges[i][0]
		next_end := ranges[i][1]
		if next_start <= end {
			if next_end > end {
				end = next_end
			}
		} else {
			fresh += end - start + 1
			start = next_start
			end = next_end
		}
	}
	fresh += end - start + 1
	fmt.Printf("Part 2: %v\n", fresh)
}
