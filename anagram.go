package main

import (
	"fmt"
	"os"
	"bufio"
	"github.com/bartash/stingo/runesort"
	"bytes"
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
	target := os.Args[1]
	sortedTarget := runesort.SortString(target)
	fmt.Printf("Anagrams of %v\n", target)


	file, err := os.Open("c:/cygwin64/usr/dict/words")
	check(err)

	scanner := bufio.NewScanner(file)
	count := 0;
	sortToOriginal := make(map[string][]string)
	for scanner.Scan() {
		count++;
		text := scanner.Text()
		sortedText := runesort.SortString(text)
		sortToOriginal[sortedText] = append(sortToOriginal[sortedText], text)
	}
	fmt.Printf("Total number of strings was %v map contains %v\n", count, len(sortToOriginal))
	check(scanner.Err())

	answers := sortToOriginal[sortedTarget];
	for _, element := range answers {
		fmt.Printf("answer: %v \n", element)
	}


	sort2ToOriginal := make(map[string] bool)
	for key1, _ := range sortToOriginal {
		for key2, _ := range sortToOriginal {
			var buffer bytes.Buffer
			buffer.WriteString(key1)
			buffer.WriteString(key2)
			sorted2Text := runesort.SortString(buffer.String());
			sort2ToOriginal[sorted2Text] = true;
		}
	}
	fmt.Printf("Total number of strings in double map %v\n", len(sort2ToOriginal))
}
