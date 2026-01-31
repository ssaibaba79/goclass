package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	words := make(map[string]int)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println("unique words:", len(words))

	type kv struct {
		word  string
		count int
	}
	var wordCounts []kv

	for k, v := range words {
		wordCounts = append(wordCounts, kv{k, v})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count < wordCounts[j].count
	})

	for _, v := range wordCounts {
		fmt.Fprintln(os.Stdout, "Word:", v.word, "Count:", v.count)
	}

}
