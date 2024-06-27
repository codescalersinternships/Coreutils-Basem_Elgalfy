package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printFileContents(reader io.Reader, n bool) {

	scanner := bufio.NewScanner(reader)
	lineNumber := 1

	for scanner.Scan() {
		if n {
			fmt.Printf("    %d ", lineNumber)
		}
		fmt.Println(scanner.Text())
		lineNumber++
	}

}
func main() {

	var n bool
	flag.BoolVar(&n, "n", false, "Numbers output lines")
	flag.Parse()

	if len(flag.Args()) == 0 {
		printFileContents(os.Stdin, n)
		return
	}

	for _, arg := range flag.Args() {
		file, err := os.Open(arg)
		checkError(err)

		printFileContents(file, n)
		defer file.Close()
	}

}
