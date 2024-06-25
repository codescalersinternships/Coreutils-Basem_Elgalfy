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

func countWords(line string) int {

	var numberOfWords int
	var isWord bool

	for _, char := range line {

		if char == ' ' || char == '\n' || char == '\t' {
			isWord = false
		} else {
			if !isWord {
				numberOfWords++
			}
			isWord = true
		}
	}

	return numberOfWords
}

func wc(file io.Reader) (w int, c int, l int, e error) {

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines, words, bytes int

	for scanner.Scan() {

		currentLine := scanner.Text()
		lines++
		bytes += len(currentLine) + 1
		words += countWords(currentLine)

	}

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, err
	}

	return lines, words, bytes, nil

}

func main() {

	var l, w, c bool
	flag.BoolVar(&l, "l", false, "Print the number of lines in each input file.")
	flag.BoolVar(&w, "w", false, "Print the number of words in each input file.")
	flag.BoolVar(&c, "c", false, "Print the number of bytes in each input file.")
	flag.Parse()

	filename := flag.Args()[0]

	file, err := os.Open(filename)
	checkError(err)
	defer file.Close()

	lines, words, bytes, err := wc(file)
	checkError(err)

	if l {
		fmt.Printf("Lines: %d\n", lines)
	}
	if w {
		fmt.Printf("Words: %d\n", words)
	}
	if c {
		fmt.Printf("Bytes: %d\n", bytes)
	}

	if len(flag.Args()) == 1 {
		fmt.Printf("%d %d %d ", lines, words, bytes)
	}

	fmt.Printf("%s \n", filename)

}
