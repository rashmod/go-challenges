package main

import (
	"fmt"
	"os"
)

func main() {
	str := os.Args[1:]
	freq := make(map[rune]int)

	for _, s := range str {
		for _, r := range s {
			freq[r]++
		}
	}

	for k, v := range freq {
		fmt.Printf("%c: %v\n", k, v)
	}
}
