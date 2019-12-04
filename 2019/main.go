package main

import (
	"flag"
	"fmt"
	"log"
)

var days = map[int]func(){}

var verbose = flag.Bool("v", false, "verbose mode")
var p2 = flag.Bool("p2", false, "part 2")
var day = flag.Int("d", 0, "day to run")

func main() {
	flag.Parse()
	if *day == 0 {
		flag.PrintDefaults()
		return
	}
	f := days[*day]
	if f == nil {
		log.Fatalf("Day %d does not exist", *day)
	}
	f()
}

func v(i ...interface{}) {
	if *verbose {
		fmt.Println(i...)
	}
}
