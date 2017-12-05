package main

import (
	"strconv"
	"strings"
)

/*

 */
/*

 */
var _ = d(5, func(part2 bool, input string) int {
	arr := []int{}
	for _, line := range strings.Split(input, "\r\n") {
		i, _ := strconv.Atoi(line)
		arr = append(arr, i)
	}
	i := 0
	steps := 0
	for {
		if i < 0 || i >= len(arr) {
			break
		}
		jmp := arr[i]
		if part2 && jmp > 2 {
			arr[i]--
		} else {
			arr[i]++
		}
		i += jmp
		steps++
	}
	return steps
})
