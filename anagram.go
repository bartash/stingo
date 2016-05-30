package main

import (
	"fmt"
	"os"
	"bufio"
	"github.com/bartash/stingo/runesort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("Anagrams....\n")

	file, err := os.Open("c:/cygwin64/usr/dict/words")
	check(err)

	scanner := bufio.NewScanner(file)
	count := 0;
	for scanner.Scan() {
		count++;
		text := scanner.Text()
		sortedText := runesort.SortString(text)
		if (count % 1000 == 0) {
			fmt.Printf("text %v maps to %v\n", text, sortedText)
		}
	}
	fmt.Printf("Total number of strings was %v", count)
	check(scanner.Err())
}
