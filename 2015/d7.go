package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
--- Day 7: Some Assembly Required ---

This year, Santa brought little Bobby Tables a set of wires and bitwise logic gates! Unfortunately, little Bobby is a little under the recommended age range, and he needs help assembling the circuit.

Each wire has an identifier (some lowercase letters) and can carry a 16-bit signal (a number from 0 to 65535). A signal is provided to each wire by a gate, another wire, or some specific value. Each wire can only get a signal from one source, but can provide its signal to multiple destinations. A gate provides no signal until all of its inputs have a signal.

The included instructions booklet describes how to connect the parts together: x AND y -> z means to connect wires x and y to an AND gate, and then connect its output to wire z.

For example:

123 -> x means that the signal 123 is provided to wire x.
x AND y -> z means that the bitwise AND of wire x and wire y is provided to wire z.
p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 and then provided to wire q.
NOT e -> f means that the bitwise complement of the value from wire e is provided to wire f.
Other possible gates include OR (bitwise OR) and RSHIFT (right-shift). If, for some reason, you'd like to emulate the circuit instead, almost all programming languages (for example, C, JavaScript, or Python) provide operators for these gates.

For example, here is a simple circuit:

123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
After it is run, these are the signals on the wires:

d: 72
e: 507
f: 492
g: 114
h: 65412
i: 65079
x: 123
y: 456
In little Bobby's kit's instructions booklet (provided as your puzzle input), what signal is ultimately provided to wire a?
*/
/*
--- Part Two ---

Now, take the signal you got on wire a, override wire b to that signal, and reset the other wires (including wire a). What new signal is ultimately provided to wire a?
*/
var _ = d(7, func(part2 bool, input string) int {
	wires := map[string]string{}
	memos := map[string]uint16{}

	var eval func(s string) uint16
	eval = func(s string) (v uint16) {
		if val, ok := memos[s]; ok {
			return val
		}
		defer func() {
			memos[s] = v
		}()
		parts := strings.Split(s, " ")
		if len(parts) == 1 {
			if i, err := strconv.Atoi(s); err == nil {
				return uint16(i)
			}
			return eval(wires[s])
		}
		if len(parts) == 3 {
			l := eval(parts[0])
			r := eval(parts[2])
			switch parts[1] {
			case "LSHIFT":
				return l << r
			case "RSHIFT":
				return l >> r
			case "OR":
				return l | r
			case "AND":
				return l & r
			default:
				log.Fatal("binop ", parts[1])
			}
		}
		if parts[0] == "NOT" {
			return ^eval(parts[1])
		}
		log.Fatal("Don't know how to eval")
		return 0
	}
	for _, line := range strings.Split(input, "\n") {
		sides := strings.Split(line, " -> ")
		result := sides[1]
		if wires[result] != "" {
			log.Fatal("wire defined twice:", result)
		}
		wires[result] = sides[0]
	}
	fmt.Println(wires)
	res := eval(wires["a"])
	if part2 {
		memos = map[string]uint16{}
		wires["b"] = fmt.Sprint(res)
		res = eval(wires["a"])
	}
	return int(res)
})
