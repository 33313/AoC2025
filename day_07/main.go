package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TachyonManifold struct {
	Splitters []map[int]bool
	BeamIdxs  map[int]bool
	Hits      int
}
type QuantumTachyonManifold struct {
	Splitters []map[int]bool
	BeamIdxs  map[int]int
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	tm := TachyonManifold{
		BeamIdxs: make(map[int]bool),
		Hits:     0,
	}
	found_start := false
	start_idx := -1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if !found_start {
			found_start = true
			i := strings.Index(scanner.Text(), "S")
			tm.BeamIdxs[i-1] = true
			start_idx = i - 1
			continue
		}
		if i := strings.Index(scanner.Text(), "^"); i == -1 {
			continue
		}
		text := strings.Split(scanner.Text(), "")
		si := make(map[int]bool)
		for i, v := range text[1:] {
			if v == "^" {
				si[i] = true
			}
		}
		tm.Splitters = append(tm.Splitters, si)
	}
	qtm := QuantumTachyonManifold{
		Splitters: tm.Splitters,
		BeamIdxs:  make(map[int]int),
	}
	qtm.BeamIdxs[start_idx] = 1
	part_01(tm)
	part_02(qtm)
}

func part_01(tm TachyonManifold) {
	for _, splitters := range tm.Splitters {
		newBeamIdxs := make(map[int]bool)
		for beam_idx := range tm.BeamIdxs {
			if splitters[beam_idx] {
				tm.Hits++
				newBeamIdxs[beam_idx-1] = true
				newBeamIdxs[beam_idx+1] = true
			} else {
				newBeamIdxs[beam_idx] = true
			}
		}
		tm.BeamIdxs = newBeamIdxs
	}
	fmt.Printf("Part 1: %v\n", tm.Hits)
}

func part_02(qtm QuantumTachyonManifold) {
	for _, splitters := range qtm.Splitters {
		newBeamIdxs := make(map[int]int)
		for col, count := range qtm.BeamIdxs {
			if splitters[col] {
				newBeamIdxs[col-1] += count
				newBeamIdxs[col+1] += count
			} else {
				newBeamIdxs[col] += count
			}
		}
		qtm.BeamIdxs = newBeamIdxs
	}
	timelines := 0
	for _, n := range qtm.BeamIdxs {
		timelines += n
	}
	fmt.Printf("Part 2: %v\n", timelines)
}
