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
	var n int

	flag.IntVar(&n, "n", 10, "Number of lines to print")
	flag.Parse()

	// Check if the number of arguments is correct
	if len(flag.Args()) != 1 {
		log.Fatal("Usage: tail <filename> or tail -n <number> <filename>")
	}

	fileName = flag.Args()[0]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	start := len(lines) - n
	if start < 0 {
		start = 0
	}

	for _, line := range lines[start:] {
		fmt.Println(line)
	}
}
