package main

import (
	"bufio"
	"fmt"
	"github.com/bartash/stingo/runesort"
	"github.com/bartash/stingo/partitionString"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// find simple anagrams
func main() {
	if len(os.Args) < 2 {
		panic("usage: anagram word")
	}
	target := strings.ToLower(os.Args[1])
	sortedTarget := runesort.SortString(target)
	fmt.Printf("Anagrams of %v\n", target)

	file, err := os.Open("c:/cygwin64/usr/dict/words")
	check(err)

	scanner := bufio.NewScanner(file)
	count := 0
	sortToOriginal := make(map[string][]string)
	for scanner.Scan() {
		count++
		text := scanner.Text()
		sortedText := runesort.SortString(text)
		sortToOriginal[sortedText] = append(sortToOriginal[sortedText], text)
	}
	fmt.Printf("Total number of strings was %v map contains %v\n", count, len(sortToOriginal))
	check(scanner.Err())

	answers := sortToOriginal[sortedTarget]
	for _, element := range answers {
		fmt.Printf("answer: %v \n", element)
	}

	second := partitionString.NewPartitionString(target)
	allAnswers := make(map[partitionString.Pair]bool)
	pair := second.Next()
	for pair != nil {
		// fmt.Printf("got pair '%v' & '%v'\n", pair.First, pair.Second)
		sortedFirst := runesort.SortString(pair.First)
		sortedSecond := runesort.SortString(pair.Second)

		sortedFirstAnswers := sortToOriginal[sortedFirst]
		sortedSecondAnswers := sortToOriginal[sortedSecond]

		if len(sortedFirstAnswers) > 0 && len (sortedSecondAnswers) > 0 {
			for _, element := range sortedFirstAnswers {
				for _, element2 := range sortedSecondAnswers {
					pair := partitionString.NewPair(element, element2)
					allAnswers[*pair] = true
				}
			}
		}

		pair = second.Next();
	}
	for key, _ := range allAnswers {
		fmt.Printf("%v %v\n", key.First, key.Second)
	}
}
