package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Defining a boolean flag -l to count lines
	lines := flag.Bool("l", false, "Count lines")

	// Defining a boolean flag -b to count bytes
	bytes := flag.Bool("b", false, "Count bytes")

	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of words
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the countLines flag is not set, we check to see if
	// the countBytes flag is set. If it isn't set we count by words. If it is set, then we count by bytes. If countLines flag is set, none of the other statements are run and the Split function follows its default of counting by lines.

	if !countLines {
		if !countBytes {
			scanner.Split(bufio.ScanWords)
		} else {
			scanner.Split(bufio.ScanBytes)
		}
	}

	// Defining a counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}
