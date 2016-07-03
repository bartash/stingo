package main

import (
	"bufio"
	"fmt"
	"github.com/bartash/stingo/runesort"
	"github.com/bartash/stingo/partitionString"
	"os"
	"strings"
	"flag"
	"unicode"
	"net/http"
	"io/ioutil"
	"bytes"
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
	surnames := flag.Bool("surnames", false, "use long list of surnames")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Printf("#args= %v\n", len(flag.Args()))
		panic("usage: anagram word")
	}
	target := strings.ToLower(flag.Arg(0))
	sortedTarget := runesort.SortString(target)

	sortToOriginal := make(map[string][]string)

	count := addFileContents(sortToOriginal, "c:/cygwin64/usr/dict/words")

	// urls from https://stackoverflow.com/questions/1803628/raw-list-of-person-names
	count += addHttpContents(sortToOriginal, "http://deron.meranda.us/data/census-dist-male-first.txt")
	count += addHttpContents(sortToOriginal, "http://deron.meranda.us/data/census-dist-female-first.txt")

	if *surnames {
		count += addHttpContents(sortToOriginal, "http://www2.census.gov/topics/genealogy/1990surnames/dist.all.last")
	}

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
		fmt.Printf("%v %v\n", upperCaseName(key.First), upperCaseName(key.Second))
	}
}

func addHttpContents( hash map[string][]string, url string) int {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// FIXME remove duplicates
	scanner := bufio.NewScanner(bytes.NewReader(body))
	count := 0

	for scanner.Scan() {
		count++
		text := scanner.Text()
		words := strings.Fields(text)
		lower := strings.ToLower(words[0])
		sortedText := runesort.SortString(lower)
		hash[sortedText] = append(hash[sortedText], lower)
	}
	check(scanner.Err())
	return count
}


func addFileContents( hash map[string][]string, fileName string) int {
	file, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		count++
		text := scanner.Text()
		sortedText := runesort.SortString(text)
		hash[sortedText] = append(hash[sortedText], text)
	}
	check(scanner.Err())
	return count
}

func upperCaseName(name string) string {
	nameRune := []rune(name)
	nameRune[0] = unicode.ToUpper(nameRune[0])
	nameWithUpper := string(nameRune)
	return nameWithUpper
}
