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
	if (len(os.Args) < 2) {
		panic("usage: anagram word")
	}
	arg := os.Args[1]
	fmt.Printf("Anagrams of %v\n", arg)


	file, err := os.Open("c:/cygwin64/usr/dict/words")
	check(err)

	scanner := bufio.NewScanner(file)
	count := 0;
	sortToOriginal := make(map[string]string)
	for scanner.Scan() {
		count++;
		text := scanner.Text()
		sortedText := runesort.SortString(text)
		sortToOriginal[sortedText] = text;
		if (count % 10000 == 0) {
			fmt.Printf("text %v maps to %v\n", text, sortedText)
		}
	}
	fmt.Printf("Total number of strings was %v map contains %v", count, len(sortToOriginal))
	check(scanner.Err())
}
