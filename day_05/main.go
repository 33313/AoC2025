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
	var ranges []string
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
			ranges = append(ranges, text)
		}
	}
	part_01(ranges, products)
	part_02(ranges, products)
}

func part_01(ranges []string, products []string) {
	fresh := 0
	var rngs [][]int
	for _, v := range ranges {
		r := strings.Split(v, "-")
		start, _ := strconv.Atoi(r[0])
		stop, _ := strconv.Atoi(r[1])
		rngs = append(rngs, []int{start, stop})
	}
	for _, p := range products {
		product, _ := strconv.Atoi(p)
		for _, v := range rngs {
			if v[0] <= product && product <= v[1] {
				fresh++
				break
			}
		}
	}
	fmt.Printf("Part 1: %v\n", fresh)
}

func part_02(ranges []string, products []string) {
	fresh := 0
	fmt.Printf("Part 2: %v\n", fresh)
}
