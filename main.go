package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var values string = os.Args[1]
	if values == "" {
		fmt.Println(0)
	} else {
		fmt.Println(len(strings.Split(values, " ")))
	}
}

