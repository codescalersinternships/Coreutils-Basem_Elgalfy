package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var fileName string
	var n int

	// Check if the number of arguments is correct
	if len(os.Args) != 4 && len(os.Args) != 2 {
		log.Fatal("Usage: tail <filename> or tail -n <filename> <number>")
	}

	if len(os.Args) == 4 {
		if os.Args[1] != "-n" {
			log.Fatal("Invalid option provided " + os.Args[1])
		}

		fileName = os.Args[2]
		numberOfLines, err := strconv.Atoi(os.Args[3])
		checkError(err)
		n = numberOfLines
	} else {
		fileName = os.Args[1]
		n = 10
	}

	file, err := os.Open(fileName)
	checkError(err)
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
