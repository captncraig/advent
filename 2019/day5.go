package main

func init() {
	days[5] = d5
}

func d5() {
	digs := inputCommaInts()
	//v(digs)
	prog := intProg{data: digs}
	//v(prog)
	for {
		op := prog.step()
		//v(op, prog)
		if op == 99 {
			break
		}
	}
}
