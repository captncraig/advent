package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var days = make([]func(bool, string) int, 26)

func d(i int, f func(bool, string) int) bool {
	days[i] = f
	return true
}

// run as `go run *.go dayNum partNum`
func main() {
	if len(os.Args) != 2 {
		log.Fatal("USAGE: go run *.go dayNumber[+]")
	}
	num := os.Args[1]
	p2 := false
	if num[len(num)-1] == '+' {
		p2 = true
		num = num[:len(num)-1]
	}
	n, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	input, err := ioutil.ReadFile(inpfile(n))
	if err != nil {
		log.Fatal(err)
	}
	sin := strings.TrimSpace(string(input))
	fmt.Println("The answer is.... ", days[n](p2, sin))
}

func inpfile(i int) string {
	return filepath.Join("inputs", fmt.Sprintf("day%d", i))
}
func codefile(i int) string {
	return fmt.Sprintf("d%d.go", i)
}
