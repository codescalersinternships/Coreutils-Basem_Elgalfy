package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var n bool
	flag.BoolVar(&n, "n", false, "If specified omits the trailing newline")
	flag.Parse()

	out := strings.Join(flag.Args(), " ")
	fmt.Print(out)

	if !n {
		fmt.Println()
	}

}
