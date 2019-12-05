package main

func init() {
	days[2] = d2
}

func d2() {
	digs := inputCommaInts()
	v(digs)
	prog := intProg{data: digs}
	v(prog)
	for {
		op := prog.step()
		v(op, prog)
		if op == 99 {
			break
		}
	}
}
