package main

import (
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

func init() {
	if err := os.Mkdir("inputs", 0644); err != nil {
		if !os.IsExist(err) {
			log.Fatal(err)
		}
	}
	for i := 1; i <= 25; i++ {
		fname := inpfile(i)
		if _, err := os.Stat(fname); err != nil {
			if os.IsNotExist(err) {
				u := fmt.Sprintf("https://adventofcode.com/2017/day/%d/input", i)
				req, err := http.NewRequest("GET", u, nil)
				req.AddCookie(&http.Cookie{
					Name: "session",
				})
				log.Fatal("TODO: add cookie above and remove this line")
				if err != nil {
					log.Fatal(err)
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Fatal(err)
				}
				if resp.StatusCode == 404 {
					break
				}
				if resp.StatusCode == 400 {
					log.Fatal("BAD LOGIN")
				}
				f, err := os.Create(fname)
				if err != nil {
					log.Fatal(err)
				}

				if _, err = io.Copy(f, resp.Body); err != nil {
					log.Fatal(err)
				}
				f.Close()
				resp.Body.Close()
			} else {
				log.Fatal(err)
			}
		}
		cfname := codefile(i)
		if _, err := os.Stat(cfname); err != nil {
			if os.IsNotExist(err) {
				tmpl := `package main
				import("fmt")
				/*

				*/
				/*

				*/
				 var _ = d(%d, func(part2 bool, input string) int {
					 	fmt.Println(%d, part2, input)
					 	return 0
					 })
				`
				code := fmt.Sprintf(tmpl, i, i)
				cbytes, _ := format.Source([]byte(code))
				ioutil.WriteFile(cfname, cbytes, 0644)
			} else {
				log.Fatal(err)
			}
		}
	}
}
