package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	var fileName string
	var n int

	flag.IntVar(&n, "n", 10, "Number of lines to print")
	flag.Parse()

	// Check if the number of arguments is correct
	if len(os.Args) != 4 && len(os.Args) != 2 {
		log.Fatal("Usage: head <filename> or head -n <filename> <number>")
	}

	if len(os.Args) == 4 {
		if os.Args[1] != "-n" {
			log.Fatal("Invalid option provided " + os.Args[1])
		}

		fileName = os.Args[3]
	} else {
		fileName = os.Args[1]
	}

	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; i < n && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
