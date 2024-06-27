package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	var fileName string
	var numberOfLines int

	flag.IntVar(&numberOfLines, "n", 10, "Number of lines to print")
	flag.Parse()

	// Check if the number of arguments is correct
	if len(flag.Args()) != 1 {
		log.Fatal("Usage: head <filename> or head -n <filename> <number>")
	}

	fileName = flag.Args()[0]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; i < numberOfLines && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
