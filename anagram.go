package main

import (
	"fmt"
	"os"
	"bufio"
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
		if (count % 1000 == 0) {
			fmt.Println(text)
		}
	}
	fmt.Printf("Total number of strings was %v", count)
	check(scanner.Err())
}
