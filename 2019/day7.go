package main

import "fmt"

func init() {
	days[7] = d7
}

func d7() {
	digs := inputCommaInts()
	base := intProg{data: digs, input: []int{0, 0}}

	create := func(seqs ...int) []*intProg {
		progs := make([]*intProg, len(seqs))
		for i, s := range seqs {
			progs[i] = base.copy()
			progs[i].input[0] = s
		}
		return progs
	}
	runInChain := func(progs []*intProg) int {
		out := 0
		for _, p := range progs {
			p.input[1] = out
			p.run()
			v(p)
			out = p.output[0]
		}
		return out
	}
	maxGain := 0
	tryPermutation := func(p []int) {
		progs := create(p...)
		gain := runInChain(progs)
		if gain > maxGain {
			maxGain = gain
			fmt.Println(p, gain)
		}
	}
	permute(nil, map[int]bool{}, tryPermutation, 0, 1, 2, 3, 4)
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
