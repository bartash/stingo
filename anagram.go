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
	i := 0;
	for scanner.Scan() {
		i++;
		if (i % 1000 == 0) {
			fmt.Println(scanner.Text())
		}
	}
	fmt.Printf("Total number of strings was %v", i)
	check(scanner.Err())
}
