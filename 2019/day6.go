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
	parents := map[string]string{}

	for _, orb := range dat {
		parts := strings.Split(orb, ")")
		a, b := parts[0], parts[1]
		//v(a, b)
		if bodies[a] == nil {
			bodies[a] = map[string]bool{}
		}
		bodies[a][b] = true
		parents[b] = a
	}
	if !*p2 {
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
		return
	}
	//v(parents)
	current := parents["YOU"]
	i := 1
	weights := map[string]int{current: 0}
	for {
		v(current)
		current = parents[current]
		weights[current] = i
		i++
		if current == "COM" {
			break
		}
	}
	v(weights)
	current = parents["SAN"]
	j := 0
	for j < 1000 {
		v("?", current)
		if weights[current] != 0 {
			fmt.Println("!!!!!", j+weights[current])
			return
		}
		current = parents[current]
		j++
	}
	panic("WTF??")
}
