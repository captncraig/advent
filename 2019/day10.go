package main

import (
	"fmt"
	"math"
)

func init() {
	days[10] = d10
}

type point struct {
	x, y int
}

func d10() {
	lines := inputLines()
	fmt.Println(lines)

	asteroids := []point{}

	for y, l := range lines {
		for x, c := range l {
			if c != '.' {
				asteroids = append(asteroids, point{x, y})
			}
		}
	}
	v(asteroids)
	max := 0
	for i, p := range asteroids {
		uniq := map[float64]bool{}
		for j, p2 := range asteroids {
			if i == j {
				continue
			}
			angle := math.Atan2(float64(p2.y-p.y), float64(p2.x-p.x))
			//v(p, p2, angle)
			uniq[angle] = true
		}
		if len(uniq) > max {
			max = len(uniq)
			v(p, max)
		}
	}
	fmt.Println(max)
}

func colinear(a, b, c point) bool {
	y1, y2, y3 := a.y, b.y, c.y
	x1, x2, x3 := a.x, b.x, c.x
	return (y1-y2)*(x1-x3) == (y1-y3)*(x1-x2)
}
