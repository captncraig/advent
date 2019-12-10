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
		for i, p := range progs {
			p.input = append(p.input, out)
			op := p.run(*p2)
			v(i, op, p.output)
			out = p.output[len(p.output)-1]
		}
		return out
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
	// if !*p2 {
	// 	permute(nil, map[int]bool{}, tryPermutation, 0, 1, 2, 3, 4)
	// } else {
	// 	permute(nil, map[int]bool{}, tryPermutation, 5, 6, 7, 8, 9)
	// }
	tryPermutation([]int{4, 3, 2, 1, 0})
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
