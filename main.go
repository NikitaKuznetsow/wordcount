package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var u bool
	flag.BoolVar(&u, "u", false, "display in uppercase")
	flag.Parse()

	values := flag.Args()
	if values[0] == "" {
		fmt.Println(0)
	} else {
		fmt.Println(len(strings.Split(values[0], " ")))
	}
}

