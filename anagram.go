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
		fmt.Println(scanner.Text())
		i++;
		if (i > 20) {
			break;
		}
	}

	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}

	check(scanner.Err())
}
