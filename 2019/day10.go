package main

import (
	"fmt"
	"math"
	"sort"
)

func init() {
	days[10] = d10
}

type point struct {
	x, y int
}

func (p point) D(p2 point) int {
	return int(math.Abs(float64(p2.y-p.y)) + math.Abs(float64(p2.x-p.x)))
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
	//v(asteroids)
	max := 0
	var maxUniq map[float64][]point
	for _, p := range asteroids {
		uniq := map[float64][]point{}
		a2 := make([]point, len(asteroids))
		copy(a2, asteroids)
		sort.Slice(a2, func(i, j int) bool { return a2[i].D(p) < a2[j].D(p) })
		for _, p2 := range a2 {
			if p == p2 {
				continue
			}
			angle := math.Atan2(float64(p2.y-p.y), float64(p2.x-p.x))
			angle *= (180 / math.Pi)
			angle += 90
			if angle < 0 {
				angle += 360
			}
			uniq[angle] = append(uniq[angle], p2)
		}
		if len(uniq) > max {
			max = len(uniq)
			maxUniq = uniq
			v(p, max, uniq)
		}
	}
	v(math.Pi / 2)
	fmt.Println(max)
	angles := []float64{}
	for k := range maxUniq {
		angles = append(angles, k)
	}
	sort.Float64s(angles)
	v(angles)
	n := 0
	for {
		for _, angle := range angles {
			list := maxUniq[angle]
			if len(list) > 0 {
				//v(angle, list)
				fmt.Println(n+1, list[0])
				maxUniq[angle] = list[1:]
				n++
				if n > 199 {
					return
				}
			}
		}
	}

}

func colinear(a, b, c point) bool {
	y1, y2, y3 := a.y, b.y, c.y
	x1, x2, x3 := a.x, b.x, c.x
	return (y1-y2)*(x1-x3) == (y1-y3)*(x1-x2)
}
