package main

import (
	"fmt"
	"log"
)

type intProg struct {
	data []int
	pc   int
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
	v(thisMode, "???", rel, op)
	if thisMode == "0" {
		return i.data[i.data[i.pc+rel]]
	} else if thisMode == "1" {
		return i.data[i.pc+rel]
	}
	log.Fatal("BAD MODE", thisMode)
	return 0
}

func (i *intProg) step() int {
	op := i.data[i.pc]
	adv := 1
	v(op%10, "!!!", op)
	switch op % 10 {
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
		i.data[r] = 1 // hard coded
	case 4: //OUT
		a1 := i.param(1, op)
		fmt.Println(a1, "OOOOOOOOOOOOOO")
		adv = 2
	case 99:
	default:
		log.Fatalf("unkown opcode %d", op)
	}
	i.pc += adv
	return op
}
