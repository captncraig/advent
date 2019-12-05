package main

import "fmt"

func init() {
	days[2] = d2
}

func d2() {
	digs := inputCommaInts()
	newSlice := make([]int, len(digs))
	exe := func(n, v int) int {
		copy(newSlice, digs)
		prog := intProg{data: newSlice}
		prog.data[1] = n
		prog.data[2] = v
		for {
			op := prog.step()
			if op == 99 {
				break
			}
		}
		return prog.data[0]
	}
	if !*p2 {
		fmt.Println(">>>>>", exe(12, 2))
		return
	}
	goal := 19690720
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			x := exe(n, v)
			if x == goal {
				fmt.Println(">>>>>", n, v, 100*n+v)
				return
			}
		}
	}
}
