package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	values := os.Args
	if values[0] == "" {
		fmt.Println(0)
	} else {
		fmt.Println(len(strings.Split(values[0], " ")))
	}
}

