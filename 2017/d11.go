package main

import (
	"log"
	"strings"
)

/*
--- Day 11: Hex Ed ---
Crossing the bridge, you've barely reached the other side of the stream when a program comes up to you, clearly in distress. "It's my child process," she says, "he's gotten lost in an infinite grid!"

Fortunately for her, you have plenty of experience with infinite grids.

Unfortunately for you, it's a hex grid.

The hexagons ("hexes") in this grid are aligned such that adjacent hexes can be found to the north, northeast, southeast, south, southwest, and northwest:

  \ n  /
nw +--+ ne
  /    \
-+      +-
  \    /
sw +--+ se
  / s  \
You have the path the child process took. Starting where he started, you need to determine the fewest number of steps required to reach him. (A "step" means to move from the hex you are in to any adjacent hex.)

For example:

ne,ne,ne is 3 steps away.
ne,ne,sw,sw is 0 steps away (back where you started).
ne,ne,s,s is 2 steps away (se,se).
se,sw,se,sw,sw is 3 steps away (s,s,sw).
*/
/*

 */

//https://stackoverflow.com/questions/2459402/hexagonal-grid-coordinates-to-pixel-coordinates for coordinate system
var _ = d(11, func(part2 bool, input string) int {
	r, g, b := 0, 0, 0
	dist := func() int {
		r2 := r
		g2 := g
		b2 := b
		if r < 0 {
			r2 = -r
		}
		if g < 0 {
			g2 = -g
		}
		if b < 0 {
			b2 = -b
		}
		return (r2 + g2 + b2) / 2
	}
	maxDist := 0
	for _, step := range strings.Split(input, ",") {
		switch step {
		case "se":
			r++
			g--
		case "ne":
			g--
			b++
		case "n":
			r--
			b++
		case "nw":
			r--
			g++
		case "sw":
			g++
			b--
		case "s":
			r++
			b--
		default:
			log.Fatal(step)
		}
		if dist() > maxDist {
			maxDist = dist()
		}
	}
	if part2 {
		return maxDist
	}
	return dist()
})
