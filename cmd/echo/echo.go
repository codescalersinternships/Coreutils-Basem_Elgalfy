package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var omitTrailingLine bool
	flag.BoolVar(&omitTrailingLine, "n", false, "If specified omits the trailing newline")
	flag.Parse()

	out := strings.Join(flag.Args(), " ")
	fmt.Print(out)

	if !omitTrailingLine {
		fmt.Println()
	}

}
