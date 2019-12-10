package main

import (
	"fmt"
	"log"
)

type intProg struct {
	data   []int
	pc     int
	input  []int
	output []int
}

func (i *intProg) copy() *intProg {
	i2 := &intProg{
		data:  make([]int, len(i.data)),
		pc:    i.pc,
		input: make([]int, len(i.input)),
	}
	copy(i2.data, i.data)
	copy(i2.input, i.input)
	return i2
}

func (i *intProg) param(rel int, op int) int {
	mode := op / 100
	thisMode := "0"
	if mode != 0 {
		modeStr := fmt.Sprint(mode)
		idx := len(modeStr) - rel
		if idx >= 0 {
			thisMode = string(modeStr[idx])
		}
	}
	//v(thisMode, "???", rel, op)
	if thisMode == "0" {
		return i.data[i.data[i.pc+rel]]
	} else if thisMode == "1" {
		return i.data[i.pc+rel]
	}
	log.Fatal("BAD MODE", thisMode)
	return 0
}

func (i *intProg) run(haltOnOut bool) int {
	for {
		op := i.step()
		if op == 4 && haltOnOut {
			return op
		}
		if op == 99 {
			return op
		}
	}
}

func (i *intProg) step() int {
	op := i.data[i.pc]
	adv := 1
	//v(op%10, "!!!", op)
	switch op % 100 {
	case 1: //+
		a1 := i.param(1, op)
		a2 := i.param(2, op)
		r := i.data[i.pc+3]
		i.data[r] = a1 + a2
		adv = 4
	case 2: //*
		a1 := i.param(1, op)
		a2 := i.param(2, op)
		r := i.data[i.pc+3]
		i.data[r] = a1 * a2
		adv = 4
	case 3: //IN
		r := i.data[i.pc+1]
		adv = 2
		i.data[r] = i.input[0]
		i.input = i.input[1:]
		if *p2 {
			i.data[r] = 5
		}
	case 4: //OUT
		a1 := i.param(1, op)
		i.output = append(i.output, a1)
		adv = 2

	case 5: // JNZ
		// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer
		// to the value from the second parameter. Otherwise, it does nothing.
		a1 := i.param(1, op)
		r := i.param(2, op)
		adv = 3
		if a1 != 0 {
			i.pc = r
			adv = 0
		}
	case 6: // JZ
		// Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer
		// to the value from the second parameter. Otherwise, it does nothing.
		a1 := i.param(1, op)
		r := i.param(2, op)
		adv = 3
		if a1 == 0 {
			i.pc = r
			adv = 0
		}
	case 7: // LT
		// Opcode 7 is less than: if the first parameter is less than the second parameter,
		// it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
		a1 := i.param(1, op)
		a2 := i.param(2, op)
		r := i.data[i.pc+3]
		if a1 < a2 {
			i.data[r] = 1
		} else {
			i.data[r] = 0
		}
		adv = 4
	case 8:
		// Opcode 8 is equals: if the first parameter is equal to the second parameter,
		// it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
		a1 := i.param(1, op)
		a2 := i.param(2, op)
		r := i.data[i.pc+3]
		if a1 == a2 {
			i.data[r] = 1
		} else {
			i.data[r] = 0
		}
		adv = 4
	case 99:
		//v("HALT")
	default:
		log.Fatalf("unkown opcode %d", op)
	}
	i.pc += adv
	return op
}
