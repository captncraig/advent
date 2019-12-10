package main

import (
	"fmt"
	"log"
)

type intProg struct {
	data         []int
	pc           int
	input        []int
	output       []int
	relativeBase int
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
	thisMode := i.mode(rel, op)
	i.expand(i.pc + rel)
	actual := i.data[i.pc+rel]
	if thisMode == "0" {
		i.expand(actual)
		return i.data[actual]
	} else if thisMode == "1" {
		return actual
	} else if thisMode == "2" {
		return i.data[i.relativeBase+actual]
	}
	log.Fatal("BAD MODE", thisMode)
	return 0
}

func (i *intProg) mode(rel int, op int) string {
	mode := op / 100
	thisMode := "0"
	if mode != 0 {
		modeStr := fmt.Sprint(mode)
		idx := len(modeStr) - rel
		if idx >= 0 {
			thisMode = string(modeStr[idx])
		}
	}
	return thisMode
}

func (i *intProg) rAddr(rel int, op int) int {
	thisMode := i.mode(rel, op)
	i.expand(i.pc + rel)
	actual := i.data[i.pc+rel]
	if thisMode == "0" {
		return actual
	} else if thisMode == "2" {
		return i.relativeBase + actual
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

func (i *intProg) expand(size int) {
	for len(i.data) < size+1 {
		i.data = append(i.data, 0)
	}
}

func (i *intProg) step() int {
	//v("---", i.pc, i.data, i.input, i.output)
	//defer v(">>>", i.pc, i.data, i.input, i.output)
	op := i.data[i.pc]
	adv := 1
	//v(op%10, "!!!", op)
	switch op % 100 {
	case 1: //+
		a1 := i.param(1, op)
		a2 := i.param(2, op)
		r := i.rAddr(3, op)
		i.expand(r)
		i.data[r] = a1 + a2
		adv = 4
	case 2: //*
		a1 := i.param(1, op)
		a2 := i.param(2, op)
		r := i.rAddr(3, op)
		i.expand(r)
		i.data[r] = a1 * a2
		adv = 4
	case 3: //IN
		r := i.rAddr(1, op)
		adv = 2
		i.expand(r)
		i.data[r] = i.input[0]
		i.input = i.input[1:]
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
		r := i.rAddr(3, op)
		i.expand(r)
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
		r := i.rAddr(3, op)
		i.expand(r)
		if a1 == a2 {
			i.data[r] = 1
		} else {
			i.data[r] = 0
		}
		adv = 4
	case 9:
		// adjust relative offset
		a1 := i.param(1, op)
		i.relativeBase += a1
		adv = 2
	case 99:
		//v("HALT")
	default:
		log.Fatalf("unkown opcode %d", op)
	}
	i.pc += adv
	return op
}
