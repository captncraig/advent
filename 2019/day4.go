package main

import (
	"fmt"
)

func init() {
	days[4] = d4
}

func d4() {
	count := 0
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
		ok = false
		for _, v := range counts {
			if v == 2 || (v >= 2 && !*p2) {
				ok = true
				break
			}
		}
		if !ok {
			continue
		}
		count++
		v(i)
	}
	fmt.Println(count)
}
