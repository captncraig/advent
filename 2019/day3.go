package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func init() {
	days[3] = d3
}

func d3() {
	getDxDy := func(dir string) (int, int) {
		switch dir {
		case "R":
			return 1, 0
		case "L":
			return -1, 0
		case "U":
			return 0, 1
		case "D":
			return 0, -1
		}
		log.Fatal("Unknown dir")
		return 0, 0
	}
	lines := inputLinesN(2)
	a := strings.Split(lines[0], ",")
	b := strings.Split(lines[1], ",")
	traverse := func(instructions []string, f func(int, int, string)) {
		x, y := 0, 0
		for _, instruction := range instructions {
			dir := instruction[:1]
			count, _ := strconv.Atoi(instruction[1:])
			v(dir, count)
			dx, dy := getDxDy(dir)
			for i := 0; i < count; i++ {
				x += dx
				y += dy
				f(x, y, fmt.Sprintf("%d,%d", x, y))
			}
		}
	}
	aPoints := map[string]bool{}
	traverse(a, func(x, y int, coord string) {
		v(coord)
		aPoints[coord] = true
	})
	min := 999999
	traverse(b, func(x, y int, coord string) {
		if x < 0 {
			x = -x
		}
		if y < 0 {
			y = -y
		}
		if aPoints[coord] {
			v(coord, "!!!", x+y)
			if x+y < min {
				min = x + y
			}
		}
	})
	fmt.Println(min, "<<<<<<")
}
