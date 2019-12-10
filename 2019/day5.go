package main

import "fmt"

func init() {
	days[5] = d5
}

func d5() {
	digs := inputCommaInts()
	//v(digs)
	inp := 1
	if *p2 {
		inp = 5
	}
	prog := intProg{data: digs, input: []int{inp}}
	//v(prog)
	for {
		op := prog.step()
		//v(op, prog)
		if op == 99 {
			break
		}
	}
	fmt.Println(prog.output)
}
