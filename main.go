package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var begin = flag.String("b", "", "begin of slice")
var end = flag.String("e", "", "end of slice")
var printfile = flag.Bool("p", false, "print filename")
var inclusive = flag.Bool("i", false, "print the end delimiter")

func main() {
	flag.Usage = usage
	flag.Parse()

	for _, file := range flag.Args() {
		if *printfile {
			fmt.Println(file)
		}
		slicef(file, *begin, *end, func(line string) {
			fmt.Println(line)
		})
	}
}

func slicef(filename, begin, end string, fn func(string)) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	var inBlock bool
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, begin):
			inBlock = true
			fn(line)
			continue
		case end == ":EOF:":
			continue
		case strings.HasPrefix(line, end):
			if *inclusive {
				fn(line)
			}
			return
		case inBlock:
			fn(line)
		}
	}
}

func usage() {
	fmt.Println(strings.TrimSpace(`
slicef slices a file from a begin line to an end line.

usage:
  slicef [OPTIONS] [FILE ...]

example:
  slicef -b 'func main' -e '}' -i -p main.go
  main.go
  func main() {
    // code in func main
  }

options:
  -b delim
	  prefix of delimiter of the start of the block. The first line matching
	  this prefix will mark the start and will be printed.

  -e delim
	  prefix of delimiter of the end of the block. The first line matching this
	  prefix will mark the end of the block and will NOT be printed unless -i
	  flag is set. If special delim ':EOF:' is given it'll print the entire file
	  starting from the beginning delimiter.

  -i (default: false)
      prints the end delimiter as well

  -p (default: false)
      print the filename
	`))
}
