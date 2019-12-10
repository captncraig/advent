package main

import "fmt"

func init() {
	days[9] = d9
}

func d9() {
	digs := inputCommaInts()
	prog := intProg{data: digs, input: []int{1}}
	prog.run(false)
	fmt.Println(prog.output)
}
