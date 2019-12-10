package main

import (
	"fmt"
	"strings"
)

func init() {
	days[8] = d8
}

func d8() {
	data := strings.TrimSpace(inputRaw())
	const width = 25
	const height = 6
	// data := "123456789012"
	// const width = 3
	// const height = 2
	pixPer := width * height
	numLayers := len(data) / pixPer
	if len(data)%pixPer != 0 {
		panic("SIZE MATCH")
	}
	layers := make([]string, numLayers)
	min := 9999
	for i := 0; i < numLayers; i++ {
		layers[i] = data[i*pixPer : i*pixPer+pixPer]
		if !*p2 {
			counts := map[rune]int{}
			for _, c := range layers[i] {
				counts[c]++
			}
			zeros := counts['0']
			prod := counts['1'] * counts['2']
			if zeros < min {
				v(zeros, prod)
				min = zeros
			}
		}
	}
	if !*p2 {
		return
	}
	for _, l := range layers {
		fmt.Println(l[:25])
	}
	for y := 0; y < height; y++ {
		fmt.Println()
		for x := 0; x < width; x++ {
			idx := y*width + x
			for _, l := range layers {
				ch := l[idx]
				if ch == '0' {
					fmt.Print(" ")
					break
				}
				if ch == '1' {
					fmt.Print("X")
					break
				}
			}
		}
	}
}
