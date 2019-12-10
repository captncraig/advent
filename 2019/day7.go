package main

import "fmt"

func init() {
	days[7] = d7
}

func d7() {
	digs := inputCommaInts()
	base := intProg{data: digs, input: []int{0}}

	create := func(seqs ...int) []*intProg {
		progs := make([]*intProg, len(seqs))
		for i, s := range seqs {
			progs[i] = base.copy()
			progs[i].input[0] = s
		}
		return progs
	}
	runChain := func(progs []*intProg) int {
		out := 0
		for {
			//v(progs[0].data, out)
			for i, p := range progs {
				p.input = append(p.input, out)
				//v(i, p.input)
				op := p.run(*p2)
				//v(p.output)
				out = p.output[len(p.output)-1]
				if op != 4 {
					//v("HHH", i)
					if i == 4 {
						return out
					}
				}
				//v(i, p.pc, op, p.output, p.input)

			}
		}
	}
	maxGain := 0
	tryPermutation := func(p []int) {
		progs := create(p...)
		var gain int
		gain = runChain(progs)
		if gain > maxGain {
			maxGain = gain
			fmt.Println(p, gain)
		}
	}
	if !*p2 {
		//tryPermutation([]int{4, 3, 2, 1, 0})
		permute(nil, map[int]bool{}, tryPermutation, 0, 1, 2, 3, 4)
	} else {
		//tryPermutation([]int{9, 8, 7, 6, 5})
		permute(nil, map[int]bool{}, tryPermutation, 5, 6, 7, 8, 9)
	}

}

func permute(path []int, taken map[int]bool, f func([]int), possible ...int) {
	if len(path) == len(possible) {
		v(path)
		f(path)
		return
	}
	for _, i := range possible {
		if taken[i] {
			continue
		}
		taken[i] = true
		permute(append(path, i), taken, f, possible...)
		taken[i] = false
	}
}
