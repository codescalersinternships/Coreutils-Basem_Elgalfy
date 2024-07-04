package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

func printTree(maxLevel int, path string, currentLevel int, indent string) (directoryCount int, fileCount int) {

	if currentLevel > maxLevel {
		return
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	dcount, fcount := 0, 0

	for i, file := range files {
		isLast := i == len(files)-1
		var prefix string
		if isLast {
			prefix = "└──"
		} else {
			prefix = "├──"
		}
		fmt.Println(indent + prefix + " " + file.Name())

		if file.IsDir() {
			dcount++
			newIndent := indent
			if isLast {
				newIndent += "    "
			} else {
				newIndent += "│   "
			}
			subdcount, subfcount := printTree(maxLevel, path+"/"+file.Name(), currentLevel+1, newIndent)
			dcount += subdcount
			fcount += subfcount
		} else {
			fcount++
		}
	}

	return dcount, fcount

}

func main() {

	var numberOfLevels int
	flag.IntVar(&numberOfLevels, "L", math.MaxInt, "Number of Levels to traverse")
	flag.Parse()

	if numberOfLevels <= 0 {
		log.Fatal("Number of levels must be greater than 0")
	}

	var root string
	if flag.NArg() > 1 {
		log.Fatal("Only one root directory is allowed")
	}

	if flag.NArg() == 0 {
		root = "."
	} else {
		root = flag.Arg(0)
	}

	directoryCount, fileCount := printTree(numberOfLevels, root, 1, "")
	fmt.Printf("\n%d directories, %d files\n", directoryCount, fileCount)

}
