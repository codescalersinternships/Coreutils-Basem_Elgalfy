package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		for {
			fmt.Println("y")
		}
	} else {

		out := strings.Join(os.Args[1:], " ")
		for {
			fmt.Println(out)
		}
	}
}
