package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

var days = map[int]func(){}

var verbose = flag.Bool("v", false, "verbose mode")
var p2 = flag.Bool("p2", false, "part 2")
var day = flag.Int("d", 0, "day to run")
var inFile = flag.String("in", "", "file to use for input. Default is d#")

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

func inputRaw() string {
	f := *inFile
	if f == "" {
		f = fmt.Sprintf("d%d", *day)
	}
	f = filepath.Join("inputs", f)
	v("Opening", f)
	dat, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	return string(dat)
}

func inputLines() []string {
	in := strings.Replace(inputRaw(), "\r\n", "\n", -1)
	return strings.Split(in, "\n")
}
func inputLinesN(n int) []string {
	in := inputLines()
	if len(in) != n {
		log.Fatalf("Have %d lines, expect %s", len(in), n)
	}
	return in
}
func inputCommaInts() []int {
	parts := strings.Split(strings.TrimSpace(inputRaw()), ",")
	out := make([]int, len(parts))
	for i, p := range parts {
		n, _ := strconv.Atoi(p)
		out[i] = n
	}
	return out
}
