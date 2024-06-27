package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

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

	var numberOfLines bool
	flag.BoolVar(&numberOfLines, "n", false, "Numbers output lines")
	flag.Parse()

	if len(flag.Args()) == 0 {
		printFileContents(os.Stdin, numberOfLines)
		return
	}

	for _, arg := range flag.Args() {
		file, err := os.Open(arg)
		if err != nil {
			log.Fatal(err)
		}

		printFileContents(file, numberOfLines)
		defer file.Close()
	}

}
