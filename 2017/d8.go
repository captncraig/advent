package main

import (
	"log"
	"strconv"
	"strings"
)

/*
--- Day 8: I Heard You Like Registers ---
You receive a signal directly from the CPU. Because of your recent assistance with jump instructions, it would like you to compute the result of a series of unusual register instructions.

Each instruction consists of several parts: the register to modify, whether to increase or decrease that register's value, the amount by which to increase or decrease it, and a condition. If the condition fails, skip the instruction without modifying the register. The registers all start at 0. The instructions look like this:

b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10
These instructions would be processed as follows:

Because a starts at 0, it is not greater than 1, and so b is not modified.
a is increased by 1 (to 1) because b is less than 5 (it is 0).
c is decreased by -10 (to 10) because a is now greater than or equal to 1 (it is 1).
c is increased by -20 (to -10) because c is equal to 10.
After this process, the largest value in any register is 1.

You might also encounter <= (less than or equal to) or != (not equal to). However, the CPU doesn't have the bandwidth to tell you what all the registers are named, and leaves that to you to determine.

What is the largest value in any register after completing the instructions in your puzzle input?
*/
/*
--- Part Two ---
To be safe, the CPU also needs to know the highest value held in any register during this process so that it can decide how much memory to allocate to these operations.
For example, in the above instructions, the highest value ever held was 10 (in register c after the third instruction was evaluated).
*/
var _ = d(8, func(part2 bool, input string) int {
	registers := map[string]int{}
	anyMax := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " if ")
		lhs, rhs := strings.Split(parts[0], " "), strings.Split(parts[1], " ")
		b, _ := strconv.Atoi(rhs[2])
		a := registers[rhs[0]]
		doLeft := false
		switch rhs[1] {
		case "!=":
			doLeft = a != b
		case ">":
			doLeft = a > b
		case ">=":
			doLeft = a >= b
		case "<":
			doLeft = a < b
		case "<=":
			doLeft = a <= b
		case "==":
			doLeft = a == b
		default:
			log.Fatalf("Unknown rhs comparision %s", rhs[1])
		}
		if doLeft {
			b, _ = strconv.Atoi(lhs[2])
			if lhs[1] == "dec" {
				b = -b
			}
			newVal := registers[lhs[0]] + b
			registers[lhs[0]] = newVal
			if newVal > anyMax {
				anyMax = newVal
			}
		}
	}
	max := 0
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	if part2 {
		return anyMax
	}
	return max
})
