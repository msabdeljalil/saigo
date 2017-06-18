package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	//	"sort"
	"strings"
)

func main() {
	file_name := os.Args[1]
	file, _ := os.Open(file_name)
	// Create Scanner object to read file contents
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	rgxp, _ := regexp.Compile("[[:punct:]]")
	words_hash := make(map[string]int)
	count := 0

	for scanner.Scan() {
		// Replace all punctuation with empty string
		word := string(rgxp.ReplaceAll(scanner.Bytes(), []byte("")))
		word = strings.ToLower(word)
		words_hash[word]++
		println(word)
		count++
	}
	check(scanner.Err(), "reading input")

	// sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	fmt.Printf("%d\n", count)
	fmt.Printf("%d\n", len(words_hash))
}

func count_words() {
	words_hash := make(map[string]int)
	println(words_hash)
}

func check(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg, err)
	}
}

func print_words_hash(words_hash map[string]int) {
	for key, val := range words_hash {
		println(key, val)
	}
}
