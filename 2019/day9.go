package main

import "fmt"

func init() {
	days[9] = d9
}

func d9() {
	digs := inputCommaInts()
	prog := intProg{data: digs, input: []int{1}}
	if *p2 {
		prog.input[0] = 2
	}
	prog.run(false)
	fmt.Println(prog.output)
}
