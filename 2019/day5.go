package main

import "fmt"

func init() {
	days[5] = d5
}

func d5() {
	digs := inputCommaInts()
	//v(digs)
	inp := make(chan int, 1)
	if *p2 {
		inp <- 5
	} else {
		inp <- 1
	}
	out := make(chan int, 100)
	prog := intProg{data: digs, input: inp, output: out}
	prog.run()
	close(out)
	for o := range out {
		fmt.Println(o)
	}
}
