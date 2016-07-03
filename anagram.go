package main

import (
	"bufio"
	"fmt"
	"github.com/bartash/stingo/runesort"
	"github.com/bartash/stingo/partitionString"
	"os"
	"strings"
	"flag"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// find simple anagrams
func main() {
	verbose := flag.Bool("verbose", false, "print verbose output")
	simple := flag.Bool("simple", false, "print simple angrams as well as pairs")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Printf("#args= %v\n", len(flag.Args()))
		panic("usage: anagram word")
	}
	target := strings.ToLower(flag.Arg(0))
	sortedTarget := runesort.SortString(target)


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
	check(scanner.Err())

	if *verbose {
		fmt.Printf("Total number of strings was %v map contains %v\n", count, len(sortToOriginal))
	}

	if *simple {
		fmt.Printf("Simple Anagrams of %v\n", target)
		answers := sortToOriginal[sortedTarget]
		for _, element := range answers {
			fmt.Printf("answer: %v \n", element)
		}
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
