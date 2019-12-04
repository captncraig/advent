package main

import (
	"flag"
	"fmt"
)

var verbose = flag.Bool("v", false, "verbose mode")

func main() {
	flag.Parse()
	countA := 0
	countB := 0
	for i := 165432; i <= 707912; i++ {
		si := fmt.Sprint(i)
		ok := true
		counts := map[byte]int{si[0]: 1}
		for j := 1; j < len(si); j++ {
			counts[si[j]]++
			if si[j] < si[j-1] {
				ok = false
				break
			}
		}

		if !ok {
			continue
		}
		okA := false
		okB := false
		for _, v := range counts {
			if v >= 2 {
				okA = true
			}
			if v == 2 {
				okB = true
				break
			}
		}
		if okA {
			countA++
		}
		if okB {
			countB++
		}
		if !okA && !okB {
			continue
		}
		v(i, okA, okB)
	}
	fmt.Println(countA, countB)
}

func v(i ...interface{}) {
	if *verbose {
		fmt.Println(i...)
	}
}
