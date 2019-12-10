package main

import (
	"fmt"
	"strings"
)

func init() {
	days[6] = d6
}

func d6() {
	dat := inputLines()

	bodies := map[string]map[string]bool{}

	for _, orb := range dat {
		parts := strings.Split(orb, ")")
		a, b := parts[0], parts[1]
		//v(a, b)
		if bodies[a] == nil {
			bodies[a] = map[string]bool{}
		}
		bodies[a][b] = true
	}
	var walk func(string, int) int
	walk = func(key string, depth int) int {
		//v(key, depth)
		children := bodies[key]
		total := depth
		for c := range children {
			total += walk(c, depth+1)
		}
		return total
	}
	fmt.Println(walk("COM", 0))
}
